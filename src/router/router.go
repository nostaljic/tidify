package router

import (
	"log"
	"tidify/interactor"
	"tidify/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//SetupRoutes : all the routes are defined here
func SetupRoutes(db *gorm.DB) {
	httpRouter := gin.Default()

	folderRepository := repository.NewFolderRepository(db)
	if err := folderRepository.Migrate(); err != nil {
		log.Fatal("Folder migrate err", err)
	}
	folderGroup := httpRouter.Group("/folders")
	folderInteractor := &interactor.FolderInteractor{FolderRepository: folderRepository}
	setUpFolder(folderGroup, folderInteractor)

	httpRouter.Run(":8888")

}
func setUpFolder(group *gin.RouterGroup,
	folderInteractor *interactor.FolderInteractor,
) {
	group.GET("", folderInteractor.GetFolder)
	group.POST("", folderInteractor.CreateFolder)
}
