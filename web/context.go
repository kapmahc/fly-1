package web

import (
	"net"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

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

// Redirect redirect
func (p *Context) Redirect(code int, url string) {
	http.Redirect(p.Writer, p.Request, url, code)
}

// Header get header
func (p *Context) Header(k string) string {
	return p.Request.Header.Get(k)
}

// Param the value of the URL param.
func (p *Context) Param(k string) string {
	return mux.Vars(p.Request)[k]
}

// ClientIP client ip
func (p *Context) ClientIP() string {
	// -------------
	if ip := strings.TrimSpace(p.Header("X-Real-Ip")); ip != "" {
		return ip
	}
	// -------------
	ip := p.Header("X-Forwarded-For")
	if idx := strings.IndexByte(ip, ','); idx >= 0 {
		ip = ip[0:idx]
	}
	ip = strings.TrimSpace(ip)
	if ip != "" {
		return ip
	}
	// -------------
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(p.Request.RemoteAddr)); err == nil {
		return ip
	}
	// -----------
	return ""
}
