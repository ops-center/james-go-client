package inbox

import (
	"errors"
	"fmt"
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

// A general Email/get method call must be made for new users
// in order to generate the default mailboxes
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

// TODO: Remove this
func (j *JMAPClient) GetTestRecipient() (string, error) {
	j.muJmap.RLock()
	defer j.muJmap.RUnlock()
	return "recipient.acc@cloud.appscode.com", nil
}

type emailData struct {
	subject                                   string
	recipients, cc, bcc                       []*mail.Address
	customHeaders                             []*email.Header
	inReplyTo, references, textBody, htmlBody []string
	attachments                               []*email.BodyPart
	client                                    *JMAPClient
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

	if len(myMailData.recipients) == 0 && len(myMailData.bcc) == 0 && len(myMailData.cc) == 0 {
		return nil, errors.New("no recipients defined")
	}

	bodyValues := make(map[string]*email.BodyValue)
	var textBodyParts, htmlBodyParts []*email.BodyPart

	for id, textBodyData := range myMailData.textBody {
		partID := fmt.Sprintf("textbody%d", id)

		bodyValues[partID] = &email.BodyValue{
			Value: textBodyData,
		}
		textBodyParts = append(textBodyParts, &email.BodyPart{
			PartID: partID,
			Type:   "text/plain",
		})
	}

	for id, htmlBodyData := range myMailData.htmlBody {
		partID := fmt.Sprintf("htmlbody%d", id)

		bodyValues[partID] = &email.BodyValue{
			Value: htmlBodyData,
		}
		htmlBodyParts = append(htmlBodyParts, &email.BodyPart{
			PartID: partID,
			Type:   "text/html",
		})
	}

	myMail := email.Email{
		Subject: myMailData.subject,
		From: []*mail.Address{
			{
				Name:  j.userEmail,
				Email: j.userEmail,
			},
		},
		To:         myMailData.recipients,
		CC:         myMailData.cc,
		BCC:        myMailData.bcc,
		InReplyTo:  myMailData.inReplyTo,
		References: myMailData.references,

		Keywords:   map[string]bool{"$draft": true},
		MailboxIDs: map[jmap.ID]bool{draftMailboxID: true},
		BodyValues: bodyValues,
		TextBody:   textBodyParts,
		HTMLBody:   htmlBodyParts,

		HasAttachment: len(myMailData.attachments) > 0,
		Attachments:   myMailData.attachments,
		CustomHeaders: myMailData.customHeaders,
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

// WithReferences references must contain the messageID(s) of the email it's replying to
func WithReferences(references []string) Option {
	return func(e *emailData) error {
		e.references = append(e.references, references...)
		return nil
	}
}

func WithTextBody(textBody string) Option {
	return func(e *emailData) error {
		e.textBody = append(e.textBody, textBody)
		return nil
	}
}

func WithHTMLBody(htmlBody string) Option {
	return func(e *emailData) error {
		e.htmlBody = append(e.htmlBody, htmlBody)
		return nil
	}
}

func WithCC(ccList []string) Option {
	return func(e *emailData) error {
		for _, cc := range ccList {
			e.cc = append(e.cc, &mail.Address{
				Name:  cc,
				Email: cc,
			})
		}
		return nil
	}
}

func WithBCC(bccList []string) Option {
	return func(e *emailData) error {
		for _, bcc := range bccList {
			e.bcc = append(e.bcc, &mail.Address{
				Name:  bcc,
				Email: bcc,
			})
		}
		return nil
	}
}

func WithRecipients(recipients []string) Option {
	return func(e *emailData) error {
		for _, recipient := range recipients {
			e.recipients = append(e.recipients, &mail.Address{
				Name:  recipient,
				Email: recipient,
			})
		}
		return nil
	}
}

func WithCustomHeader(key string, value string) Option {
	return func(e *emailData) error {
		e.customHeaders = append(e.customHeaders, &email.Header{
			Name:  key,
			Value: value,
		})
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

	return &email.BodyPart{
		BlobID:      resp.ID,
		Size:        resp.Size,
		Type:        resp.Type,
		Name:        attachmentName,
		Disposition: "attachment",
	}, nil
}
