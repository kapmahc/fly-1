package web

import "net/http"

// Context context
type Context struct {
	Writer  http.ResponseWriter
	Request *http.Request
}
