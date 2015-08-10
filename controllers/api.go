package controllers

import (
	"net/http"

	"crypto/rand"
	"encoding/base64"

	"fmt"
	"time"

	log "github.com/Sirupsen/logrus"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"

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

type GithubOauthUser struct {
	Login             string    `json:"login"`
	ID                int       `json:"id"`
	AvatarURL         string    `json:"avatar_url"`
	HTMLURL           string    `json:"html_url"`
	GravatarID        string    `json:"gravatar_id"`
	Name              string    `json:"name"`
	Company           string    `json:"company"`
	Blog              string    `json:"blog"`
	Location          string    `json:"location"`
	Email             string    `json:"email"`
	Hireable          bool      `json:"hireable"`
	PublicRepos       int       `json:"public_repos"`
	PublicGists       int       `json:"public_gists"`
	Followers         int       `json:"followers"`
	Following         int       `json:"following"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	Type              string    `json:"type"`
	SiteAdmin         bool      `json:"site_admin"`
	URL               string    `json:"url"`
	EventsURL         string    `json:"events_url"`
	FollowingURL      string    `json:"following_url"`
	FollowersURL      string    `json:"followers_url"`
	GistsURL          string    `json:"gists_url"`
	OrganizationsURL  string    `json:"organizations_url"`
	ReceivedEventsURL string    `json:"received_events_url"`
	ReposURL          string    `json:"repos_url"`
	StarredURL        string    `json:"starred_url"`
	SubscriptionsURL  string    `json:"subscriptions_url"`
}

type GithubOauthRepositories []struct {
	ID    int `json:"id"`
	Owner struct {
		Login             string `json:"login"`
		ID                int    `json:"id"`
		AvatarURL         string `json:"avatar_url"`
		HTMLURL           string `json:"html_url"`
		GravatarID        string `json:"gravatar_id"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
		URL               string `json:"url"`
		EventsURL         string `json:"events_url"`
		FollowingURL      string `json:"following_url"`
		FollowersURL      string `json:"followers_url"`
		GistsURL          string `json:"gists_url"`
		OrganizationsURL  string `json:"organizations_url"`
		ReceivedEventsURL string `json:"received_events_url"`
		ReposURL          string `json:"repos_url"`
		StarredURL        string `json:"starred_url"`
		SubscriptionsURL  string `json:"subscriptions_url"`
	} `json:"owner"`
	Name            string    `json:"name"`
	FullName        string    `json:"full_name"`
	Description     string    `json:"description"`
	Homepage        string    `json:"homepage"`
	DefaultBranch   string    `json:"default_branch"`
	CreatedAt       time.Time `json:"created_at"`
	PushedAt        time.Time `json:"pushed_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	HTMLURL         string    `json:"html_url"`
	CloneURL        string    `json:"clone_url"`
	GitURL          string    `json:"git_url"`
	SSHURL          string    `json:"ssh_url"`
	SvnURL          string    `json:"svn_url"`
	Language        string    `json:"language"`
	Fork            bool      `json:"fork"`
	ForksCount      int       `json:"forks_count"`
	OpenIssuesCount int       `json:"open_issues_count"`
	StargazersCount int       `json:"stargazers_count"`
	WatchersCount   int       `json:"watchers_count"`
	Size            int       `json:"size"`
	Permissions     struct {
		Admin bool `json:"admin"`
		Pull  bool `json:"pull"`
		Push  bool `json:"push"`
	} `json:"permissions"`
	Private          bool        `json:"private"`
	HasIssues        bool        `json:"has_issues"`
	HasWiki          bool        `json:"has_wiki"`
	HasDownloads     bool        `json:"has_downloads"`
	TeamID           interface{} `json:"team_id"`
	URL              string      `json:"url"`
	ArchiveURL       string      `json:"archive_url"`
	AssigneesURL     string      `json:"assignees_url"`
	BlobsURL         string      `json:"blobs_url"`
	BranchesURL      string      `json:"branches_url"`
	CollaboratorsURL string      `json:"collaborators_url"`
	CommentsURL      string      `json:"comments_url"`
	CommitsURL       string      `json:"commits_url"`
	CompareURL       string      `json:"compare_url"`
	ContentsURL      string      `json:"contents_url"`
	ContributorsURL  string      `json:"contributors_url"`
	DownloadsURL     string      `json:"downloads_url"`
	EventsURL        string      `json:"events_url"`
	ForksURL         string      `json:"forks_url"`
	GitCommitsURL    string      `json:"git_commits_url"`
	GitRefsURL       string      `json:"git_refs_url"`
	GitTagsURL       string      `json:"git_tags_url"`
	HooksURL         string      `json:"hooks_url"`
	IssueCommentURL  string      `json:"issue_comment_url"`
	IssueEventsURL   string      `json:"issue_events_url"`
	IssuesURL        string      `json:"issues_url"`
	KeysURL          string      `json:"keys_url"`
	LabelsURL        string      `json:"labels_url"`
	LanguagesURL     string      `json:"languages_url"`
	MergesURL        string      `json:"merges_url"`
	MilestonesURL    string      `json:"milestones_url"`
	NotificationsURL string      `json:"notifications_url"`
	PullsURL         string      `json:"pulls_url"`
	ReleasesURL      string      `json:"releases_url"`
	StargazersURL    string      `json:"stargazers_url"`
	StatusesURL      string      `json:"statuses_url"`
	SubscribersURL   string      `json:"subscribers_url"`
	SubscriptionURL  string      `json:"subscription_url"`
	TagsURL          string      `json:"tags_url"`
	TreesURL         string      `json:"trees_url"`
	TeamsURL         string      `json:"teams_url"`
}

var globalSessions *session.Manager

func init() {
	globalSessions, _ = session.NewManager("memory", `{"cookieName":"gosessionid", "enableSetCookie,omitempty": true, "gclifetime":3600, "maxLifetime": 3600, "secure": false, "sessionIDHashFunc": "sha1", "sessionIDHashKey": "", "cookieLifeTime": 3600, "providerConfig": ""}`)
	go globalSessions.GC()
}

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

	w := c.Ctx.ResponseWriter

	sess, _ := globalSessions.SessionStart(w, r)
	defer sess.SessionRelease(w)

	//token := sess.Get("token")
	//fmt.Println(token)

	sess.Set("token", tkn)

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

// GetAccountInfo return the info of github user.
func (c *ApiController) GetAccountInfo() {
	fmt.Println("Start get account info")

	//user := GithubOauthUser{Name:"tobe", HTMLURL: "https://github.com/tobegit3hub", Company: "UnitedStack Inc"}

	client := github.NewClient(oauthCfg.Client(oauth2.NoContext, tkn))

	user, _, err := client.Users.Get("")
	if err != nil {
		fmt.Println("error getting name")
		//return
	}

	/*
		{
		"login": "tobegit3hub",
		"id": 2715000,
		"avatar_url": "https://avatars.githubusercontent.com/u/2715000?v=3",
		"html_url": "https://github.com/tobegit3hub",
		"gravatar_id": "",
		"name": "tobe",
		"company": "UnitedStack Inc",
		"blog": "http://chendihao.cn",
		"location": "Beijing",
		"email": "tobeg3oogle@gmail.com",
		"hireable": false,
		"public_repos": 161,
		"public_gists": 1,
		"followers": 155,
		"following": 68,
		"created_at": "2012-11-03T13:17:18Z",
		"updated_at": "2015-08-04T19:01:39Z",
		"type": "User",
		"site_admin": false,
		"url": "https://api.github.com/users/tobegit3hub",
		"events_url": "https://api.github.com/users/tobegit3hub/events{/privacy}",
		"following_url": "https://api.github.com/users/tobegit3hub/following{/other_user}",
		"followers_url": "https://api.github.com/users/tobegit3hub/followers",
		"gists_url": "https://api.github.com/users/tobegit3hub/gists{/gist_id}",
		"organizations_url": "https://api.github.com/users/tobegit3hub/orgs",
		"received_events_url": "https://api.github.com/users/tobegit3hub/received_events",
		"repos_url": "https://api.github.com/users/tobegit3hub/repos",
		"starred_url": "https://api.github.com/users/tobegit3hub/starred{/owner}{/repo}",
		"subscriptions_url": "https://api.github.com/users/tobegit3hub/subscriptions"
		}
	*/

	c.Data["json"] = user
	c.ServeJson()

}

// GetAccountOrganizations return the organizations of github user.
func (c *ApiController) GetAccountOrganizations() {
	fmt.Println("Start get account organizations")

	client := github.NewClient(oauthCfg.Client(oauth2.NoContext, tkn))

	organizations, _, err := client.Organizations.List("", nil)
	if err != nil {
		fmt.Println("error getting organizations")
		fmt.Println(nil)
	}

	c.Data["json"] = organizations
	c.ServeJson()
}

// GetAccountProjects returns the projects of github user.
func (c *ApiController) GetAccountProjects() {

	//projects := GithubOauthRepositories{{FullName:"ArchCI/archci", HTMLURL:"https://github.com/ArchCI/archci", Language: "Go", Description: "The best CI system"}, {FullName:"ArchCI/simple-worker", Homepage:"https://github.com/ArchCI/simple-worker", Language:"Go", Description: "The client of ArchCI"}}

	/*
		r := c.Ctx.Request
		w := c.Ctx.ResponseWriter

		token := tkn

		sess, _ := globalSessions.SessionStart(w, r)
		defer sess.SessionRelease(w)

		if sess.Get("token") == nil {
			token = sess.Get("token").(*oauth2.Token)

			fmt.Println("hit token")

		}
	*/

	client := github.NewClient(oauthCfg.Client(oauth2.NoContext, tkn))

	projects, _, err := client.Repositories.List("", nil)
	if err != nil {
		fmt.Println("error getting repos")
		//return
	}

	/*
		[
		{
		"id": 36988569,
		"owner": {
		"login": "ArchCI",
		"id": 12673804,
		"avatar_url": "https://avatars.githubusercontent.com/u/12673804?v=3",
		"html_url": "https://github.com/ArchCI",
		"gravatar_id": "",
		"type": "Organization",
		"site_admin": false,
		"url": "https://api.github.com/users/ArchCI",
		"events_url": "https://api.github.com/users/ArchCI/events{/privacy}",
		"following_url": "https://api.github.com/users/ArchCI/following{/other_user}",
		"followers_url": "https://api.github.com/users/ArchCI/followers",
		"gists_url": "https://api.github.com/users/ArchCI/gists{/gist_id}",
		"organizations_url": "https://api.github.com/users/ArchCI/orgs",
		"received_events_url": "https://api.github.com/users/ArchCI/received_events",
		"repos_url": "https://api.github.com/users/ArchCI/repos",
		"starred_url": "https://api.github.com/users/ArchCI/starred{/owner}{/repo}",
		"subscriptions_url": "https://api.github.com/users/ArchCI/subscriptions"
		},
		"name": "aci",
		"full_name": "ArchCI/aci",
		"description": "Command-line client for ArchCI",
		"homepage": "http://archci.com",
		"default_branch": "master",
		"created_at": "2015-06-06T17:21:16Z",
		"pushed_at": "2015-06-14T01:00:11Z",
		"updated_at": "2015-06-28T13:58:55Z",
		"html_url": "https://github.com/ArchCI/aci",
		"clone_url": "https://github.com/ArchCI/aci.git",
		"git_url": "git://github.com/ArchCI/aci.git",
		"ssh_url": "git@github.com:ArchCI/aci.git",
		"svn_url": "https://github.com/ArchCI/aci",
		"language": "Go",
		"fork": false,
		"forks_count": 0,
		"open_issues_count": 0,
		"stargazers_count": 1,
		"watchers_count": 1,
		"size": 168,
		"permissions": {
		"admin": true,
		"pull": true,
		"push": true
		},
		"private": false,
		"has_issues": true,
		"has_wiki": true,
		"has_downloads": true,
		"team_id": null,
		"url": "https://api.github.com/repos/ArchCI/aci",
		"archive_url": "https://api.github.com/repos/ArchCI/aci/{archive_format}{/ref}",
		"assignees_url": "https://api.github.com/repos/ArchCI/aci/assignees{/user}",
		"blobs_url": "https://api.github.com/repos/ArchCI/aci/git/blobs{/sha}",
		"branches_url": "https://api.github.com/repos/ArchCI/aci/branches{/branch}",
		"collaborators_url": "https://api.github.com/repos/ArchCI/aci/collaborators{/collaborator}",
		"comments_url": "https://api.github.com/repos/ArchCI/aci/comments{/number}",
		"commits_url": "https://api.github.com/repos/ArchCI/aci/commits{/sha}",
		"compare_url": "https://api.github.com/repos/ArchCI/aci/compare/{base}...{head}",
		"contents_url": "https://api.github.com/repos/ArchCI/aci/contents/{+path}",
		"contributors_url": "https://api.github.com/repos/ArchCI/aci/contributors",
		"downloads_url": "https://api.github.com/repos/ArchCI/aci/downloads",
		"events_url": "https://api.github.com/repos/ArchCI/aci/events",
		"forks_url": "https://api.github.com/repos/ArchCI/aci/forks",
		"git_commits_url": "https://api.github.com/repos/ArchCI/aci/git/commits{/sha}",
		"git_refs_url": "https://api.github.com/repos/ArchCI/aci/git/refs{/sha}",
		"git_tags_url": "https://api.github.com/repos/ArchCI/aci/git/tags{/sha}",
		"hooks_url": "https://api.github.com/repos/ArchCI/aci/hooks",
		"issue_comment_url": "https://api.github.com/repos/ArchCI/aci/issues/comments{/number}",
		"issue_events_url": "https://api.github.com/repos/ArchCI/aci/issues/events{/number}",
		"issues_url": "https://api.github.com/repos/ArchCI/aci/issues{/number}",
		"keys_url": "https://api.github.com/repos/ArchCI/aci/keys{/key_id}",
		"labels_url": "https://api.github.com/repos/ArchCI/aci/labels{/name}",
		"languages_url": "https://api.github.com/repos/ArchCI/aci/languages",
		"merges_url": "https://api.github.com/repos/ArchCI/aci/merges",
		"milestones_url": "https://api.github.com/repos/ArchCI/aci/milestones{/number}",
		"notifications_url": "https://api.github.com/repos/ArchCI/aci/notifications{?since,all,participating}",
		"pulls_url": "https://api.github.com/repos/ArchCI/aci/pulls{/number}",
		"releases_url": "https://api.github.com/repos/ArchCI/aci/releases{/id}",
		"stargazers_url": "https://api.github.com/repos/ArchCI/aci/stargazers",
		"statuses_url": "https://api.github.com/repos/ArchCI/aci/statuses/{sha}",
		"subscribers_url": "https://api.github.com/repos/ArchCI/aci/subscribers",
		"subscription_url": "https://api.github.com/repos/ArchCI/aci/subscription",
		"tags_url": "https://api.github.com/repos/ArchCI/aci/tags",
		"trees_url": "https://api.github.com/repos/ArchCI/aci/git/trees{/sha}",
		"teams_url": "https://api.github.com/repos/ArchCI/aci/teams"
		}
		]
	*/

	c.Data["json"] = projects
	c.ServeJson()

}

// GetAccountToken returns github oauth token as string.
func (c *ApiController) GetAccountToken() {
	result := tkn.AccessToken
	c.Ctx.WriteString(result)
}
