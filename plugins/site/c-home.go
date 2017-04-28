package site

import "github.com/gin-gonic/gin"

func (p *Plugin) getDashboard(c *gin.Context, l string) (gin.H, error) {
	return gin.H{"title": p.I18n.T(l, "header.dashboard")}, nil
}

func (p *Plugin) getHome(c *gin.Context, l string) (gin.H, error) {
	return gin.H{"title": p.I18n.T(l, "header.home")}, nil
}
