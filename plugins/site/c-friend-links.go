package site

import (
	"net/http"

	"github.com/kapmahc/h2o"
)

func (p *Plugin) indexFriendLinks(c *h2o.Context) error {
	var items []FriendLink
	if err := p.Db.Order("updated_at DESC").Find(&items).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, items)

}

type fmFriendLink struct {
	Title string `form:"title" validate:"required,max=255"`
	Home  string `form:"home" validate:"required,max=255"`
	Logo  string `form:"logo" validate:"required,max=255"`
}

func (p *Plugin) createFriendLink(c *h2o.Context) error {
	var fm fmFriendLink
	if err := c.Bind(&fm); err != nil {
		return err
	}
	item := FriendLink{
		Title: fm.Title,
		Logo:  fm.Logo,
		Home:  fm.Home,
	}
	if err := p.Db.Create(&item).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, item)
}

func (p *Plugin) showFriendLink(c *h2o.Context) error {
	var item FriendLink
	if err := p.Db.Where("id = ?", c.Param("id")).First(&item).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, item)
}

func (p *Plugin) updateFriendLink(c *h2o.Context) error {
	var fm fmFriendLink
	if err := c.Bind(&fm); err != nil {
		return err
	}
	if err := p.Db.Model(&FriendLink{}).
		Where("id = ?", c.Param("id")).
		Updates(map[string]interface{}{
			"home":  fm.Home,
			"title": fm.Title,
			"logo":  fm.Logo,
		}).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, h2o.H{})
}

func (p *Plugin) destroyFriendLink(c *h2o.Context) error {
	if err := p.Db.
		Where("id = ?", c.Param("id")).
		Delete(FriendLink{}).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, h2o.H{})
}
