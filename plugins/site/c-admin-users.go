package site

import (
	"github.com/gin-gonic/gin"
	"github.com/kapmahc/fly/plugins/auth"
)

func (p *Plugin) getAdminUsers(c *gin.Context, l string) (interface{}, error) {
	var items []auth.User
	err := p.Db.
		Order("last_sign_in_at DESC").Find(&items).Error
	return items, err
}
