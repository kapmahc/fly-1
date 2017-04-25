package root

import (
	"strconv"

	"github.com/astaxie/beego"
	"github.com/beego/i18n"
	"golang.org/x/text/language"
)

// Layout layout
type Layout struct {
	beego.Controller
	Lang string
}

// Abort abort
func (p *Layout) Abort(code int) {
	p.Controller.Abort(strconv.Itoa(code))
}

func (p *Layout) setLangVer() {
	const key = "locale"
	write := false

	// 1. Check URL arguments.
	lang := p.Input().Get(key)

	// 2. Get language information from cookies.
	if len(lang) == 0 {
		lang = p.Ctx.GetCookie(key)
	} else {
		write = true
	}

	// 3. Get language information from 'Accept-Language'.
	if len(lang) == 0 {
		write = true

		al := p.Ctx.Request.Header.Get("Accept-Language")
		if len(al) > 4 {
			lang = al[:5] // Only compare first 5 letters.
		}
	}

	// 4. Default language is English.
	if len(lang) == 0 || !i18n.IsExist(lang) {
		lang = language.AmericanEnglish.String()
	}

	// Save language information in cookies.
	if write {
		p.Ctx.SetCookie(key, lang, 1<<31-1, "/")
	}

	// Set language properties.
	p.Lang = lang
	p.Data["lang"] = lang
	p.Data["languages"] = i18n.ListLangs()
}
