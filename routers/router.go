package routers

import (
	"github.com/ArchCI/archci/controllers"
	"github.com/astaxie/beego"
)

// Init registries all the APIs when startup.
func init() {
	// ArchCI web pages.
	beego.Router("/", &controllers.MainController{})
	beego.Router("/builds", &controllers.MainController{})
	beego.Router("/builds/:buildId", &controllers.MainController{})
	beego.Router("/projects", &controllers.MainController{})
	beego.Router("/projects/:projectId", &controllers.MainController{})
	beego.Router("/workers", &controllers.MainController{})
	beego.Router("/account", &controllers.MainController{})

	// Login api.
	beego.Router("/v1/login/github", &controllers.ApiController{}, "get:LoginGithub")
	beego.Router("/v1/login/github/callback", &controllers.ApiController{}, "get:LoginGithubCallback")

	// Builds api.
	beego.Router("/v1/builds/new", &controllers.ApiController{}, "post:NewBuild")
	beego.Router("/v1/builds/all", &controllers.ApiController{}, "get:GetBuildsAll")
	beego.Router("/v1/builds/all/project/:projectName", &controllers.ApiController{}, "get:GetBuildsWithProjectName")
	beego.Router("/v1/builds/:buildId", &controllers.ApiController{}, "get:GetBuild")
	beego.Router("/v1/builds/active", &controllers.ApiController{}, "get:GetActiveBuilds")
	beego.Router("/v1/builds/search", &controllers.ApiController{}, "get:GetSearchBuilds")
	beego.Router("/v1/builds/:buildId/logs/:index", &controllers.ApiController{}, "get:GetBuildLogsIndex")
	beego.Router("/v1/builds/:buildId/logs/all", &controllers.ApiController{}, "get:GetBuildLogsAll")

	// Projects api.
	beego.Router("/v1/projects/new", &controllers.ApiController{}, "post:NewProject")
	beego.Router("/v1/projects/all", &controllers.ApiController{}, "get:GetProjectsAll")
	beego.Router("/v1/projects/:projectId", &controllers.ApiController{}, "get:GetProject")

	// Workers api.
	beego.Router("/v1/workers/all", &controllers.ApiController{}, "get:GetWorkersAll")
	// TODO(tobe): this is not really RESTful
	beego.Router("/v1/workers/all/status/:status", &controllers.ApiController{}, "get:GetWorkersAllStatus")

	// account api.
	beego.Router("/v1/account/info", &controllers.ApiController{}, "get:GetAccountInfo")
	beego.Router("/v1/account/projects", &controllers.ApiController{}, "get:GetAccountProjects")
	beego.Router("/v1/account/token", &controllers.ApiController{}, "get:GetAccountToken")

	// Hooks api.
	beego.Router("/v1/hook/github/push", &controllers.ApiController{}, "post:TriggerGithubPushHook")
	beego.Router("/v1/hook/gitlab/push", &controllers.ApiController{}, "post:TriggerGitlabPushHook")

	// Badge api.
	beego.Router("/v1/badge/:projectId", &controllers.ApiController{}, "get:GetProjectBadge")
	beego.Router("/v1/badge/:projectId/url", &controllers.ApiController{}, "get:GetProjectBadgeUrl")
	beego.Router("/v1/badge/:projectId/markdown", &controllers.ApiController{}, "get:GetProjectBadgeMarkdown")
}
