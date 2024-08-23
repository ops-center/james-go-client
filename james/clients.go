package james

import (
	"context"
	"fmt"
	"git.sr.ht/~rockorager/go-jmap"
	"git.sr.ht/~rockorager/go-jmap/mail"
	"git.sr.ht/~rockorager/go-jmap/mail/email"
	openapi "github.com/ops-center/james-go-client"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/util/wait"
	"strings"
	"sync"
	"time"
)

type Service struct {
	WebAdminClient
	tokenService
}

func New(c *Config) (*Service, error) {
	if err := c.checkValidity(); err != nil {
		return nil, err
	}

	rsaPrivateKey, err := c.getRsaPrivateKey()
	if err != nil {
		return nil, err
	}

	return &Service{
		WebAdminClient: c.getJamesWebAdminApiClient(),
		tokenService:   newTokenService(rsaPrivateKey),
	}, nil
}

type ClientsHubConf struct {
	WebAdminConf
	JMAPConf
}

func (ch *ClientsHubConf) checkValidity() (err error) {
	err = ch.WebAdminConf.checkValidity()
	err = errors.Wrapf(err, ch.JMAPConf.checkValidity().Error())

	return err
}

type ClientsHubService struct {
	WebAdminClient
	JMAPClient
}

type WebAdminClient struct {
	openapi.APIClient
	mu              sync.RWMutex
	authRenewChan   chan chan error
	started         bool
	done            chan struct{}
	sessionEndpoint string
}

func NewWebAdminClient(wc *WebAdminConf) (*WebAdminClient, error) {
	client := WebAdminClient{
		sessionEndpoint: wc.WebAdminSessionEndpoint,
		authRenewChan:   make(chan chan error),
		started:         false,
	}
	if err := client.renew(); err != nil {
		return nil, err
	}

	client.Start()

	return &client, nil
}

func (w *WebAdminClient) Start() {
	w.mu.Lock()
	defer w.mu.Unlock()

	if w.started {
		return
	}
	w.started = true
	w.done = make(chan struct{})

	go func() {
		for {
			select {
			case <-w.done:
				return
			case errChan := <-w.authRenewChan:
				w.mu.Lock()
				err := w.renew()
				w.mu.Unlock()
				time.Sleep(20 * time.Millisecond)
				errChan <- err
				chanlen := len(w.authRenewChan)
				for i := 0; i < chanlen; i++ {
					c := <-w.authRenewChan
					c <- err
				}
			}
		}
	}()
}

func (w *WebAdminClient) Close() {
	w.mu.Lock()
	defer w.mu.Unlock()

	if w.started {
		return
	}
	w.started = false

	close(w.done)
}

func (w *WebAdminClient) RenewSession() error {
	w.mu.RLock()
	if !w.started {
		return fmt.Errorf("webadmin client auth channel closed")
	}
	w.mu.RUnlock()

	errChan := make(chan error)
	w.authRenewChan <- errChan
	return <-errChan
}

func (w *WebAdminClient) renew() error {
	var lastErr error
	err := wait.PollUntilContextTimeout(
		context.TODO(),
		100*time.Millisecond,
		5*time.Second,
		true,
		func(ctx context.Context) (bool, error) {
			token, err := getJamesToken()
			if err != nil {
				lastErr = err
				return false, nil
			}

			configuration := openapi.NewConfiguration().WithAccessToken(token)
			configuration.Servers = []openapi.ServerConfiguration{
				{
					URL:         w.sessionEndpoint,
					Description: "Appscode James WebAdmin endpoint",
				},
			}
			w.APIClient = *openapi.NewAPIClient(configuration)

			return true, nil
		},
	)

	if err != nil {
		return errors.Wrap(lastErr, err.Error())
	}
	return nil
}

type JMAPClient struct {
	jmap.Client
	muJmap         sync.RWMutex
	userId         jmap.ID
	userEmail      string
	mailboxIds     map[string]jmap.ID
	forceBasicAuth bool
	basicAuthCreds BasicAuthCredentials
	authRenewChan  chan chan error
	done           chan struct{}
	started        bool
}

const (
	Recipient = "recipient.acc@cloud.appscode.com"
)

func NewJMAPClient(jc *JMAPConf) (*JMAPClient, error) {
	client := &JMAPClient{
		Client: jmap.Client{
			SessionEndpoint: jc.JMAPSessionEndpoint,
		},
		authRenewChan:  make(chan chan error),
		started:        false,
		forceBasicAuth: false,
	}

	if jc.ForceBasicAuth {
		client.basicAuthCreds = jc.BasicAuthCreds
		client.forceBasicAuth = true
	}

	if err := client.renew(); err != nil {
		return nil, err
	}

	client.Start()

	return client, nil
}

func (j *JMAPClient) Start() {
	j.muJmap.Lock()
	defer j.muJmap.Unlock()

	if j.started {
		return
	}
	j.started = true
	j.done = make(chan struct{})

	go func() {
		for {
			select {
			case <-j.done:
				return
			case errChan := <-j.authRenewChan:
				j.muJmap.Lock()
				err := j.renew()
				j.muJmap.Unlock()
				time.Sleep(20 * time.Millisecond)
				errChan <- err
				chanlen := len(j.authRenewChan)
				for i := 0; i < chanlen; i++ {
					c := <-j.authRenewChan
					c <- err
				}
			}
		}
	}()
}

func (j *JMAPClient) Close() {
	j.muJmap.Lock()
	defer j.muJmap.Unlock()

	if !j.started {
		return
	}
	j.started = false

	close(j.done)
}

func (j *JMAPClient) RenewSession() error {
	j.muJmap.RLock()
	if !j.started {
		return fmt.Errorf("jmap client auth channel closed")
	}
	j.muJmap.RUnlock()

	errChan := make(chan error, 1)
	j.authRenewChan <- errChan
	return <-errChan
}

func (j *JMAPClient) renew() error {
	var lastErr error
	err := wait.PollUntilContextTimeout(
		context.TODO(),
		100*time.Millisecond,
		5*time.Second,
		true,
		func(ctx context.Context) (bool, error) {

			if j.forceBasicAuth {
				j.WithBasicAuth(j.basicAuthCreds.Username, j.basicAuthCreds.Password)
			} else {
				if token, err := getJamesToken(); err == nil {
					j.WithAccessToken(token)
				} else {
					lastErr = err
					return false, nil
				}
			}

			if err := j.Client.Authenticate(); err != nil {
				lastErr = err
				return false, nil
			}

			if userId, ok := j.Session.PrimaryAccounts[mail.URI]; ok {
				j.userId = userId
			} else {
				lastErr = fmt.Errorf("user id not found in session")
				return false, nil
			}

			if account, ok := j.Session.Accounts[j.userId]; ok {
				j.userEmail = account.Name
			} else {
				lastErr = fmt.Errorf("user email not found in session")
				return false, nil
			}

			if err := j.initializeMailboxIds(); err != nil {
				lastErr = err
				return false, nil
			}

			return true, nil
		})

	if err != nil {
		return errors.Wrap(lastErr, err.Error())
	}
	return nil
}

type JamesMailboxName string

const (
	DraftMailbox   JamesMailboxName = "Drafts"
	SentMailbox    JamesMailboxName = "Sent"
	InboxMailbox   JamesMailboxName = "INBOX"
	ArchiveMailbox JamesMailboxName = "Archive"
	OutboxMailbox  JamesMailboxName = "Outbox"
	TrashMailbox   JamesMailboxName = "Trash"
	SpamMailbox    JamesMailboxName = "Spam"
)

func (j *JMAPClient) getMailboxId(mailboxName JamesMailboxName) (jmap.ID, error) {
	trimmedMailboxName := strings.ToLower(strings.TrimSpace(string(mailboxName)))
	if mailboxid, ok := j.mailboxIds[trimmedMailboxName]; ok {
		return mailboxid, nil
	}
	return "", fmt.Errorf("no mailbox found with name: %s", trimmedMailboxName)
}

type JmapSortProperty string

const (
	//Sort Properties supported by james distributed 3.9
	ReceivedAt JmapSortProperty = "receivedAt"
	SentAt     JmapSortProperty = "sentAt"
	Size       JmapSortProperty = "size"
	From       JmapSortProperty = "from"
	To         JmapSortProperty = "to"
	Subject    JmapSortProperty = "subject"
)

type JmapEmailFilter struct {
	email.FilterCondition
	HasProperties       []string
	MaxEmails           uint64
	SortProperty        JmapSortProperty
	AscendingOrder      bool
	Position            int64
	FetchAllBodyValues  bool
	FetchHTMLBodyValues bool
	InMailboxWithName   JamesMailboxName
}
