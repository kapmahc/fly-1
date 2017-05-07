package forms

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kapmahc/fly/web"
	"github.com/kapmahc/fly/web/widgets"
)

func (p *Plugin) indexForms(c *gin.Context, l string) (gin.H, error) {
	var items []Form
	err := p.Db.Order("updated_at DESC").Find(&items).Error
	return gin.H{
		"items": items,
		"title": p.I18n.T(l, "forms.index.title"),
	}, err
}

func (p *Plugin) newForm(c *gin.Context, lang string) (gin.H, error) {
	title := p.I18n.T(lang, "buttons.new")
	deadline := widgets.NewTextField("deadline", p.I18n.T(lang, "attributes.shutDown"), "")
	deadline.Helper(p.I18n.T(lang, "helpers.date"))
	fields := widgets.NewTextarea("fields", p.I18n.T(lang, "forms.attributes.form.fields"), "", 16)
	fields.Helper(p.I18n.T(lang, "helpers.json"))
	fm := widgets.NewForm(
		c.Request,
		lang,
		"/forms",
		"/forms/manage",
		title,
		widgets.NewTextField("title", p.I18n.T(lang, "attributes.title"), ""),
		widgets.NewHiddenField("type", web.TypeMARKDOWN),
		widgets.NewTextarea("body", p.I18n.T(lang, "attributes.body"), "", 8),
		deadline,
		fields,
	)
	return gin.H{"form": fm, "title": title}, nil
}

func (p *Plugin) _parseFields(l, s string) ([]Field, error) {
	var fields []Field
	var items []interface{}
	if err := json.Unmarshal([]byte(s), &items); err != nil {
		return nil, err
	}

	for _, item := range items {
		bad := fmt.Sprintf(
			"%s: %+v",
			p.I18n.T(l, "errors.bad-format"),
			item,
		)
		val, ok := item.(map[string]interface{})
		if !ok {
			return nil, errors.New(bad)
		}
		field := Field{}
		if field.Name, ok = val["name"].(string); !ok {
			return nil, errors.New(bad + " name")
		}
		if field.Label, ok = val["label"].(string); !ok {
			return nil, errors.New(bad + " label")
		}
		if field.Type, ok = val["type"].(string); !ok {
			return nil, errors.New(bad + " type")
		}
		if field.Value, ok = val["value"].(string); !ok {
			return nil, errors.New(bad + " value")
		}

		switch {
		case field.Type == widgets.TypeCheckboxs || field.Type == widgets.TypeRadios:
			// log.Printf("%v\n", reflect.TypeOf(val["body"]))
			options, ok := val["body"].(map[string]interface{})
			if !ok {
				return nil, errors.New(bad + " body")
			}
			buf, err := json.Marshal(options)
			if err != nil {
				return nil, err
			}
			field.Body = string(buf)
		}

		fields = append(fields, field)
	}
	return fields, nil
}

func (p *Plugin) createForm(c *gin.Context, l string, o interface{}) (interface{}, error) {
	fm := o.(*fmForm)

	deadline, err := time.Parse(web.FormatDateInput, fm.Deadline)
	if err != nil {
		return nil, err
	}
	fields, err := p._parseFields(l, fm.Fields)
	if err != nil {
		return nil, err
	}
	// log.Printf("FIELDS: %+v", fields)

	item := Form{
		Title:    fm.Title,
		Deadline: deadline,
		Media: web.Media{
			Body: fm.Body,
			Type: fm.Type,
		},
	}
	if err := p.Db.Create(&item).Error; err != nil {
		return nil, err
	}
	for _, field := range fields {
		field.FormID = item.ID
		if err := p.Db.Create(&field).Error; err != nil {
			return nil, err
		}
	}

	return gin.H{}, nil
}

func (p *Plugin) _buildFields(fm *Form) ([]gin.H, error) {
	var items []gin.H
	for _, f := range fm.Fields {
		it := gin.H{
			"name":  f.Name,
			"label": f.Label,
			"type":  f.Type,
			"value": f.Value,
		}
		switch {
		case f.Type == widgets.TypeCheckboxs || f.Type == widgets.TypeRadios:
			options := make(map[string]interface{})
			if err := json.Unmarshal([]byte(f.Body), &options); err != nil {
				return nil, err
			}
			it["body"] = options
		}
		items = append(items, it)
	}
	return items, nil
}

func (p *Plugin) _mustSelectForm(c *gin.Context, l string) error {
	var item Form
	if err := p.Db.Where("id = ?", c.Param("id")).First(&item).Error; err != nil {
		return err
	}
	if err := p.Db.Model(&item).Association("Fields").Find(&item.Fields).Error; err != nil {
		return err
	}
	if err := p.Db.Model(&item).Association("Records").Find(&item.Records).Error; err != nil {
		return err
	}
	c.Set("item", &item)
	return nil
}

func (p *Plugin) editForm(c *gin.Context, lang string) (gin.H, error) {
	item := c.MustGet("item").(*Form)

	title := p.I18n.T(lang, "buttons.edit")
	deadline := widgets.NewTextField("deadline", p.I18n.T(lang, "attributes.shutDown"), item.Deadline.Format(web.FormatDateInput))
	deadline.Helper(p.I18n.T(lang, "helpers.date"))

	options, err := p._buildFields(item)
	if err != nil {
		return nil, err
	}
	// log.Printf("%+v\n", options)
	buf, err := json.MarshalIndent(options, "", "  ")
	if err != nil {
		return nil, err
	}

	fields := widgets.NewTextarea("fields", p.I18n.T(lang, "forms.attributes.form.fields"), string(buf), 16)
	fields.Helper(p.I18n.T(lang, "helpers.json"))

	fm := widgets.NewForm(
		c.Request,
		lang,
		fmt.Sprintf("/forms/edit/%d", item.ID),
		"/forms/manage",
		title,
		widgets.NewTextField("title", p.I18n.T(lang, "attributes.title"), item.Title),
		widgets.NewHiddenField("type", web.TypeMARKDOWN),
		widgets.NewTextarea("body", p.I18n.T(lang, "attributes.body"), item.Body, 8),
		deadline,
		fields,
	)
	return gin.H{"form": fm, "title": title}, nil
}

type fmForm struct {
	Title    string `form:"title" binding:"required,max=255"`
	Deadline string `form:"deadline" binding:"required"`
	Body     string `form:"body" binding:"required"`
	Type     string `form:"type" binding:"required,max=8"`
	Fields   string `form:"fields" binding:"required"`
}

func (p *Plugin) updateForm(c *gin.Context, l string, o interface{}) (interface{}, error) {
	item := c.MustGet("item").(*Form)
	fm := o.(*fmForm)

	deadline, err := time.Parse(web.FormatDateInput, fm.Deadline)
	if err != nil {
		return nil, err
	}
	fields, err := p._parseFields(l, fm.Fields)
	if err != nil {
		return nil, err
	}
	// log.Printf("FIELDS: %+v", fields)

	if err := p.Db.Model(&item).Updates(map[string]interface{}{
		"title":    fm.Title,
		"type":     fm.Type,
		"body":     fm.Body,
		"deadline": deadline,
	}).Error; err != nil {
		return nil, err
	}
	if err := p.Db.Model(item).Association("Fields").Clear().Error; err != nil {
		return nil, err
	}
	for _, f := range fields {
		f.FormID = item.ID
		if err := p.Db.Create(&f).Error; err != nil {
			return nil, err
		}
	}

	return gin.H{}, nil
}

func (p *Plugin) destroyForm(c *gin.Context, l string) (interface{}, error) {
	id := c.Param("id")
	if err := p.Db.Where("form_id = ?", id).Delete(Field{}).Error; err != nil {
		return nil, err
	}
	if err := p.Db.Where("form_id = ?", id).Delete(Record{}).Error; err != nil {
		return nil, err
	}
	if err := p.Db.Where("id = ?", id).Delete(Form{}).Error; err != nil {
		return nil, err
	}

	return gin.H{}, nil
}
