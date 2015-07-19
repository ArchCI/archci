package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	"github.com/ArchCI/archci/models"
	_ "github.com/ArchCI/archci/routers"
)

const (
	_MYSQL_DRIVER = "mysql"
	_DATASOURCE   = "root:@/my_db?charset=utf8"
)

func init() {

	models.RegisterModels()

	orm.RegisterDriver(_MYSQL_DRIVER, orm.DR_MySQL)
	orm.RegisterDataBase("default", _MYSQL_DRIVER, _DATASOURCE, 30)
	orm.RunSyncdb("default", false, true)
}

func main() {

	/*
		models.AddProject("tobegit3hub/seagull", "https://github.com/tobegit3hub/seagull")
		models.AddProject("tobegit3hub/note", "https://github.com/tobegit3hub/note")
		models.AddProject("ArchCI/archci", "https://github.com/ArchCI/archci")

		commitTime := time.Now()
		//models.AddBuild("tobegit3hub/seagull", "master", "a34dbad42", commitTime, "tobegit3hub")
		models.AddBuild("ArchCI/archci", "master", "ba888d42", commitTime, "tobegit3hub")
		models.AddBuild("ArchCI/simple-worker", "master", "ffdbad42", commitTime, "tobegit3hub")
	*/

	beego.Run()
}
