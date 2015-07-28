package models

import (
	"github.com/astaxie/beego/orm"
)

// RegisterModels registries the models of archci.
func RegisterModels() {
	orm.RegisterModel(new(Project), new(Build), new(Worker))
}
