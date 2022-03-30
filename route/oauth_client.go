package route

import (
	apiV1 "oauth2-server-go/api/v1"

	"github.com/gin-gonic/gin"
)

func OauthClientV1(r *gin.Engine) {
	v1Auth := r.Group("/v1/oauth/clients")

	v1Auth.POST("/", func(c *gin.Context) {
		apiV1.AddOauthClient(c)
	})
}
