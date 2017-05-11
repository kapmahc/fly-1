package auth

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kapmahc/fly/web"
	"github.com/kapmahc/fly/web/i18n"
)

type fmAttachmentNew struct {
	Type string `form:"type" binding:"required,max=255"`
	ID   uint   `form:"id"`
}

func (p *Plugin) createAttachment(c *web.Context, o interface{}) (interface{}, error) {
	user := c.Get(CurrentUser).(*User)
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

type fmAttachmentEdit struct {
	Title string `form:"title" binding:"required,max=255"`
}

func (p *Plugin) updateAttachment(c *web.Context, o interface{}) (interface{}, error) {
	a := c.Get("item").(*Attachment)
	fm := o.(*fmAttachmentEdit)
	if err := p.Db.Model(a).Update("title", fm.Title).Error; err != nil {
		return nil, err
	}

	return web.H{}, nil
}

func (p *Plugin) destroyAttachment(c *web.Context) (interface{}, error) {
	a := c.Get("item").(*Attachment)
	err := p.Db.Delete(a).Error
	if err != nil {
		return nil, err
	}
	return a, p.Uploader.Remove(a.URL)
}

func (p *Plugin) showAttachment(c *web.Context) (interface{}, error) {
	a := c.Get("item").(*Attachment)
	return a, nil
}

func (p *Plugin) indexAttachments(c *web.Context) (interface{}, error) {
	user := c.Get(CurrentUser).(*User)
	isa := c.Get(IsAdmin).(bool)
	var items []Attachment
	qry := p.Db
	if !isa {
		qry = qry.Where("user_id = ?", user.ID)
	}
	if err := qry.Order("updated_at DESC").Find(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}

func (p *Plugin) canEditAttachment(w http.ResponseWriter, r *http.Request, n http.HandlerFunc) {
	user := r.Context().Value(web.K(CurrentUser)).(*User)
	lng := r.Context().Value(web.K(i18n.LOCALE)).(string)
	vars := mux.Vars(r)

	var a Attachment
	err := p.Db.Where("id = ?", vars["id"]).First(&a).Error
	if err == nil {
		if user.ID == a.UserID || r.Context().Value(IsAdmin).(bool) {
			n(w, r.WithContext(context.WithValue(r.Context(), web.K("item"), &a)))
		} else {
			http.Error(w, p.I18n.T(lng, "auth.errors.not-allow"), http.StatusInternalServerError)
		}
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
