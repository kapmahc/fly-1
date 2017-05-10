package web

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
)

// Handle handle
func Handle(f func(*Context) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if e := f(&Context{Writer: w, Request: r}); e != nil {
			c := http.StatusInternalServerError
			if he, ok := e.(*HTTPError); ok {
				c = he.Code
			}
			http.Error(w, e.Error(), c)
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
