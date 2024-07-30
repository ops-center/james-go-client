package james

import (
	"fmt"
	"github.com/pkg/errors"
	"io"
	"strings"

	"git.sr.ht/~rockorager/go-jmap"
	_ "git.sr.ht/~rockorager/go-jmap/core"
	"git.sr.ht/~rockorager/go-jmap/mail"
	"git.sr.ht/~rockorager/go-jmap/mail/email"
	"git.sr.ht/~rockorager/go-jmap/mail/emailsubmission"
	"git.sr.ht/~rockorager/go-jmap/mail/mailbox"
)

func (j *JMAPClient) SendEmail(myMail *email.Email) error {
	j.muJmap.RLock()
	defer j.muJmap.RUnlock()

	const draftEmailId = "draft"
	const sendIt = "sendIt"

	var draftMailboxId, sentMailboxId jmap.ID
	{
		var err error
		if draftMailboxId, err = j.getMailboxId(DraftMailbox); err != nil {
			return err
		}
		if sentMailboxId, err = j.getMailboxId(SentMailbox); err != nil {
			return err
		}
	}

	req := &jmap.Request{
		Using: []jmap.URI{"urn:ietf:params:jmap:core", "urn:ietf:params:jmap:mail"},
	}

	myMap := map[jmap.ID]*email.Email{
		draftEmailId: myMail,
	}

	req.Invoke(&email.Set{
		Account: j.userId,
		Create:  myMap,
	})

	myEmailSubmission := emailsubmission.EmailSubmission{
		EmailID: "#" + draftEmailId,
	}

	req.Invoke(&emailsubmission.Set{
		Account: j.userId,

		Create: map[jmap.ID]*emailsubmission.EmailSubmission{
			sendIt: &myEmailSubmission,
		},

		OnSuccessUpdateEmail: map[jmap.ID]jmap.Patch{
			"#" + sendIt: {
				"mailboxIds/" + string(draftMailboxId): nil,
				"mailboxIds/" + string(sentMailboxId):  true,
				"keywords/$seen":                       nil,
				"keywords/$draft":                      nil,
			},
		},
	})

	resp, err := j.Do(req)
	if err != nil {
		return err
	}

	for _, inv := range resp.Responses {
		switch r := inv.Args.(type) {
		case *jmap.MethodError:
			return r
		case *email.SetResponse:
			if emailResponse, ok := r.Created[draftEmailId]; ok && emailResponse.ID != "" {
				return nil
			}
		}
	}

	return fmt.Errorf("email submission failed / not created")
}

func (j *JMAPClient) GetEmails(filter JmapEmailFilter) ([]*email.Email, error) {
	j.muJmap.RLock()
	defer j.muJmap.RUnlock()

	if !isEmptyString(string(filter.InMailboxWithName)) {
		if mailboxId, err := j.getMailboxId(filter.InMailboxWithName); err == nil {
			filter.InMailbox = mailboxId
		} else {
			return nil, err
		}
	}

	req := &jmap.Request{
		Using: []jmap.URI{"urn:ietf:params:jmap:core", "urn:ietf:params:jmap:mail"},
	}

	mySortComparator := email.SortComparator{
		Property:    string(filter.SortProperty),
		IsAscending: filter.AscendingOrder,
	}

	callID := req.Invoke(&email.Query{
		Account:  j.userId,
		Filter:   &filter.FilterCondition,
		Sort:     []*email.SortComparator{&mySortComparator},
		Limit:    filter.MaxEmails,
		Position: filter.Position,
	})

	req.Invoke(&email.Get{
		Account: j.userId,
		ReferenceIDs: &jmap.ResultReference{
			ResultOf: callID,
			Name:     "Email/query",
			Path:     "/ids",
		},
		Properties:          filter.HasProperties,
		FetchAllBodyValues:  filter.FetchAllBodyValues,
		FetchHTMLBodyValues: filter.FetchHTMLBodyValues,
	})

	resp, err := j.Do(req)
	if err != nil {
		return nil, err
	}

	for _, inv := range resp.Responses {
		switch r := inv.Args.(type) {
		case *jmap.MethodError:
			return nil, r
		case *email.GetResponse:
			if r != nil {
				return r.List, nil
			}
		}
	}

	return nil, fmt.Errorf("email/get response not found")
}

// DestroyAllEmails will require multiple calls if the number of emails is large enough
func (j *JMAPClient) DestroyAllEmails() (destroyed []jmap.ID, notDestroyed map[jmap.ID]*jmap.SetError, err error) {
	j.muJmap.Lock()
	defer j.muJmap.Unlock()

	req := &jmap.Request{
		Using: []jmap.URI{"urn:ietf:params:jmap:core", "urn:ietf:params:jmap:mail"},
	}

	callID0 := req.Invoke(&email.Query{
		Account: j.userId,
	})

	callID1 := req.Invoke(&email.Get{
		Account: j.userId,
		ReferenceIDs: &jmap.ResultReference{
			ResultOf: callID0,
			Name:     "Email/query",
			Path:     "/ids",
		},
		Properties: []string{"id"},
	})

	req.Invoke(&email.Set{
		Account: j.userId,
		ReferenceDestroyIDs: &jmap.ResultReference{
			ResultOf: callID1,
			Name:     "Email/get",
			Path:     "list/*/id",
		},
	})

	resp, err := j.Do(req)
	if err != nil {
		return nil, nil, err
	}

	for _, inv := range resp.Responses {
		switch r := inv.Args.(type) {
		case *jmap.MethodError:
			return nil, nil, r
		case *email.SetResponse:
			if r != nil {
				return r.Destroyed, r.NotDestroyed, nil
			}
		}
	}

	return nil, nil, fmt.Errorf("email/set response not found")
}

// James Distributed Server 3.9 doesn't generate mailboxes
// for new users until a "mailbox/get" query is made
func (j *JMAPClient) initializeMailboxIds() error {
	req := &jmap.Request{
		Using: []jmap.URI{"urn:ietf:params:jmap:core", "urn:ietf:params:jmap:mail"},
	}

	req.Invoke(&mailbox.Get{
		Account: j.userId,
	})

	resp, err := j.Do(req)
	if err != nil {
		return err
	}

	var MailboxList []*mailbox.Mailbox

	for _, inv := range resp.Responses {
		switch r := inv.Args.(type) {
		case *jmap.MethodError:
			return r
		case *mailbox.GetResponse:
			if r != nil {
				MailboxList = r.List
				break
			}
		}
	}

	j.mailboxIds = make(map[string]jmap.ID)

	for _, m := range MailboxList {
		j.mailboxIds[strings.ToLower(string(m.Role))] = m.ID
	}

	return nil
}

func (j *JMAPClient) GetTestRecipient() (string, error) {
	j.muJmap.RLock()
	defer j.muJmap.RUnlock()
	return Recipient, nil
}

type emailData struct {
	recipient, subject, bodyValue string
	customHeaders                 []*email.Header
	inReplyTo, references         []string
	attachments                   []*email.BodyPart
	client                        *JMAPClient
}

type Option func(data *emailData) error

func (j *JMAPClient) NewEmail(options ...Option) (*email.Email, error) {
	j.muJmap.RLock()
	defer j.muJmap.RUnlock()

	myMailData := &emailData{client: j}

	draftMailboxID, err := j.getMailboxId(DraftMailbox)
	if err != nil {
		return nil, err
	}

	for _, opt := range options {
		err = opt(myMailData)
		if err != nil {
			return nil, err
		}
	}

	if len(myMailData.recipient) == 0 {
		return nil, errors.New("no recipient defined")
	}

	from := mail.Address{
		Name:  j.userEmail,
		Email: j.userEmail,
	}

	to := mail.Address{
		Name:  myMailData.recipient,
		Email: myMailData.recipient,
	}

	myBodyValue := email.BodyValue{
		Value: myMailData.bodyValue,
	}

	myBodyPart := email.BodyPart{
		PartID: "body",
		Type:   "text/plain",
	}

	myMail := email.Email{
		CustomHeaders: myMailData.customHeaders,
		From: []*mail.Address{
			&from,
		},
		To: []*mail.Address{
			&to,
		},
		InReplyTo:     myMailData.inReplyTo,
		References:    myMailData.references,
		Subject:       myMailData.subject,
		Keywords:      map[string]bool{"$draft": true},
		MailboxIDs:    map[jmap.ID]bool{draftMailboxID: true},
		BodyValues:    map[string]*email.BodyValue{"body": &myBodyValue},
		TextBody:      []*email.BodyPart{&myBodyPart},
		HasAttachment: true,
		Attachments:   myMailData.attachments,
	}

	return &myMail, nil
}

func WithSubject(subject string) Option {
	return func(e *emailData) error {
		e.subject = subject
		return nil
	}
}

func WithInReplyTo(inReplyTo []string) Option {
	return func(e *emailData) error {
		e.inReplyTo = append(e.inReplyTo, inReplyTo...)
		return nil
	}
}

func WithReference(reference []string) Option {
	return func(e *emailData) error {
		e.references = append(e.references, reference...)
		return nil
	}
}

func WithBodyValue(body string) Option {
	return func(e *emailData) error {
		e.bodyValue = body
		return nil
	}
}

func WithRecipient(recipient string) Option {
	return func(e *emailData) error {
		e.recipient = recipient
		return nil
	}
}

func WithCustomHeader(key string, value string) Option {
	return func(e *emailData) error {
		customHeader := email.Header{
			Name:  key,
			Value: value,
		}

		e.customHeaders = append(e.customHeaders, &customHeader)
		return nil
	}
}

func WithAttachment(attachmentName string, blob io.Reader) Option {
	return func(e *emailData) error {
		myAttachment, err := e.client.uploadAttachment(attachmentName, blob)
		if err != nil {
			return err
		}

		e.attachments = append(e.attachments, myAttachment)
		return nil
	}
}

func (j *JMAPClient) uploadAttachment(attachmentName string, blob io.Reader) (*email.BodyPart, error) {
	resp, err := j.Upload(j.userId, blob)
	if err != nil {
		return nil, err
	}

	myAttachment := email.BodyPart{
		BlobID:      resp.ID,
		Size:        resp.Size,
		Type:        resp.Type,
		Name:        attachmentName,
		Disposition: "attachment",
	}

	return &myAttachment, nil
}
