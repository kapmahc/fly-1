package auth

import (
	"time"

	"github.com/SermoDigital/jose/jws"
	"github.com/gin-gonic/gin"
	"github.com/kapmahc/fly/web"
	"github.com/kapmahc/fly/web/widgets"
)

func (p *Plugin) getUsersSignUp(c *gin.Context, lang string) (gin.H, error) {
	title := p.I18n.T(lang, "auth.users.sign-up.title")
	fm := widgets.NewForm(
		c.Request,
		lang,
		"/users/sign-up",
		"/users/sign-in",
		title,
		widgets.NewTextField("name", p.I18n.T(lang, "attributes.fullName"), ""),
		widgets.NewEmailField("email", p.I18n.T(lang, "attributes.email"), ""),
		widgets.NewPasswordField("password", p.I18n.T(lang, "attributes.password"), p.I18n.T(lang, "helpers.password")),
		widgets.NewPasswordField("passwordConfirmation", p.I18n.T(lang, "attributes.passwordConfirmation"), p.I18n.T(lang, "helpers.passwordConfirmation")),
	)
	return gin.H{"form": fm, "title": title}, nil
}

type fmSignUp struct {
	Name                 string `form:"name" binding:"required,max=255"`
	Email                string `form:"email" binding:"required,email"`
	Password             string `form:"password" binding:"min=6,max=32"`
	PasswordConfirmation string `form:"passwordConfirmation" binding:"eqfield=Password"`
}

func (p *Plugin) postUsersSignUp(c *gin.Context, l string, o interface{}) (interface{}, error) {
	fm := o.(*fmSignUp)

	var count int
	if err := p.Db.
		Model(&User{}).
		Where("email = ?", fm.Email).
		Count(&count).Error; err != nil {
		return nil, err
	}

	if count > 0 {
		return nil, p.I18n.E(l, "auth.errors.email-already-exists")
	}

	user, err := p.Dao.AddEmailUser(fm.Name, fm.Email, fm.Password)
	if err != nil {
		return nil, err
	}

	p.Dao.Log(user.ID, c.ClientIP(), p.I18n.T(l, "auth.logs.sign-up"))
	p.sendEmail(l, user, actConfirm)

	return gin.H{"message": p.I18n.T(l, "auth.messages.email-for-confirm")}, nil
}

func (p *Plugin) getUsersSignIn(c *gin.Context, lang string) (gin.H, error) {
	title := p.I18n.T(lang, "auth.users.sign-in.title")
	fm := widgets.NewForm(
		c.Request,
		lang,
		"/users/sign-in",
		"/dashboard",
		title,
		widgets.NewEmailField("email", p.I18n.T(lang, "attributes.email"), ""),
		widgets.NewPasswordField("password", p.I18n.T(lang, "attributes.password"), ""),
	)
	return gin.H{"form": fm, "title": title}, nil
}

type fmSignIn struct {
	Email      string `form:"email" binding:"required,email"`
	Password   string `form:"password" binding:"required"`
	RememberMe bool   `form:"rememberMe"`
}

func (p *Plugin) postUsersSignIn(c *gin.Context, lang string, o interface{}) (interface{}, error) {
	fm := o.(*fmSignIn)

	user, err := p.Dao.SignIn(lang, fm.Email, fm.Password, c.ClientIP())
	if err != nil {
		return nil, err
	}

	cm := jws.Claims{}
	cm.Set(UID, user.UID)
	tkn, err := p.Jwt.Sum(cm, time.Hour*24*7)
	if err != nil {
		return nil, err
	}
	c.SetCookie(TOKEN, string(tkn), 0, "/", "", web.Secure(), true)
	return gin.H{}, nil
}

type fmEmail struct {
	Email string `form:"email" binding:"required,email"`
}

func (p *Plugin) getUsersEmail(action string) func(*gin.Context, string) (gin.H, error) {
	return func(c *gin.Context, lang string) (gin.H, error) {
		title := p.I18n.T(lang, "auth.users."+action+".title")
		fm := widgets.NewForm(
			c.Request,
			lang,
			"/users/"+action,
			"/users/sign-in",
			title,
			widgets.NewEmailField("email", p.I18n.T(lang, "attributes.email"), ""),
		)
		return gin.H{"form": fm, "title": title}, nil
	}
}

func (p *Plugin) getUsersConfirmToken(c *gin.Context, lang string) (string, error) {
	token := c.Param("token")
	user, err := p.parseToken(lang, token, actConfirm)
	if err != nil {
		return "", err
	}
	if user.IsConfirm() {
		return "", p.I18n.E(lang, "auth.errors.user-already-confirm")
	}
	p.Db.Model(user).Update("confirmed_at", time.Now())
	p.Dao.Log(user.ID, c.ClientIP(), p.I18n.T(lang, "auth.logs.confirm"))
	return "/users/sign-in", nil
}

func (p *Plugin) postUsersConfirm(c *gin.Context, lang string, o interface{}) (interface{}, error) {
	fm := o.(*fmEmail)
	user, err := p.Dao.GetByEmail(fm.Email)
	if err != nil {
		return nil, err
	}

	if user.IsConfirm() {
		return nil, p.I18n.E(lang, "auth.errors.user-already-confirm")
	}

	p.sendEmail(lang, user, actConfirm)

	return gin.H{"message": p.I18n.T(lang, "auth.messages.email-for-confirm")}, nil
}

func (p *Plugin) getUsersUnlockToken(c *gin.Context, lang string) (string, error) {
	token := c.Param("token")
	user, err := p.parseToken(lang, token, actUnlock)
	if err != nil {
		return "", err
	}
	if !user.IsLock() {
		return "", p.I18n.E(lang, "auth.errors.user-not-lock")
	}

	p.Db.Model(user).Update(map[string]interface{}{"locked_at": nil})
	p.Dao.Log(user.ID, c.ClientIP(), p.I18n.T(lang, "auth.logs.unlock"))
	return "/users/sign-in", nil
}

func (p *Plugin) postUsersUnlock(c *gin.Context, lang string, o interface{}) (interface{}, error) {
	fm := o.(*fmEmail)
	user, err := p.Dao.GetByEmail(fm.Email)
	if err != nil {
		return nil, err
	}
	if !user.IsLock() {
		return nil, p.I18n.E(lang, "auth.errors.user-not-lock")
	}
	p.sendEmail(lang, user, actUnlock)
	return gin.H{"message": p.I18n.T(lang, "auth.messages.email-for-unlock")}, nil
}

func (p *Plugin) postUsersForgotPassword(c *gin.Context, lang string, o interface{}) (interface{}, error) {
	fm := o.(*fmEmail)
	var user *User
	user, err := p.Dao.GetByEmail(fm.Email)
	if err != nil {
		return nil, err
	}
	p.sendEmail(lang, user, actResetPassword)

	return gin.H{"message": p.I18n.T(lang, "auth.messages.email-for-reset-password")}, nil
}

type fmResetPassword struct {
	Token                string `form:"token" binding:"required"`
	Password             string `form:"password" binding:"min=6,max=32"`
	PasswordConfirmation string `form:"passwordConfirmation" binding:"eqfield=Password"`
}

func (p *Plugin) getUsersResetPassword(c *gin.Context, lang string) (gin.H, error) {
	title := p.I18n.T(lang, "auth.users.reset-password.title")
	fm := widgets.NewForm(
		c.Request,
		lang,
		"/users/reset-password",
		"/users/sign-in",
		title,
		widgets.NewHiddenField("token", c.Param("token")),
		widgets.NewPasswordField("password", p.I18n.T(lang, "attributes.password"), p.I18n.T(lang, "helpers.password")),
		widgets.NewPasswordField("passwordConfirmation", p.I18n.T(lang, "attributes.passwordConfirmation"), p.I18n.T(lang, "helpers.passwordConfirmation")),
	)
	return gin.H{"form": fm, "title": title}, nil
}

func (p *Plugin) postUsersResetPassword(c *gin.Context, lang string, o interface{}) (interface{}, error) {
	fm := o.(*fmResetPassword)
	user, err := p.parseToken(lang, fm.Token, actResetPassword)
	if err != nil {
		return nil, err
	}
	p.Db.Model(user).Update("password", p.Hmac.Sum([]byte(fm.Password)))
	p.Dao.Log(user.ID, c.ClientIP(), p.I18n.T(lang, "auth.logs.reset-password"))
	return gin.H{"message": p.I18n.T(lang, "auth.messages.reset-password-success")}, nil
}

func (p *Plugin) deleteUsersSignOut(c *gin.Context, lang string) (interface{}, error) {
	user := c.MustGet(CurrentUser).(*User)
	p.Dao.Log(user.ID, c.ClientIP(), p.I18n.T(lang, "auth.logs.sign-out"))
	c.SetCookie(TOKEN, "", -1, "/", "", web.Secure(), true)
	return gin.H{}, nil
}

func (p *Plugin) getUsersInfo(c *gin.Context, lang string) (gin.H, error) {
	user := c.MustGet(CurrentUser).(*User)
	title := p.I18n.T(lang, "auth.users.info.title")
	email := widgets.NewEmailField("email", p.I18n.T(lang, "attributes.email"), user.Email)
	email.Readonly()
	fm := widgets.NewForm(
		c.Request,
		lang,
		"/users/info",
		"",
		title,
		widgets.NewTextField("name", p.I18n.T(lang, "attributes.fullName"), user.Name),
		email,
	)
	return gin.H{"form": fm, "title": title}, nil
}

type fmInfo struct {
	Name string `form:"name" binding:"required,max=255"`
	// Home string `form:"home" binding:"max=255"`
	// Logo string `form:"logo" binding:"max=255"`
}

func (p *Plugin) postUsersInfo(c *gin.Context, lang string, o interface{}) (interface{}, error) {
	user := c.MustGet(CurrentUser).(*User)
	fm := o.(*fmInfo)

	if err := p.Db.Model(user).Updates(map[string]interface{}{
		// "home": fm.Home,
		// "logo": fm.Logo,
		"name": fm.Name,
	}).Error; err != nil {
		return nil, err
	}
	return gin.H{}, nil
}
func (p *Plugin) getUsersChangePassword(c *gin.Context, lang string) (gin.H, error) {
	title := p.I18n.T(lang, "auth.users.change-password.title")
	fm := widgets.NewForm(
		c.Request,
		lang,
		"/users/change-password",
		"",
		title,
		widgets.NewPasswordField("currentPassword", p.I18n.T(lang, "attributes.currentPassword"), ""),
		widgets.NewPasswordField("newPassword", p.I18n.T(lang, "attributes.newPassword"), p.I18n.T(lang, "helpers.password")),
		widgets.NewPasswordField("passwordConfirmation", p.I18n.T(lang, "attributes.passwordConfirmation"), p.I18n.T(lang, "helpers.passwordConfirmation")),
	)
	return gin.H{"form": fm, "title": title}, nil
}

type fmChangePassword struct {
	CurrentPassword      string `form:"currentPassword" binding:"required"`
	NewPassword          string `form:"newPassword" binding:"min=6,max=32"`
	PasswordConfirmation string `form:"passwordConfirmation" binding:"eqfield=NewPassword"`
}

func (p *Plugin) postUsersChangePassword(c *gin.Context, lang string, o interface{}) (interface{}, error) {
	user := c.MustGet(CurrentUser).(*User)
	fm := o.(*fmChangePassword)
	if !p.Hmac.Chk([]byte(fm.CurrentPassword), user.Password) {
		return nil, p.I18n.E(lang, "auth.errors.bad-password")
	}
	if err := p.Db.Model(user).
		Update("password", p.Hmac.Sum([]byte(fm.NewPassword))).Error; err != nil {
		return nil, err
	}

	return gin.H{}, nil
}

func (p *Plugin) getUsersLogs(c *gin.Context, lang string) (gin.H, error) {
	user := c.MustGet(CurrentUser).(*User)
	data := gin.H{}
	data["title"] = p.I18n.T(lang, "auth.users.logs.title")
	var logs []Log
	err := p.Db.
		Select([]string{"ip", "message", "created_at"}).
		Where("user_id = ?", user.ID).
		Order("id DESC").Limit(120).
		Find(&logs).Error
	data["logs"] = logs
	return data, err
}

func (p *Plugin) indexUsers(c *gin.Context, l string) (gin.H, error) {
	var users []User
	data := gin.H{}
	data["title"] = p.I18n.T(l, "auth.users.index.title")
	if err := p.Db.
		Select([]string{"name", "logo", "home"}).
		Order("last_sign_in_at DESC").
		Find(&users).Error; err != nil {
		return nil, err
	}
	data["items"] = users
	return data, nil
}
