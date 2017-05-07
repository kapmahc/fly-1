package forms

import (
	"github.com/gin-gonic/gin"
	"github.com/kapmahc/fly/web"
	"github.com/kapmahc/fly/web/widgets"
)

func (p *Plugin) indexForms(c *gin.Context, l string) (gin.H, error) {
	var items []Form
	err := p.Db.Order("updated_at DESC").Find(&items).Error
	return gin.H{
		"items": items,
		"title": p.I18n.T(l, "forms.index.title"),
	}, err
}

func (p *Plugin) newForm(c *gin.Context, lang string) (gin.H, error) {
	title := p.I18n.T(lang, "buttons.new")

	fm := widgets.NewForm(
		c.Request,
		lang,
		"/forms",
		"",
		title,
		widgets.NewTextField("title", p.I18n.T(lang, "attributes.title"), ""),
		widgets.NewHiddenField("type", web.TypeMARKDOWN),
		widgets.NewTextarea("body", p.I18n.T(lang, "attributes.body"), "", 8),
		widgets.NewTextField("deadline", p.I18n.T(lang, "attributes.shutDown"), ""),
		widgets.NewTextField("fields", p.I18n.T(lang, "forms.attributes.form.fields"), ""),
	)
	return gin.H{"form": fm, "title": title}, nil
}

func (p *Plugin) createForm(c *gin.Context, l string, o interface{}) (interface{}, error) {
	return gin.H{}, nil
}

func (p *Plugin) editForm(c *gin.Context, l string) (gin.H, error) {
	return gin.H{}, nil
}

type fm struct {
	Title    string `form:"title" binding:"required,max=255"`
	Deadline string `form:"deadline" binding:"required"`
	Body     string `form:"body" binding:"required"`
	Type     string `form:"type" binding:"required,max=8"`
	Fields   string `form:"fields" binding:"required"`
}

func (p *Plugin) updateForm(c *gin.Context, l string, o interface{}) (interface{}, error) {
	return gin.H{}, nil
}

func (p *Plugin) destroyForm(c *gin.Context, l string) (interface{}, error) {
	return gin.H{}, nil
}
