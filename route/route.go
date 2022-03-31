package route

import (
	"oauth2-server-go/config"
	_ "oauth2-server-go/docs"
	"oauth2-server-go/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func Init() *gin.Engine {
	r := gin.New()

	// gin 檔案上傳body限制
	r.MaxMultipartMemory = 64 << 20 // 8 MiB

	// Middleware
	r.Use(middleware.LogRequest())
	r.Use(middleware.ErrorResponse())

	// Swagger
	if mode := gin.Mode(); mode == gin.DebugMode {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	corsConf := cors.DefaultConfig()
	corsConf.AllowCredentials = true
	corsConf.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"}
	corsConf.AllowHeaders = []string{"Origin", "X-Requested-With", "Content-Type", "Accept", "Authorization", "Bearer", "Accept-Language"}
	corsConf.AllowOriginFunc = config.GetCorsRule
	r.Use(cors.New(corsConf))

	// Oauth authentication flow
	OauthV1(r)
	OauthWebV1(r)

	// Oauth client and scope api
	OauthClientV1(r)

	return r
}
