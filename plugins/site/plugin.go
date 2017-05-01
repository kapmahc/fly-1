package site

import (
	"github.com/gin-gonic/gin"
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
	Render   *render.Render     `inject:""`
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
				"site.dashboard.title",
				widgets.NewLink("site.admin.status.title", "/admin/site/status"),
				nil,
				widgets.NewLink("site.admin.links.index.title", "/admin/links"),
				widgets.NewLink("site.admin.cards.index.title", "/admin/cards"),
				nil,
				widgets.NewLink("site.admin.info.title", "/admin/site/info"),
				widgets.NewLink("site.admin.author.title", "/admin/site/author"),
				widgets.NewLink("site.admin.seo.title", "/admin/site/seo"),
				widgets.NewLink("site.admin.smtp.title", "/admin/site/smtp"),
				nil,
				widgets.NewLink("site.admin.users.index.title", "/admin/users"),
				widgets.NewLink("site.admin.locales.index.title", "/admin/locales"),
				widgets.NewLink("site.admin.notices.index.title", "/admin/notices"),
				widgets.NewLink("site.admin.leave-words.index.title", "/admin/leave-words"),
				widgets.NewLink("site.admin.friend-links.index.title", "/admin/friend-links"),
				widgets.NewLink("site.admin.posts.index.title", "/admin/posts"),
			),
		)
	}
	return items
}

// Atom rss.atom
func (p *Plugin) Atom(lang string) ([]*atom.Entry, error) {
	return []*atom.Entry{}, nil
}

// Sitemap sitemap.xml.gz
func (p *Plugin) Sitemap() ([]stm.URL, error) {
	var items []stm.URL
	items = append(
		items,
		stm.URL{"loc": "/leave-words/new"},
		stm.URL{"loc": "/notices"},
		stm.URL{"loc": "/posts"},
	)

	var posts []Post
	if err := p.Db.Select([]string{"name", "updated_at"}).Find(&posts).Error; err != nil {
		return nil, err
	}
	for _, it := range posts {
		items = append(items, stm.URL{"loc": "/posts/show/" + it.Name, "lastmod": it.UpdatedAt})
	}
	return items, nil
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
