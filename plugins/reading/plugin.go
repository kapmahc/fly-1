package reading

import (
	"github.com/facebookgo/inject"
	"github.com/ikeikeikeike/go-sitemap-generator/stm"
	"github.com/jinzhu/gorm"
	"github.com/kapmahc/fly/plugins/auth"
	"github.com/kapmahc/fly/web"
	"github.com/kapmahc/fly/web/i18n"
	"github.com/kapmahc/fly/web/job"
	"github.com/kapmahc/fly/web/settings"
	"github.com/urfave/cli"
	"golang.org/x/tools/blog/atom"
)

// Plugin plugin
type Plugin struct {
	Db       *gorm.DB           `inject:""`
	Jwt      *auth.Jwt          `inject:""`
	I18n     *i18n.I18n         `inject:""`
	Settings *settings.Settings `inject:""`
	Server   *job.Server        `inject:""`
}

// Init load config
func (p *Plugin) Init() {}

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

// Mount mount web points
func (p *Plugin) Mount() {

}

// Workers job handler
func (p *Plugin) Workers() map[string]job.Handler {
	return map[string]job.Handler{}
}

// Console console commands
func (p *Plugin) Console() []cli.Command {
	return []cli.Command{}
}

func init() {
	web.Register(&Plugin{})
}
