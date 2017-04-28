package i18n

import (
	"math"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/language"
)

const (
	// LOCALE locale key
	LOCALE = "locale"
)

// Middleware locale-middleware
func (p *I18n) Middleware() (gin.HandlerFunc, error) {
	langs, err := p.Store.Languages()
	if err != nil {
		return nil, err
	}
	var tags []language.Tag
	for _, l := range langs {
		tags = append(tags, language.Make(l))
	}
	matcher := language.NewMatcher(tags)

	return func(c *gin.Context) {
		write := false
		// 1. Check URL arguments.
		lang := c.Request.URL.Query().Get(LOCALE)

		// 2. Get language information from cookies.
		if len(lang) == 0 {
			if ck, er := c.Request.Cookie(LOCALE); er == nil {
				lang = ck.Value
			} else {
				write = true
			}
		} else {
			write = true
		}

		// 3. Get language information from 'Accept-Language'.
		if len(lang) == 0 {
			write = true
			al := c.Request.Header.Get("Accept-Language")
			if len(al) > 4 {
				lang = al[:5] // Only compare first 5 letters.
			}
		}

		tag, _, _ := matcher.Match(language.Make(lang))
		if lang != tag.String() {
			write = true
		}

		if write {
			c.SetCookie(LOCALE, tag.String(), math.MaxInt32-1, "/", "", false, false)
		}
		c.Set(LOCALE, tag.String())
	}, nil
}
