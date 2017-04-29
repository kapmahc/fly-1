package site

import (
	"github.com/gin-gonic/gin"
	"github.com/kapmahc/fly/web/widgets"
)

func (p *Plugin) indexCards(c *gin.Context, lang string) (gin.H, error) {
	var items []Card
	err := p.Db.Order("loc ASC, sort_order ASC").Find(&items).Error
	return gin.H{
		"items": items,
		"title": p.I18n.T(lang, "site.admin.cards.index.title"),
	}, err
}

type fmCard struct {
	Loc       string `form:"loc" binding:"required,max=32"`
	Title     string `form:"title" binding:"required,max=255"`
	Summary   string `form:"loc" binding:"required,max=800"`
	Href      string `form:"href" binding:"required,max=255"`
	Logo      string `form:"logo" binding:"required,max=255"`
	SortOrder int    `form:"sortOrder"`
}

func (p *Plugin) newCard(c *gin.Context, lang string) (gin.H, error) {
	title := p.I18n.T(lang, "buttons.new")

	fm := widgets.NewForm(
		c.Request,
		lang,
		"/admin/cards",
		"",
		title,
		widgets.NewTextField("loc", p.I18n.T(lang, "attributes.loc"), ""),
		widgets.NewTextField("href", p.I18n.T(lang, "attributes.href"), ""),
		widgets.NewTextField("title", p.I18n.T(lang, "attributes.title"), ""),
		widgets.NewTextarea("summary", p.I18n.T(lang, "attributes.summary"), "", 8),
		widgets.NewTextField("logo", p.I18n.T(lang, "attributes.logo"), ""),
		widgets.NewSortSelect("sortOrder", p.I18n.T(lang, "attributes.sortOrder"), 0, -10, 10),
	)
	return gin.H{"form": fm, "title": title}, nil
}

func (p *Plugin) createCard(c *gin.Context, lang string, o interface{}) (interface{}, error) {
	fm := o.(*fmCard)
	item := Card{
		Title:     fm.Title,
		Logo:      fm.Logo,
		Href:      fm.Href,
		Summary:   fm.Summary,
		SortOrder: fm.SortOrder,
		Loc:       fm.Loc,
		Action:    "buttons.view",
	}
	if err := p.Db.Create(&item).Error; err != nil {
		return nil, err
	}
	return gin.H{}, nil
}

func (p *Plugin) editCard(c *gin.Context, lang string) (gin.H, error) {
	title := p.I18n.T(lang, "buttons.edit")
	id := c.Param("id")
	var item Card
	if err := p.Db.Where("id = ?", id).First(&item).Error; err != nil {
		return nil, err
	}
	fm := widgets.NewForm(
		c.Request,
		lang,
		"/admin/cards/"+id,
		"/admin/cards",
		title,
		widgets.NewTextField("loc", p.I18n.T(lang, "attributes.loc"), item.Loc),
		widgets.NewTextField("href", p.I18n.T(lang, "attributes.href"), item.Href),
		widgets.NewTextField("title", p.I18n.T(lang, "attributes.title"), item.Title),
		widgets.NewTextarea("summary", p.I18n.T(lang, "attributes.summary"), item.Summary, 8),
		widgets.NewTextField("logo", p.I18n.T(lang, "attributes.logo"), item.Logo),
		widgets.NewSortSelect("sortOrder", p.I18n.T(lang, "attributes.sortOrder"), item.SortOrder, -10, 10),
	)
	return gin.H{"form": fm, "title": title}, nil
}

func (p *Plugin) updateCard(c *gin.Context, lang string, o interface{}) (interface{}, error) {
	fm := o.(*fmCard)
	if err := p.Db.Model(&Card{}).
		Where("id = ?", c.Param("id")).
		Updates(map[string]interface{}{
			"href":       fm.Href,
			"title":      fm.Title,
			"logo":       fm.Logo,
			"sort_order": fm.SortOrder,
			"loc":        fm.Loc,
			"summary":    fm.Summary,
		}).Error; err != nil {
		return nil, err
	}
	return gin.H{}, nil
}

func (p *Plugin) destroyCard(c *gin.Context, l string) (interface{}, error) {
	err := p.Db.
		Where("id = ?", c.Param("id")).
		Delete(Card{}).Error
	return gin.H{}, err
}
