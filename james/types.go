package james

import (
	"crypto/rsa"
	"fmt"
	"git.sr.ht/~rockorager/go-jmap"
	"git.sr.ht/~rockorager/go-jmap/mail"
	"git.sr.ht/~rockorager/go-jmap/mail/email"
	"strings"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
	openapi "github.com/ops-center/james-go-client"
	"github.com/pkg/errors"
)

type WebAdminConf struct {
	WebAdminServiceAddr string
	WebAdminServicePort string
	WebAdminAuthnToken  string
}

func (wc *WebAdminConf) checkValidity() (err error) {
	if isEmptyString(wc.WebAdminServiceAddr) {
		err = errors.Wrapf(err, "web admin address is required")
	}
	if isEmptyString(wc.WebAdminServicePort) {
		err = errors.Wrapf(err, "web admin port is required")
	}

	_, _, perr := jwt.NewParser().ParseUnverified(wc.WebAdminAuthnToken, jwt.MapClaims{})
	if perr != nil {
		err = errors.Wrapf(err, perr.Error())
	}

	return err
}

func (wc *WebAdminConf) getJamesWebAdminApiClient() WebAdminClient {
	configuration := openapi.NewConfiguration().WithAccessToken(wc.WebAdminAuthnToken)
	configuration.Servers[0] = openapi.ServerConfiguration{
		URL: fmt.Sprintf("%v:%v", wc.WebAdminServiceAddr, wc.WebAdminServicePort),
	}
	apiClient := openapi.NewAPIClient(configuration)

	return WebAdminClient(*apiClient)
}

type JmapConf struct {
	JmapServerAddr      string
	JmapServerPort      string
	JmapSessionEndpoint string
	JmapAuthnToken      string
}

func (jc *JmapConf) checkValidity() (err error) {
	if isEmptyString(jc.JmapServerAddr) {
		err = errors.Wrapf(err, "jmap server address is required")
	}
	if isEmptyString(jc.JmapServerPort) {
		err = errors.Wrapf(err, "jmap port address is required")
	}
	if isEmptyString(jc.JmapSessionEndpoint) {
		err = errors.Wrapf(err, "jmap session endpoint is required")
	}

	return err
}

type Config struct {
	WebAdminConf
	PrivateKey string
}

func (c *Config) checkValidity() (err error) {
	err = c.WebAdminConf.checkValidity()
	if isEmptyString(c.PrivateKey) {
		err = errors.Wrapf(err, "private key is required")
	}

	return
}

func (c *Config) getRsaPrivateKey() (*rsa.PrivateKey, error) {
	rsaPk, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(c.PrivateKey))
	if err != nil {
		return nil, err
	}

	return rsaPk, nil
}

type WebAdminClient openapi.APIClient

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
	JmapConf
}

func (ch *ClientsHubConf) checkValidity() (err error) {
	err = ch.WebAdminConf.checkValidity()
	err = errors.Wrapf(err, ch.JmapConf.checkValidity().Error())

	return err
}

type JmapClient struct {
	jmap.Client
	muJmap     sync.RWMutex
	userId     jmap.ID
	userEmail  string
	mailboxIds map[string]jmap.ID
}

func NewJmapClient(jc *JmapConf) (*JmapClient, error) {
	jmapClient := &JmapClient{
		Client: jmap.Client{
			SessionEndpoint: jc.JmapSessionEndpoint,
		},
	}

	jmapClient.WithAccessToken(jc.JmapAuthnToken)
	if err := jmapClient.Client.Authenticate(); err != nil {
		return nil, err
	}

	if userId, ok := jmapClient.Session.PrimaryAccounts[mail.URI]; ok {
		jmapClient.userId = userId
	} else {
		return nil, fmt.Errorf("user id not found in session")
	}

	if account, ok := jmapClient.Session.Accounts[jmapClient.userId]; ok {
		jmapClient.userEmail = account.Name
	} else {
		return nil, fmt.Errorf("user account not found in session")
	}

	if err := jmapClient.initializeMailboxIds(); err != nil {
		return nil, err
	}

	return jmapClient, nil
}

func (j *JmapClient) GetUserId() jmap.ID {
	j.muJmap.RLock()
	defer j.muJmap.RUnlock()
	return j.getUserId()
}

func (j *JmapClient) getUserId() jmap.ID {
	return j.userId
}

func (j *JmapClient) GetUserEmail() string {
	j.muJmap.RLock()
	defer j.muJmap.RUnlock()
	return j.getUserEmail()
}

func (j *JmapClient) getUserEmail() string {
	return j.userEmail
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

func (j *JmapClient) GetMailboxId(mailboxName JamesMailboxName) (jmap.ID, error) {
	j.muJmap.RLock()
	defer j.muJmap.RUnlock()
	return j.getMailboxId(mailboxName)
}

func (j *JmapClient) getMailboxId(mailboxName JamesMailboxName) (jmap.ID, error) {
	trimmedMailboxName := strings.ToLower(strings.TrimSpace(string(mailboxName)))
	if mailboxid, ok := j.mailboxIds[trimmedMailboxName]; ok {
		return mailboxid, nil
	}
	return "", fmt.Errorf("no mailbox found with name: %s", trimmedMailboxName)
}

type JmapSortProperty string

const (
	//Properties supported by james distributed 3.9
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

type ClientsHubService struct {
	WebAdminClient
	JmapClient
}

const (
	DurationDay             = time.Hour * 24
	DefaultTokenExpDuration = DurationDay
	AdminTokenExpDuration   = DurationDay * 10 * 365

	TokenCookieName     = "_james_token"
	TokenCookieDuration = DurationDay

	GlobalMailDomain = "cloud.appscode.com"
)

type ObjectTypeIdentifier int

const (
	UserType ObjectTypeIdentifier = iota + 1
	OrgType
	TeamType
	GroupType
	ClusterType
	NamespaceType
	DbType
	AgentType
	WorkloadType
)

func (i ObjectTypeIdentifier) String() string {
	return []string{"usr", "org", "tm", "grp", "cluster", "ns", "db", "agent", "wl"}[i-1]
}

func (i ObjectTypeIdentifier) EnumIndex() int {
	return int(i)
}

// Object interface represents an entity within the Apache James server,
// encompassing both accounts and groups. It facilitates the generation of
// account addresses and management of object relationships.
type Object interface {
	// GetName returns the primary identifier or name of the object. This name
	// is used for various operations, such as generating email addresses,
	// identifying objects within the server, and displaying user-facing names.
	//
	// It is typically a concise and unique identifier, often automatically
	// generated based on the object's type and internal properties. However,
	// some implementations might allow setting a custom name.
	//
	// Return value:
	//   - string: The name of the object.
	//
	// Example:
	//   - For a user object, GetName() might return a username like "john.doe".
	//   - For a product object, GetName() might return a product code like "ABC123".
	GetName() string

	// GetUniqueID returns a globally unique identifier for the object. This identifier
	// is distinct from the name returned by GetName(), providing a guaranteed unique
	// reference to the object within the system.
	//
	// The unique ID is typically generated using a UUID or similar mechanism to ensure
	// its uniqueness across the entire system.
	//
	// Return value:
	//   - string: The globally unique identifier for the object.
	GetUniqueID() string

	// GetType returns the type of the object, indicating whether it's
	// an account or a group. This is important to decide what
	// actions and behaviors are allowed for the object.
	GetType() ObjectTypeIdentifier

	// HasParentObject signifies whether the object belongs to a parent
	// group. This information is essential for understanding object
	// hierarchy and managing group memberships.
	HasParentObject() bool

	// GetParentObject retrieves the parent Object if it exists, providing
	// a direct link to the group that the object belongs to. This enables
	// efficient navigation and manipulation of object relationships.
	GetParentObject() (Object, error)

	// IsGroup explicitly states whether the object is a group. This
	// distinction is crucial for applying group-specific operations and
	// enforcing appropriate access controls.
	IsGroup() bool

	// AdditionalTokenClaims returns additional JWT claims to be included when
	// generating a token for the object. These additional claims can provide custom
	// metadata or attributes associated with the object, supplementing the standard
	// claims like issuer, subject, expiration time, etc.
	//
	// This method is useful for extending the standard JWT token with domain-specific information,
	// allowing for more context-rich authentication and authorization mechanisms.
	AdditionalTokenClaims() *jwt.MapClaims

	// GetBoundedUserIdentity returns key-value pairs representing the system user
	// associated with the object, in cases where the object is not a user or organization.
	// This can include details such as userID, userName, etc.
	//
	// This method is useful for identifying the system user context associated with
	// non-user objects, facilitating better tracking and management of object-user
	// relationships within the system.
	//
	// Return value:
	//   - map[string]string: Key-value pairs representing the bounded user identity.
	GetBoundedUserIdentity() UserIdentity
}

type ObjectIdentifier struct {
	ObjectName          string
	ObjectUniqueID      string
	ObjectType          ObjectTypeIdentifier
	IsGroupType         bool
	ParentObject        *ObjectIdentifier
	AdditionalClaims    *jwt.MapClaims
	BoundedUserIdentity UserIdentity
}

// Implements the Object interface
var _ Object = (*ObjectIdentifier)(nil)

func (o ObjectIdentifier) GetName() string {
	return o.ObjectName
}

func (o ObjectIdentifier) GetUniqueID() string {
	return o.ObjectUniqueID
}

func (o ObjectIdentifier) GetType() ObjectTypeIdentifier {
	return o.ObjectType
}

func (o ObjectIdentifier) HasParentObject() bool {
	return o.ParentObject != nil
}

func (o ObjectIdentifier) GetParentObject() (Object, error) {
	return o.ParentObject, nil
}

func (o ObjectIdentifier) IsGroup() bool {
	return o.IsGroupType
}

func (o ObjectIdentifier) AdditionalTokenClaims() *jwt.MapClaims {
	return o.AdditionalClaims
}

func (o ObjectIdentifier) GetBoundedUserIdentity() UserIdentity {
	return o.BoundedUserIdentity
}

type UserIdentity struct {
	OwnerName string
	OwnerID   string
	OwnerType string
}

func (u UserIdentity) String() string {
	return fmt.Sprintf("nm-%s&id-%s&typ-%s", u.OwnerName, u.OwnerID, u.OwnerType)
}
