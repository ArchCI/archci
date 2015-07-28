package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

const (
	PROJECT_STATUS_NEED_TEST = 0
	PROJECT_STATUS_TESTING   = 1
	PROJECT_STATUS_SUCCESS   = 2
	PROJECT_STATUS_FAIL      = 3
)

// TODO(tobe): add more restriction in http://beego.me/docs/mvc/model/models.md

// Project contains all the information of project.
type Project struct {
	Id          int64  `orm:"pk;auto"`
	UserName    string `orm:"size(1024);null"`
	ProjectName string `orm:"size(1024);null"`
	RepoUrl     string `orm:"size(1024);null"`
	Status      int    `orm:"null"`
}

// GetAllProjects return all projects from database.
func GetAllProjects() []*Project {
	o := orm.NewOrm()

	var projects []*Project
	num, err := o.QueryTable("project").All(&projects)
	fmt.Printf("Returned Rows Num: %s, %s", num, err)
	return projects
}

// GetProjectWithId takes id to return the project.
func GetProjectWithId(projectId int64) Project {
	o := orm.NewOrm()

	project := Project{Id: projectId}
	err := o.Read(&project)
	fmt.Printf("ERR: %v\n", err)
	return project
}

// ReadOrCreateProject tries to get the project and create it if it doesn't exist.
func ReadOrCreateProject(userName string, projectName string, repoUrl string) (int64, error) {
	o := orm.NewOrm()

	project := Project{
		UserName:    userName,
		ProjectName: projectName,
		RepoUrl:     repoUrl,
	}

	created, id, err := o.ReadOrCreate(&project, "UserName", "ProjectName", "RepoUrl")
	if err == nil {
		if created {
			fmt.Println("New Insert an object. Id:", id)
			return id, nil
		} else {
			fmt.Println("Get an object. Id:", id)
			return id, nil
		}
	} else {
		return 0, err
	}
}

// AddProject adds the project in database.
func AddProject(project Project) error {
	o := orm.NewOrm()

	_, err := o.Insert(&project)
	fmt.Println(err)
	return err
}
