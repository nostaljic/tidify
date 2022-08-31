package router

import (
	"log"
	"tidify/auth"
	"tidify/interactor"
	"tidify/repository"

	apauth "tidify/apple"
	goauth "tidify/google"
	kaauth "tidify/kakao"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupRoutes : all the routes are defined here
func SetupRoutes(db *gorm.DB) {
	httpRouter := gin.Default()
	httpRouter.Use(corsConfig())
	userRepository := repository.NewUserRepository(db)
	if err := userRepository.Migrate(); err != nil {
		log.Fatal("Folder migrate err", err)
	}
	folderRepository := repository.NewFolderRepository(db)
	if err := folderRepository.Migrate(); err != nil {
		log.Fatal("Folder migrate err", err)
	}
	bookmarkRepository := repository.NewBookmarkRepository(db)
	if err := bookmarkRepository.Migrate(); err != nil {
		log.Fatal("Folder migrate err", err)
	}
	userInteractor := &interactor.UserInteractor{UserRepository: userRepository}
	folderInteractor := &interactor.FolderInteractor{FolderRepository: folderRepository}
	bookmarkInteractor := &interactor.BookmarkInteractor{BookmarkRepository: bookmarkRepository, FolderRepository: folderRepository}

	userGroup := httpRouter.Group("/signin")
	oauthGroup := httpRouter.Group("/auth")
	folderGroup := httpRouter.Group("/folders")
	bookmarkGroup := httpRouter.Group("/bookmarks")

	setUpUser(userGroup, userInteractor)
	setUpOauth(oauthGroup, userInteractor)
	setUpFolder(folderGroup, folderInteractor)
	setUpBookmark(bookmarkGroup, bookmarkInteractor)

	httpRouter.Run(":8888")
}
func corsConfig() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization", "Origin", "*"},
		AllowCredentials: true,
		AllowAllOrigins:  true,
	})
}
func setUpFolder(group *gin.RouterGroup,
	folderInteractor *interactor.FolderInteractor,
) {
	group.Use(auth.JwtCheckMiddleware())
	group.GET("", folderInteractor.GetFolder)
	group.POST("", folderInteractor.CreateFolder)
	group.DELETE("", folderInteractor.DeleteFolder)
	group.PUT("", folderInteractor.UpdateFolder)
}

func setUpBookmark(group *gin.RouterGroup,
	bookmarkInteractor *interactor.BookmarkInteractor,
) {
	group.Use(auth.JwtCheckMiddleware())
	group.GET("", bookmarkInteractor.GetBookmark)
	group.POST("", bookmarkInteractor.CreateBookmark)
	group.DELETE("", bookmarkInteractor.DeleteBookmark)
	group.PUT("", bookmarkInteractor.UpdateBookmark)
}

func setUpUser(group *gin.RouterGroup,
	userInteractor *interactor.UserInteractor,
) {
	//group.POST("", userInteractor.CreateUser)
	group.GET("", userInteractor.SignAgain)
}

func setUpOauth(group *gin.RouterGroup, userInteractor *interactor.UserInteractor) {
	group.GET("")
	group.GET("/google", goauth.GoogleLoginHandler)
	group.GET("/google/callback", goauth.GoogleAuthCallback(userInteractor))
	group.POST("/apple", apauth.AppleLoginHandler(userInteractor))
	group.GET("/kakao", kaauth.KakaoLoginHandler)
	group.GET("/kakao/callback", kaauth.KakaoAuthCallback(userInteractor))
}
