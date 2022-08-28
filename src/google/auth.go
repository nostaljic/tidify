package goauth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"os"
	"tidify/devlog"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	config = &oauth2.Config{
		ClientID:     "",
		ClientSecret: "",
		RedirectURL:  "",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
	googleToken = ""
)

type UserGoogle struct {
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Hd            string `json:"hd"`
	Picture       string `json:"picture"`
	Sub           string `json:"sub"`
}

func setConfig() {
	config.ClientID = os.Getenv("GOOGLE_CLIENT_ID")
	config.ClientSecret = os.Getenv("GOOGLE_SECRET_KEY")
	config.RedirectURL = "http://localhost:8888/auth/google/callback"
	config.Scopes = []string{"https://www.googleapis.com/auth/userinfo.email"}
	config.Endpoint = google.Endpoint
}

func GoogleLoginHandler(ctx *gin.Context) {
	devlog.Debug(config)
	setConfig()
	state := generateStateOauthCookie(ctx)
	url := config.AuthCodeURL(state)
	ctx.Redirect(http.StatusFound, url)
}

func GoogleAuthCallback(ctx *gin.Context) {
	oauthstate, _ := ctx.Cookie("oauthstate")
	if ctx.Request.FormValue("state") != oauthstate {
		devlog.Debug("[GoogleAuthCallback] Invalid googla oauth state - cookie:", oauthstate, ctx.Request.FormValue("state"))
		ctx.Redirect(http.StatusFound, "http://localhost:8888")
		return
	}
	data, err := getGoogleUserInfo(ctx.Request.FormValue("code"))
	if err != nil {
		devlog.Debug("[GoogleAuthCallback] Invalid googla oauth code ", err.Error())
		ctx.Redirect(http.StatusFound, "http://localhost:8888")
		return
	}
	devlog.Debug(string(data), googleToken)
}

func generateStateOauthCookie(ctx *gin.Context) string {
	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	ctx.SetCookie("oauthstate", state, 24*60*60, "", "", false, false)
	return state
}

func getGoogleUserInfo(code string) ([]byte, error) {
	token, err := config.Exchange(context.Background(), code)
	if err != nil {
		devlog.Fatal("[getGoogleUserInfo] Failed to Exchange", err.Error())
	}
	googleToken = token.AccessToken
	client := config.Client(oauth2.NoContext, token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		devlog.Fatal("[getGoogleUserInfo] Failed to get response", err.Error())
	}
	return ioutil.ReadAll(resp.Body)
}
