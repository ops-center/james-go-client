package james

import (
	"fmt"
	"git.sr.ht/~rockorager/go-jmap"
	"git.sr.ht/~rockorager/go-jmap/mail/email"
	"git.sr.ht/~rockorager/go-jmap/mail/emailsubmission"
	"git.sr.ht/~rockorager/go-jmap/mail/mailbox"
	"strings"
)

func (j *JmapClient) SendEmail(myMail *email.Email) error {
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
		Account: j.getUserId(),
		Create:  myMap,
	})

	myEmailSubmission := emailsubmission.EmailSubmission{
		EmailID: "#" + draftEmailId,
	}

	req.Invoke(&emailsubmission.Set{
		Account: j.getUserId(),

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

func (j *JmapClient) GetEmails(filter JmapEmailFilter) ([]*email.Email, error) {
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
		Account:  j.getUserId(),
		Filter:   &filter.FilterCondition,
		Sort:     []*email.SortComparator{&mySortComparator},
		Limit:    filter.MaxEmails,
		Position: filter.Position,
	})

	req.Invoke(&email.Get{
		Account: j.getUserId(),
		ReferenceIDs: &jmap.ResultReference{
			ResultOf: callID,        // The CallID of the referenced method
			Name:     "Email/query", // The name of the referenced method
			Path:     "/ids",        // JSON pointer to the location of the reference
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
func (j *JmapClient) DestroyAllEmails() (destroyed []jmap.ID, notDestroyed map[jmap.ID]*jmap.SetError, err error) {
	j.muJmap.Lock()
	defer j.muJmap.Unlock()

	req := &jmap.Request{
		Using: []jmap.URI{"urn:ietf:params:jmap:core", "urn:ietf:params:jmap:mail"},
	}

	callID0 := req.Invoke(&email.Query{
		Account: j.getUserId(),
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
		Account: j.getUserId(),
		ReferenceDestroyIDs: &jmap.ResultReference{
			ResultOf: callID1,     // The CallID of the referenced method
			Name:     "Email/get", // The name of the referenced method
			Path:     "list/*/id", // JSON pointer to the location of the reference
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
func (j *JmapClient) initializeMailboxIds() error {
	req := &jmap.Request{
		Using: []jmap.URI{"urn:ietf:params:jmap:core", "urn:ietf:params:jmap:mail"},
	}

	req.Invoke(&mailbox.Get{
		Account: j.getUserId(),
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
