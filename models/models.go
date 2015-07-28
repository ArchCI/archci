package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func RegisterModels() {
	orm.RegisterModel(new(Project), new(Build), new(Worker))
}
