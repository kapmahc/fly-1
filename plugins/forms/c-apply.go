package forms

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kapmahc/fly/web/widgets"
)

func (p *Plugin) _parseOptions(f *Field) ([]widgets.Option, error) {
	var options []widgets.Option
	buf := make(map[string]interface{})
	if err := json.Unmarshal([]byte(f.Body), &buf); err != nil {
		return nil, err
	}
	for k, v := range buf {
		options = append(options, widgets.NewOption(k, v))
	}
	return options, nil
}

func (p *Plugin) _parseValues(f *Field) []interface{} {
	var items []interface{}
	for _, s := range strings.Split(f.Value, ";") {
		items = append(items, s)
	}
	return items
}

func (p *Plugin) getFormApply(c *gin.Context, lang string) (gin.H, error) {
	item := c.MustGet("item").(*Form)
	title := p.I18n.T(lang, "buttons.apply") + "-" + item.Title

	fm := widgets.NewForm(
		c.Request,
		lang,
		fmt.Sprintf("/forms/apply/%d", item.ID),
		"/forms",
		title,
	)
	for _, f := range item.Fields {
		switch f.Type {
		case widgets.TypeCheckboxs:
			options, err := p._parseOptions(&f)
			if err != nil {
				return nil, err
			}

			fm.Append(widgets.NewCheckboxs(f.Name, f.Label, p._parseValues(&f), options...))
		case widgets.TypeRadios:
			options, err := p._parseOptions(&f)
			if err != nil {
				return nil, err
			}
			fm.Append(widgets.NewRadios(f.Name, f.Label, f.Value, options...))
		case "text":
			fm.Append(widgets.NewTextField(f.Name, f.Label, f.Value))
		case "email":
			fm.Append(widgets.NewEmailField(f.Name, f.Label, f.Value))
		}
	}
	// log.Printf("%+v", fm)
	return gin.H{
		"expired": item.Expire(),
		"title":   title,
		"form":    fm,
	}, nil
}

type fmApply struct {
	Username string `form:"username" binding:"required,max=255"`
	Email    string `form:"email" binding:"required,max=255"`
	Phone    string `form:"phone" binding:"required,max=255"`
}

func (p *Plugin) postFormApply(c *gin.Context, l string, o interface{}) (interface{}, error) {
	item := c.MustGet("item").(*Form)
	if item.Expire() {
		return nil, p.I18n.E(l, "forms.errors.expired")
	}
	fm := o.(*fmApply)
	var count int
	if err := p.Db.Model(&Record{}).Where("form_id = ? AND (phone = ? OR email = ?)", item.ID, fm.Phone, fm.Email).Count(&count).Error; err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, p.I18n.E(l, "forms.errors.already-apply")
	}

	data := c.Request.Form
	data.Del("email")
	data.Del("phone")
	data.Del("username")
	data.Del("authenticity_token")
	val, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	record := Record{
		Email:    fm.Email,
		Phone:    fm.Phone,
		Username: fm.Username,
		Value:    string(val),
		FormID:   item.ID,
	}
	if err := p.Db.Create(&record).Error; err != nil {
		return nil, err
	}
	p._sendEmail(l, item, &record, actApply)
	return gin.H{"message": p.I18n.T(l, "forms.messages.apply-success")}, nil
}
