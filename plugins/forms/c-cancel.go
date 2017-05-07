package forms

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kapmahc/fly/web/widgets"
)

func (p *Plugin) getFormCancel(c *gin.Context, lang string) (gin.H, error) {
	item := c.MustGet("item").(*Form)
	title := p.I18n.T(lang, "buttons.cancel") + "-" + item.Title

	fm := widgets.NewForm(
		c.Request,
		lang,
		fmt.Sprintf("/forms/cancel/%d", item.ID),
		"/forms",
		title,
		widgets.NewTextField("who", p.I18n.T(lang, "forms.attributes.form.phone-or-email"), ""),
	)
	return gin.H{
		"expired": item.Expire(),
		"title":   title,
		"form":    fm,
	}, nil
}

type fmCancel struct {
	Who string `form:"who" binding:"required,max=255"`
}

func (p *Plugin) postFormCancel(c *gin.Context, l string, o interface{}) (interface{}, error) {
	item := c.MustGet("item").(*Form)
	if item.Expire() {
		return nil, p.I18n.E(l, "forms.errors.expired")
	}
	fm := o.(*fmCancel)
	var record Record
	if err := p.Db.Where("form_id = ? AND (phone = ? OR email = ?)", item.ID, fm.Who, fm.Who).First(&record).Error; err != nil {
		return nil, err
	}

	if err := p.Db.Delete(&record).Error; err != nil {
		return nil, err
	}
	p._sendEmail(l, item, &record, actCancel)
	return gin.H{}, nil
}
