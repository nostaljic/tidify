package router

import (

	"log"
	"tidify/interactor"
	"tidify/repository"

	goauth "tidify/google"
	//kaauth "tidify/kakao"
	//apauth "tidify/apple"


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


	OauthGroup := httpRouter.Group("/auth")

	httpRouter.Run(":8888")

}
func setUpFolder(group *gin.RouterGroup,
	folderInteractor *interactor.FolderInteractor,
) {
	group.GET("", folderInteractor.GetFolder)
	group.POST("", folderInteractor.CreateFolder)
}

func setUpOauth(group *gin.RouterGroup) {
	group.GET("")
	group.GET("/google", goauth.GoogleLoginHandler)
	group.GET("/google/callback", goauth.GoogleAuthCallback)
	//TODO : APPLE
	//group.GET("/apple", apauth.AppleLoginHandler)
	//group.GET("/apple/callback", apauth.AppleAuthCallback)
	//TODO : KAKAO
	//group.GET("/kakao", kaauth.KakaoLoginHandler)
	//group.GET("/kakao/callback", kaauth.KakaoAuthCallback)
}
