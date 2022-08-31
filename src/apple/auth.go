package apauth

import (
	"encoding/json"
	"io/ioutil"
	"tidify/devlog"
	"tidify/interactor"

	"github.com/Timothylock/go-signin-with-apple/apple"
	"github.com/gin-gonic/gin"
)

type AppleToken struct {
	IDToken string `json:"id_token"`
}
type User struct {
	Name  string
	EMail string
	UUID  string
}

func getAppleUserInfo(idToken string) (*User, error) {
	devlog.Debug("[getAppleUserInfo] Parse Token", idToken)
	uid, err := apple.GetUniqueID(idToken)
	if err != nil {
		devlog.Debug("[getAppleUserInfo] Error GetUniqueID")
		return nil, err
	}

	claim, err := apple.GetClaims(idToken)
	if err != nil {
		devlog.Debug("[getAppleUserInfo] Error GetClaims")
		return nil, err
	}
	email := (*claim)["email"]
	user := &User{
		UUID:  uid,
		EMail: email.(string),
	}

	return user, nil
}
func AppleLoginHandler(u *interactor.UserInteractor) gin.HandlerFunc {
	return func(c *gin.Context) {
		reqData := AppleToken{}
		body, _ := ioutil.ReadAll(c.Request.Body)
		if err := json.Unmarshal(body, &reqData); err != nil {
			devlog.Debug("[AppleLoginHandler] Error Parse Body")
			c.Abort()
			return
		}
		userinfo, err := getAppleUserInfo(reqData.IDToken)
		if err != nil {
			devlog.Debug("[getAppleUserInfo] Error getAppleUserInfo")
			c.Abort()
			return
		}
		u.CreateUser(c, userinfo.EMail, "apple")
	}
}
