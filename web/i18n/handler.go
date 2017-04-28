package i18n

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/text/language"
)

const (
	// LOCALE locale key
	LOCALE = "locale"
)

// Middleware locale-middleware
func (p *I18n) Middleware(c *gin.Context) {
	// 1. Check URL arguments.
	lang := c.Request.URL.Query().Get(LOCALE)

	// 2. Get language information from cookies.
	if len(lang) == 0 {
		if ck, er := c.Request.Cookie(LOCALE); er == nil {
			lang = ck.Value
		}
	}

	// 3. Get language information from 'Accept-Language'.
	if len(lang) == 0 {
		al := c.Request.Header.Get("Accept-Language")
		if len(al) > 4 {
			lang = al[:5] // Only compare first 5 letters.
		}
	}

	tag, _, _ := p.Matcher.Match(language.Make(lang))
	c.Set(LOCALE, tag.String())
}
