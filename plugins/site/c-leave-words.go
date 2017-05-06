package site

import (
	"github.com/gin-gonic/gin"
	"github.com/kapmahc/fly/web/widgets"
)

func (p *Plugin) indexLeaveWords(c *gin.Context, lang string) (gin.H, error) {
	var items []LeaveWord
	err := p.Db.Order("created_at DESC").Find(&items).Error
	return gin.H{
		"items": items,
		"title": p.I18n.T(lang, "site.admin.leave-words.index.title"),
	}, err
}

type fmLeaveWord struct {
	Body string `form:"body" binding:"required,max=2048"`
}

func (p *Plugin) newLeaveWord(c *gin.Context, lang string) (gin.H, error) {
	title := p.I18n.T(lang, "site.leave-words.new.title")
	body := widgets.NewTextarea("body", p.I18n.T(lang, "attributes.body"), "", 12)
	body.Helper(p.I18n.T(lang, "site.helpers.leave-word.body"))

	fm := widgets.NewForm(
		c.Request,
		lang,
		"/leave-words",
		"/leave-words/new",
		title,
		body,
	)
	return gin.H{"form": fm, "title": title}, nil
}

func (p *Plugin) createLeaveWord(c *gin.Context, lang string, o interface{}) (interface{}, error) {
	fm := o.(*fmLeaveWord)
	item := LeaveWord{
		Body: fm.Body,
	}
	if err := p.Db.Create(&item).Error; err != nil {
		return nil, err
	}
	return gin.H{"message": p.I18n.T(lang, "success")}, nil
}

func (p *Plugin) destroyLeaveWord(c *gin.Context, l string) (interface{}, error) {
	err := p.Db.
		Where("id = ?", c.Param("id")).
		Delete(LeaveWord{}).Error
	return gin.H{}, err
}
