package inbox

import (
	"context"
	go_errs "errors"
	"fmt"
	"strings"
	"sync"
	"time"

	"git.sr.ht/~rockorager/go-jmap"
	"git.sr.ht/~rockorager/go-jmap/mail"
	"git.sr.ht/~rockorager/go-jmap/mail/email"
	openapi "go.opscenter.dev/james-go-client"
	"golang.org/x/oauth2"
	"k8s.io/apimachinery/pkg/util/wait"
)

type Service struct {
	*WebAdminClient
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

	client, err := c.getJamesWebAdminApiClient()
	if err != nil {
		return nil, err
	}

	return &Service{
		WebAdminClient: client,
		tokenService:   newTokenService(rsaPrivateKey, c.EmailDomain),
	}, nil
}

type ClientsHubConf struct {
	WebAdminConf
	JMAPConf
}

//nolint:unused
func (ch *ClientsHubConf) checkValidity() (err error) {
	if err := ch.JMAPConf.checkValidity(); err != nil {
		return err
	}
	return ch.WebAdminConf.checkValidity()
}

type ClientsHubService struct {
	WebAdminClient
	JMAPClient
}

const authRenewalCacheTTL = int64(1 * time.Second)

type WebAdminClient struct {
	openapi.APIClient
	tokenGetter     TokenGetterFunc
	mu              sync.RWMutex
	serviceEndpoint string
	emailDomain     string

	lastComputedCacheAtUnixTime int64
	cachedRenewalErr            error
}

func NewWebAdminClient(wc *WebAdminConf) (*WebAdminClient, error) {
	client := WebAdminClient{
		serviceEndpoint: wc.WebAdminServiceEndpoint,
		tokenGetter:     wc.TokenGetter,
	}
	if err := client.renew(); err != nil {
		return nil, err
	}

	client.emailDomain = wc.EmailDomain
	client.lastComputedCacheAtUnixTime = time.Now().UnixNano()
	client.cachedRenewalErr = nil

	return &client, nil
}

func (w *WebAdminClient) RenewSession() error {
	w.mu.RLock()

	if !w.refreshRequired() {
		err := w.cachedRenewalErr
		w.mu.RUnlock()
		return err
	}
	w.mu.RUnlock()

	w.mu.Lock()
	defer w.mu.Unlock()

	if w.cachedRenewalErr != nil || w.refreshRequired() {
		w.cachedRenewalErr = w.renew()
		w.lastComputedCacheAtUnixTime = time.Now().UnixNano()
	}

	err := w.cachedRenewalErr
	return err
}

func (w *WebAdminClient) refreshRequired() bool {
	return time.Now().UnixNano() > authRenewalCacheTTL+w.lastComputedCacheAtUnixTime
}

func (w *WebAdminClient) renew() error {
	var lastErr error
	err := wait.PollUntilContextTimeout(context.TODO(), 3*time.Second, 1*time.Minute, true,
		func(ctx context.Context) (bool, error) {
			hc, token, err := w.tokenGetter()
			if err != nil {
				lastErr = fmt.Errorf("failed to get inbox service token: %w", err)
				return false, nil
			}

			ctx2 := context.WithValue(context.Background(), oauth2.HTTPClient, hc)
			configuration := openapi.NewConfiguration().WithAccessToken(ctx2, token)
			configuration.Servers = []openapi.ServerConfiguration{
				{
					URL:         w.serviceEndpoint,
					Description: "Appscode Inbox Service WebAdmin endpoint",
				},
			}
			w.APIClient = *openapi.NewAPIClient(configuration)

			return true, nil
		},
	)
	if err != nil {
		return fmt.Errorf("failed to renew WebAdmin client: %w", go_errs.Join(lastErr, err))
	}
	return nil
}

type JMAPClient struct {
	jmap.Client
	tokenGetter                 TokenGetterFunc
	mu                          sync.RWMutex
	userId                      jmap.ID
	userEmail                   string
	mailboxIds                  map[string]jmap.ID
	lastComputedCacheAtUnixTime int64
	cachedRenewalErr            error
}

func NewJMAPClient(jc *JMAPConf) (*JMAPClient, error) {
	client := &JMAPClient{
		Client: jmap.Client{
			SessionEndpoint: jc.JMAPSessionEndpoint,
		},
		tokenGetter: jc.TokenGetter,
	}

	if err := client.renew(); err != nil {
		return nil, err
	}

	client.lastComputedCacheAtUnixTime = time.Now().UnixNano()
	client.cachedRenewalErr = nil

	return client, nil
}

func (j *JMAPClient) RenewSession() error {
	j.mu.RLock()

	if !j.refreshRequired() {
		err := j.cachedRenewalErr
		j.mu.RUnlock()
		return err
	}
	j.mu.RUnlock()

	j.mu.Lock()
	defer j.mu.Unlock()

	if j.cachedRenewalErr != nil || j.refreshRequired() {
		j.cachedRenewalErr = j.renew()
		j.lastComputedCacheAtUnixTime = time.Now().UnixNano()
	}

	err := j.cachedRenewalErr
	return err
}

func (j *JMAPClient) refreshRequired() bool {
	return time.Now().UnixNano() > authRenewalCacheTTL+j.lastComputedCacheAtUnixTime
}

func (j *JMAPClient) renew() error {
	var lastErr error
	err := wait.PollUntilContextTimeout(context.TODO(), 3*time.Second, 1*time.Minute, true,
		func(ctx context.Context) (bool, error) {
			hc, token, err := j.tokenGetter()
			if err == nil {
				ctx2 := context.WithValue(context.Background(), oauth2.HTTPClient, hc)
				j.WithAccessToken(ctx2, token)
			} else {
				lastErr = err
				return false, nil
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
		return fmt.Errorf("failed to renew JMAP client: %w", go_errs.Join(lastErr, err))
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
	// Sort Properties supported by james distributed 3.9
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
