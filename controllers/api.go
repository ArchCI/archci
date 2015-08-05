package controllers

import (
	"net/http"

	"crypto/rand"
	"encoding/base64"

	"fmt"

	log "github.com/Sirupsen/logrus"

	"github.com/astaxie/beego"

	"golang.org/x/oauth2"

	"github.com/google/go-github/github"
)

// ApiController is the custom controller to provide APIs.
type ApiController struct {
	beego.Controller
}

const (
	githubAuthorizeUrl = "https://github.com/login/oauth/authorize"
	githubTokenUrl     = "https://github.com/login/oauth/access_token"
	redirectUrl        = ""
)

var oauthCfg *oauth2.Config
var tkn *oauth2.Token

func (c *ApiController) LoginGithub() {
	log.Debug("Start to login github")

	oauthCfg = &oauth2.Config{
		ClientID:     "69028d944609d7243cbf",
		ClientSecret: "efe2f3d0a547f0840fd17411e729d20f02ece615",
		Endpoint: oauth2.Endpoint{
			AuthURL:  githubAuthorizeUrl,
			TokenURL: githubTokenUrl,
		},
		RedirectURL: redirectUrl,
		Scopes:      []string{"repo"},
	}

	b := make([]byte, 16)
	rand.Read(b)

	state := base64.URLEncoding.EncodeToString(b)

	url := oauthCfg.AuthCodeURL(state)

	http.Redirect(c.Ctx.ResponseWriter, c.Ctx.Request, url, 302)
}

func (c *ApiController) LoginGithubCallback() {
	fmt.Println("Start login github callback")

	r := c.Ctx.Request

	/*
		if r.URL.Query().Get("state") != session.Values["state"] {
			fmt.Println("no state match; possible csrf OR cookies not enabled")
			return
		}
	*/

	tkn, _ = oauthCfg.Exchange(oauth2.NoContext, r.URL.Query().Get("code"))
	/*
		if err != nil {
			fmt.Println("there was an issue getting your token")
			//return
		}
	*/

	fmt.Println(tkn)

	if !tkn.Valid() {
		fmt.Println("retreived invalid token")
		//return
	}

	client := github.NewClient(oauthCfg.Client(oauth2.NoContext, tkn))

	user, _, err := client.Users.Get("")
	if err != nil {
		fmt.Println("error getting name")
		//return
	}

	fmt.Println(user)

	url := "/account"

	http.Redirect(c.Ctx.ResponseWriter, c.Ctx.Request, url, 302)

}

func (c *ApiController) GetAccount() {
	fmt.Println("Start get account info")

	client := github.NewClient(oauthCfg.Client(oauth2.NoContext, tkn))

	user, _, err := client.Users.Get("")
	if err != nil {
		fmt.Println("error getting name")
		//return
	}

	/*
	   {
	   login: "tobegit3hub",
	   id: 2715000,
	   avatar_url: "https://avatars.githubusercontent.com/u/2715000?v=3",
	   html_url: "https://github.com/tobegit3hub",
	   gravatar_id: "",
	   name: "tobe",
	   company: "UnitedStack Inc",
	   blog: "http://chendihao.cn",
	   location: "Beijing",
	   email: "tobeg3oogle@gmail.com",
	   hireable: false,
	   public_repos: 161,
	   public_gists: 1,
	   followers: 155,
	   following: 68,
	   created_at: "2012-11-03T13:17:18Z",
	   updated_at: "2015-08-04T19:01:39Z",
	   type: "User",
	   site_admin: false,
	   url: "https://api.github.com/users/tobegit3hub",
	   events_url: "https://api.github.com/users/tobegit3hub/events{/privacy}",
	   following_url: "https://api.github.com/users/tobegit3hub/following{/other_user}",
	   followers_url: "https://api.github.com/users/tobegit3hub/followers",
	   gists_url: "https://api.github.com/users/tobegit3hub/gists{/gist_id}",
	   organizations_url: "https://api.github.com/users/tobegit3hub/orgs",
	   received_events_url: "https://api.github.com/users/tobegit3hub/received_events",
	   repos_url: "https://api.github.com/users/tobegit3hub/repos",
	   starred_url: "https://api.github.com/users/tobegit3hub/starred{/owner}{/repo}",
	   subscriptions_url: "https://api.github.com/users/tobegit3hub/subscriptions"
	   }
	*/

	c.Data["json"] = user
	c.ServeJson()

}

func (c *ApiController) GetAccountProjects() {

	client := github.NewClient(oauthCfg.Client(oauth2.NoContext, tkn))

	repos, _, err := client.Repositories.List("", nil)
	if err != nil {
		fmt.Println("error getting repos")
		//return
	}

	/*
	   [
	   {
	   id: 36988569,
	   owner: {
	   login: "ArchCI",
	   id: 12673804,
	   avatar_url: "https://avatars.githubusercontent.com/u/12673804?v=3",
	   html_url: "https://github.com/ArchCI",
	   gravatar_id: "",
	   type: "Organization",
	   site_admin: false,
	   url: "https://api.github.com/users/ArchCI",
	   events_url: "https://api.github.com/users/ArchCI/events{/privacy}",
	   following_url: "https://api.github.com/users/ArchCI/following{/other_user}",
	   followers_url: "https://api.github.com/users/ArchCI/followers",
	   gists_url: "https://api.github.com/users/ArchCI/gists{/gist_id}",
	   organizations_url: "https://api.github.com/users/ArchCI/orgs",
	   received_events_url: "https://api.github.com/users/ArchCI/received_events",
	   repos_url: "https://api.github.com/users/ArchCI/repos",
	   starred_url: "https://api.github.com/users/ArchCI/starred{/owner}{/repo}",
	   subscriptions_url: "https://api.github.com/users/ArchCI/subscriptions"
	   },
	   name: "aci",
	   full_name: "ArchCI/aci",
	   description: "Command-line client for ArchCI",
	   homepage: "http://archci.com",
	   default_branch: "master",
	   created_at: "2015-06-06T17:21:16Z",
	   pushed_at: "2015-06-14T01:00:11Z",
	   updated_at: "2015-06-28T13:58:55Z",
	   html_url: "https://github.com/ArchCI/aci",
	   clone_url: "https://github.com/ArchCI/aci.git",
	   git_url: "git://github.com/ArchCI/aci.git",
	   ssh_url: "git@github.com:ArchCI/aci.git",
	   svn_url: "https://github.com/ArchCI/aci",
	   language: "Go",
	   fork: false,
	   forks_count: 0,
	   open_issues_count: 0,
	   stargazers_count: 1,
	   watchers_count: 1,
	   size: 168,
	   permissions: {
	   admin: true,
	   pull: true,
	   push: true
	   },
	   private: false,
	   has_issues: true,
	   has_wiki: true,
	   has_downloads: true,
	   team_id: null,
	   url: "https://api.github.com/repos/ArchCI/aci",
	   archive_url: "https://api.github.com/repos/ArchCI/aci/{archive_format}{/ref}",
	   assignees_url: "https://api.github.com/repos/ArchCI/aci/assignees{/user}",
	   blobs_url: "https://api.github.com/repos/ArchCI/aci/git/blobs{/sha}",
	   branches_url: "https://api.github.com/repos/ArchCI/aci/branches{/branch}",
	   collaborators_url: "https://api.github.com/repos/ArchCI/aci/collaborators{/collaborator}",
	   comments_url: "https://api.github.com/repos/ArchCI/aci/comments{/number}",
	   commits_url: "https://api.github.com/repos/ArchCI/aci/commits{/sha}",
	   compare_url: "https://api.github.com/repos/ArchCI/aci/compare/{base}...{head}",
	   contents_url: "https://api.github.com/repos/ArchCI/aci/contents/{+path}",
	   contributors_url: "https://api.github.com/repos/ArchCI/aci/contributors",
	   downloads_url: "https://api.github.com/repos/ArchCI/aci/downloads",
	   events_url: "https://api.github.com/repos/ArchCI/aci/events",
	   forks_url: "https://api.github.com/repos/ArchCI/aci/forks",
	   git_commits_url: "https://api.github.com/repos/ArchCI/aci/git/commits{/sha}",
	   git_refs_url: "https://api.github.com/repos/ArchCI/aci/git/refs{/sha}",
	   git_tags_url: "https://api.github.com/repos/ArchCI/aci/git/tags{/sha}",
	   hooks_url: "https://api.github.com/repos/ArchCI/aci/hooks",
	   issue_comment_url: "https://api.github.com/repos/ArchCI/aci/issues/comments{/number}",
	   issue_events_url: "https://api.github.com/repos/ArchCI/aci/issues/events{/number}",
	   issues_url: "https://api.github.com/repos/ArchCI/aci/issues{/number}",
	   keys_url: "https://api.github.com/repos/ArchCI/aci/keys{/key_id}",
	   labels_url: "https://api.github.com/repos/ArchCI/aci/labels{/name}",
	   languages_url: "https://api.github.com/repos/ArchCI/aci/languages",
	   merges_url: "https://api.github.com/repos/ArchCI/aci/merges",
	   milestones_url: "https://api.github.com/repos/ArchCI/aci/milestones{/number}",
	   notifications_url: "https://api.github.com/repos/ArchCI/aci/notifications{?since,all,participating}",
	   pulls_url: "https://api.github.com/repos/ArchCI/aci/pulls{/number}",
	   releases_url: "https://api.github.com/repos/ArchCI/aci/releases{/id}",
	   stargazers_url: "https://api.github.com/repos/ArchCI/aci/stargazers",
	   statuses_url: "https://api.github.com/repos/ArchCI/aci/statuses/{sha}",
	   subscribers_url: "https://api.github.com/repos/ArchCI/aci/subscribers",
	   subscription_url: "https://api.github.com/repos/ArchCI/aci/subscription",
	   tags_url: "https://api.github.com/repos/ArchCI/aci/tags",
	   trees_url: "https://api.github.com/repos/ArchCI/aci/git/trees{/sha}",
	   teams_url: "https://api.github.com/repos/ArchCI/aci/teams"
	   }
	   ]
	*/
	c.Data["json"] = repos
	c.ServeJson()

}
