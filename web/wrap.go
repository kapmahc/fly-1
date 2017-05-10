package web

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
)

// Handle handle
func Handle(f func(c *Context) error) http.HandlerFunc {
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
func JSON(f func(*Context) (interface{}, error)) http.HandlerFunc {
	return Handle(func(c *Context) error {
		v, e := f(c)
		if e != nil {
			return e
		}
		enc := json.NewEncoder(c.Writer)
		return enc.Encode(v)
	})
}

// XML wrap XML
func XML(f func(*Context) (interface{}, error)) http.HandlerFunc {
	return Handle(func(c *Context) error {
		v, e := f(c)
		if e != nil {
			return e
		}
		enc := xml.NewEncoder(c.Writer)
		return enc.Encode(v)
	})
}
