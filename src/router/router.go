package router

import (
	"log"
	"tidify/auth"
	"tidify/interactor"
	"tidify/repository"

	goauth "tidify/google"
	kaauth "tidify/kakao"

	//apauth "tidify/apple"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupRoutes : all the routes are defined here
func SetupRoutes(db *gorm.DB) {
	httpRouter := gin.Default()

	folderRepository := repository.NewFolderRepository(db)
	if err := folderRepository.Migrate(); err != nil {
		log.Fatal("Folder migrate err", err)
	}
	userRepository := repository.NewUserRepository(db)
	if err := userRepository.Migrate(); err != nil {
		log.Fatal("Folder migrate err", err)
	}

	folderInteractor := &interactor.FolderInteractor{FolderRepository: folderRepository}
	userInteractor := &interactor.UserInteractor{UserRepository: userRepository}

	folderGroup := httpRouter.Group("/folders")
	userGroup := httpRouter.Group("/signin")
	oauthGroup := httpRouter.Group("/auth")

	setUpFolder(folderGroup, folderInteractor)
	setUpUser(userGroup, userInteractor)
	setUpOauth(oauthGroup)

	httpRouter.Run(":8888")
}
func setUpFolder(group *gin.RouterGroup,
	folderInteractor *interactor.FolderInteractor,
) {
	group.Use(auth.JwtCheckMiddleware())
	group.GET("", folderInteractor.GetFolder)
	group.POST("", folderInteractor.CreateFolder)
}

func setUpUser(group *gin.RouterGroup,
	userInteractor *interactor.UserInteractor,
) {
	group.POST("", userInteractor.CreateUser)
	group.GET("", userInteractor.SignAgain)
}

func setUpOauth(group *gin.RouterGroup) {
	group.GET("")
	group.GET("/google", goauth.GoogleLoginHandler)
	group.GET("/google/callback", goauth.GoogleAuthCallback)
	//TODO : APPLE
	//group.GET("/apple", apauth.AppleLoginHandler)
	//group.GET("/apple/callback", apauth.AppleAuthCallback)
	//TODO : KAKAO
	group.GET("/kakao", kaauth.KakaoLoginHandler)
	group.GET("/kakao/callback", kaauth.KakaoAuthCallback(userInteractor))
}
