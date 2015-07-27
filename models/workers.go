package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"

	"fmt"
	"time"
)

const (
	WORKER_STATUS_IDLE = 0
	WORKER_STATUS_BUSY = 1
	WORKER_STATUS_DIE  = 2
)


type Worker struct {
	Id         int64  `orm:"pk;auto"`
	Ip         string `orm:"size(1024)"`
	LastUpdate time.Time
	Status     int
}

func GetAllWorkers() []*Worker {
	o := orm.NewOrm()

	var workers []*Worker
	num, err := o.QueryTable("worker").All(&workers)
	fmt.Printf("Returned Rows Num: %s, %s", num, err)
	return workers
}

func GetAllWorkersWithStatus(status int) []*Worker {
	o := orm.NewOrm()

	var workers []*Worker
	num, err := o.QueryTable("worker").Filter("status", status).All(&workers)
	fmt.Printf("Returned Rows Num: %s, %s", num, err)
	return workers
}
