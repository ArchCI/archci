package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"

	"fmt"
)

const (
	BUILD_STATUS_NOT_START = 0
	BUILD_STATUS_BUILDING  = 1
	BUILD_STATUS_SUCCESS   = 2
	BUILD_STATUS_FAIL      = 3
	BUILD_STATUS_CANCELED  = 4

	PROJECT_STATUS_NEED_TEST = 0
	PROJECT_STATUS_TESTING   = 1
	PROJECT_STATUS_SUCCESS   = 2
	PROJECT_STATUS_FAIL      = 3
)

// More setting in http://beego.me/docs/mvc/model/models.md
type Project struct {
	Id          int64  `orm:"pk;auto"`
	UserName    string `orm:"size(1024)"`
	ProjectName string `orm:"size(1024)"`
	RepoUrl     string `orm:"size(1024)"`
	Status      int
}

type Build struct {
	Id          int64 `orm:"pk;auto"`
	ProjectId   int64
	UserName    string `orm:"size(1024)"`
	ProjectName string `orm:"size(1024)"`
	RepoUrl     string `orm:"size(1024)"`
	Branch      string `orm:"size(1024)"`
	Commit      string `orm:"size(1024"`
	CommitTime  time.Time
	Committer   string `orm:"size(1024)"`
	BuildTime   time.Time
	FinishTime  time.Time
	Worker      string `orm:"size(1024)"`
	Status      int
}

type Worker struct {
	Id         int64  `orm:"pk;auto"`
	Ip         string `orm:"size(1024)"`
	LastUpdate time.Time
	Status     int
}

func RegisterModels() {
	orm.RegisterModel(new(Project), new(Build), new(Worker))
}

func GetAllBuilds() []*Build {
	o := orm.NewOrm()

	var builds []*Build
	// o.QueryTable("build").Filter("name", "slene").All(&builds) to filter with build status
	num, err := o.QueryTable("build").All(&builds)
	fmt.Printf("Returned Rows Num: %s, %s", num, err)
	return builds
}

func GetBuildsWithProjectName(projectName string) []*Build {
	o := orm.NewOrm()

	var builds []*Build
	num, err := o.QueryTable("build").Filter("project_name", projectName).All(&builds)
	fmt.Printf("Returned Rows Num: %s, %s", num, err)
	return builds
}

func GetBuildWithId(buildId int64) Build {
	o := orm.NewOrm()

	build := Build{Id: buildId}
	err := o.Read(&build)
	fmt.Printf("ERR: %v\n", err)
	return build
}

func AddBuildWithProject(project Project) error {
	o := orm.NewOrm()
	build := Build{UserName: project.UserName, ProjectName: project.ProjectName, RepoUrl: project.RepoUrl, Branch: "master", BuildTime: time.Now(), CommitTime: time.Now()}
	_, err := o.Insert(&build)
	fmt.Printf("ERR: %v\n", err)
	return err
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

func GetAllWorkers() []*Worker {
	o := orm.NewOrm()

	var workers []*Worker
	num, err := o.QueryTable("worker").All(&workers)
	fmt.Printf("Returned Rows Num: %s, %s", num, err)
	return workers
}

func GetAllWorkersWithStatus(status int) []*Worker {
	o := orm.NewOrm()

	var workers []*Worker
	num, err := o.QueryTable("worker").Filter("status", status).All(&workers)
	fmt.Printf("Returned Rows Num: %s, %s", num, err)
	return workers
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

func AddBuild(projectName string, branch string, commit string, commitTime time.Time, committer string) error {
	o := orm.NewOrm()

	build := Build{ProjectName: projectName, Branch: branch, Commit: commit, CommitTime: commitTime, Committer: committer}
	_, err := o.Insert(&build)
	return err
}

func TestOrm() {
	o := orm.NewOrm()

	project := Project{ProjectName: "tobegit3hub/seagull", RepoUrl: "https://github.com/tobegit3hub/seagull", Status: 0}

	// insert
	id, err := o.Insert(&project)
	fmt.Printf("ID: %d, ERR: %v\n", id, err)

	// update
	project.ProjectName = "ArchCI/archci"
	num, err := o.Update(&project)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)

	// read one
	u := Project{Id: project.Id}
	err = o.Read(&u)
	fmt.Printf("ERR: %v\n", err)

	// delete
	num, err = o.Delete(&u)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)
}
