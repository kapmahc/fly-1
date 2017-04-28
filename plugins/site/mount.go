package site

import "github.com/gin-gonic/gin"

// Mount mount web points
func (p *Plugin) Mount(rt *gin.Engine) {
	rt.GET("/install", p.mustDatabaseEmpty, p.Wrap.HTML("form", p.getInstall))
	rt.POST("/install", p.mustDatabaseEmpty, p.Wrap.FORM(&fmInstall{}, p.postInstall))
}
