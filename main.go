package main

import (
	"fmt"
	"os"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	"github.com/ArchCI/archci/models"
	_ "github.com/ArchCI/archci/routers"
)

const (
	ENV_MYSQL_SERVER   = "MYSQL_SERVER"
	ENV_MYSQL_USERNAME = "MYSQL_USERNAME"
	ENV_MYSQL_PASSWORD = "MYSQL_PASSWORD"
	ENV_MYSQL_DATABASE = "MYSQL_DATABASE"

	MYSQL_DRIVER = "mysql"
)

// Init will initialize database to create tables automatically.
func init() {
	// Registry archci database models.
	models.RegisterModels()

	// Initialize database with environment variables.
	server := ""
	username := "root"
	password := "root"
	database := "mysql"

	if os.Getenv(ENV_MYSQL_SERVER) != "" {
		server = os.Getenv(ENV_MYSQL_SERVER)
	}
	if os.Getenv(ENV_MYSQL_USERNAME) != "" {
		username = os.Getenv(ENV_MYSQL_USERNAME)
	}
	if os.Getenv(ENV_MYSQL_PASSWORD) != "" {
		password = os.Getenv(ENV_MYSQL_PASSWORD)
	}
	if os.Getenv(ENV_MYSQL_DATABASE) != "" {
		database = os.Getenv(ENV_MYSQL_DATABASE)
	}

	// The datasource looks like "root:root@/archci?charset=utf8".
	DATASOURCE := username + ":" + password + "@" + server + "/" + database + "?charset=utf8"
	fmt.Println("Connect to database with " + DATASOURCE)

	orm.RegisterDriver(MYSQL_DRIVER, orm.DR_MySQL)
	orm.RegisterDataBase("default", MYSQL_DRIVER, DATASOURCE, 30)
	orm.RunSyncdb("default", false, true)
}

// Main is the entry to start beego application.
func main() {
	beego.Run()
}
