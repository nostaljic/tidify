package interactor

import (
	"encoding/json"
	"io/ioutil"
	"tidify/devlog"
	models "tidify/models"
	repository "tidify/repository"

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

func (u *FolderInteractor) CreateFolder(c *gin.Context) {
	devlog.Debug("[FolderInteractor] - Create")
	reqData := models.Folder{}
	body, _ := ioutil.ReadAll(c.Request.Body)
	if err := json.Unmarshal(body, &reqData); err != nil {
		return
	}
	u.FolderRepository.Create(&reqData)
	devlog.Debug("CREATE!!!!!!!!!!!!!", body, reqData)
}
func (u *FolderInteractor) GetFolder(c *gin.Context) {
	response := FolderList{
		List: []models.Folder{},
	}
	folders, err := u.FolderRepository.FindFolderList()
	response.List = folders
	response.TotalCount = 0
	response.ApiResponse = GetAPIResponse(OK)
	devlog.Debug(folders, err)
	u.returnResponse(c, response)
}

func (u *FolderInteractor) returnResponse(c *gin.Context, data interface{}) {
	switch v := data.(type) {
	case APIResponse:
		response := data.(APIResponse)
		c.JSON(GetHTTPStatusCode(response.ResultCode), response)
	case FolderList:
		response := data.(FolderList)
		c.JSON(GetHTTPStatusCode(response.ApiResponse.ResultCode), response)
	default:
		devlog.Fatal("Type error: ", v)
	}
}
