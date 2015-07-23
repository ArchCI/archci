package controllers

import (
	//"database/sql"
	"fmt"
	_ "github.com/lib/pq"

	"github.com/astaxie/beego"
	"github.com/golang/glog"

	"encoding/json"
	"github.com/ArchCI/archci/models"
	"github.com/ArchCI/archci/redisutil"
)

type ApiController struct {
	beego.Controller
}

type GetBuildLogsIndexResponse struct {
	Log  string `json:"log"`
	Next bool   `json:"next`
}

/* Create acount */
func (c *ApiController) CreateAccount() {
	glog.Info("Creat accout")

	result := "{data: 1}"
	c.Ctx.WriteString(result)
}

// New build record in database
func (c *ApiController) NewBuild() {
	glog.Info("New build record")

	project := models.Project{}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &project); err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte("empty title"))
		fmt.Println(err)
		return
	}

	models.AddBuildWithProject(project)
}

// Get all builds from database
func (c *ApiController) GetBuildsAll() {
	glog.Info("Get all builds")

	builds := models.GetAllBuilds()

	c.Data["json"] = builds
	c.ServeJson()
}

// Get builds with project name
func (c *ApiController) GetBuildsWithProjectName() {
	glog.Info("Get builds with project name")

	projectName := c.GetString(":projectName")

	builds := models.GetBuildsWithProjectName(projectName)

	c.Data["json"] = builds
	c.ServeJson()
}

// Get one build with build id
func (c *ApiController) GetBuild() {
	glog.Info("Get build")

	buildId, _ := c.GetInt64(":buildId")

	build := models.GetBuildWithId(buildId)

	c.Data["json"] = build
	c.ServeJson()
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

/* Put build log with index */
func (c *ApiController) PutBuildLogsIndex() {
	glog.Info("Put build log part")

	result := "{data: 1}"
	c.Ctx.WriteString(result)
}

/* Get build log with index */
func (c *ApiController) GetBuildLogsIndex() {
	glog.Info("Get build log with index")

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

	// Throw error if index is larger than current

	// Handle when get the index is more than current but it's not total

	response := &GetBuildLogsIndexResponse{
		Log:  log,
		Next: next}

	c.Data["json"] = response
	c.ServeJson()

}

// Get all logs of the build
func (c *ApiController) GetBuildLogsAll() {
	glog.Info("Get all build logs")

	//buildId := c.GetString(":buildId")
	//field := 0
	//result := redisutil.HgetString(buildId, field)
	//c.Ctx.WriteString(result)

	// TODO(tobe): change to hgetall  from redis
	mystruct := `{0: "apt-get install", 1: "go test"}`

	c.Data["json"] = &mystruct
	c.ServeJson()
}

// New project
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

// Get all projects from database
func (c *ApiController) GetProjectsAll() {
	glog.Info("Get all projects")

	projects := models.GetAllProjects()

	c.Data["json"] = projects
	c.ServeJson()
}

// Get one project with project id
func (c *ApiController) GetProject() {
	glog.Info("Get project")

	projectId, _ := c.GetInt64(":projectId")

	project := models.GetProjectWithId(projectId)

	c.Data["json"] = project
	c.ServeJson()
}

// Get all workers
func (c *ApiController) GetWorkersAll() {
	glog.Info("Get all workers")

	workers := models.GetAllWorkers()

	c.Data["json"] = workers
	c.ServeJson()
}

// Get all worker with this status
func (c *ApiController) GetWorkersAllStatus() {
	glog.Info("Get all workers with status")

	status, _ := c.GetInt(":status")
	workers := models.GetAllWorkersWithStatus(status)

	c.Data["json"] = workers
	c.ServeJson()
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

// Process github push hook
func (c *ApiController) HookGithubPush() {
	glog.Info("Trigger github push hook")

	/*
	{
	  "zen": "Mind your words, they are important.",
	  "hook_id": 5371014,
	  "hook": {
		"url": "https://api.github.com/repos/ArchCI/success-test/hooks/5371014",
		"test_url": "https://api.github.com/repos/ArchCI/success-test/hooks/5371014/test",
		"ping_url": "https://api.github.com/repos/ArchCI/success-test/hooks/5371014/pings",
		"id": 5371014,
		"name": "web",
		"active": true,
		"events": [
		  "push"
		],
		"config": {
		  "url": "http://192.168.1.113:10010/v1/hook/github/push",
		  "content_type": "json",
		  "insecure_ssl": "0",
		  "secret": "********"
		},
		"last_response": {
		  "code": null,
		  "status": "unused",
		  "message": null
		},
		"updated_at": "2015-07-22T14:48:22Z",
		"created_at": "2015-07-22T14:48:22Z"
	  },
	  "repository": {
		"id": 39022734,
		"name": "success-test",
		"full_name": "ArchCI/success-test",
		"owner": {
		  "login": "ArchCI",
		  "id": 12673804,
		  "avatar_url": "https://avatars.githubusercontent.com/u/12673804?v=3",
		  "gravatar_id": "",
		  "url": "https://api.github.com/users/ArchCI",
		  "html_url": "https://github.com/ArchCI",
		  "followers_url": "https://api.github.com/users/ArchCI/followers",
		  "following_url": "https://api.github.com/users/ArchCI/following{/other_user}",
		  "gists_url": "https://api.github.com/users/ArchCI/gists{/gist_id}",
		  "starred_url": "https://api.github.com/users/ArchCI/starred{/owner}{/repo}",
		  "subscriptions_url": "https://api.github.com/users/ArchCI/subscriptions",
		  "organizations_url": "https://api.github.com/users/ArchCI/orgs",
		  "repos_url": "https://api.github.com/users/ArchCI/repos",
		  "events_url": "https://api.github.com/users/ArchCI/events{/privacy}",
		  "received_events_url": "https://api.github.com/users/ArchCI/received_events",
		  "type": "Organization",
		  "site_admin": false
		},
		"private": false,
		"html_url": "https://github.com/ArchCI/success-test",
		"description": "Success test project for ArchCI",
		"fork": false,
		"url": "https://api.github.com/repos/ArchCI/success-test",
		"forks_url": "https://api.github.com/repos/ArchCI/success-test/forks",
		"keys_url": "https://api.github.com/repos/ArchCI/success-test/keys{/key_id}",
		"collaborators_url": "https://api.github.com/repos/ArchCI/success-test/collaborators{/collaborator}",
		"teams_url": "https://api.github.com/repos/ArchCI/success-test/teams",
		"hooks_url": "https://api.github.com/repos/ArchCI/success-test/hooks",
		"issue_events_url": "https://api.github.com/repos/ArchCI/success-test/issues/events{/number}",
		"events_url": "https://api.github.com/repos/ArchCI/success-test/events",
		"assignees_url": "https://api.github.com/repos/ArchCI/success-test/assignees{/user}",
		"branches_url": "https://api.github.com/repos/ArchCI/success-test/branches{/branch}",
		"tags_url": "https://api.github.com/repos/ArchCI/success-test/tags",
		"blobs_url": "https://api.github.com/repos/ArchCI/success-test/git/blobs{/sha}",
		"git_tags_url": "https://api.github.com/repos/ArchCI/success-test/git/tags{/sha}",
		"git_refs_url": "https://api.github.com/repos/ArchCI/success-test/git/refs{/sha}",
		"trees_url": "https://api.github.com/repos/ArchCI/success-test/git/trees{/sha}",
		"statuses_url": "https://api.github.com/repos/ArchCI/success-test/statuses/{sha}",
		"languages_url": "https://api.github.com/repos/ArchCI/success-test/languages",
		"stargazers_url": "https://api.github.com/repos/ArchCI/success-test/stargazers",
		"contributors_url": "https://api.github.com/repos/ArchCI/success-test/contributors",
		"subscribers_url": "https://api.github.com/repos/ArchCI/success-test/subscribers",
		"subscription_url": "https://api.github.com/repos/ArchCI/success-test/subscription",
		"commits_url": "https://api.github.com/repos/ArchCI/success-test/commits{/sha}",
		"git_commits_url": "https://api.github.com/repos/ArchCI/success-test/git/commits{/sha}",
		"comments_url": "https://api.github.com/repos/ArchCI/success-test/comments{/number}",
		"issue_comment_url": "https://api.github.com/repos/ArchCI/success-test/issues/comments{/number}",
		"contents_url": "https://api.github.com/repos/ArchCI/success-test/contents/{+path}",
		"compare_url": "https://api.github.com/repos/ArchCI/success-test/compare/{base}...{head}",
		"merges_url": "https://api.github.com/repos/ArchCI/success-test/merges",
		"archive_url": "https://api.github.com/repos/ArchCI/success-test/{archive_format}{/ref}",
		"downloads_url": "https://api.github.com/repos/ArchCI/success-test/downloads",
		"issues_url": "https://api.github.com/repos/ArchCI/success-test/issues{/number}",
		"pulls_url": "https://api.github.com/repos/ArchCI/success-test/pulls{/number}",
		"milestones_url": "https://api.github.com/repos/ArchCI/success-test/milestones{/number}",
		"notifications_url": "https://api.github.com/repos/ArchCI/success-test/notifications{?since,all,participating}",
		"labels_url": "https://api.github.com/repos/ArchCI/success-test/labels{/name}",
		"releases_url": "https://api.github.com/repos/ArchCI/success-test/releases{/id}",
		"created_at": "2015-07-13T16:06:28Z",
		"updated_at": "2015-07-22T02:12:02Z",
		"pushed_at": "2015-07-22T02:12:02Z",
		"git_url": "git://github.com/ArchCI/success-test.git",
		"ssh_url": "git@github.com:ArchCI/success-test.git",
		"clone_url": "https://github.com/ArchCI/success-test.git",
		"svn_url": "https://github.com/ArchCI/success-test",
		"homepage": null,
		"size": 120,
		"stargazers_count": 0,
		"watchers_count": 0,
		"language": "Go",
		"has_issues": true,
		"has_downloads": true,
		"has_wiki": true,
		"has_pages": false,
		"forks_count": 0,
		"mirror_url": null,
		"open_issues_count": 0,
		"forks": 0,
		"open_issues": 0,
		"watchers": 0,
		"default_branch": "master"
	  },
	  "sender": {
		"login": "tobegit3hub",
		"id": 2715000,
		"avatar_url": "https://avatars.githubusercontent.com/u/2715000?v=3",
		"gravatar_id": "",
		"url": "https://api.github.com/users/tobegit3hub",
		"html_url": "https://github.com/tobegit3hub",
		"followers_url": "https://api.github.com/users/tobegit3hub/followers",
		"following_url": "https://api.github.com/users/tobegit3hub/following{/other_user}",
		"gists_url": "https://api.github.com/users/tobegit3hub/gists{/gist_id}",
		"starred_url": "https://api.github.com/users/tobegit3hub/starred{/owner}{/repo}",
		"subscriptions_url": "https://api.github.com/users/tobegit3hub/subscriptions",
		"organizations_url": "https://api.github.com/users/tobegit3hub/orgs",
		"repos_url": "https://api.github.com/users/tobegit3hub/repos",
		"events_url": "https://api.github.com/users/tobegit3hub/events{/privacy}",
		"received_events_url": "https://api.github.com/users/tobegit3hub/received_events",
		"type": "User",
		"site_admin": false
	  }
	}
	*/

	/*
{
  "object_kind": "push",
  "before": "95790bf891e76fee5e1747ab589903a6a1f80f22",
  "after": "da1560886d4f094c3e6c9ef40349f7d38b5d27d7",
  "ref": "refs/heads/master",
  "user_id": 4,
  "user_name": "John Smith",
  "user_email": "john@example.com",
  "project_id": 15,
  "repository": {
    "name": "Diaspora",
    "url": "git@example.com:mike/diasporadiaspora.git",
    "description": "",
    "homepage": "http://example.com/mike/diaspora", 
    "git_http_url":"http://example.com/mike/diaspora.git",
    "git_ssh_url":"git@example.com:mike/diaspora.git",
    "visibility_level":0
  },
  "commits": [
    {
      "id": "b6568db1bc1dcd7f8b4d5a946b0b91f9dacd7327",
      "message": "Update Catalan translation to e38cb41.",
      "timestamp": "2011-12-12T14:27:31+02:00",
      "url": "http://example.com/mike/diaspora/commit/b6568db1bc1dcd7f8b4d5a946b0b91f9dacd7327",
      "author": {
        "name": "Jordi Mallach",
        "email": "jordi@softcatala.org"
      }
    },
    {
      "id": "da1560886d4f094c3e6c9ef40349f7d38b5d27d7",
      "message": "fixed readme",
      "timestamp": "2012-01-03T23:36:29+02:00",
      "url": "http://example.com/mike/diaspora/commit/da1560886d4f094c3e6c9ef40349f7d38b5d27d7",
      "author": {
        "name": "GitLab dev user",
        "email": "gitlabdev@dv6700.(none)"
      }
    }
  ],
  "total_commits_count": 4
}
*/
	result := "{github_hook: 111}"
	c.Ctx.WriteString(result)
}
