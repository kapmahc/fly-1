package reading

import "github.com/kapmahc/fly/plugins/root"

// Controller controller
type Controller struct {
	root.Layout
}

// GetHome home page
// @router / [get]
func (p *Controller) GetHome() {

}
