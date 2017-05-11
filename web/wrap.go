package web

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/go-playground/form"
	"github.com/urfave/negroni"
	validator "gopkg.in/go-playground/validator.v9"
)

// Wrapper wrapper
type Wrapper struct {
	Validate *validator.Validate `inject:""`
	Decoder  *form.Decoder       `inject:""`
}

// Wrap wrap
func (p *Wrapper) Wrap(hnd http.HandlerFunc, mids ...func(http.ResponseWriter, *http.Request, http.HandlerFunc)) http.Handler {
	var items []negroni.Handler
	for _, m := range mids {
		items = append(items, negroni.HandlerFunc(m))
	}
	items = append(items, negroni.Wrap(hnd))
	return negroni.New(items...)
}

// Form form handler
func (p *Wrapper) Form(o interface{}, f func(c *Context, o interface{}) (interface{}, error)) http.HandlerFunc {
	return p.JSON(func(c *Context) (interface{}, error) {
		if e := c.Request.ParseForm(); e != nil {
			return nil, e
		}
		if e := p.Decoder.Decode(o, c.Request.Form); e != nil {
			return nil, e
		}
		if e := p.Validate.Struct(o); e != nil {
			return nil, e
		}
		return f(c, o)
	})
}

// Handle handle
func (p *Wrapper) Handle(f func(c *Context) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(&Context{
			Request: r,
			Writer:  w,
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
