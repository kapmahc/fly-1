package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kapmahc/fly/web/widgets"
)

type fmAttachmentNew struct {
	// Type string `form:"type" binding:"required,max=255"`
	// ID   uint   `form:"id"`
}

func (p *Plugin) newAttachment(c *gin.Context, lang string) (gin.H, error) {

	return gin.H{
		"type":  c.Query("type"),
		"id":    c.Query("id"),
		"title": p.I18n.T(lang, "buttons.upload"),
	}, nil
}

func (p *Plugin) createAttachment(c *gin.Context, l string, o interface{}) (interface{}, error) {
	user := c.MustGet(CurrentUser).(*User)
	// fm := o.(*fmAttachmentNew)

	file, header, err := c.Request.FormFile("file")
	if err != nil {
		return nil, err
	}

	url, size, err := p.Uploader.Save(header)
	if err != nil {
		return nil, err
	}

	// http://golang.org/pkg/net/http/#DetectContentType
	buf := make([]byte, 512)
	file.Seek(0, 0)
	if _, err = file.Read(buf); err != nil {
		return nil, err
	}

	a := Attachment{
		Title:        header.Filename,
		URL:          url,
		UserID:       user.ID,
		MediaType:    http.DetectContentType(buf),
		Length:       size / 1024,
		ResourceType: DefaultResourceType, //fm.Type,
		ResourceID:   DefaultResourceID,   //fm.ID,
	}
	if err := p.Db.Create(&a).Error; err != nil {
		return nil, err
	}
	return a, nil
}

func (p *Plugin) editAttachment(c *gin.Context, lang string) (gin.H, error) {
	title := p.I18n.T(lang, "buttons.edit")
	id := c.Param("id")
	item := c.MustGet("attachment").(*Attachment)
	fm := widgets.NewForm(
		c.Request,
		lang,
		"/attachments/"+id,
		"/attachments",
		title,
		widgets.NewTextField("title", p.I18n.T(lang, "attributes.title"), item.Title),
	)
	return gin.H{"form": fm, "title": title}, nil
}

type fmAttachmentEdit struct {
	Title string `form:"title" binding:"required,max=255"`
}

func (p *Plugin) updateAttachment(c *gin.Context, l string, o interface{}) (interface{}, error) {
	a := c.MustGet("attachment").(*Attachment)
	fm := o.(*fmAttachmentEdit)
	if err := p.Db.Model(a).Update("title", fm.Title).Error; err != nil {
		return nil, err
	}

	return gin.H{}, nil
}

func (p *Plugin) destroyAttachment(c *gin.Context, l string) (interface{}, error) {
	a := c.MustGet("attachment").(*Attachment)
	err := p.Db.Delete(a).Error
	if err != nil {
		return nil, err
	}
	return a, p.Uploader.Remove(a.URL)
}

func (p *Plugin) indexAttachments(c *gin.Context, l string) (gin.H, error) {
	user := c.MustGet(CurrentUser).(*User)
	isa := c.MustGet(IsAdmin).(bool)
	var items []Attachment
	qry := p.Db
	if !isa {
		qry = qry.Where("user_id = ?", user.ID)
	}
	if err := qry.Order("updated_at DESC").Find(&items).Error; err != nil {
		return nil, err
	}

	return gin.H{
		"title": p.I18n.T(l, "auth.attachments.index.title"),
		"items": items,
	}, nil
}

func (p *Plugin) canEditAttachment(c *gin.Context) {
	user := c.MustGet(CurrentUser).(*User)

	var a Attachment
	err := p.Db.Where("id = ?", c.Param("id")).First(&a).Error
	if err == nil {
		if user.ID == a.UserID || c.MustGet(IsAdmin).(bool) {
			c.Set("attachment", &a)
			return
		}
		c.AbortWithStatus(http.StatusForbidden)
	} else {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
}
