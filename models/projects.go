package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"

	"fmt"
)

const (
	PROJECT_STATUS_NEED_TEST = 0
	PROJECT_STATUS_TESTING   = 1
	PROJECT_STATUS_SUCCESS   = 2
	PROJECT_STATUS_FAIL      = 3
)

// More setting in http://beego.me/docs/mvc/model/models.md
type Project struct {
	Id          int64  `orm:"pk;auto"`
	UserName    string `orm:"size(1024);null"`
	ProjectName string `orm:"size(1024);null"`
	RepoUrl     string `orm:"size(1024);null"`
	Status      int    `orm:"null"`
}

// For advanced usage in http://beego.me/docs/mvc/model/query.md#all
func GetAllProjects() []*Project {
	o := orm.NewOrm()

	var projects []*Project
	num, err := o.QueryTable("project").All(&projects)
	fmt.Printf("Returned Rows Num: %s, %s", num, err)
	return projects
}

func GetProjectWithId(projectId int64) Project {
	o := orm.NewOrm()

	project := Project{Id: projectId}
	err := o.Read(&project)
	fmt.Printf("ERR: %v\n", err)
	return project
}

// Read or create the project in database
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

// For more usage in http://beego.me/docs/mvc/model/overview.md
func AddProjectWithNameUrl(projectName string, repoUrl string) error {
	o := orm.NewOrm()

	project := Project{ProjectName: projectName, RepoUrl: repoUrl}
	_, err := o.Insert(&project)
	return err
}

func AddProject(project Project) error {
	o := orm.NewOrm()

	_, err := o.Insert(&project)
	fmt.Println(err)
	return err
}
