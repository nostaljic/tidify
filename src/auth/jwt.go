package auth

import (
	"os"
	"tidify/devlog"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type Claims struct {
	Email string
	Sns   string
	jwt.StandardClaims
}

func CreateJWT(Email string, Sns string) (string, error) {
	mySigningKey := []byte(os.Getenv("SECRET_KEY"))
	//claims := &Claims{Email: Email, Sns: Sns, StandardClaims: jwt.StandardClaims{ExpiresAt: time.Now().Add(time.Minute * 60 * 24).Unix()}}
	claims := &Claims{Email: Email, Sns: Sns, StandardClaims: jwt.StandardClaims{ExpiresAt: time.Now().Add(time.Minute * 120).Unix()}}
	aToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tk, err := aToken.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}
	return tk, nil
}
func CreateRefreshJWT(Email string, Sns string) (string, error) {
	mySigningKey := []byte(os.Getenv("SECRET_REFRESH_KEY"))
	claims := &Claims{Email: Email, Sns: Sns, StandardClaims: jwt.StandardClaims{ExpiresAt: time.Now().Add(time.Minute * 60 * 24 * 7).Unix()}}
	aToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tk, err := aToken.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}
	return tk, nil
}
func RefreshAccessToken(c *gin.Context) (string, error) {
	token, err := c.Request.Cookie("refresh-token")
	if err != nil {
		devlog.Debug("[RefreshAccessToken] Refresh Token error", err)
		return "", err
	}
	tknStr := token.Value
	if tknStr == "" {
		devlog.Debug("[RefreshAccessToken] Refresh Token string error")
		return "", err
	}
	claims := &Claims{}
	key := func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_REFRESH_KEY")), nil
	}
	_, err = jwt.ParseWithClaims(tknStr, claims, key)
	if err != nil {
		if claims.ExpiresAt <= time.Now().Unix() {
			devlog.Debug("[RefreshAccessToken] Refresh Token Time Expired", claims.ExpiresAt, time.Now().Unix())
			return "", err
		} else {
			devlog.Debug("[RefreshAccessToken] Refresh Token Parse Error", tknStr, claims)
			return "", err
		}
	}
	return CreateJWT(claims.Email, claims.Sns)
}

// DBTransactionMiddleware : to setup the database transaction middleware
func JwtCheckMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Request.Cookie("access-token")
		if err != nil {
			devlog.Debug("[JwtCheckMiddleware] Access Token error", err)
			c.JSON(GetHTTPStatusCode(TOKEN_AUTHENTICATION_ERROR), GetAPIResponse(TOKEN_AUTHENTICATION_ERROR))
			c.Abort()
			return
		}
		tknStr := token.Value
		if tknStr == "" {
			devlog.Debug("[JwtCheckMiddleware] Access Token string error")
			c.JSON(GetHTTPStatusCode(TOKEN_AUTHENTICATION_ERROR), GetAPIResponse(TOKEN_AUTHENTICATION_ERROR))
			c.Abort()
			return
		}

		claims := &Claims{}
		key := func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET_KEY")), nil
		}
		_, err = jwt.ParseWithClaims(tknStr, claims, key)
		if err != nil {
			if claims.ExpiresAt <= time.Now().Unix() {
				devlog.Debug("[JwtCheckMiddleware] Access Token Time Expired", claims.ExpiresAt, time.Now().Unix())
				c.JSON(GetHTTPStatusCode(TOKEN_EXPIRED), GetAPIResponse(TOKEN_EXPIRED))
				c.Abort()
				return
			} else {
				devlog.Debug("[JwtCheckMiddleware] Access Token Parse Error", tknStr, claims)
				c.JSON(GetHTTPStatusCode(TOKEN_AUTHENTICATION_ERROR), GetAPIResponse(TOKEN_AUTHENTICATION_ERROR))
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
