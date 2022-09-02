package interactor

import (
	"crypto/sha512"
	"encoding/base64"
	auth "tidify/auth"
	"tidify/devlog"
	models "tidify/models"
	repository "tidify/repository"

	"github.com/gin-gonic/gin"
)

type UserInteractor struct {
	UserRepository repository.UserRepository
	UserModel      models.User
}
type UserResponse struct {
	AccessToken  string      `json:"access_token"`
	RefreshToken string      `json:"refresh_token"`
	APIResponse  APIResponse `json:"api_response"`
}

func (u *UserInteractor) CreateUser(c *gin.Context, email string, sns string) {
	reqData := models.User{UserEmail: email, SnsType: sns}
	if len(reqData.UserEmail) == 0 || len(reqData.SnsType) == 0 {
		u.returnResponse(c, GetAPIResponse(REQUEST_DATA_EMPTY))
		return
	}
	userEmail := reqData.UserEmail
	reqData.UserEmail = hashEmail(reqData.UserEmail)
	devlog.Debug("[CreateUser] User Email Hash", reqData.UserEmail)
	isAlreadyExist, err := u.UserRepository.IsUserExist(&reqData)
	devlog.Debug("[CreateUser] isAlreadyExist", isAlreadyExist)
	if isAlreadyExist {
		devlog.Debug("[CreateUser] Already Exist", err)
		u.SignIn(c, userEmail, reqData.SnsType)
		return
	}
	devlog.Debug("[CreateUser] Not Already Exist", isAlreadyExist)
	createResult, err := u.UserRepository.Create(&reqData)
	if !createResult {
		u.returnResponse(c, GetAPIResponse(ERROR_COMMUNICATE_INTERNAL_DATABASE))
	}
	u.SignIn(c, userEmail, reqData.SnsType)
	devlog.Debug("[CreateUser] Result", reqData)
	return
}
func (u *UserInteractor) SignAgain(c *gin.Context) {
	accessToken, err := auth.RefreshAccessToken(c)
	if err != nil {
		u.returnResponse(c, GetAPIResponse(TOKEN_AUTHENTICATION_ERROR))
		return
	}
	c.SetCookie("access-token", accessToken, 60*60*24, "", "", false, true)
	resp := UserResponse{}
	resp.APIResponse = GetAPIResponse(OK)
	resp.AccessToken = accessToken
	u.returnResponse(c, resp)
	return
}
func (u *UserInteractor) SignIn(c *gin.Context, email string, sns string) {
	accessToken, err := auth.CreateJWT(email, sns)
	if err != nil {
		u.returnResponse(c, GetAPIResponse(INTERNAL_SERVER_ERROR))
		return
	}
	c.SetCookie("access-token", accessToken, 60*60*24, "", "", false, true)
	refreshToken, err := auth.CreateRefreshJWT(email, sns)
	if err != nil {
		u.returnResponse(c, GetAPIResponse(INTERNAL_SERVER_ERROR))
		return
	}
	c.SetCookie("refresh-token", refreshToken, 60*60*24*7, "", "", false, true)
	resp := UserResponse{}
	resp.APIResponse = GetAPIResponse(OK)
	resp.AccessToken = accessToken
	resp.RefreshToken = refreshToken
	u.returnResponse(c, resp)
	return
}

func hashEmail(email string) string {
	sha_512 := sha512.New()
	sha_512.Write([]byte(email))
	str := base64.URLEncoding.EncodeToString(sha_512.Sum(nil))
	devlog.Debug("[hashEmail]", str)
	return str
}

func (u *UserInteractor) returnResponse(c *gin.Context, data interface{}) {
	switch v := data.(type) {
	case BasicResponse:
		response := data.(BasicResponse)
		c.JSON(GetHTTPStatusCode(response.APIResponse.ResultCode), response)
	case APIResponse:
		response := data.(APIResponse)
		resp := BasicResponse{APIResponse: response}
		c.JSON(GetHTTPStatusCode(response.ResultCode), resp)
	case UserResponse:
		response := data.(UserResponse)
		c.JSON(GetHTTPStatusCode(response.APIResponse.ResultCode), response)
	default:
		devlog.Fatal("[returnResponse] Type error: ", v)
	}
}
