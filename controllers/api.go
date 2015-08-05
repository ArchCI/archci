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

	tkn, err := oauthCfg.Exchange(oauth2.NoContext, r.URL.Query().Get("code"))
	if err != nil {
		fmt.Println("there was an issue getting your token")
		//return
	}

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
