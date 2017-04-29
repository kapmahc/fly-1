package site

import (
	"github.com/gin-gonic/gin"
	"github.com/kapmahc/fly/web/widgets"
)

func (p *Plugin) indexLinks(c *gin.Context, lang string) (gin.H, error) {
	var items []Link
	err := p.Db.Order("loc ASC, sort_order ASC").Find(&items).Error
	return gin.H{
		"items": items,
		"title": p.I18n.T(lang, "site.admin.links.index.title"),
	}, err
}

type fmLink struct {
	Label     string `form:"label" binding:"required,max=255"`
	Href      string `form:"href" binding:"required,max=255"`
	Loc       string `form:"loc" binding:"required,max=32"`
	SortOrder int    `form:"sortOrder"`
}

func (p *Plugin) newLink(c *gin.Context, lang string) (gin.H, error) {
	title := p.I18n.T(lang, "buttons.new")

	fm := widgets.NewForm(
		c.Request,
		lang,
		"/admin/links",
		"",
		title,
		widgets.NewTextField("loc", p.I18n.T(lang, "attributes.loc"), ""),
		widgets.NewTextField("href", p.I18n.T(lang, "attributes.href"), ""),
		widgets.NewTextField("label", p.I18n.T(lang, "attributes.label"), ""),
		widgets.NewSortSelect("sortOrder", p.I18n.T(lang, "attributes.sortOrder"), 0, -10, 10),
	)
	return gin.H{"form": fm, "title": title}, nil
}

func (p *Plugin) createLink(c *gin.Context, lang string, o interface{}) (interface{}, error) {
	fm := o.(*fmLink)
	item := Link{
		Label:     fm.Label,
		Href:      fm.Href,
		Loc:       fm.Loc,
		SortOrder: fm.SortOrder,
	}
	if err := p.Db.Create(&item).Error; err != nil {
		return nil, err
	}
	return gin.H{}, nil
}

func (p *Plugin) editLink(c *gin.Context, lang string) (gin.H, error) {
	title := p.I18n.T(lang, "buttons.edit")
	id := c.Param("id")
	var item Link
	if err := p.Db.Where("id = ?", id).First(&item).Error; err != nil {
		return nil, err
	}
	fm := widgets.NewForm(
		c.Request,
		lang,
		"/admin/links/"+id,
		"/admin/links",
		title,

		widgets.NewTextField("loc", p.I18n.T(lang, "attributes.loc"), item.Loc),
		widgets.NewTextField("href", p.I18n.T(lang, "attributes.href"), item.Href),
		widgets.NewTextField("label", p.I18n.T(lang, "attributes.label"), item.Label),
		widgets.NewSortSelect("sortOrder", p.I18n.T(lang, "attributes.sortOrder"), item.SortOrder, -10, 10),
	)
	return gin.H{"form": fm, "title": title}, nil
}

func (p *Plugin) updateLink(c *gin.Context, lang string, o interface{}) (interface{}, error) {
	fm := o.(*fmLink)
	if err := p.Db.Model(&Link{}).
		Where("id = ?", c.Param("id")).
		Updates(map[string]interface{}{
			"loc":        fm.Loc,
			"label":      fm.Label,
			"href":       fm.Href,
			"sort_order": fm.SortOrder,
		}).Error; err != nil {
		return nil, err
	}
	return gin.H{}, nil
}

func (p *Plugin) destroyLink(c *gin.Context, l string) (interface{}, error) {
	err := p.Db.
		Where("id = ?", c.Param("id")).
		Delete(Link{}).Error
	return gin.H{}, err
}
