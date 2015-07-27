package controllers

import (
	"fmt"
	_ "github.com/lib/pq"

	"github.com/golang/glog"

	"encoding/json"
	"github.com/ArchCI/archci/githubutil"
	"github.com/ArchCI/archci/gitlabutil"
	"github.com/ArchCI/archci/models"
)

// Process github push hook.
func (c *ApiController) TriggerGithubPushHook() {
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

	hook := githubutil.GithubPushHook{}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &hook); err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte("empty title"))
		fmt.Println(err)
	}

	projectId, _ := models.ReadOrCreateProject(hook.Repository.Owner.Login, hook.Repository.Name, hook.Repository.URL)

	models.AddGithubBuild(projectId, hook)

	return

}

// Trigger gitlab push hook.
func (c *ApiController) TriggerGitlabPushHook() {
	glog.Info("Trigger gitlab push hook")

	/*
		{
		"before":"dd80657a63bf284fc1b62c5242addac10c8ee68b",
		"after":"f9ccf83322e36793a244cf2da379e358154bdd62",
		"ref":"refs/heads/master",
		"user_id":72,
		"user_name":"chendihao",
		"project_id":412,
		"repository":{
			"name":"gitlab_hook",
			"url":"git@git.ustack.com:chendihao/gitlab_hook.git",
			"description":"","homepage":"http://git.ustack.com/chendihao/gitlab_hook"
		},
		"commits":[
			{
				"id":"f9ccf83322e36793a244cf2da379e358154bdd62",
				"message":"Add introduction in readme",
				"timestamp":"2015-07-24T02:00:34+00:00",
				"url":"http://git.ustack.com/chendihao/gitlab_hook/commit/f9ccf83322e36793a244cf2da379e358154bdd62",
				"author":{
					"name":"tobe",
					"email":"tobeg3oogle@gmail.com"
				}
			},{
				"id":"dd80657a63bf284fc1b62c5242addac10c8ee68b",
				"message":"Add readme",
				"timestamp":"2015-07-23T02:47:30+00:00",
				"url":"http://git.ustack.com/chendihao/gitlab_hook/commit/dd80657a63bf284fc1b62c5242addac10c8ee68b",
				"author":{
					"name":"tobe",
					"email":"tobeg3oogle@gmail.com"
				}
			}
		],
		"total_commits_count":2
		}
	*/

	hook := gitlabutil.GitlabPushHook{}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &hook); err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte("empty title"))
		fmt.Println(err)
	}

	projectId, _ := models.ReadOrCreateProject(hook.UserName, hook.Repository.Name, hook.Repository.URL)

	models.AddGitlabBuild(projectId, hook)

	return
}
