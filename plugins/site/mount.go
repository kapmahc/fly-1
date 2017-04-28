package site

import "github.com/gin-gonic/gin"

// Mount mount web points
func (p *Plugin) Mount(rt *gin.Engine) {
	rt.GET("/install", p.Wrap.HTML("site/install", p.getInstall))
	rt.POST("/install", p.Wrap.FORM(&fmInstall{}, p.postInstall))
}
