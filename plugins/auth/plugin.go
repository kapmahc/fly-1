package auth

import (
	"github.com/facebookgo/inject"
	"github.com/gin-gonic/gin"
	"github.com/ikeikeikeike/go-sitemap-generator/stm"
	"github.com/jinzhu/gorm"
	"github.com/kapmahc/fly/web"
	"github.com/kapmahc/fly/web/i18n"
	"github.com/kapmahc/fly/web/job"
	"github.com/kapmahc/fly/web/security"
	"github.com/kapmahc/fly/web/settings"
	"github.com/kapmahc/fly/web/uploader"
	"github.com/kapmahc/fly/web/widgets"
	"golang.org/x/tools/blog/atom"
)

// Plugin plugin
type Plugin struct {
	Db       *gorm.DB           `inject:""`
	Jwt      *Jwt               `inject:""`
	Dao      *Dao               `inject:""`
	I18n     *i18n.I18n         `inject:""`
	Settings *settings.Settings `inject:""`
	Server   *job.Server        `inject:""`
	Wrap     *web.Wrap          `inject:""`
	Hmac     *security.Hmac     `inject:""`
	Uploader uploader.Store     `inject:""`
}

// Init load config
func (p *Plugin) Init() {}

// Dashboard dashboard nav
func (p *Plugin) Dashboard(c *gin.Context) []*widgets.Dropdown {
	var items []*widgets.Dropdown
	if _, ok := c.Get(CurrentUser); ok {
		items = append(items, widgets.NewDropdown(
			"auth.dashboard.title",
			widgets.NewLink("auth.users.info.title", "/users/info"),
			widgets.NewLink("auth.users.change-password.title", "/users/change-password"),
			nil,
			widgets.NewLink("auth.users.logs.title", "/users/logs"),
		))
	}
	return items
}

// Open open beans
func (p *Plugin) Open(*inject.Graph) error {
	return nil
}

// Atom rss.atom
func (p *Plugin) Atom(lang string) ([]*atom.Entry, error) {
	return []*atom.Entry{}, nil
}

// Sitemap sitemap.xml.gz
func (p *Plugin) Sitemap() ([]stm.URL, error) {
	return []stm.URL{}, nil
}

func init() {
	web.Register(&Plugin{})
}
