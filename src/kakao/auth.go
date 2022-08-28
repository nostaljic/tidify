package kaauth

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
)

var (
	state  = "YOUR_RANDOM_STATE"
	config = &oauth2.Config{
		ClientID:     "",
		ClientSecret: "",
		Endpoint: oauth2.Endpoint{
			AuthURL:  "",
			TokenURL: "",
		},
		RedirectURL: "",
	}
	kakaoToken = ""
)

func setConfig() {
	config.ClientID = os.Getenv("KAKAO_CLIENT_ID")
	config.ClientSecret = os.Getenv("KAKAO_SECRET_KEY")
	config.RedirectURL = "http://localhost:8888/auth/kakao/callback"
	config.Endpoint = oauth2.Endpoint{
		AuthURL:  "https://kauth.kakao.com/oauth/authorize",
		TokenURL: "https://kauth.kakao.com/oauth/token",
	}
}

func KakaoLoginHandler(ctx *gin.Context) {
	devlog.Debug(config)
	setConfig()
	state := generateStateOauthCookie(ctx)
	url := config.AuthCodeURL(state)
	ctx.Redirect(http.StatusFound, url)
}
func generateStateOauthCookie(ctx *gin.Context) string {
	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	ctx.SetCookie("oauthstate", state, 24*60*60, "", "", false, false)
	return state
}
func KakaoAuthCallback(ctx *gin.Context) {
	oauthstate, _ := ctx.Cookie("oauthstate")
	if ctx.Request.FormValue("state") != oauthstate {
		devlog.Debug("[KakaoAuthCallback] Invalid Kakao oauth state - cookie:", oauthstate, ctx.Request.FormValue("state"))
		ctx.Redirect(http.StatusFound, "http://localhost:8888")
		return
	}
	data, err := getKakaoUserInfo(ctx.Request.FormValue("code"))
	if err != nil {
		devlog.Debug("[KakaoAuthCallback] Invalid Kakao oauth code ", err.Error())
		ctx.Redirect(http.StatusFound, "http://localhost:8888")
		return
	}
	devlog.Debug(string(data), kakaoToken)
}
func getKakaoUserInfo(code string) ([]byte, error) {
	token, err := config.Exchange(context.Background(), code)
	if err != nil {
		devlog.Fatal("[getKakaoUserInfo] Failed to Exchange", err.Error())
	}
	kakaoToken = token.AccessToken
	client := config.Client(oauth2.NoContext, token)
	resp, err := client.Get("https://kapi.kakao.com/v1/api/talk/profile")
	if err != nil {
		devlog.Fatal("[getKakaoUserInfo] Failed to get response", err.Error())
	}
	return ioutil.ReadAll(resp.Body)
}

//TODO: KAKAO HANDLERS
