package root

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// InstallController /install
type InstallController struct {
	beego.Controller
}

// Get install
// @router / [get]
func (p *InstallController) Get() {

}

// Post install
// @router / [post]
func (p *InstallController) Post() {

}

func (p *InstallController) checkDatabaseEmpty() error {
	count, err := orm.NewOrm().QueryTable(new(Host)).Count()
	if err != nil {
		return err
	}
	if count > 0 {
		// todo
	}
	return nil
}
