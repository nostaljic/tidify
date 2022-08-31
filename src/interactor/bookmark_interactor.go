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

type BookmarkInteractor struct {
	FolderRepository   repository.FolderRepository
	BookmarkRepository repository.BookmarkRepository
	BookmarkModel      models.Bookmark
}
type BookmarkList struct {
	List        []models.Bookmark `json:"list"`
	TotalCount  int64             `json:"total_count"`
	APIResponse APIResponse       `json:"api_response"`
}
type GetBookmarkConditions struct {
	Start   int    `form:"start" binding:"gte=0"`
	Count   int    `form:"count" binding:"required,gte=1"`
	Folder  int    `form:"folder" binding:"gte=0"`
	Keyword string `form:"keyword"`
}
type DeleteBookmarkBody struct {
	BookmarkID int `json:"Bookmark_id"`
}

func (u *BookmarkInteractor) CreateBookmark(c *gin.Context) {
	reqData := models.Bookmark{}
	body, _ := ioutil.ReadAll(c.Request.Body)
	if err := json.Unmarshal(body, &reqData); err != nil {
		u.returnResponse(c, GetAPIResponse(INVALID_REQUEST_DATAS))
		return
	}
	if reqData.BookmarkTitle == "" || reqData.BookmarkUrl == "" {
		u.returnResponse(c, GetAPIResponse(INVALID_REQUEST_DATAS))
		return
	}
	userEmail, err := auth.ParseEmailFromToken(c)
	isMyFolder := u.isMyFolder(int(reqData.FolderID), userEmail)
	if !isMyFolder {
		u.returnResponse(c, GetAPIResponse(NO_PERMISSION))
		return
	}
	devlog.Debug("[CreateBookmark] Permission Test", err, userEmail)
	reqData.UserEmail = userEmail
	if err := u.BookmarkRepository.Create(&reqData); err != nil {
		u.returnResponse(c, GetAPIResponse(INTERNAL_SERVER_ERROR))
		return
	}
	u.returnResponse(c, GetAPIResponse(OK))
	return
}

func (u *BookmarkInteractor) DeleteBookmark(c *gin.Context) {
	reqData := DeleteBookmarkBody{}
	body, _ := ioutil.ReadAll(c.Request.Body)
	if err := json.Unmarshal(body, &reqData); err != nil {
		u.returnResponse(c, GetAPIResponse(INTERNAL_SERVER_ERROR))
		return
	}
	userEmail, err := auth.ParseEmailFromToken(c)
	devlog.Debug("[DeleteBookmark] Permission Test", err, userEmail)
	available := u.isMyBookmark(int(reqData.BookmarkID), userEmail)
	if !available {
		u.returnResponse(c, GetAPIResponse(NO_PERMISSION))
		return
	}
	if err := u.BookmarkRepository.Delete(reqData.BookmarkID); err != nil {
		u.returnResponse(c, GetAPIResponse(INTERNAL_SERVER_ERROR))
		return
	}
	u.returnResponse(c, GetAPIResponse(OK))
	return
}

func (u *BookmarkInteractor) UpdateBookmark(c *gin.Context) {
	reqData := models.Bookmark{}
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
	isMyFolder := u.isMyFolder(int(reqData.FolderID), userEmail)
	if !isMyFolder {
		u.returnResponse(c, GetAPIResponse(NO_PERMISSION))
		return
	}
	devlog.Debug("[UpdateBookmark] Permission Test", err, userEmail)
	available := u.isMyBookmark(int(reqData.BookmarkID), userEmail)
	if !available {
		u.returnResponse(c, GetAPIResponse(NO_PERMISSION))
		return
	}
	if err := u.BookmarkRepository.Update(&reqData); err != nil {
		u.returnResponse(c, GetAPIResponse(INTERNAL_SERVER_ERROR))
		return
	}
	u.returnResponse(c, GetAPIResponse(OK))
	return
}

func (u *BookmarkInteractor) GetBookmark(c *gin.Context) {
	response := BookmarkList{}
	req := GetBookmarkConditions{}
	if err := c.ShouldBindQuery(&req); err != nil {
		devlog.Debug("[GetBookmark]", err)
		u.returnResponse(c, GetAPIResponse(INVALID_REQUEST_QUERIES))
		return
	}
	userEmail, err := auth.ParseEmailFromToken(c)

	Bookmarks, err := u.BookmarkRepository.FindBookmarkList(userEmail, req.Start, req.Count, req.Keyword, req.Folder)
	totalCount, err := u.BookmarkRepository.FindBookmarkListCount(userEmail, req.Keyword, req.Folder)
	response.List = Bookmarks
	response.TotalCount = totalCount
	response.APIResponse = GetAPIResponse(OK)
	devlog.Debug("[GetBookmark]", Bookmarks, err)
	u.returnResponse(c, response)
}

func (u *BookmarkInteractor) isMyBookmark(BookmarkId int, email string) bool {
	Bookmark := u.BookmarkRepository.GetBookmarkByID(BookmarkId)
	if Bookmark == nil {
		return false
	}
	devlog.Debug("[isMyBookmark] result", Bookmark, Bookmark.UserEmail, email)
	if Bookmark.UserEmail != email {
		return false
	}
	return true
}

func (u *BookmarkInteractor) isMyFolder(folderId int, email string) bool {
	if folderId == 0 {
		return true
	}
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

func (u *BookmarkInteractor) returnResponse(c *gin.Context, data interface{}) {
	switch v := data.(type) {
	case BasicResponse:
		response := data.(BasicResponse)
		c.JSON(GetHTTPStatusCode(response.APIResponse.ResultCode), response)
	case APIResponse:
		response := data.(APIResponse)
		resp := BasicResponse{APIResponse: response}
		c.JSON(GetHTTPStatusCode(response.ResultCode), resp)
	case BookmarkList:
		response := data.(BookmarkList)
		c.JSON(GetHTTPStatusCode(response.APIResponse.ResultCode), response)
	default:
		devlog.Fatal("[returnResponse] Type error: ", v)
	}
}
