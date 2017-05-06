package site

import (
	"github.com/gin-gonic/gin"
	"github.com/kapmahc/fly/web"
	"github.com/kapmahc/fly/web/widgets"
)

func (p *Plugin) indexAdminNotices(c *gin.Context, lang string) (gin.H, error) {
	var items []Notice
	err := p.Db.Order("updated_at DESC").Find(&items).Error
	return gin.H{
		"items": items,
		"title": p.I18n.T(lang, "site.admin.notices.index.title"),
	}, err
}

func (p *Plugin) indexNotices(c *gin.Context, lang string) (gin.H, error) {
	var items []Notice
	err := p.Db.Order("updated_at DESC").Limit(60).Find(&items).Error
	return gin.H{
		"items": items,
		"title": p.I18n.T(lang, "site.notices.index.title"),
	}, err
}

type fmNotice struct {
	Body string `form:"body" binding:"required"`
	Type string `form:"type" binding:"required,max=8"`
}

func (p *Plugin) newNotice(c *gin.Context, lang string) (gin.H, error) {
	title := p.I18n.T(lang, "buttons.new")

	fm := widgets.NewForm(
		c.Request,
		lang,
		"/admin/notices",
		"",
		title,
		widgets.NewHiddenField("type", web.TypeMARKDOWN),
		widgets.NewTextarea("body", p.I18n.T(lang, "attributes.body"), "", 12),
	)
	return gin.H{"form": fm, "title": title}, nil
}

func (p *Plugin) createNotice(c *gin.Context, lang string, o interface{}) (interface{}, error) {
	fm := o.(*fmNotice)
	item := Notice{
		Media: web.Media{Type: fm.Type, Body: fm.Body},
	}
	if err := p.Db.Create(&item).Error; err != nil {
		return nil, err
	}
	return gin.H{}, nil
}

func (p *Plugin) editNotice(c *gin.Context, lang string) (gin.H, error) {
	title := p.I18n.T(lang, "buttons.edit")
	id := c.Param("id")
	var item Notice
	if err := p.Db.Where("id = ?", id).First(&item).Error; err != nil {
		return nil, err
	}
	fm := widgets.NewForm(
		c.Request,
		lang,
		"/admin/notices/"+id,
		"/admin/notices",
		title,
		widgets.NewHiddenField("type", item.Type),
		widgets.NewTextarea("body", p.I18n.T(lang, "attributes.body"), item.Body, 12),
	)
	return gin.H{"form": fm, "title": title}, nil
}

func (p *Plugin) updateNotice(c *gin.Context, lang string, o interface{}) (interface{}, error) {
	fm := o.(*fmNotice)
	if err := p.Db.Model(&Notice{}).
		Where("id = ?", c.Param("id")).
		Updates(map[string]interface{}{
			"body": fm.Body,
			"type": fm.Type,
		}).Error; err != nil {
		return nil, err
	}
	return gin.H{}, nil
}

func (p *Plugin) destroyNotice(c *gin.Context, l string) (interface{}, error) {
	err := p.Db.
		Where("id = ?", c.Param("id")).
		Delete(Notice{}).Error
	return gin.H{}, err
}
