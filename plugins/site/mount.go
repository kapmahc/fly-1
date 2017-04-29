package site

import "github.com/gin-gonic/gin"

// Mount mount web points
func (p *Plugin) Mount(rt *gin.Engine) {

	rt.GET("/", p.Wrap.HTML("site/home", p.getHome))
	rt.GET("/dashboard", p.Jwt.MustSignInMiddleware, p.Wrap.HTML("site/dashboard", p.getDashboard))
	rt.GET("/install", p.mustDatabaseEmpty, p.Wrap.HTML("form", p.getInstall))
	rt.POST("/install", p.mustDatabaseEmpty, p.Wrap.FORM(&fmInstall{}, p.postInstall))

	ag := rt.Group("/admin", p.Jwt.MustAdminMiddleware)

	ag.GET("/users", p.Wrap.HTML("site/admin/users/index", p.indexAdminUsers))

	ag.GET("/locales", p.Wrap.HTML("site/admin/locales/index", p.getAdminLocales))
	ag.GET("/locales/edit", p.Wrap.HTML("form", p.editAdminLocales))
	ag.POST("/locales", p.Wrap.FORM(&fmLocale{}, p.saveAdminLocales))
	ag.DELETE("/locales/:code", p.Wrap.JSON(p.deleteAdminLocales))

	asg := ag.Group("/site")
	asg.GET("/status", p.Wrap.HTML("site/admin/status", p.getAdminSiteStatus))
	asg.GET("/info", p.Wrap.HTML("form", p.getAdminSiteInfo))
	asg.POST("/info", p.Wrap.FORM(&fmSiteInfo{}, p.postAdminSiteInfo))
	asg.GET("/author", p.Wrap.HTML("form", p.getAdminSiteAuthor))
	asg.POST("/author", p.Wrap.FORM(&fmSiteAuthor{}, p.postAdminSiteAuthor))
	asg.GET("/seo", p.Wrap.HTML("site/admin/seo", p.getAdminSiteSeo))
	asg.POST("/seo", p.Wrap.FORM(&fmSiteSeo{}, p.postAdminSiteSeo))
	asg.GET("/smtp", p.Wrap.HTML("site/admin/smtp", p.getAdminSiteSMTP))
	asg.POST("/smtp", p.Wrap.FORM(&fmSiteSMTP{}, p.postAdminSiteSMTP))
}
