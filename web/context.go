package web

import "net/http"

// K key type
type K string

// H hash
type H map[string]interface{}

// Context context
type Context struct {
	Writer  http.ResponseWriter
	Request *http.Request
}

// Get get
func (p *Context) Get(k string) interface{} {
	return p.Request.Context().Value(K(k))
}
