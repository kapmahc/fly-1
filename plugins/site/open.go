package site

import (
	"crypto/aes"
	"errors"
	"fmt"
	"html/template"
	"path"
	"reflect"
	"strings"
	"time"

	"github.com/SermoDigital/jose/crypto"
	log "github.com/Sirupsen/logrus"
	"github.com/facebookgo/inject"
	_redis "github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/kapmahc/fly/web"
	"github.com/kapmahc/fly/web/cache/redis"
	i_orm "github.com/kapmahc/fly/web/i18n/orm"
	"github.com/kapmahc/fly/web/job/rabbitmq"
	s_orm "github.com/kapmahc/fly/web/settings/orm"
	"github.com/kapmahc/fly/web/uploader/fs"
	"github.com/spf13/viper"
	"github.com/unrolled/render"
	"golang.org/x/text/language"
)

// Open init beans
func (p *Plugin) Open(g *inject.Graph) error {
	db, err := p.openDatabase()
	if err != nil {
		return err
	}
	// -------------------
	cip, err := aes.NewCipher([]byte(viper.GetString("secrets.aes")))
	if err != nil {
		return err
	}
	// -------------------
	var tags []language.Tag
	for _, l := range viper.GetStringSlice("languages") {
		lng, er := language.Parse(l)
		if er != nil {
			return er
		}
		tags = append(tags, lng)
	}
	// -------------------
	theme := viper.GetString("server.theme")
	up, err := fs.NewFileSystemStore(
		path.Join("public", "files"),
		web.Home()+"/files",
	)
	if err != nil {
		return err
	}
	// -------------------
	return g.Provide(
		&inject.Object{Value: []byte(viper.GetString("secrets.hmac")), Name: "hmac.key"},
		&inject.Object{Value: []byte(viper.GetString("secrets.jwt")), Name: "jwt.key"},
		&inject.Object{Value: crypto.SigningMethodHS512, Name: "jwt.method"},

		&inject.Object{Value: language.NewMatcher(tags)},
		&inject.Object{Value: cip},
		&inject.Object{Value: db},
		&inject.Object{Value: p.openRedis()},
		&inject.Object{Value: p.openRender(theme)},
		&inject.Object{Value: up},

		&inject.Object{Value: &redis.Store{}},
		&inject.Object{Value: rabbitmq.New(
			web.Name(),
			viper.GetString("rabbitmq.host"),
			viper.GetInt("rabbitmq.port"),
			viper.GetString("rabbitmq.user"),
			viper.GetString("rabbitmq.password"),
			viper.GetString("rabbitmq.virtual"),
		)},
		&inject.Object{Value: &i_orm.Store{}},
		&inject.Object{Value: &s_orm.Store{}},
	)
}

func (p *Plugin) openDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(viper.GetString("database.driver"), web.DataSource())
	if err != nil {
		return nil, err
	}
	if !web.IsProduction() {
		db.LogMode(true)
	}

	if err := db.DB().Ping(); err != nil {
		return nil, err
	}

	db.DB().SetMaxIdleConns(viper.GetInt("database.pool.max_idle"))
	db.DB().SetMaxOpenConns(viper.GetInt("database.pool.max_open"))
	return db, nil

}
func (p *Plugin) openRender(theme string) *render.Render {
	fm := template.FuncMap{
		"t": func(lang, code string, args ...interface{}) string {
			return p.I18n.T(lang, code, args...)
		},
		"tn": func(v interface{}) string {
			return reflect.TypeOf(v).String()
		},
		"dict": func(values ...interface{}) (gin.H, error) {
			dict := gin.H{}
			for i := 0; i < len(values); i += 2 {
				key, ok := values[i].(string)
				if !ok {
					return nil, errors.New("dict key must be string")
				}
				dict[key] = values[i+1]
			}
			return dict, nil
		},
		"even": func(i interface{}) bool {
			if i != nil {
				switch i.(type) {
				case int:
					return i.(int)%2 == 0
				case uint:
					return i.(uint)%2 == 0
				case int64:
					return i.(int64)%2 == 0
				case uint64:
					return i.(uint64)%2 == 0
				}
			}
			return false
		},
		"fmt": fmt.Sprintf,
		"eq": func(arg1, arg2 interface{}) bool {
			return arg1 == arg2
		},
		"str2htm": func(s string) template.HTML {
			return template.HTML(s)
		},
		"dtf": func(t interface{}) string {
			if t != nil {
				f := "Mon Jan _2 15:04:05 2006"
				switch t.(type) {
				case time.Time:
					return t.(time.Time).Format(f)
				case *time.Time:
					if t != (*time.Time)(nil) {
						return t.(*time.Time).Format(f)
					}
				}
			}
			return ""
		},
		"df": func(t interface{}) string {
			if t != nil {
				f := "Mon Jan _2 2006"
				switch t.(type) {
				case time.Time:
					return t.(time.Time).Format(f)
				case *time.Time:
					if t != (*time.Time)(nil) {
						return t.(*time.Time).Format(f)
					}
				}
			}
			return ""
		},
		"in": func(o interface{}, args []interface{}) bool {
			for _, v := range args {
				if o == v {
					return true
				}
			}
			return false
		},
		"starts": func(s string, b string) bool {
			return strings.HasPrefix(s, b)
		},
		"links": func(lang, loc string) []Link {
			var items []Link
			if err := p.Db.
				Where("loc = ?", lang, loc).
				Order("sort DESC").
				Find(&items).Error; err != nil {
				log.Error(err)
			}
			return items
		},
		"cards": func(loc string) []Card {
			var items []Card
			if err := p.Db.
				Where("loc = ?", loc).
				Order("sort_order DESC").
				Find(&items).Error; err != nil {
				log.Error(err)
			}
			return items
		},
	}
	return render.New(render.Options{
		Directory:     path.Join("themes", theme, "views"),
		Layout:        "application",
		Extensions:    []string{".html"},
		Funcs:         []template.FuncMap{fm},
		IsDevelopment: !web.IsProduction(),
	})
}
func (p *Plugin) openRedis() *_redis.Pool {
	return &_redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (_redis.Conn, error) {
			c, e := _redis.Dial(
				"tcp",
				fmt.Sprintf(
					"%s:%d",
					viper.GetString("redis.host"),
					viper.GetInt("redis.port"),
				),
			)
			if e != nil {
				return nil, e
			}
			if _, e = c.Do("SELECT", viper.GetInt("redis.db")); e != nil {
				c.Close()
				return nil, e
			}
			return c, nil
		},
		TestOnBorrow: func(c _redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}
