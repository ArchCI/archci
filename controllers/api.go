package controllers

import (
	"net/http"

	"crypto/rand"
	"encoding/base64"

	"fmt"

	log "github.com/Sirupsen/logrus"

	"github.com/astaxie/beego"

	"golang.org/x/oauth2"

	//"github.com/google/go-github/github"

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


func (c *ApiController) LoginGithub() {
	log.Debug("Start to login github")

	oauthCfg := &oauth2.Config{
		ClientID:     "893d4b7aaab4c8c7bf52",
		ClientSecret: "5df2d172df4e497f3eb66ef3a1ff3b6eb719a949",
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

	oauthCfg := &oauth2.Config{
		ClientID:     "893d4b7aaab4c8c7bf52",
		ClientSecret: "5df2d172df4e497f3eb66ef3a1ff3b6eb719a949",
		Endpoint: oauth2.Endpoint{
			AuthURL:  githubAuthorizeUrl,
			TokenURL: githubTokenUrl,
		},
		RedirectURL: redirectUrl,
		Scopes:      []string{"repo"},
	}

	/*
	if r.URL.Query().Get("state") != session.Values["state"] {
		fmt.Println("no state match; possible csrf OR cookies not enabled")
		return
	}
	*/

	tkn, err := oauthCfg.Exchange(oauth2.NoContext, r.URL.Query().Get("code"))
	if err != nil {
		fmt.Println("there was an issue getting your token")
		return
	}

	if !tkn.Valid() {
		fmt.Println("retreived invalid token")
		return
	}

	/*
	client := github.NewClient(oauthCfg.Client(oauth2.NoContext, tkn))
	*/

	/*
	user, _, err := client.Users.Get("")
	if err != nil {
		fmt.Println("error getting name")
		return
	}
	*/

	url := "/account"

	http.Redirect(c.Ctx.ResponseWriter, c.Ctx.Request, url, 302)

}

