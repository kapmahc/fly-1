package i18n

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"golang.org/x/text/language"
)

// Lang detect language from http request
func (p *I18n) Lang(r *http.Request) language.Tag {
	lang := p.detect(r)
	if lang == "" {
		return language.AmericanEnglish
	}
	langs, err := p.Store.Languages()
	if err != nil {
		log.Error(err)
		return language.AmericanEnglish
	}
	var tags []language.Tag
	for _, l := range langs {
		tags = append(tags, language.Make(l))
	}
	matcher := language.NewMatcher(tags)
	tag, _, _ := matcher.Match(language.Make(lang))
	return tag
}

func (p *I18n) detect(r *http.Request) string {
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

	return ""
}
