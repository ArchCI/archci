package controllers

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"

	"github.com/golang/glog"
	"github.com/astaxie/beego"
)

type ApiController struct {
	beego.Controller
}

func usedb() {
	fmt.Println("Start to use postgrel")

	//db, err := sql.Open("postgres", "user=archci dbname=pqgotest sslmode=verify-full")
	db, err := sql.Open("postgres", "postgres://archci:archci@192.168.1.103/arch")
	if err != nil {
		glog.Fatal(err)
	}

	username := "tobe"
	rows, err2 := db.Query("SELECT * FROM account WHERE username = $1", username)
	if err2 != nil {
		glog.Fatal(err2)
	}

	fmt.Println(rows)

	fmt.Println("End of usedb")
}

/* Create acount */
func (c *ApiController) CreateAccount() {
	glog.Info("Creat accout")

	result := "{data: 1}"
	c.Ctx.WriteString(result)
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

/* Put build log part by part id */
func (c *ApiController) PutBuildLogPart() {
        glog.Info("Put build log part")

        result := "{data: 1}"
        c.Ctx.WriteString(result)
}


/* Get build log part by part id */
func (c *ApiController) GetBuildLogPart() {
	glog.Info("Get build log part")

	result := "{data: 1}"
	c.Ctx.WriteString(result)
}

/* Create project */
func (c *ApiController) CreateProject() {
	glog.Info("Creat project")

	result := "{data: 1}"
	c.Ctx.WriteString(result)
}

/* Get projects */
func (c *ApiController) GetProjects() {
	glog.Info("Get projects")

	result := "{data: 1}"
	c.Ctx.WriteString(result)
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





