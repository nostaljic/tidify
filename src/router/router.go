package router

import (
	goauth "tidify/google"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//SetupRoutes : all the routes are defined here
func SetupRoutes(db *gorm.DB) {
	httpRouter := gin.Default()

	OauthGroup := httpRouter.Group("/auth")
	setUpOauth(OauthGroup)
	httpRouter.Run(":8888")
}

func setUpOauth(group *gin.RouterGroup) {
	group.GET("")
	group.GET("/google", goauth.GoogleLoginHandler)
	group.GET("/google/callback", goauth.GoogleAuthCallback)
	//TODO : APPLE
	//TODO : KAKAO
}
