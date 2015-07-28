package controllers

import (
	log "github.com/Sirupsen/logrus"

	"github.com/ArchCI/archci/models"
)

// GetWorkerAll returns all workers from database.
func (c *ApiController) GetWorkersAll() {
	log.Info("Get all workers")

	workers := models.GetAllWorkers()

	c.Data["json"] = workers
	c.ServeJson()
}

// GetWorkersAllStatus take the parameter of status and return the workers.
func (c *ApiController) GetWorkersAllStatus() {
	log.Info("Get all workers with status")

	status, _ := c.GetInt(":status")
	workers := models.GetAllWorkersWithStatus(status)

	c.Data["json"] = workers
	c.ServeJson()
}
