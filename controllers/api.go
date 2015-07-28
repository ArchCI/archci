package controllers

import (
	_ "github.com/lib/pq"

	"github.com/astaxie/beego"
	log "github.com/Sirupsen/logrus"
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
