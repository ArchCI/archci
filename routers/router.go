package routers

import (
	"github.com/ArchCI/archci/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
