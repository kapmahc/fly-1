package auth

import (
	"strconv"

	"github.com/astaxie/beego"
)

// Layout layout
type Layout struct {
	beego.Controller
}

// Abort abort
func (p *Layout) Abort(code int) {
	p.Controller.Abort(strconv.Itoa(code))
}
