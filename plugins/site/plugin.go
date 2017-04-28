package site

import (
	"github.com/ikeikeikeike/go-sitemap-generator/stm"
	"github.com/jinzhu/gorm"
	"github.com/kapmahc/fly/plugins/auth"
	"github.com/kapmahc/fly/web"
	"github.com/kapmahc/fly/web/cache"
	"github.com/kapmahc/fly/web/i18n"
	"github.com/kapmahc/fly/web/job"
	"github.com/kapmahc/fly/web/settings"
	"github.com/kapmahc/fly/web/widgets"
	"github.com/spf13/viper"
	"github.com/unrolled/render"
	"golang.org/x/text/language"
	"golang.org/x/tools/blog/atom"
)

// Plugin plugin
type Plugin struct {
	Db       *gorm.DB           `inject:""`
	Jwt      *auth.Jwt          `inject:""`
	Dao      *auth.Dao          `inject:""`
	I18n     *i18n.I18n         `inject:""`
	Settings *settings.Settings `inject:""`
	Server   *job.Server        `inject:""`
	Cache    *cache.Cache       `inject:""`
	Wrap     *web.Wrap          `inject:""`
	Matcher  language.Matcher   `inject:""`
	Render   *render.Render     `inject:""`
}

// Init load config
func (p *Plugin) Init() {}

// Dashboard dashboard nav
func (p *Plugin) Dashboard() []*widgets.Dropdown {
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

// Workers register workers
func (p *Plugin) Workers() map[string]job.Handler {
	return map[string]job.Handler{}
}

func init() {
	viper.SetDefault("redis", map[string]interface{}{
		"host": "localhost",
		"port": 6379,
		"db":   8,
	})

	viper.SetDefault("rabbitmq", map[string]interface{}{
		"user":     "guest",
		"password": "guest",
		"host":     "localhost",
		"port":     "5672",
		"virtual":  "fly-dev",
	})

	viper.SetDefault("database", map[string]interface{}{
		"driver": "postgres",
		"args": map[string]interface{}{
			"host":     "localhost",
			"port":     5432,
			"user":     "postgres",
			"password": "",
			"dbname":   "fly_dev",
			"sslmode":  "disable",
		},
		"pool": map[string]int{
			"max_open": 180,
			"max_idle": 6,
		},
	})

	viper.SetDefault("server", map[string]interface{}{
		"port":  8080,
		"ssl":   false,
		"name":  "localhost",
		"theme": "bootstrap",
	})

	viper.SetDefault("secrets", map[string]interface{}{
		"jwt":     web.Random(32),
		"aes":     web.Random(32),
		"csrf":    web.Random(32),
		"session": web.Random(32),
		"hmac":    web.Random(32),
	})

	viper.SetDefault("elasticsearch", map[string]interface{}{
		"host": "localhost",
		"port": 9200,
	})

	web.Register(&Plugin{})
}
