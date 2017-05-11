package auth

import (
	"net/http"
	"time"

	"github.com/SermoDigital/jose/jws"
	"github.com/kapmahc/fly/web"
	"github.com/kapmahc/fly/web/i18n"
)

type fmSignUp struct {
	Name                 string `form:"name" binding:"required,max=255"`
	Email                string `form:"email" binding:"required,email"`
	Password             string `form:"password" binding:"min=6,max=32"`
	PasswordConfirmation string `form:"passwordConfirmation" binding:"eqfield=Password"`
}

func (p *Plugin) postUsersSignUp(c *web.Context, o interface{}) (interface{}, error) {
	l := c.Get(i18n.LOCALE).(string)
	fm := o.(*fmSignUp)

	var count int
	if err := p.Db.
		Model(&User{}).
		Where("email = ?", fm.Email).
		Count(&count).Error; err != nil {
		return nil, err
	}

	if count > 0 {
		return nil, p.I18n.E(http.StatusInternalServerError, l, "auth.errors.email-already-exists")
	}

	user, err := p.Dao.AddEmailUser(fm.Name, fm.Email, fm.Password)
	if err != nil {
		return nil, err
	}

	p.Dao.Log(user.ID, c.ClientIP(), p.I18n.T(l, "auth.logs.sign-up"))
	p.sendEmail(l, user, actConfirm)

	return web.H{"message": p.I18n.T(l, "auth.messages.email-for-confirm")}, nil
}

type fmSignIn struct {
	Email      string `form:"email" binding:"required,email"`
	Password   string `form:"password" binding:"required"`
	RememberMe bool   `form:"rememberMe"`
}

func (p *Plugin) postUsersSignIn(c *web.Context, o interface{}) (interface{}, error) {
	l := c.Get(i18n.LOCALE).(string)
	fm := o.(*fmSignIn)

	user, err := p.Dao.SignIn(l, fm.Email, fm.Password, c.ClientIP())
	if err != nil {
		return nil, err
	}

	cm := jws.Claims{}
	cm.Set(UID, user.UID)
	tkn, err := p.Jwt.Sum(cm, time.Hour*24*7)
	if err != nil {
		return nil, err
	}
	return web.H{
		"token": string(tkn),
	}, nil
}

type fmEmail struct {
	Email string `form:"email" binding:"required,email"`
}

func (p *Plugin) getUsersConfirm(c *web.Context) error {
	l := c.Get(i18n.LOCALE).(string)
	token := c.Param("token")
	user, err := p.parseToken(l, token, actConfirm)
	if err != nil {
		return err
	}
	if user.IsConfirm() {
		return p.I18n.E(http.StatusForbidden, l, "auth.errors.user-already-confirm")
	}
	p.Db.Model(user).Update("confirmed_at", time.Now())
	p.Dao.Log(user.ID, c.ClientIP(), p.I18n.T(l, "auth.logs.confirm"))

	c.Redirect(http.StatusFound, p._signInURL())
	return nil
}

func (p *Plugin) postUsersConfirm(c *web.Context, o interface{}) (interface{}, error) {
	l := c.Get(i18n.LOCALE).(string)
	fm := o.(*fmEmail)
	user, err := p.Dao.GetByEmail(fm.Email)
	if err != nil {
		return nil, err
	}

	if user.IsConfirm() {
		return nil, p.I18n.E(http.StatusForbidden, l, "auth.errors.user-already-confirm")
	}

	p.sendEmail(l, user, actConfirm)

	return web.H{"message": p.I18n.T(l, "auth.messages.email-for-confirm")}, nil
}

func (p *Plugin) getUsersUnlock(c *web.Context) error {
	l := c.Get(i18n.LOCALE).(string)
	token := c.Param("token")
	user, err := p.parseToken(l, token, actUnlock)
	if err != nil {
		return err
	}
	if !user.IsLock() {
		return p.I18n.E(http.StatusForbidden, l, "auth.errors.user-not-lock")
	}

	p.Db.Model(user).Update(map[string]interface{}{"locked_at": nil})
	p.Dao.Log(user.ID, c.ClientIP(), p.I18n.T(l, "auth.logs.unlock"))

	c.Redirect(http.StatusFound, p._signInURL())
	return nil
}

func (p *Plugin) postUsersUnlock(c *web.Context, o interface{}) (interface{}, error) {
	l := c.Get(i18n.LOCALE).(string)

	fm := o.(*fmEmail)
	user, err := p.Dao.GetByEmail(fm.Email)
	if err != nil {
		return nil, err
	}
	if !user.IsLock() {
		return nil, p.I18n.E(http.StatusForbidden, l, "auth.errors.user-not-lock")
	}
	p.sendEmail(l, user, actUnlock)
	return web.H{"message": p.I18n.T(l, "auth.messages.email-for-unlock")}, nil
}

func (p *Plugin) postUsersForgotPassword(c *web.Context, o interface{}) (interface{}, error) {
	l := c.Get(i18n.LOCALE).(string)

	fm := o.(*fmEmail)
	var user *User
	user, err := p.Dao.GetByEmail(fm.Email)
	if err != nil {
		return nil, err
	}
	p.sendEmail(l, user, actResetPassword)

	return web.H{"message": p.I18n.T(l, "auth.messages.email-for-reset-password")}, nil
}

type fmResetPassword struct {
	Token                string `form:"token" binding:"required"`
	Password             string `form:"password" binding:"min=6,max=32"`
	PasswordConfirmation string `form:"passwordConfirmation" binding:"eqfield=Password"`
}

func (p *Plugin) postUsersResetPassword(c *web.Context, o interface{}) (interface{}, error) {
	l := c.Get(i18n.LOCALE).(string)

	fm := o.(*fmResetPassword)
	user, err := p.parseToken(l, fm.Token, actResetPassword)
	if err != nil {
		return nil, err
	}
	p.Db.Model(user).Update("password", p.Hmac.Sum([]byte(fm.Password)))
	p.Dao.Log(user.ID, c.ClientIP(), p.I18n.T(l, "auth.logs.reset-password"))
	return web.H{"message": p.I18n.T(l, "auth.messages.reset-password-success")}, nil
}

func (p *Plugin) deleteUsersSignOut(c *web.Context) (interface{}, error) {
	l := c.Get(i18n.LOCALE).(string)
	user := c.Get(CurrentUser).(*User)
	p.Dao.Log(user.ID, c.ClientIP(), p.I18n.T(l, "auth.logs.sign-out"))
	return web.H{}, nil
}

func (p *Plugin) getUsersInfo(c *web.Context) (interface{}, error) {
	user := c.Get(CurrentUser).(*User)
	return web.H{"name": user.Name, "email": user.Email}, nil
}

type fmInfo struct {
	Name string `form:"name" binding:"required,max=255"`
	// Home string `form:"home" binding:"max=255"`
	// Logo string `form:"logo" binding:"max=255"`
}

func (p *Plugin) postUsersInfo(c *web.Context, o interface{}) (interface{}, error) {
	user := c.Get(CurrentUser).(*User)
	fm := o.(*fmInfo)

	if err := p.Db.Model(user).Updates(map[string]interface{}{
		// "home": fm.Home,
		// "logo": fm.Logo,
		"name": fm.Name,
	}).Error; err != nil {
		return nil, err
	}
	return web.H{}, nil
}

type fmChangePassword struct {
	CurrentPassword      string `form:"currentPassword" binding:"required"`
	NewPassword          string `form:"newPassword" binding:"min=6,max=32"`
	PasswordConfirmation string `form:"passwordConfirmation" binding:"eqfield=NewPassword"`
}

func (p *Plugin) postUsersChangePassword(c *web.Context, o interface{}) (interface{}, error) {
	l := c.Get(i18n.LOCALE).(string)

	user := c.Get(CurrentUser).(*User)
	fm := o.(*fmChangePassword)
	if !p.Hmac.Chk([]byte(fm.CurrentPassword), user.Password) {
		return nil, p.I18n.E(http.StatusForbidden, l, "auth.errors.bad-password")
	}
	if err := p.Db.Model(user).
		Update("password", p.Hmac.Sum([]byte(fm.NewPassword))).Error; err != nil {
		return nil, err
	}

	return web.H{}, nil
}

func (p *Plugin) getUsersLogs(c *web.Context) (interface{}, error) {
	user := c.Get(CurrentUser).(*User)
	var logs []Log
	err := p.Db.
		Select([]string{"ip", "message", "created_at"}).
		Where("user_id = ?", user.ID).
		Order("id DESC").Limit(120).
		Find(&logs).Error
	return logs, err
}

func (p *Plugin) indexUsers(c *web.Context) (interface{}, error) {
	var users []User
	err := p.Db.
		Select([]string{"name", "logo", "home"}).
		Order("last_sign_in_at DESC").
		Find(&users).Error
	return users, err
}

func (p *Plugin) _signInURL() string {
	return web.Frontend() + "/users/sign-in"
}
