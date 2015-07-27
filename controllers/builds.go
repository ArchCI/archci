package controllers

import (
	"fmt"
	_ "github.com/lib/pq"

	"github.com/golang/glog"

	"encoding/json"
	"github.com/ArchCI/archci/models"
	"github.com/ArchCI/archci/redisutil"
)

// GetBuildLogsIndexResponse is the json struct to return to archci server.
type GetBuildLogsIndexResponse struct {
	Log  string `json:"log"`
	Next bool   `json:"next`
}

// New build record in database.
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

// Get all builds from database.
func (c *ApiController) GetBuildsAll() {
	glog.Info("Get all builds")

	builds := models.GetAllBuilds()

	c.Data["json"] = builds
	c.ServeJson()
}

// Get builds with project name.
func (c *ApiController) GetBuildsWithProjectName() {
	glog.Info("Get builds with project name")

	projectName := c.GetString(":projectName")

	builds := models.GetBuildsWithProjectName(projectName)

	c.Data["json"] = builds
	c.ServeJson()
}

// Get one build with build id.
func (c *ApiController) GetBuild() {
	glog.Info("Get build")

	buildId, _ := c.GetInt64(":buildId")

	build := models.GetBuildWithId(buildId)

	c.Data["json"] = build
	c.ServeJson()
}

// Get active builds.
func (c *ApiController) GetActiveBuilds() {
	glog.Info("Get active builds")

	result := "Not implemented"
	c.Ctx.WriteString(result)
}

// Get search builds.
func (c *ApiController) GetSearchBuilds() {
	glog.Info("Get search builds")

	result := "Not implemented"
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

// Get all logs of the build.
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
