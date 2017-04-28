package web

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/csrf"
	"github.com/kapmahc/fly/web/i18n"
	"github.com/kapmahc/fly/web/widgets"
	"github.com/unrolled/render"
)

// Wrap wrap
type Wrap struct {
	Render *render.Render `inject:""`
	I18n   *i18n.I18n     `inject:""`
}

// Redirect wrap redirect
func (p *Wrap) Redirect(f func(*gin.Context, string) (string, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		if u, e := f(c, c.MustGet(i18n.LOCALE).(string)); e == nil {
			c.Redirect(http.StatusFound, u)
		} else {
			log.Error(e)
			c.String(http.StatusInternalServerError, e.Error())
		}
	}
}

// FORM wrap form handler
func (p *Wrap) FORM(fm interface{}, fn func(*gin.Context, string, interface{}) (interface{}, error)) gin.HandlerFunc {
	return p.JSON(func(c *gin.Context, l string) (interface{}, error) {
		if err := c.Bind(fm); err != nil {
			return nil, err
		}
		return fn(c, l, fm)
	})
}

// HTML wrap html render
func (p *Wrap) HTML(t string, f func(*gin.Context, string) (gin.H, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		lang := c.MustGet(i18n.LOCALE).(string)
		if d, e := f(c, lang); e == nil {
			// -------------
			for k, v := range c.Keys {
				d[k] = v
			}
			// ------------
			d["lang"] = lang
			d["languages"], _ = p.I18n.Store.Languages()
			// -----------
			var dashboard []*widgets.Dropdown
			for _, en := range plugins {
				items := en.Dashboard(c)
				dashboard = append(dashboard, items...)
			}
			d["dashboard"] = dashboard
			// -----------
			d["csrf"] = csrf.Token(c.Request)
			// -----------
			p.Render.HTML(c.Writer, http.StatusOK, t, d)
		} else {
			log.Error(e)
			c.String(http.StatusInternalServerError, e.Error())
		}
	}
}

// XML wrap xml render
func (p *Wrap) XML(f func(*gin.Context, string) (interface{}, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		if v, e := f(c, c.MustGet(i18n.LOCALE).(string)); e == nil {
			c.XML(http.StatusOK, v)
		} else {
			log.Error(e)
			c.String(http.StatusInternalServerError, e.Error())
		}
	}
}

// JSON wrap json render
func (p *Wrap) JSON(f func(*gin.Context, string) (interface{}, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		if v, e := f(c, c.MustGet(i18n.LOCALE).(string)); e == nil {
			c.JSON(http.StatusOK, v)
		} else {
			log.Error(e)
			c.String(http.StatusInternalServerError, e.Error())
		}
	}
}
