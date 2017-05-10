package web

import (
	"net/http"

	"github.com/kapmahc/fly/web/i18n"
)

// Context context
type Context struct {
	Lang    string
	Writer  http.ResponseWriter
	Request *http.Request

	I18n   *i18n.I18n
	Values map[string]interface{}
}

// E create http error
func (p *Context) E(code int, format string, args ...interface{}) error {
	return &HTTPError{
		Message: p.T(format, args...),
		Code:    code,
	}
}

// T translate
func (p *Context) T(format string, args ...interface{}) string {
	return p.I18n.T(p.Lang, format, args...)
}

// Set set
func (p *Context) Set(k string, v interface{}) {
	p.Values[k] = v
}

// Get get
func (p *Context) Get(k string) (interface{}, bool) {
	val, ok := p.Values[k]
	return val, ok
}
