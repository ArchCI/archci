package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Worker struct {
	Id         int64  `orm:"pk;auto"`
	Ip         string `orm:"size(1024)"`
	LastUpdate time.Time
	Status     int
}

func RegisterModels() {
	orm.RegisterModel(new(Project), new(Build), new(Worker))
}
