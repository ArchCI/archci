package controllers

import (
	_ "github.com/lib/pq"

	log "github.com/Sirupsen/logrus"
	"github.com/astaxie/beego"
)

// ApiController is the custom controller to provide APIs.
type ApiController struct {
	beego.Controller
}

// CreateAccount is used to create an new account.
func (c *ApiController) CreateAccount() {
	// TODO(tobe): this is not implemented until we integrate with Github or Gitlab.
	log.Info("Creat accout")

	result := "Not implemented"
	c.Ctx.WriteString(result)
}
