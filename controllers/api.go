package controllers

import (
	//"database/sql"
	"fmt"
	_ "github.com/lib/pq"

	"github.com/astaxie/beego"
	"github.com/golang/glog"

	"encoding/json"
	"github.com/ArchCI/archci/models"
	"github.com/ArchCI/archci/redisutil"
)

type ApiController struct {
	beego.Controller
}

type GetBuildLogsIndexResponse struct {
	Log  string `json:"log"`
	Next bool   `json:"next`
}

/* Create acount */
func (c *ApiController) CreateAccount() {
	glog.Info("Creat accout")

	result := "{data: 1}"
	c.Ctx.WriteString(result)
}

// New build record in database
func (c *ApiController) NewBuild() {
	glog.Info("New build record")

	project := models.Project{}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &project); err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte("empty title"))
		fmt.Println(err)
		return
	}

	models.AddBuildWithProject(project)
}

// Get all builds from database
func (c *ApiController) GetBuildsAll() {
	glog.Info("Get all builds")

	builds := models.GetAllBuilds()

	c.Data["json"] = builds
	c.ServeJson()
}

// Get builds with project name
func (c *ApiController) GetBuildsWithProjectName() {
	glog.Info("Get builds with project name")

	projectName := c.GetString(":projectName")

	builds := models.GetBuildsWithProjectName(projectName)

	c.Data["json"] = builds
	c.ServeJson()
}

// Get one build with build id
func (c *ApiController) GetBuild() {
	glog.Info("Get build")

	buildId, _ := c.GetInt64(":buildId")

	build := models.GetBuildWithId(buildId)

	c.Data["json"] = build
	c.ServeJson()
}

/* Get active builds */
func (c *ApiController) GetActiveBuilds() {
	glog.Info("Get active builds")

	result := "{data: 1}"
	c.Ctx.WriteString(result)
}

/* Get search builds */
func (c *ApiController) GetSearchBuilds() {
	glog.Info("Get search builds")

	result := "{data: 1}"
	c.Ctx.WriteString(result)
}

/* Get build log by id */
func (c *ApiController) GetBuildLog() {
	glog.Info("Get build log")

	result := "{data: 1}"
	c.Ctx.WriteString(result)
}

/* Put build log with index */
func (c *ApiController) PutBuildLogsIndex() {
	glog.Info("Put build log part")

	result := "{data: 1}"
	c.Ctx.WriteString(result)
}

/* Get build log with index */
func (c *ApiController) GetBuildLogsIndex() {
	glog.Info("Get build log with index")

	buildId := c.GetString(":buildId")
	index, _ := c.GetInt(":index")

	log := redisutil.HgetString(buildId, index)

	next := false

	finish := redisutil.HgetBool(buildId, "finish")
	if finish == false {
		next = true
		fmt.Println("finish == false")
	}

	current := redisutil.HgetInt(buildId, "current")
	if index < current {
		next = true
		fmt.Println("index != current")
	}

	// Throw error if index is larger than current

	// Handle when get the index is more than current but it's not total

	response := &GetBuildLogsIndexResponse{
		Log:  log,
		Next: next}

	c.Data["json"] = response
	c.ServeJson()

}

// Get all logs of the build
func (c *ApiController) GetBuildLogsAll() {
	glog.Info("Get all build logs")

	//buildId := c.GetString(":buildId")
	//field := 0
	//result := redisutil.HgetString(buildId, field)
	//c.Ctx.WriteString(result)

	// TODO(tobe): change to hgetall  from redis
	mystruct := `{0: "apt-get install", 1: "go test"}`

	c.Data["json"] = &mystruct
	c.ServeJson()
}

/* Create project */
func (c *ApiController) CreateProject() {
	glog.Info("Creat project")

	result := "{data: 1}"
	c.Ctx.WriteString(result)
}

// Get all projects from database
func (c *ApiController) GetProjectsAll() {
	glog.Info("Get all projects")

	projects := models.GetAllProjects()

	c.Data["json"] = projects
	c.ServeJson()
}

// Get one project with project id
func (c *ApiController) GetProject() {
	glog.Info("Get project")

	projectId, _ := c.GetInt64(":projectId")

	project := models.GetProjectWithId(projectId)

	c.Data["json"] = project
	c.ServeJson()
}

// Get all workers
func (c *ApiController) GetWorkersAll() {
	glog.Info("Get all workers")

	workers := models.GetAllWorkers()

	c.Data["json"] = workers
	c.ServeJson()
}

// Get all worker with this status
func (c *ApiController) GetWorkersAllStatus() {
	glog.Info("Get all workers with status")

	status, _ := c.GetInt(":status")
	workers := models.GetAllWorkersWithStatus(status)

	c.Data["json"] = workers
	c.ServeJson()
}

/* Create image */
func (c *ApiController) CreateImage() {
	glog.Info("Create image")

	result := "{data: 1}"
	c.Ctx.WriteString(result)
}

/* Get images */
func (c *ApiController) GetImages() {
	glog.Info("Get images")

	result := "{data: 1}"
	c.Ctx.WriteString(result)
}

/* Get image by id */
func (c *ApiController) GetImage() {
	glog.Info("Get image")

	result := "{data: 1}"
	c.Ctx.WriteString(result)
}

/* Get tasks */
func (c *ApiController) GetTasks() {
	glog.Info("Get tasks")

	result := "{data: 1}"
	c.Ctx.WriteString(result)
}

/* Finish task */
func (c *ApiController) FinishTask() {
	glog.Info("Finish task")

	result := "{data: 1}"
	c.Ctx.WriteString(result)
}

/* Get workers */
func (c *ApiController) GetWorkers() {
	glog.Info("Get workers")

	result := "{data: 1}"
	c.Ctx.WriteString(result)
}
