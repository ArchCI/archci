package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"

	"fmt"

	"github.com/ArchCI/archci/githubutil"
	"github.com/ArchCI/archci/gitlabutil"
)

const (
	BUILD_STATUS_NOT_START = 0
	BUILD_STATUS_BUILDING  = 1
	BUILD_STATUS_SUCCESS   = 2
	BUILD_STATUS_FAIL      = 3
	BUILD_STATUS_CANCELED  = 4
)

type Build struct {
	Id          int64 `orm:"pk;auto"`
	ProjectId   int64
	UserName    string    `orm:"size(1024);null"`
	ProjectName string    `orm:"size(1024);null"`
	RepoUrl     string    `orm:"size(1024);null"`
	Branch      string    `orm:"size(1024);null"`
	Commit      string    `orm:"size(1024);null"`
	CommitTime  time.Time `orm:"null"`
	Committer   string    `orm:"size(1024);null"`
	BuildTime   time.Time `orm:"null"`
	FinishTime  time.Time `orm:"null"`
	Worker      string    `orm:"size(1024);null"`
	Status      int       `orm:"null"`
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

func AddGithubBuild(projectId int64, data githubutil.GithubPushHook) error {
	o := orm.NewOrm()

	// TODO(tobe): could not get commit id, commit time and committer

	build := Build{
		UserName:    data.Repository.Owner.Login,
		ProjectId:   projectId,
		ProjectName: data.Repository.Name,
		RepoUrl:     data.Repository.URL,
		Branch:      data.Repository.DefaultBranch,
		Commit:      "ffffffffff",
		CommitTime:  time.Now(),
		Committer:   "unknown",
		BuildTime:   time.Now(),
		FinishTime:  time.Now(),
		Status:      BUILD_STATUS_NOT_START}

	_, err := o.Insert(&build)
	fmt.Println("ERR: %v\n", err)
	return err
}

func AddGitlabBuild(projectId int64, data gitlabutil.GitlabPushHook) error {
	o := orm.NewOrm()

	// TODO(tobe): need to find project id by user name and project name
	// TODO(tobe): no branch data from gitlab webhook
	// TODO(tobe): should choose the latest commit
	// TODO(tobe): build time and finish time should be null

	build := Build{
		UserName:    data.UserName,
		ProjectId:   projectId,
		ProjectName: data.Repository.Name,
		RepoUrl:     data.Repository.URL,
		Branch:      "master",
		Commit:      data.Commits[0].ID,
		CommitTime:  data.Commits[0].Timestamp,
		Committer:   data.Commits[0].Author.Name,
		BuildTime:   time.Now(),
		FinishTime:  time.Now(),
		Status:      BUILD_STATUS_NOT_START}

	_, err := o.Insert(&build)
	fmt.Println("ERR: %v\n", err)
	return err
}

func AddBuild(projectName string, branch string, commit string, commitTime time.Time, committer string) error {
	o := orm.NewOrm()

	build := Build{ProjectName: projectName, Branch: branch, Commit: commit, CommitTime: commitTime, Committer: committer}
	_, err := o.Insert(&build)
	return err
}
