package site

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kapmahc/fly/web/widgets"
)

func (p *Plugin) getAdminSiteInfo(c *gin.Context, lang string) (gin.H, error) {
	title := p.I18n.T(lang, "site.admin.info.title")
	fm := widgets.NewForm(
		c.Request,
		lang,
		"/admin/site/info",
		"",
		title,
		widgets.NewTextField("title", p.I18n.T(lang, "site.attributes.title"), p.I18n.T(lang, "site.title")),
		widgets.NewTextField("subTitle", p.I18n.T(lang, "site.attributes.subTitle"), p.I18n.T(lang, "site.subTitle")),
		widgets.NewTextField("keywords", p.I18n.T(lang, "site.attributes.keywords"), p.I18n.T(lang, "site.keywords")),
		widgets.NewTextarea("description", p.I18n.T(lang, "site.attributes.description"), p.I18n.T(lang, "site.description"), 6),
		widgets.NewTextField("copyright", p.I18n.T(lang, "site.attributes.copyright"), p.I18n.T(lang, "site.copyright")),
	)
	return gin.H{"form": fm, "title": title}, nil
}

type fmSiteInfo struct {
	Title       string `form:"title"`
	SubTitle    string `form:"subTitle"`
	Keywords    string `form:"keywords"`
	Description string `form:"description"`
	Copyright   string `form:"copyright"`
}

func (p *Plugin) postAdminSiteInfo(c *gin.Context, l string, o interface{}) (interface{}, error) {
	fm := o.(*fmSiteInfo)

	for k, v := range map[string]string{
		"title":       fm.Title,
		"subTitle":    fm.SubTitle,
		"keywords":    fm.Keywords,
		"description": fm.Description,
		"copyright":   fm.Copyright,
	} {
		if err := p.I18n.Set(l, "site."+k, v); err != nil {
			return nil, err
		}
	}

	return gin.H{}, nil
}

func (p *Plugin) getAdminSiteAuthor(c *gin.Context, lang string) (gin.H, error) {
	title := p.I18n.T(lang, "site.admin.author.title")
	fm := widgets.NewForm(
		c.Request,
		lang,
		"/admin/site/author",
		"",
		title,
		widgets.NewTextField("name", p.I18n.T(lang, "attributes.fullName"), p.I18n.T(lang, "site.author.name")),
		widgets.NewEmailField("email", p.I18n.T(lang, "attributes.email"), p.I18n.T(lang, "site.author.email")),
	)
	return gin.H{"form": fm, "title": title}, nil
}

type fmSiteAuthor struct {
	Name  string `form:"name"`
	Email string `form:"email"`
}

func (p *Plugin) postAdminSiteAuthor(c *gin.Context, l string, o interface{}) (interface{}, error) {
	fm := o.(*fmSiteAuthor)
	for k, v := range map[string]string{
		"name":  fm.Name,
		"email": fm.Email,
	} {
		if err := p.I18n.Set(l, "site.author."+k, v); err != nil {
			return nil, err
		}
	}

	return gin.H{}, nil
}

func (p *Plugin) getAdminSiteSeo(c *gin.Context, lang string) (gin.H, error) {
	var gc string
	var bc string
	p.Settings.Get("site.google.verify.code", &gc)
	p.Settings.Get("site.baidu.verify.code", &bc)

	title := p.I18n.T(lang, "site.admin.seo.title")
	fm := widgets.NewForm(
		c.Request,
		lang,
		"/admin/site/seo",
		"",
		title,
		widgets.NewTextField("googleVerifyCode", p.I18n.T(lang, "site.admin.seo.googleVerifyCode"), gc),
		widgets.NewTextField("baiduVerifyCode", p.I18n.T(lang, "site.admin.seo.baiduVerifyCode"), bc),
	)

	links := []string{"robots.txt", "sitemap.xml.gz", "google" + gc + ".html", "baidu_verify_" + bc + ".html"}
	langs, err := p.I18n.Store.Languages()
	if err != nil {
		return nil, err
	}
	for _, l := range langs {
		links = append(links, "rss-"+l+".atom")
	}

	return gin.H{"form": fm, "title": title, "links": links}, nil
}

type fmSiteSeo struct {
	GoogleVerifyCode string `form:"googleVerifyCode"`
	BaiduVerifyCode  string `form:"baiduVerifyCode"`
}

func (p *Plugin) postAdminSiteSeo(c *gin.Context, l string, o interface{}) (interface{}, error) {
	fm := o.(*fmSiteSeo)

	for k, v := range map[string]string{
		"google.verify.code": fm.GoogleVerifyCode,
		"baidu.verify.code":  fm.BaiduVerifyCode,
	} {
		if err := p.Settings.Set("site."+k, v, true); err != nil {
			return nil, err
		}
	}
	return gin.H{}, nil
}

type fmSiteSMTP struct {
	Host                 string `form:"host"`
	Port                 int    `form:"port"`
	Ssl                  string `form:"ssl"`
	Username             string `form:"username"`
	Password             string `form:"password" binding:"min=6,max=32"`
	PasswordConfirmation string `form:"passwordConfirmation" binding:"eqfield=Password"`
}

func (p *Plugin) getAdminSiteSMTP(c *gin.Context, lang string) (gin.H, error) {
	smtp := make(map[string]interface{})
	if err := p.Settings.Get("site.smtp", &smtp); err == nil {
		smtp["password"] = ""
	} else {
		smtp["host"] = "localhost"
		smtp["port"] = 25
		smtp["ssl"] = false
		smtp["username"] = "no-reply@change-me.com"
		smtp["password"] = ""
	}

	var options []widgets.Option
	for _, v := range []int{25, 465, 587} {
		options = append(options, widgets.NewOption(strconv.Itoa(v), v))
	}

	title := p.I18n.T(lang, "site.admin.smtp.title")
	fm := widgets.NewForm(
		c.Request,
		lang,
		"/admin/site/smtp",
		"",
		title,
		widgets.NewTextField("host", p.I18n.T(lang, "attributes.host"), smtp["host"].(string)),
		widgets.NewSelect("port", p.I18n.T(lang, "attributes.port"), smtp["port"].(int), options...),
		widgets.NewEmailField("username", p.I18n.T(lang, "site.admin.smtp.sender"), smtp["username"].(string)),
		widgets.NewPasswordField("password", p.I18n.T(lang, "attributes.password"), p.I18n.T(lang, "helpers.password")),
		widgets.NewPasswordField("passwordConfirmation", p.I18n.T(lang, "attributes.passwordConfirmation"), p.I18n.T(lang, "helpers.passwordConfirmation")),
		widgets.NewCheckbox("ssl", p.I18n.T(lang, "attributes.ssl"), smtp["ssl"].(bool)),
	)
	return gin.H{"form": fm, "title": title}, nil
}

func (p *Plugin) postAdminSiteSMTP(c *gin.Context, l string, o interface{}) (interface{}, error) {
	fm := o.(*fmSiteSMTP)
	val := map[string]interface{}{
		"host":     fm.Host,
		"port":     fm.Port,
		"username": fm.Username,
		"password": fm.Password,
		"ssl":      fm.Ssl == "on",
	}
	err := p.Settings.Set("site.smtp", val, true)
	return gin.H{}, err
}
