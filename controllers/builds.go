package controllers

import (
	"fmt"
	"encoding/json"

	log "github.com/Sirupsen/logrus"

	"github.com/ArchCI/archci/models"
	"github.com/ArchCI/archci/redisutil"
)

// GetBuildLogsIndexResponse is the json struct to return to archci server.
type GetBuildLogsIndexResponse struct {
	Log  string `json:"log"`
	Next bool   `json:"next`
}

// Newbuild processes POST data to build record and add it in database.
func (c *ApiController) NewBuild() {
	log.Info("New build record")

	project := models.Project{}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &project); err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte("empty title"))
		fmt.Println(err)
		return
	}

	models.AddBuildWithProject(project)
}

// GetBuildsAll read all builds from database.
func (c *ApiController) GetBuildsAll() {
	log.Info("Get all builds")

	builds := models.GetAllBuilds()

	c.Data["json"] = builds
	c.ServeJson()
}

// GetBuildsWithProjectName gets project name in url and return the project object.
func (c *ApiController) GetBuildsWithProjectName() {
	log.Info("Get builds with project name")

	projectName := c.GetString(":projectName")

	builds := models.GetBuildsWithProjectName(projectName)

	c.Data["json"] = builds
	c.ServeJson()
}

// GetBuild gets the build id and return the build object.
func (c *ApiController) GetBuild() {
	log.Info("Get build")

	buildId, _ := c.GetInt64(":buildId")

	build := models.GetBuildWithId(buildId)

	c.Data["json"] = build
	c.ServeJson()
}

// GetActiveBuilds get build data from database.
func (c *ApiController) GetActiveBuilds() {
	// TODO(tobe): this is not implemented because nobody uses it yet.
	log.Info("Get active builds")

	result := "Not implemented"
	c.Ctx.WriteString(result)
}

// GetSearchBuilds get search string and return related builds.
func (c *ApiController) GetSearchBuilds() {
	log.Info("Get search builds")

	result := "Not implemented"
	c.Ctx.WriteString(result)
}

// GetBuildLogsIndex return the single build log with the index.
func (c *ApiController) GetBuildLogsIndex() {
	log.Info("Get build log with index")

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

	// TODO(tobe): throw error if index is larger than current
	// TODO(tobe): handle when get the index is more than current but it's not total

	response := &GetBuildLogsIndexResponse{
		Log:  log,
		Next: next}

	c.Data["json"] = response
	c.ServeJson()

}

// GetBuildLogsAll return all logs of the build.
func (c *ApiController) GetBuildLogsAll() {
	log.Info("Get all build logs")

	//buildId := c.GetString(":buildId")
	//field := 0
	//result := redisutil.HgetString(buildId, field)
	//c.Ctx.WriteString(result)

	// TODO(tobe): change to hgetall  from redis
	mystruct := `{0: "apt-get install", 1: "go test"}`

	c.Data["json"] = &mystruct
	c.ServeJson()
}
