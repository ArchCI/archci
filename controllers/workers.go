package controllers

import (
	_ "github.com/lib/pq"

	"github.com/golang/glog"

	"github.com/ArchCI/archci/models"
)

// Get all workers.
func (c *ApiController) GetWorkersAll() {
	glog.Info("Get all workers")

	workers := models.GetAllWorkers()

	c.Data["json"] = workers
	c.ServeJson()
}

// Get all worker with this status.
func (c *ApiController) GetWorkersAllStatus() {
	glog.Info("Get all workers with status")

	status, _ := c.GetInt(":status")
	workers := models.GetAllWorkersWithStatus(status)

	c.Data["json"] = workers
	c.ServeJson()
}
