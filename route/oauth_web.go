package route

import (
	"oauth2-server-go/web"

	"github.com/gin-gonic/gin"
)

func OauthWebV1(r *gin.Engine) {
	oauth := r.Group("/oauth")

	// Oauth2 登入頁
	oauth.GET("/login", func(c *gin.Context) {
		web.LoginHandler(c)
	})

	// Oauth2 授權頁
	oauth.GET("/auth", func(c *gin.Context) {
		web.AuthHandler(c)
	})
}
