package site

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kapmahc/fly/plugins/auth"
	"github.com/kapmahc/fly/web/i18n"
)

func (p *Plugin) getInstall(c *gin.Context, l string) (gin.H, error) {
	return gin.H{}, nil
}

type fmInstall struct {
	Title                string `form:"title" binding:"required"`
	SubTitle             string `form:"subTitle" binding:"required"`
	Email                string `form:"email" binding:"required,email"`
	Password             string `form:"password" binding:"min=6,max=32"`
	PasswordConfirmation string `form:"passwordConfirmation" binding:"eqfield=Password"`
}

func (p *Plugin) postInstall(c *gin.Context, lang string, o interface{}) (interface{}, error) {
	fm := o.(*fmInstall)
	p.I18n.Set(lang, "site.title", fm.Title)
	p.I18n.Set(lang, "site.subTitle", fm.SubTitle)
	user, err := p.Dao.AddEmailUser("root", fm.Email, fm.Password)
	if err != nil {
		return nil, err
	}
	for _, r := range []string{auth.RoleAdmin, auth.RoleRoot} {
		role, er := p.Dao.Role(r, auth.DefaultResourceType, auth.DefaultResourceID)
		if er != nil {
			return nil, er
		}
		if err = p.Dao.Allow(role.ID, user.ID, 50, 0, 0); err != nil {
			return nil, err
		}
	}
	if err = p.Db.Model(user).UpdateColumn("confirmed_at", time.Now()).Error; err != nil {
		return nil, err
	}
	return gin.H{}, nil
}

func (p *Plugin) mustDatabaseEmpty(c *gin.Context) {
	lang := c.MustGet(i18n.LOCALE).(string)
	var count int
	if err := p.Db.Model(&auth.User{}).Count(&count).Error; err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	if count > 0 {
		c.String(http.StatusForbidden, p.I18n.T(lang, "errors.forbidden"))
		return
	}
	c.Next()
}
