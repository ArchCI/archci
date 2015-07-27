package controllers

import (
	_ "github.com/lib/pq"

	"github.com/astaxie/beego"
	"github.com/golang/glog"
)

type ApiController struct {
	beego.Controller
}

// Create account.
func (c *ApiController) CreateAccount() {
	glog.Info("Creat accout")

	result := "Not implemented"
	c.Ctx.WriteString(result)
}
