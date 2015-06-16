package routers

import (
	"github.com/ArchCI/archci/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})

	beego.Router("/v1/account", &controllers.ApiController{}, "post:CreateAccount")
	// login

	beego.Router("/v1/builds/active", &controllers.ApiController{}, "get:GetActiveBuilds")
	beego.Router("/v1/builds/search", &controllers.ApiController{}, "get:GetSearchBuilds")
	beego.Router("/v1/builds/:id/logs", &controllers.ApiController{}, "get:GetBuildLog")
	beego.Router("/v1/builds/:id/logs/:part", &controllers.ApiController{}, "post:PutBuildLogPart")
	beego.Router("/v1/builds/:id/logs/:part", &controllers.ApiController{}, "get:GetBuildLogPart")

	beego.Router("/v1/projects", &controllers.ApiController{}, "post:CreateProject")
	beego.Router("/v1/projects", &controllers.ApiController{}, "get:GetProjects")

	beego.Router("/v1/images", &controllers.ApiController{}, "post:CreateImage")
	beego.Router("/v1/images", &controllers.ApiController{}, "get:GetImages")
	beego.Router("/v1/images/:id", &controllers.ApiController{}, "get:GetImage")

	beego.Router("/v1/tasks", &controllers.ApiController{}, "get:GetTasks")
	beego.Router("/v1/tasks/:id", &controllers.ApiController{}, "put:FinishTask")

	beego.Router("/v1/workers", &controllers.ApiController{}, "get:GetWorkers")
}
