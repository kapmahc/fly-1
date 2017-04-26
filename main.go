package main

import (
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/toolbox"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kapmahc/fly/plugins/root"
	_ "github.com/kapmahc/fly/routers"
	_ "github.com/lib/pq"
)

func main() {
	orm.RegisterDataBase(
		"default",
		beego.AppConfig.String("databasedriver"),
		beego.AppConfig.String("databasesource"),
	)
	if err := root.DatabaseMigrate(); err != nil {
		beego.Error(err)
	}
	// -----------
	toolbox.StartTask()
	defer toolbox.StopTask()
	// ---------
	beego.Run()
}
