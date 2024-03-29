package goauth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	auth "tidify/auth"
	"tidify/devlog"
	"tidify/interactor"

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
	config.RedirectURL = "https://regal-crostata-9cf44e.netlify.app/google"
	//config.RedirectURL = "http://localhost:8888/auth/google/callback"
	//config.RedirectURL = "https://localhost:8081/google"
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

func GoogleAuthCallback(u *interactor.UserInteractor) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		oauthstate, _ := ctx.Cookie("oauthstate")
		if ctx.Request.FormValue("state") != oauthstate {
			devlog.Debug("[GoogleAuthCallback] Invalid googla oauth state - cookie:", oauthstate, ctx.Request.FormValue("state"))
			ctx.JSON(auth.GetHTTPStatusCode(auth.TOKEN_AUTHENTICATION_ERROR), auth.GetAPIResponse(auth.TOKEN_AUTHENTICATION_ERROR))
			return

		}
		data, err := getGoogleUserInfo(ctx.Request.FormValue("code"))
		if err != nil {
			devlog.Debug("[GoogleAuthCallback] Invalid googla oauth code ", err.Error())
			ctx.JSON(auth.GetHTTPStatusCode(auth.TOKEN_AUTHENTICATION_ERROR), auth.GetAPIResponse(auth.TOKEN_AUTHENTICATION_ERROR))
			return
		}
		userData := &UserGoogle{}
		json.Unmarshal(data, &userData)
		u.CreateUser(ctx, userData.Email, "google")
		devlog.Debug(string(data), googleToken, userData)
	}

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
