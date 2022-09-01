package interactor

import (
	"encoding/json"
	"io/ioutil"
	auth "tidify/auth"
	"tidify/devlog"
	models "tidify/models"
	repository "tidify/repository"
	"time"

	"github.com/gin-gonic/gin"
)

type FolderInteractor struct {
	FolderRepository repository.FolderRepository
	FolderModel      models.Folder
}
type FolderList struct {
	List        []models.Folder `json:"list"`
	TotalCount  int64           `json:"total_count"`
	ApiResponse APIResponse     `json:"api_response"`
}
type GetFolderConditions struct {
	Start   int    `form:"start" binding:"gte=0"`
	Count   int    `form:"count" binding:"required,gte=1"`
	Keyword string `form:"keyword"`
}
type DeleteFolderBody struct {
	FolderID int `json:"folder_id"`
}

func (u *FolderInteractor) CreateFolder(c *gin.Context) {
	reqData := models.Folder{}
	body, _ := ioutil.ReadAll(c.Request.Body)
	if err := json.Unmarshal(body, &reqData); err != nil {
		u.returnResponse(c, GetAPIResponse(INVALID_REQUEST_DATAS))
		return
	}
	if reqData.FolderTitle == "" || reqData.FolderColor == "" {
		u.returnResponse(c, GetAPIResponse(INVALID_REQUEST_DATAS))
		return
	}
	userEmail, err := auth.ParseEmailFromToken(c)
	devlog.Debug("[CreateFolder] Permission Test", err, userEmail)
	reqData.UserEmail = userEmail
	devlog.Debug("[CreateFolder] Create DB Model", reqData)
	if err := u.FolderRepository.Create(&reqData); err != nil {
		u.returnResponse(c, GetAPIResponse(INTERNAL_SERVER_ERROR))
		return
	}
	u.returnResponse(c, GetAPIResponse(OK))
	return
}

func (u *FolderInteractor) DeleteFolder(c *gin.Context) {
	reqData := DeleteFolderBody{}
	body, _ := ioutil.ReadAll(c.Request.Body)
	if err := json.Unmarshal(body, &reqData); err != nil {
		u.returnResponse(c, GetAPIResponse(INTERNAL_SERVER_ERROR))
		return
	}
	userEmail, err := auth.ParseEmailFromToken(c)
	devlog.Debug("[DeleteFolder] Permission Test", err, userEmail)
	available := u.isMyFolder(int(reqData.FolderID), userEmail)
	if !available {
		u.returnResponse(c, GetAPIResponse(NO_PERMISSION))
		return
	}
	if err := u.FolderRepository.Delete(reqData.FolderID, userEmail); err != nil {
		u.returnResponse(c, GetAPIResponse(INTERNAL_SERVER_ERROR))
		return
	}
	u.returnResponse(c, GetAPIResponse(OK))
	return
}

func (u *FolderInteractor) UpdateFolder(c *gin.Context) {
	reqData := models.Folder{}
	body, _ := ioutil.ReadAll(c.Request.Body)
	if err := json.Unmarshal(body, &reqData); err != nil {
		u.returnResponse(c, GetAPIResponse(INTERNAL_SERVER_ERROR))
		return
	}
	if reqData.UserEmail != "" {
		u.returnResponse(c, GetAPIResponse(INVALID_REQUEST_DATAS))
		return
	}

	reqData.UpdatedAt = time.Now()
	userEmail, err := auth.ParseEmailFromToken(c)
	devlog.Debug("[UpdateFolder] Permission Test", err, userEmail)
	available := u.isMyFolder(int(reqData.FolderID), userEmail)
	if !available {
		u.returnResponse(c, GetAPIResponse(NO_PERMISSION))
		return
	}
	if err := u.FolderRepository.Update(&reqData); err != nil {
		u.returnResponse(c, GetAPIResponse(INTERNAL_SERVER_ERROR))
		return
	}
	u.returnResponse(c, GetAPIResponse(OK))
	return
}

func (u *FolderInteractor) GetFolder(c *gin.Context) {
	response := FolderList{}
	req := GetFolderConditions{}
	if c.ShouldBindQuery(&req) != nil {
		u.returnResponse(c, GetAPIResponse(INVALID_REQUEST_QUERIES))
		return
	}

	userEmail, err := auth.ParseEmailFromToken(c)

	folders, err := u.FolderRepository.FindFolderList(userEmail, req.Start, req.Count, req.Keyword)
	totalCount, err := u.FolderRepository.FindFolderListCount(userEmail, req.Keyword)
	response.List = folders
	response.TotalCount = totalCount
	response.ApiResponse = GetAPIResponse(OK)
	devlog.Debug("[GetFolder]", folders, err)
	u.returnResponse(c, response)
}

func (u *FolderInteractor) isMyFolder(folderId int, email string) bool {
	folder := u.FolderRepository.GetFolderByID(folderId)
	if folder == nil {
		return false
	}
	devlog.Debug("[isMyFolder] result", folder, folder.UserEmail, email)
	if folder.UserEmail != email {
		return false
	}
	return true
}

func (u *FolderInteractor) returnResponse(c *gin.Context, data interface{}) {
	switch v := data.(type) {
	case BasicResponse:
		response := data.(BasicResponse)
		c.JSON(GetHTTPStatusCode(response.APIResponse.ResultCode), response)
	case APIResponse:
		response := data.(APIResponse)
		resp := BasicResponse{APIResponse: response}
		c.JSON(GetHTTPStatusCode(response.ResultCode), resp)
	case FolderList:
		response := data.(FolderList)
		c.JSON(GetHTTPStatusCode(response.ApiResponse.ResultCode), response)
	default:
		devlog.Fatal("[returnResponse] Type error: ", v)
	}
}
