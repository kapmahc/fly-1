package pos

import (
	"github.com/facebookgo/inject"
	"github.com/gin-gonic/gin"
	"github.com/ikeikeikeike/go-sitemap-generator/stm"
	"github.com/jinzhu/gorm"
	"github.com/kapmahc/fly/web"
	"github.com/kapmahc/fly/web/job"
	"github.com/kapmahc/fly/web/widgets"
	"github.com/urfave/cli"
	"golang.org/x/tools/blog/atom"
)

// Plugin plugin
type Plugin struct {
	Db *gorm.DB `inject:""`
}

// Init load config
func (p *Plugin) Init() {}

// Mount mount web points
func (p *Plugin) Mount(*gin.Engine) {

}

// Dashboard dashboard nav
func (p *Plugin) Dashboard() []*widgets.Dropdown {
	return nil
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
	return []stm.URL{}, nil
}

// Workers register workers
func (p *Plugin) Workers() map[string]job.Handler {
	return map[string]job.Handler{}
}

func init() {
	web.Register(&Plugin{})
}
