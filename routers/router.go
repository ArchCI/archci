package routers

import (
	"github.com/ArchCI/archci/controllers"
	"github.com/astaxie/beego"
)

func init() {
	// archci web pages
	beego.Router("/", &controllers.MainController{})
	beego.Router("/builds", &controllers.MainController{})
	beego.Router("/builds/:buildId", &controllers.MainController{})
	beego.Router("/projects", &controllers.MainController{})
	beego.Router("/projects/:projectId", &controllers.MainController{})
	beego.Router("/workers", &controllers.MainController{})

	// account api
	beego.Router("/v1/account", &controllers.ApiController{}, "post:CreateAccount")

	// builds api
	beego.Router("/v1/builds/new", &controllers.ApiController{}, "post:NewBuild")
	beego.Router("/v1/builds/all", &controllers.ApiController{}, "get:GetBuildsAll")
	beego.Router("/v1/builds/all/project/:projectName", &controllers.ApiController{}, "get:GetBuildsWithProjectName")
	beego.Router("/v1/builds/:buildId", &controllers.ApiController{}, "get:GetBuild")
	beego.Router("/v1/builds/active", &controllers.ApiController{}, "get:GetActiveBuilds")
	beego.Router("/v1/builds/search", &controllers.ApiController{}, "get:GetSearchBuilds")
	beego.Router("/v1/builds/:buildId/logs/:index", &controllers.ApiController{}, "get:GetBuildLogsIndex")
	beego.Router("/v1/builds/:buildId/logs/all", &controllers.ApiController{}, "get:GetBuildLogsAll")

	// projects api
	beego.Router("/v1/projects/new", &controllers.ApiController{}, "post:NewProject")
	beego.Router("/v1/projects/all", &controllers.ApiController{}, "get:GetProjectsAll")
	beego.Router("/v1/projects/:projectId", &controllers.ApiController{}, "get:GetProject")

	// workers api
	beego.Router("/v1/workers/all", &controllers.ApiController{}, "get:GetWorkersAll")
	// TODO(tobe): this is not really RESTful
	beego.Router("/v1/workers/all/status/:status", &controllers.ApiController{}, "get:GetWorkersAllStatus")

	// hooks api
	beego.Router("/v1/hook/github/push", &controllers.ApiController{}, "post:TriggerGithubPushHook")
	beego.Router("/v1/hook/gitlab/push", &controllers.ApiController{}, "post:TriggerGitlabPushHook")
}
