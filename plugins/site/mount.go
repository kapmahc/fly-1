package site

import "github.com/gin-gonic/gin"

// Mount mount web points
func (p *Plugin) Mount(rt *gin.Engine) {

	rt.GET("/", p.Wrap.HTML("site/home", p.getHome))
	rt.GET("/dashboard", p.Jwt.MustSignInMiddleware, p.Wrap.HTML("site/dashboard", p.getDashboard))
	rt.GET("/install", p.mustDatabaseEmpty, p.Wrap.HTML("form", p.getInstall))
	rt.POST("/install", p.mustDatabaseEmpty, p.Wrap.FORM(&fmInstall{}, p.postInstall))

	ag := rt.Group("/admin", p.Jwt.MustAdminMiddleware)

	asg := ag.Group("/site")
	asg.GET("/status", p.Wrap.HTML("site/admin/status", p.getAdminSiteStatus))
}
