package site

import (
	"crypto/aes"
	"fmt"
	"path"
	"time"

	"github.com/SermoDigital/jose/crypto"
	"github.com/facebookgo/inject"
	_redis "github.com/garyburd/redigo/redis"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/kapmahc/fly/web"
	"github.com/kapmahc/fly/web/cache/redis"
	i_orm "github.com/kapmahc/fly/web/i18n/orm"
	"github.com/kapmahc/fly/web/job/rabbitmq"
	s_orm "github.com/kapmahc/fly/web/settings/orm"
	"github.com/kapmahc/fly/web/uploader/fs"
	"github.com/spf13/viper"
	"github.com/urfave/negroni"
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
	up, err := fs.NewFileSystemStore(
		path.Join("public", "files"),
		web.Home()+"/files",
	)
	if err != nil {
		return err
	}
	// -------------------
	return g.Provide(
		&inject.Object{Value: negroni.New()},
		&inject.Object{Value: mux.NewRouter()},

		&inject.Object{Value: []byte(viper.GetString("secrets.hmac")), Name: "hmac.key"},
		&inject.Object{Value: []byte(viper.GetString("secrets.jwt")), Name: "jwt.key"},
		&inject.Object{Value: crypto.SigningMethodHS512, Name: "jwt.method"},

		&inject.Object{Value: cip},
		&inject.Object{Value: db},
		&inject.Object{Value: p.openRedis()},
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
