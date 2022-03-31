package route

import (
	apiV1 "oauth2-server-go/api/v1"
	"oauth2-server-go/middleware"

	"github.com/gin-gonic/gin"
)

func UserV1(r *gin.Engine) {
	v1Oauth := r.Group("/v1")
	v1Oauth.Use(middleware.OauthMiddleware())

	v1Oauth.GET("/users", func(c *gin.Context) {
		apiV1.GetUser(c)
	})
}
