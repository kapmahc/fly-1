package site

import (
	"github.com/gin-gonic/gin"
	"github.com/kapmahc/fly/web/widgets"
)

func (p *Plugin) getAdminLocales(c *gin.Context, lang string) (gin.H, error) {
	title := p.I18n.T(lang, "site.admin.locales.index.title")
	items, err := p.I18n.Store.All(lang)
	if err != nil {
		return nil, err
	}
	return gin.H{
		"title": title,
		"items": items,
	}, nil
}

func (p *Plugin) editAdminLocales(c *gin.Context, lang string) (gin.H, error) {
	title := p.I18n.T(lang, "buttons.new")
	code := c.Query("code")
	message := ""
	if code != "" {
		title = p.I18n.T(lang, "buttons.edit")
		message = p.I18n.T(lang, code)
	}
	fm := widgets.NewForm(
		c.Request,
		lang,
		"/admin/locales",
		"",
		title,
		widgets.NewTextField("code", p.I18n.T(lang, "site.attributes.locale.code"), code),
		widgets.NewTextarea("message", p.I18n.T(lang, "site.attributes.locale.message"), message, 6),
	)
	return gin.H{"form": fm, "title": title}, nil
}

func (p *Plugin) deleteAdminLocales(c *gin.Context, lang string) (interface{}, error) {
	err := p.I18n.Store.Del(lang, c.Param("code"))
	return gin.H{}, err
}

type fmLocale struct {
	Code    string `form:"code" binding:"required,max=255"`
	Message string `form:"message" binding:"required"`
}

func (p *Plugin) saveAdminLocales(c *gin.Context, lang string, o interface{}) (interface{}, error) {
	fm := o.(*fmLocale)
	err := p.I18n.Set(lang, fm.Code, fm.Message)
	return gin.H{}, err
}
