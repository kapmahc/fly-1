package site

import (
	"github.com/gin-gonic/gin"
	"github.com/kapmahc/fly/web"
	"github.com/kapmahc/fly/web/widgets"
)

func (p *Plugin) indexAdminPosts(c *gin.Context, lang string) (gin.H, error) {
	var items []Post
	err := p.Db.Order("updated_at DESC").Find(&items).Error
	return gin.H{
		"items": items,
		"title": p.I18n.T(lang, "site.admin.posts.index.title"),
	}, err
}

func (p *Plugin) indexPosts(c *gin.Context, lang string) (gin.H, error) {
	var items []Post
	err := p.Db.Order("updated_at DESC").Find(&items).Error
	return gin.H{
		"items": items,
		"title": p.I18n.T(lang, "site.posts.index.title"),
	}, err
}

func (p *Plugin) showPost(c *gin.Context, lang string) (gin.H, error) {
	name := c.Param("name")
	if name == "" {
		return nil, p.I18n.E(lang, "errors.not-found")
	}
	var item Post
	err := p.Db.Where("name = ?", name[1:]).First(&item).Error
	return gin.H{
		"item":  item,
		"title": item.Title,
	}, err
}

type fmPost struct {
	Name  string `form:"name" binding:"required,max=255"`
	Title string `form:"body" binding:"required,max=255"`
	Body  string `form:"body" binding:"required"`
	Type  string `form:"type" binding:"required,max=8"`
}

func (p *Plugin) newPost(c *gin.Context, lang string) (gin.H, error) {
	title := p.I18n.T(lang, "buttons.new")

	fm := widgets.NewForm(
		c.Request,
		lang,
		"/admin/posts",
		"",
		title,
		widgets.NewHiddenField("type", web.TypeMARKDOWN),
		widgets.NewTextField("name", p.I18n.T(lang, "attributes.name"), ""),
		widgets.NewTextField("title", p.I18n.T(lang, "attributes.title"), ""),
		widgets.NewTextarea("body", p.I18n.T(lang, "attributes.body"), "", 12),
	)
	return gin.H{"form": fm, "title": title}, nil
}

func (p *Plugin) createPost(c *gin.Context, lang string, o interface{}) (interface{}, error) {
	fm := o.(*fmPost)
	item := Post{
		Media: web.Media{Type: fm.Type, Body: fm.Body},
		Title: fm.Title,
		Name:  fm.Name,
	}
	if err := p.Db.Create(&item).Error; err != nil {
		return nil, err
	}
	return gin.H{}, nil
}

func (p *Plugin) editPost(c *gin.Context, lang string) (gin.H, error) {
	title := p.I18n.T(lang, "buttons.edit")
	id := c.Param("id")
	var item Post
	if err := p.Db.Where("id = ?", id).First(&item).Error; err != nil {
		return nil, err
	}
	fm := widgets.NewForm(
		c.Request,
		lang,
		"/admin/posts/"+id,
		"/admin/posts",
		title,
		widgets.NewTextField("name", p.I18n.T(lang, "attributes.name"), item.Name),
		widgets.NewTextField("title", p.I18n.T(lang, "attributes.title"), item.Title),
		widgets.NewHiddenField("type", item.Type),
		widgets.NewTextarea("body", p.I18n.T(lang, "attributes.body"), item.Body, 12),
	)
	return gin.H{"form": fm, "title": title}, nil
}

func (p *Plugin) updatePost(c *gin.Context, lang string, o interface{}) (interface{}, error) {
	fm := o.(*fmPost)
	if err := p.Db.Model(&Post{}).
		Where("id = ?", c.Param("id")).
		Updates(map[string]interface{}{
			"body":  fm.Body,
			"type":  fm.Type,
			"name":  fm.Name,
			"title": fm.Title,
		}).Error; err != nil {
		return nil, err
	}
	return gin.H{}, nil
}

func (p *Plugin) destroyPost(c *gin.Context, l string) (interface{}, error) {
	err := p.Db.
		Where("id = ?", c.Param("id")).
		Delete(Post{}).Error
	return gin.H{}, err
}
