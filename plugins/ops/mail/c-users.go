package mail

import (
	"net/http"

	"github.com/kapmahc/h2o"
	"github.com/kapmahc/h2o/i18n"
)

func (p *Plugin) indexUsers(c *h2o.Context) error {
	var items []User
	if err := p.Db.Order("updated_at DESC").Find(&items).Error; err != nil {
		return err
	}
	var domains []Domain
	if err := p.Db.Select([]string{"id", "name"}).Find(&domains).Error; err != nil {
		return err
	}
	for i := range items {
		u := &items[i]
		for _, d := range domains {
			if d.ID == u.DomainID {
				u.Domain = d
				break
			}
		}
	}
	return c.JSON(http.StatusOK, items)
}

type fmUserNew struct {
	FullName             string `form:"fullName" validate:"required,max=255"`
	Email                string `form:"email" validate:"required,email"`
	Password             string `form:"password" validate:"min=6,max=32"`
	PasswordConfirmation string `form:"passwordConfirmation" validate:"eqfield=Password"`
	Enable               bool   `form:"enable"`
	DomainID             uint   `form:"domainId"`
}

func (p *Plugin) createUser(c *h2o.Context) error {

	var fm fmUserNew
	if err := c.Bind(&fm); err != nil {
		return err
	}
	user := User{
		FullName: fm.FullName,
		Email:    fm.Email,
		Enable:   fm.Enable,
		DomainID: fm.DomainID,
	}
	if err := user.SetPassword(fm.Password); err != nil {
		return err
	}
	if err := p.Db.Create(&user).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}

func (p *Plugin) showUser(c *h2o.Context) error {
	var item User
	if err := p.Db.Where("id = ?", c.Param("id")).First(&item).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, item)
}

type fmUserEdit struct {
	FullName string `form:"fullName" validate:"required,max=255"`
	Enable   bool   `form:"enable"`
}

func (p *Plugin) updateUser(c *h2o.Context) error {

	var fm fmUserEdit
	if err := c.Bind(&fm); err != nil {
		return err
	}

	var item User
	if err := p.Db.Where("id = ?", c.Param("id")).First(&item).Error; err != nil {
		return err
	}

	if err := p.Db.Model(&item).
		Updates(map[string]interface{}{
			"enable":    fm.Enable,
			"full_name": fm.FullName,
		}).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, h2o.H{})
}

type fmUserResetPassword struct {
	Password             string `form:"password" validate:"min=6,max=32"`
	PasswordConfirmation string `form:"passwordConfirmation" validate:"eqfield=Password"`
}

func (p *Plugin) postResetUserPassword(c *h2o.Context) error {

	var fm fmUserResetPassword
	if err := c.Bind(&fm); err != nil {
		return err
	}

	var item User
	if err := p.Db.Where("id = ?", c.Param("id")).First(&item).Error; err != nil {
		return err
	}

	if err := item.SetPassword(fm.Password); err != nil {
		return err
	}
	if err := p.Db.Model(&item).
		Updates(map[string]interface{}{
			"password": item.Password,
		}).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, h2o.H{})
}

type fmUserChangePassword struct {
	Email                string `form:"email" validate:"required,email"`
	CurrentPassword      string `form:"currentPassword" validate:"required"`
	NewPassword          string `form:"newPassword" validate:"min=6,max=32"`
	PasswordConfirmation string `form:"passwordConfirmation" validate:"eqfield=NewPassword"`
}

func (p *Plugin) postChangeUserPassword(c *h2o.Context) error {
	lng := c.Get(i18n.LOCALE).(string)
	var fm fmUserChangePassword
	if err := c.Bind(&fm); err != nil {
		return err
	}
	var user User
	if err := p.Db.Where("email = ?", fm.Email).First(&user).Error; err != nil {
		return err
	}
	if !user.ChkPassword(fm.CurrentPassword) {
		return p.I18n.E(http.StatusBadRequest, lng, "ops.mail.users.email-password-not-match")
	}
	if err := user.SetPassword(fm.NewPassword); err != nil {
		return err
	}

	if err := p.Db.Model(user).
		Updates(map[string]interface{}{
			"password": user.Password,
		}).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, h2o.H{})
}

func (p *Plugin) destroyUser(c *h2o.Context) error {
	lng := c.Get(i18n.LOCALE).(string)
	var user User
	if err := p.Db.
		Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		return err
	}
	var count int
	if err := p.Db.Model(&Alias{}).Where("destination = ?", user.Email).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return p.I18n.E(http.StatusForbidden, lng, "errors.in-use")
	}
	if err := p.Db.Delete(&user).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, h2o.H{})
}
