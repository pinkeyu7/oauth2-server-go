package web

import (
	"net/http"
	"oauth2-server-go/api"
	"oauth2-server-go/config"
	"oauth2-server-go/dto/apires"
	clientRepo "oauth2-server-go/internal/oauth/client/repository"
	oauthLib "oauth2-server-go/internal/oauth/library"
	userRepo "oauth2-server-go/internal/user/repository"
	userSrv "oauth2-server-go/internal/user/service"
	"oauth2-server-go/pkg/er"

	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	env := api.GetEnv()
	ocr := clientRepo.NewRepository(env.Orm)

	// 驗證 Oauth 資訊
	_, client, redirectUri, err := oauthLib.Validation(c, ocr)
	if err != nil {
		clientErr := er.NewAppErr(http.StatusBadRequest, er.OauthClientDataError, "", err)
		_ = c.Error(clientErr)
		return
	}

	pageData := apires.OauthLoginPage{
		ClientName:   client.Name,
		AccountError: false,
		RedirectUrl:  redirectUri,
		BasePath:     config.GetHtmlBasePath(),
	}

	c.HTML(http.StatusOK, "login.tmpl", pageData)
}

func AuthHandler(c *gin.Context) {
	env := api.GetEnv()
	ocr := clientRepo.NewRepository(env.Orm)
	ur := userRepo.NewRepository(env.Orm)
	us := userSrv.NewService(ur)

	// 驗證 Oauth 資訊
	store, client, redirectUri, err := oauthLib.Validation(c, ocr)
	if err != nil {
		clientErr := er.NewAppErr(http.StatusBadRequest, er.OauthClientDataError, "", err)
		_ = c.Error(clientErr)
		return
	}

	// 驗證是否完成登入，並取得電話號碼
	account, isPass := oauthLib.ValidateLogin(store)
	if !isPass {
		c.Redirect(http.StatusFound, oauthLib.LoginUrl)
		return
	}

	_, err = us.Get(account)
	if err != nil {
		_ = c.Error(err)
		return
	}

	scopes := []*apires.Scope{
		{
			Name:   "Name",
			Detail: "",
		},
		{
			Name:   "Public Profile",
			Detail: "",
		},
		{
			Name:   "Email Address",
			Detail: "",
		},
	}

	pageData := apires.OauthAuthPage{
		ClientIconPath: client.IconPath,
		ClientName:     client.Name,
		RedirectUrl:    redirectUri,
		Scopes:         scopes,
		UserName:       "the user name",
		BasePath:       config.GetHtmlBasePath(),
	}

	c.HTML(http.StatusOK, "auth.tmpl", pageData)
}
