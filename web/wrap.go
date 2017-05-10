package web

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/kapmahc/fly/web/i18n"
)

// Wrapper wrapper
type Wrapper struct {
	I18n *i18n.I18n `inject:""`
}

// Handle handle
func (p *Wrapper) Handle(f func(*Context) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(&Context{
			Writer:  w,
			Request: r,
			Lang:    p.I18n.Lang(r).String(),
			I18n:    p.I18n,
		}); err != nil {
			code := http.StatusInternalServerError
			if her, ok := err.(*HTTPError); ok {
				code = her.Code
			}
			http.Error(w, err.Error(), code)
		}
	}
}

// JSON wrap JSON
func (p *Wrapper) JSON(f func(*Context) (interface{}, error)) http.HandlerFunc {
	return p.Handle(func(c *Context) error {
		v, e := f(c)
		if e != nil {
			return e
		}
		enc := json.NewEncoder(c.Writer)
		return enc.Encode(v)
	})
}

// XML wrap XML
func (p *Wrapper) XML(f func(*Context) (interface{}, error)) http.HandlerFunc {
	return p.Handle(func(c *Context) error {
		v, e := f(c)
		if e != nil {
			return e
		}
		enc := xml.NewEncoder(c.Writer)
		return enc.Encode(v)
	})
}
