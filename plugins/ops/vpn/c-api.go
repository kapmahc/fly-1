package vpn

import (
	"net/http"
	"time"

	"github.com/kapmahc/h2o"
	"github.com/kapmahc/h2o/i18n"
)

type fmSignIn struct {
	Email    string `form:"username" validate:"required,email"`
	Password string `form:"password" validate:"min=6,max=32"`
}

func (p *Plugin) apiAuth(c *h2o.Context) error {
	var fm fmSignIn
	if err := c.Bind(&fm); err != nil {
		return err
	}
	lng := c.Get(i18n.LOCALE).(string)
	var user User
	if err := p.Db.Where("email = ?", fm.Email).First(&user).Error; err != nil {
		return err
	}
	now := time.Now()
	if user.Enable && user.StartUp.Before(now) && user.ShutDown.After(now) {
		return c.JSON(http.StatusOK, h2o.H{})
	}
	return p.I18n.E(http.StatusForbidden, lng, "ops.vpn.errors.user-is-not-available")
}

type fmStatus struct {
	Email       string  `form:"common_name" validate:"required,email"`
	TrustedIP   string  `form:"trusted_ip" validate:"required"`
	TrustedPort uint    `form:"trusted_port" validate:"required"`
	RemoteIP    string  `form:"ifconfig_pool_remote_ip" validate:"required"`
	RemotePort  uint    `form:"remote_port_1" validate:"required"`
	Received    float64 `form:"bytes_received" validate:"required"`
	Send        float64 `form:"bytes_sent" validate:"required"`
}

func (p *Plugin) apiConnect(c *h2o.Context) error {
	var fm fmStatus
	if err := c.Bind(&fm); err != nil {
		return err
	}
	var user User
	if err := p.Db.Where("email = ?", fm.Email).First(&user).Error; err != nil {
		return err
	}
	if err := p.Db.Create(&Log{
		UserID:      user.ID,
		RemoteIP:    fm.RemoteIP,
		RemotePort:  fm.RemotePort,
		TrustedIP:   fm.TrustedIP,
		TrustedPort: fm.TrustedPort,
		Received:    fm.Received,
		Send:        fm.Send,
		StartUp:     time.Now(),
	}).Error; err != nil {
		return err
	}
	if err := p.Db.Model(&User{}).
		Where("id = ?", user.ID).
		UpdateColumn("online", true).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, h2o.H{})
}

func (p *Plugin) apiDisconnect(c *h2o.Context) error {
	var fm fmStatus
	if err := c.Bind(&fm); err != nil {
		return err
	}
	var user User
	if err := p.Db.Where("email = ?", fm.Email).First(&user).Error; err != nil {
		return err
	}
	if err := p.Db.Model(&User{}).
		Where("id = ?", user.ID).
		UpdateColumn("online", false).Error; err != nil {
		return err
	}

	if err := p.Db.
		Model(&Log{}).
		Where(
			"trusted_ip = ? AND trusted_port = ? AND user_id = ? AND shut_down IS NULL",
			fm.TrustedIP,
			fm.TrustedPort,
			user.ID,
		).Update(map[string]interface{}{
		"shut_down": time.Now(),
		"received":  fm.Received,
		"send":      fm.Send,
	}).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, h2o.H{})
}
