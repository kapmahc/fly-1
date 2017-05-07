package forms

import (
	"fmt"

	"github.com/facebookgo/inject"
	"github.com/gin-gonic/gin"
	"github.com/ikeikeikeike/go-sitemap-generator/stm"
	"github.com/jinzhu/gorm"
	"github.com/kapmahc/fly/plugins/auth"
	"github.com/kapmahc/fly/web"
	"github.com/kapmahc/fly/web/i18n"
	"github.com/kapmahc/fly/web/job"
	"github.com/kapmahc/fly/web/widgets"
	"github.com/urfave/cli"
	"golang.org/x/tools/blog/atom"
)

// Plugin plugin
type Plugin struct {
	Db   *gorm.DB   `inject:""`
	Jwt  *auth.Jwt  `inject:""`
	Wrap *web.Wrap  `inject:""`
	I18n *i18n.I18n `inject:""`
}

// Init load config
func (p *Plugin) Init() {}

// Dashboard dashboard nav
func (p *Plugin) Dashboard(c *gin.Context) []*widgets.Dropdown {
	var items []*widgets.Dropdown
	if admin, ok := c.Get(auth.IsAdmin); ok && admin.(bool) {
		items = append(
			items,
			widgets.NewDropdown(
				"forms.dashboard.title",
				widgets.NewLink("forms.index.title", "/forms/manage"),
			),
		)
	}
	return items

}

// Open open beans
func (p *Plugin) Open(*inject.Graph) error {
	return nil
}

// Console console commands
func (p *Plugin) Console() []cli.Command {
	return []cli.Command{}
}

// Atom rss.atom
func (p *Plugin) Atom(lang string) ([]*atom.Entry, error) {
	return []*atom.Entry{}, nil
}

// Sitemap sitemap.xml.gz
func (p *Plugin) Sitemap() ([]stm.URL, error) {
	items := []stm.URL{
		stm.URL{"loc": "/forms"},
	}

	var forms []Form
	if err := p.Db.Select([]string{"id", "updated_at"}).Find(&forms).Error; err != nil {
		return nil, err
	}
	for _, it := range forms {
		items = append(
			items,
			stm.URL{"loc": fmt.Sprintf("/forms/apply/%d", it.ID), "lastmod": it.UpdatedAt},
			stm.URL{"loc": fmt.Sprintf("/forms/cancel/%d", it.ID), "lastmod": it.UpdatedAt},
		)
	}

	return []stm.URL{}, nil
}

// Workers register workers
func (p *Plugin) Workers() map[string]job.Handler {
	return map[string]job.Handler{}
}

func init() {
	web.Register(&Plugin{})
}
