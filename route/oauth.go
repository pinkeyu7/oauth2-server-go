package route

import (
	apiV1 "oauth2-server-go/api/v1"

	"github.com/gin-gonic/gin"
)

func OauthV1(r *gin.Engine) {
	oauth := r.Group("/oauth")

	// Oauth2 授權
	oauth.GET("/authorize", func(c *gin.Context) {
		apiV1.UserAuthorizeHandler(c)
	})

	oauth.POST("/login", func(c *gin.Context) {
		apiV1.LoginHandler(c)
	})

	oauth.POST("/auth", func(c *gin.Context) {
		apiV1.UserAuthHandler(c)
	})

	oauth.POST("/authorize", func(c *gin.Context) {
		apiV1.UserAuthorizeHandler(c)
	})

	oauth.POST("/token", func(c *gin.Context) {
		apiV1.TokenHandler(c)
	})
}
