package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	"github.com/ArchCI/archci/models"
	_ "github.com/ArchCI/archci/routers"
)

const (
	_MYSQL_DRIVER = "mysql"
	_DATASOURCE   = "root:root@/archci?charset=utf8"
)

func init() {
	models.RegisterModels()

	orm.RegisterDriver(_MYSQL_DRIVER, orm.DR_MySQL)
	orm.RegisterDataBase("default", _MYSQL_DRIVER, _DATASOURCE, 30)
	orm.RunSyncdb("default", false, true)
}

func main() {

	beego.Run()
}
