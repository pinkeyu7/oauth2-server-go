package middleware

import (
	"net/http"
	"oauth2-server-go/api"
	clientRepo "oauth2-server-go/internal/oauth/client/repository"
	scopeRepo "oauth2-server-go/internal/oauth/scope/repository"
	scopeSrv "oauth2-server-go/internal/oauth/scope/service"
	"oauth2-server-go/pkg/er"
	"oauth2-server-go/pkg/logr"

	"go.uber.org/zap"

	tokenLib "oauth2-server-go/internal/token/library"

	"github.com/gin-gonic/gin"
)

func OauthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Bearer")
		if token == "" {
			authErr := er.NewAppErr(http.StatusUnauthorized, er.UnauthorizedError, "Token is required", nil)
			c.AbortWithStatusJSON(authErr.GetStatus(), authErr.GetMsg())
			return
		}

		claims, err := tokenLib.ParseOauthToken(token)
		if err != nil {
			authErr := er.NewAppErr(http.StatusUnauthorized, er.UnauthorizedError, "Token is not valid", nil)
			c.AbortWithStatusJSON(authErr.GetStatus(), authErr.GetMsg())
			return
		}

		env := api.GetEnv()
		isPass, err := oauthCheckScope(c, env, claims.Audience)
		if err != nil {
			permissionErr := er.NewAppErr(http.StatusForbidden, er.ForbiddenError, "Permission Error", err)
			c.AbortWithStatusJSON(permissionErr.GetStatus(), permissionErr.GetMsg())
			return
		}
		if !isPass {
			permissionErr := er.NewAppErr(http.StatusForbidden, er.ForbiddenError, "Permission Error", nil)
			c.AbortWithStatusJSON(permissionErr.GetStatus(), permissionErr.GetMsg())
			return
		}

		c.Set("sub", claims.Subject)
		c.Set("aud", claims.Audience)
	}
}

func oauthCheckScope(c *gin.Context, env *api.Env, clientId string) (bool, error) {
	ocr := clientRepo.NewRepository(env.Orm)
	osc := scopeRepo.NewRedis(env.RedisCluster)
	osr := scopeRepo.NewRepository(env.Orm, osc)
	oss := scopeSrv.NewService(osr, osc, ocr)

	// 取得 API 對應之 scope 名稱
	scope, err := oss.GetScope(c.Request.URL.Path, c.Request.Method)
	if err != nil {
		logr.L.Error("get scope error.", zap.String("error", err.Error()))
		return false, err
	}

	// 驗證 API 是否在授權名單內
	isPass, err := oss.VerifyScope(clientId, scope)
	if err != nil {
		logr.L.Error("verify scope error.", zap.String("error", err.Error()))
		return false, err
	}

	return isPass, nil
}
