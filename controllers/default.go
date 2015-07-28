package controllers

import (
	"github.com/astaxie/beego"
)

// MainController is generated as the default controller.
type MainController struct {
	beego.Controller
}

// Get will render the single-page application.
func (c *MainController) Get() {
	c.TplNames = "index.html"
	c.Render()
}
