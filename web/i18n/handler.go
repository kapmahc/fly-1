package i18n

import (
	"net/http"

	"golang.org/x/text/language"
)

// Detect detect language from http request
func Detect(r *http.Request) string {
	const key = "locale"

	// 1. Check URL arguments.
	if lang := r.URL.Query().Get(key); lang != "" {
		return lang
	}

	// 2. Get language information from cookies.
	if ck, er := r.Cookie(key); er == nil {
		return ck.Value
	}

	// 3. Get language information from 'Accept-Language'.
	if al := r.Header.Get("Accept-Language"); len(al) > 4 {
		return al[:5] // Only compare first 5 letters.
	}

	return language.AmericanEnglish.String()
}
