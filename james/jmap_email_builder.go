package james

import (
	"errors"
	"git.sr.ht/~rockorager/go-jmap"
	_ "git.sr.ht/~rockorager/go-jmap/core"
	"git.sr.ht/~rockorager/go-jmap/mail"
	"git.sr.ht/~rockorager/go-jmap/mail/email"
	"io"
)

type emailData struct {
	recipient, subject, bodyValue string
	customHeaders                 []*email.Header
	inReplyTo, references         []string
	attachments                   []*email.BodyPart
	client                        *JmapClient
}

type Option func(data *emailData) error

func (j *JmapClient) NewEmail(options ...Option) (*email.Email, error) {
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
		Name:  j.getUserEmail(),
		Email: j.getUserEmail(),
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

		InReplyTo: myMailData.inReplyTo,

		References: myMailData.references,

		Subject: myMailData.subject,

		Keywords: map[string]bool{"$draft": true},

		MailboxIDs: map[jmap.ID]bool{draftMailboxID: true},

		BodyValues: map[string]*email.BodyValue{"body": &myBodyValue},

		TextBody: []*email.BodyPart{&myBodyPart},

		HasAttachment: true,

		Attachments: myMailData.attachments,
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

func (j *JmapClient) uploadAttachment(attachmentName string, blob io.Reader) (*email.BodyPart, error) {
	resp, err := j.Upload(j.userId, blob)
	if err != nil {
		return nil, err
	}

	myAttachment := email.BodyPart{
		BlobID: resp.ID,

		Size: resp.Size,

		Type: resp.Type,

		Name: attachmentName,

		Disposition: "attachment",
	}

	return &myAttachment, nil
}
