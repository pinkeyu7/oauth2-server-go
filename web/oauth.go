package web

import (
	"net/http"
	"oauth2-server-go/api"
	"oauth2-server-go/config"
	"oauth2-server-go/dto/apires"
	oauthClientRepo "oauth2-server-go/internal/oauth/client/repository"
	oauthLib "oauth2-server-go/internal/oauth/library"
	"oauth2-server-go/pkg/er"

	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	env := api.GetEnv()
	ocr := oauthClientRepo.NewRepository(env.Orm)

	// 驗證 Oauth 資訊
	_, client, redirectUri, err := oauthLib.Validation(c, ocr)
	if err != nil {
		clientErr := er.NewAppErr(http.StatusBadRequest, er.OauthClientDataError, "", err)
		_ = c.Error(clientErr)
		return
	}

	pageData := apires.OauthLoginPage{
		HostIconPath: client.IconPath,
		ClientName:   client.Name,
		AccountError: false,
		RedirectUrl:  redirectUri,
		BasePath:     config.GetHtmlBasePath(),
	}

	c.HTML(http.StatusOK, "login.tmpl", pageData)
}
