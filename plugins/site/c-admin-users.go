package site

import (
	"github.com/gin-gonic/gin"
	"github.com/kapmahc/fly/plugins/auth"
)

func (p *Plugin) indexAdminUsers(c *gin.Context, l string) (gin.H, error) {
	var items []auth.User
	if err := p.Db.
		Order("last_sign_in_at DESC").Find(&items).Error; err != nil {
		return nil, err
	}

	return gin.H{
		"title": p.I18n.T(l, "site.admin.users.index.title"),
		"items": items,
	}, nil
}
