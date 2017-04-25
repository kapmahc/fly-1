package shop

import "github.com/kapmahc/fly/plugins/auth"

// Controller controller
type Controller struct {
	auth.Layout
}

// GetHome home page
// @router / [get]
func (p *Controller) GetHome() {

}
