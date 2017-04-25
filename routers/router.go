package routers

import (
	"github.com/astaxie/beego"
	"github.com/kapmahc/fly/plugins/auth"
	"github.com/kapmahc/fly/plugins/forum"
	"github.com/kapmahc/fly/plugins/ops/mail"
	"github.com/kapmahc/fly/plugins/ops/vpn"
	"github.com/kapmahc/fly/plugins/reading"
	"github.com/kapmahc/fly/plugins/shop"
	"github.com/kapmahc/fly/plugins/site"
)

func init() {
	beego.Include(&auth.Controller{}, &site.Controller{})
	for k, v := range map[string]beego.ControllerInterface{
		"/forum":    &forum.Controller{},
		"/reading":  &reading.Controller{},
		"/shop":     &shop.Controller{},
		"/ops/mail": &mail.Controller{},
		"/ops/vpn":  &vpn.Controller{},
	} {
		ns := beego.NewNamespace(k, beego.NSInclude(v))
		beego.AddNamespace(ns)
	}
}
