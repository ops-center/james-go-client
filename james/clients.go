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

const authRenewalCacheTTL = 5 * time.Second

type WebAdminClient struct {
	openapi.APIClient
	mu              sync.RWMutex
	sessionEndpoint string

	lastComputedCacheAt time.Time
	cachedRenewalErr    error
}

func NewWebAdminClient(wc *WebAdminConf) (*WebAdminClient, error) {
	client := WebAdminClient{
		sessionEndpoint: wc.WebAdminSessionEndpoint,
	}
	if err := client.renew(); err != nil {
		return nil, err
	}

	client.lastComputedCacheAt = time.Now()
	client.cachedRenewalErr = nil

	return &client, nil
}

func (w *WebAdminClient) RenewSession() error {
	w.mu.Lock()
	defer w.mu.Unlock()

	if time.Now().After(w.lastComputedCacheAt.Add(authRenewalCacheTTL)) {
		w.cachedRenewalErr = w.renew()
		w.lastComputedCacheAt = time.Now()
	}

	return w.cachedRenewalErr
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
	muJmap     sync.RWMutex
	userId     jmap.ID
	userEmail  string
	mailboxIds map[string]jmap.ID

	forceBasicAuth      bool
	basicAuthCreds      BasicAuthCredentials
	lastComputedCacheAt time.Time
	cachedRenewalErr    error
}

func NewJMAPClient(jc *JMAPConf) (*JMAPClient, error) {
	client := &JMAPClient{
		Client: jmap.Client{
			SessionEndpoint: jc.JMAPSessionEndpoint,
		},
		forceBasicAuth: false,
	}

	if jc.ForceBasicAuth {
		client.basicAuthCreds = jc.BasicAuthCreds
		client.forceBasicAuth = true
	}

	if err := client.renew(); err != nil {
		return nil, err
	}

	client.lastComputedCacheAt = time.Now()
	client.cachedRenewalErr = nil

	return client, nil
}

func (j *JMAPClient) RenewSession() error {
	j.muJmap.Lock()
	defer j.muJmap.Unlock()

	if time.Now().After(j.lastComputedCacheAt.Add(authRenewalCacheTTL)) {
		j.cachedRenewalErr = j.renew()
		j.lastComputedCacheAt = time.Now()
	}

	return j.cachedRenewalErr
}

func (j *JMAPClient) renew() error {
	var lastErr error
	err := wait.PollUntilContextTimeout(
		context.TODO(),
		100*time.Millisecond,
		3*time.Second,
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

type MailboxName string

const (
	DraftMailbox   MailboxName = "Drafts"
	SentMailbox    MailboxName = "Sent"
	InboxMailbox   MailboxName = "INBOX"
	ArchiveMailbox MailboxName = "Archive"
	OutboxMailbox  MailboxName = "Outbox"
	TrashMailbox   MailboxName = "Trash"
	SpamMailbox    MailboxName = "Spam"
)

func (j *JMAPClient) getMailboxId(mailboxName MailboxName) (jmap.ID, error) {
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
	InMailboxWithName   MailboxName
}
