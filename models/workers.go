package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

const (
	WORKER_STATUS_IDLE = 0
	WORKER_STATUS_BUSY = 1
	WORKER_STATUS_DIE  = 2
)

// Worker contains all the information of the worker.
type Worker struct {
	Id         int64  `orm:"pk;auto"`
	Ip         string `orm:"size(1024)"`
	LastUpdate time.Time
	Status     int
}

// GetAllWorkers return all the workers from database.
func GetAllWorkers() []*Worker {
	o := orm.NewOrm()

	var workers []*Worker
	num, err := o.QueryTable("worker").All(&workers)
	fmt.Printf("Returned Rows Num: %s, %s", num, err)
	return workers
}

// GetAllWorkersWithStatus takes the status and return the workers.
func GetAllWorkersWithStatus(status int) []*Worker {
	o := orm.NewOrm()

	var workers []*Worker
	num, err := o.QueryTable("worker").Filter("status", status).All(&workers)
	fmt.Printf("Returned Rows Num: %s, %s", num, err)
	return workers
}
