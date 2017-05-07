package widgets

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/csrf"
)

// NewForm new form
func NewForm(req *http.Request, lang, action, next, title string, fields ...interface{}) Form {
	if next == "" {
		next = action
	}
	return Form{
		"id":             "form-" + uuid.New().String(),
		"lang":           lang,
		"method":         http.MethodPost,
		"action":         action,
		"next":           next,
		"title":          title,
		"fields":         fields,
		csrf.TemplateTag: csrf.TemplateField(req),
	}
}

// Form form
type Form map[string]interface{}

// Append  append fields
func (p Form) Append(fields ...Field) {
	items := p["fields"].([]interface{})
	for _, f := range fields {
		items = append(items, f)
	}
	p["fields"] = items
}

// Method set method
func (p Form) Method(m string) {
	p["method"] = m
}

// Field field
type Field map[string]interface{}

// Readonly set readonly
func (p Field) Readonly() {
	p["readonly"] = true
}

// Helper set helper message
func (p Field) Helper(h string) {
	p["helper"] = h
}

// NewTextarea new textarea field
func NewTextarea(id, label, value string, rows int) Field {
	return Field{
		"id":    id,
		"type":  "textarea",
		"label": label,
		"value": value,
		"rows":  rows,
	}
}

// NewTextField new text field
func NewTextField(id, label, value string) Field {
	return Field{
		"id":    id,
		"type":  "text",
		"label": label,
		"value": value,
	}
}

// NewHiddenField new hidden field
func NewHiddenField(id string, value interface{}) Field {
	return Field{
		"id":    id,
		"type":  "hidden",
		"value": value,
	}
}

// NewEmailField new email field
func NewEmailField(id, label, value string) Field {
	return Field{
		"id":    id,
		"type":  "email",
		"label": label,
		"value": value,
	}
}

// NewPasswordField new password field
func NewPasswordField(id, label, helper string) Field {
	return Field{
		"id":     id,
		"type":   "password",
		"label":  label,
		"helper": helper,
	}
}

// NewCheckboxs new checkboxs field
func NewCheckboxs(id, label string, value []interface{}, options ...Option) Field {
	return Field{
		"id":      id,
		"type":    TypeCheckboxs,
		"label":   label,
		"value":   value,
		"options": options,
	}
}

const (
	// TypeRadios radios
	TypeRadios = "radios"
	// TypeCheckboxs checkboxs
	TypeCheckboxs = "checkboxs"
)

// NewRadios new radios filed
func NewRadios(id, label string, value interface{}, options ...Option) Field {
	return Field{
		"id":      id,
		"type":    TypeRadios,
		"label":   label,
		"value":   value,
		"options": options,
	}
}

// NewCheckbox new checkbox field
func NewCheckbox(id, label string, checked bool) Field {
	return Field{
		"id":      id,
		"type":    "checkbox",
		"label":   label,
		"checked": checked,
	}
}

// NewSelect new select field
func NewSelect(id, label string, value interface{}, options ...Option) Field {
	return Field{
		"id":       id,
		"type":     "select",
		"multiple": false,
		"label":    label,
		"value":    value,
		"options":  options,
	}
}

// Option select option
type Option gin.H

// NewOption new option
func NewOption(label string, value interface{}) Option {
	return Option{"value": value, "label": label}
}

// NewSortSelect new sort select
func NewSortSelect(id, label string, value, min, max int) Field {
	var options []Option
	for i := min; i <= max; i++ {
		options = append(options, NewOption(strconv.Itoa(i), i))
	}
	return NewSelect(id, label, value, options...)
}
