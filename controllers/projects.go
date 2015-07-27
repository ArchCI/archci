package controllers

import (
	"fmt"
	_ "github.com/lib/pq"

	"github.com/golang/glog"

	"encoding/json"
	"github.com/ArchCI/archci/models"
)

// Add new project.
func (c *ApiController) NewProject() {
	glog.Info("New build record")

	project := models.Project{}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &project); err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte("empty title"))
		fmt.Println(err)
		return
	}

	models.AddProject(project)
}

// Get all projects from database.
func (c *ApiController) GetProjectsAll() {
	glog.Info("Get all projects")

	projects := models.GetAllProjects()

	c.Data["json"] = projects
	c.ServeJson()
}

// Get one project with project id.
func (c *ApiController) GetProject() {
	glog.Info("Get project")

	projectId, _ := c.GetInt64(":projectId")

	project := models.GetProjectWithId(projectId)

	c.Data["json"] = project
	c.ServeJson()
}
