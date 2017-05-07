package forms

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

func (p *Plugin) _exportForm(f *Form) ([]string, [][]string, error) {
	var header []string
	for _, f := range f.Fields {
		header = append(header, f.Label)
	}

	var items [][]string
	for _, r := range f.Records {
		var row []string
		val := make(map[string]interface{})
		if err := json.Unmarshal([]byte(r.Value), &val); err != nil {
			return nil, nil, err
		}
		for _, f := range f.Fields {
			switch f.Name {
			case "phone":
				row = append(row, r.Phone)
			case "email":
				row = append(row, r.Email)
			case "username":
				row = append(row, r.Username)
			default:
				row = append(row, fmt.Sprintf("%+v", val[f.Name]))
			}
		}
		items = append(items, row)
	}

	return header, items, nil
}

func (p *Plugin) getFormReport(c *gin.Context, l string) (gin.H, error) {
	item := c.MustGet("item").(*Form)
	title := p.I18n.T(l, "buttons.report") + "-" + item.Title

	header, rows, err := p._exportForm(item)
	if err != nil {
		return nil, err
	}
	return gin.H{
		"title":  title,
		"header": header,
		"rows":   rows,
	}, nil
}
