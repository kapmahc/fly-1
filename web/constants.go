package web

import (
	"io"
	"path"
	"text/template"
)

const (
	// FormatDateInput  format for date input
	FormatDateInput = "2006-01-02"
)

// Template parse template
func Template(w io.Writer, n string, d interface{}) error {
	t, e := template.ParseFiles(path.Join("templates", n))
	if e != nil {
		return e
	}
	return t.Execute(w, d)
}
