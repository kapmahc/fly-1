package site

import (
	"github.com/gin-gonic/gin"
	"github.com/kapmahc/fly/web/widgets"
)

func (p *Plugin) indexFriendLinks(c *gin.Context, lang string) (gin.H, error) {
	var items []FriendLink
	err := p.Db.Order("updated_at DESC").Find(&items).Error
	return gin.H{
		"items": items,
		"title": p.I18n.T(lang, "site.admin.friend-links.index.title"),
	}, err
}

type fmFriendLink struct {
	Title string `form:"title" binding:"required,max=255"`
	Home  string `form:"home" binding:"required,max=255"`
	Logo  string `form:"logo" binding:"required,max=255"`
}

func (p *Plugin) newFriendLink(c *gin.Context, lang string) (gin.H, error) {
	title := p.I18n.T(lang, "buttons.new")

	fm := widgets.NewForm(
		c.Request,
		lang,
		"/admin/friend-links",
		"",
		title,
		widgets.NewTextField("title", p.I18n.T(lang, "attributes.title"), ""),
		widgets.NewTextField("home", p.I18n.T(lang, "attributes.home"), ""),
		widgets.NewTextField("logo", p.I18n.T(lang, "attributes.logo"), ""),
	)
	return gin.H{"form": fm, "title": title}, nil
}

func (p *Plugin) createFriendLink(c *gin.Context, lang string, o interface{}) (interface{}, error) {
	fm := o.(*fmFriendLink)
	item := FriendLink{
		Title: fm.Title,
		Logo:  fm.Logo,
		Home:  fm.Home,
	}
	if err := p.Db.Create(&item).Error; err != nil {
		return nil, err
	}
	return gin.H{}, nil
}

func (p *Plugin) editFriendLink(c *gin.Context, lang string) (gin.H, error) {
	title := p.I18n.T(lang, "buttons.edit")
	id := c.Param("id")
	var item FriendLink
	if err := p.Db.Where("id = ?", id).First(&item).Error; err != nil {
		return nil, err
	}
	fm := widgets.NewForm(
		c.Request,
		lang,
		"/admin/friend-links/"+id,
		"/admin/friend-links",
		title,
		widgets.NewTextField("title", p.I18n.T(lang, "attributes.title"), item.Title),
		widgets.NewTextField("home", p.I18n.T(lang, "attributes.home"), item.Home),
		widgets.NewTextField("logo", p.I18n.T(lang, "attributes.logo"), item.Logo),
	)
	return gin.H{"form": fm, "title": title}, nil
}

func (p *Plugin) updateFriendLink(c *gin.Context, lang string, o interface{}) (interface{}, error) {
	fm := o.(*fmFriendLink)
	if err := p.Db.Model(&FriendLink{}).
		Where("id = ?", c.Param("id")).
		Updates(map[string]interface{}{
			"home":  fm.Home,
			"title": fm.Title,
			"logo":  fm.Logo,
		}).Error; err != nil {
		return nil, err
	}
	return gin.H{}, nil
}

func (p *Plugin) destroyFriendLink(c *gin.Context, l string) (interface{}, error) {
	err := p.Db.
		Where("id = ?", c.Param("id")).
		Delete(FriendLink{}).Error
	return gin.H{}, err
}
