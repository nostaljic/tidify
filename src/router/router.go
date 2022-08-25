package router

import (
	goauth "tidify/google"
	//kaauth "tidify/kakao"
	//apauth "tidify/apple"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

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
	//group.GET("/apple", apauth.AppleLoginHandler)
	//group.GET("/apple/callback", apauth.AppleAuthCallback)
	//TODO : KAKAO
	//group.GET("/kakao", kaauth.KakaoLoginHandler)
	//group.GET("/kakao/callback", kaauth.KakaoAuthCallback)
}
