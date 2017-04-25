package root

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/google/uuid"
)

// Timestamp timestamp
type Timestamp struct {
	ID      uint      `orm:"column(id)"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
}

// Model model
type Model struct {
	Timestamp
	Updated time.Time `orm:"auto_now;type(datetime)"`
}

const (
	// MediaTypeHTML html
	MediaTypeHTML = "html"
	// MediaTypeMarkdown markdown
	MediaTypeMarkdown = "markdown"
)

// Media media data
type Media struct {
	Body string
	Type string
}

// Resource resource
type Resource struct {
	ResourceType string
	ResourceID   uint `orm:"column(resource_id)"`
}

// --------------------------------------------------------

const (
	// RoleAdmin admin role
	RoleAdmin = "admin"
	// RoleRoot root role
	RoleRoot = "root"
	// UserTypeEmail email user
	UserTypeEmail = "email"
	// UserTypeFacebook facebook
	UserTypeFacebook = "facebook"
	// UserTypeGoogle google
	UserTypeGoogle = "google"
)

// User user
type User struct {
	Model

	Name            string
	Email           string
	UID             string `orm:"column(uid)"`
	Password        string
	ProviderID      string `orm:"column(provider_id)"`
	ProviderType    string
	Home            string
	Logo            string
	SignInCount     uint
	LastSignInAt    *time.Time
	LastSignInIP    string `orm:"column(last_sign_in_ip)"`
	CurrentSignInAt *time.Time
	CurrentSignInIP string `orm:"column(current_sign_in_ip)"`
	ConfirmedAt     *time.Time
	LockedAt        *time.Time

	Logs []*Log `orm:"reverse(many)"`
}

// TableName table name
func (*User) TableName() string {
	return "users"
}

// IsConfirm is confirm?
func (p *User) IsConfirm() bool {
	return p.ConfirmedAt != nil
}

// IsLock is lock?
func (p *User) IsLock() bool {
	return p.LockedAt != nil
}

//SetGravatarLogo set logo by gravatar
func (p *User) SetGravatarLogo() {
	buf := md5.Sum([]byte(strings.ToLower(p.Email)))
	p.Logo = fmt.Sprintf("https://gravatar.com/avatar/%s.png", hex.EncodeToString(buf[:]))
}

//SetUID generate uid
func (p *User) SetUID() {
	p.UID = uuid.New().String()
}

func (p User) String() string {
	return fmt.Sprintf("%s<%s>", p.Name, p.Email)
}

// Log log
type Log struct {
	Timestamp

	Message string
	Type    string
	IP      string `orm:"column(ip)"`

	User *User `orm:"rel(fk)"`
}

// TableName table name
func (*Log) TableName() string {
	return "logs"
}

// Policy policy
type Policy struct {
	Model

	StartUp  time.Time
	ShutDown time.Time

	User *User `orm:"rel(fk)"`
	Role *Role `orm:"rel(fk)"`
}

//Enable is enable?
func (p *Policy) Enable() bool {
	now := time.Now()
	return now.After(p.StartUp) && now.Before(p.ShutDown)
}

// TableName table name
func (*Policy) TableName() string {
	return "policies"
}

// Role role
type Role struct {
	Model
	Resource

	Name string
}

// TableName table name
func (*Role) TableName() string {
	return "roles"
}

func (p *Role) String() string {
	return fmt.Sprintf("%s@%s://%d", p.Name, p.ResourceType, p.ResourceID)
}

// --------------------------------------------------------

// Attachment attachment
type Attachment struct {
	Model
	Resource

	Title     string
	URL       string `orm:"column(url)"`
	Length    int64
	MediaType string

	User *User `orm:"rel(fk)"`
}

// TableName table name
func (*Attachment) TableName() string {
	return "attachments"
}

// Vote vote
type Vote struct {
	Model
	Resource

	Point int
}

// TableName table name
func (*Vote) TableName() string {
	return "votes"
}

// --------------------------------------------------------

// Locale locale
type Locale struct {
	Model

	Code    string
	Message string
	Lang    string
}

// TableName table name
func (*Locale) TableName() string {
	return "locales"
}

// Setting setting
type Setting struct {
	Model

	Key    string
	Val    string
	Encode bool
}

// TableName table name
func (*Setting) TableName() string {
	return "settings"
}

// --------------------------------------------------------

// Host host
type Host struct {
	Model

	Name string
	Lang string

	Title       string
	SubTitle    string
	Keywords    string
	Description string
	Copyright   string

	Ssl         bool
	PublicPerm  string
	PrivatePerm string

	Author *User `orm:"rel(fk)"`
}

// TableName table name
func (*Host) TableName() string {
	return "hosts"
}

// Post post
type Post struct {
	Model

	Name  string
	Title string
	Body  string

	Host   *Host `orm:"rel(fk)"`
	Author *User `orm:"rel(fk)"`
}

// TableName table name
func (*Post) TableName() string {
	return "posts"
}

//FriendLink friend-links
type FriendLink struct {
	Model

	Name string
	Home string
	Logo string

	Host *Host `orm:"rel(fk)"`
}

// TableName table name
func (*FriendLink) TableName() string {
	return "friend_links"
}

//LeaveWord leave-word
type LeaveWord struct {
	Timestamp
	Media

	Host *Host `orm:"rel(fk)"`
}

// TableName table name
func (*LeaveWord) TableName() string {
	return "leave_words"
}

// --------------------------------------------------------

//Link link
type Link struct {
	Model

	Sort  int
	Label string
	Href  string

	Host *Host `orm:"rel(fk)"`
}

// TableName table name
func (*Link) TableName() string {
	return "links"
}

// Card card
type Card struct {
	Model

	Sort    int
	Label   string
	Href    string
	Summary string
	Logo    string

	Host *Host `orm:"rel(fk)"`
}

// TableName table name
func (*Card) TableName() string {
	return "cards"
}

// --------------------------------------------------------

func init() {
	orm.RegisterModel(
		new(Host), new(FriendLink), new(Post), new(LeaveWord),
		new(Link), new(Card),
		new(Locale), new(Setting),
		new(Vote), new(Attachment),
		new(User), new(Log), new(Role), new(Policy),
	)
}
