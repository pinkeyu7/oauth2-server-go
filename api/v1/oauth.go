package v1

import (
	"net/http"
	"net/url"
	"oauth2-server-go/api"
	"oauth2-server-go/config"
	"oauth2-server-go/dto/apires"
	oauthClientRepo "oauth2-server-go/internal/oauth/client/repository"
	oauthLib "oauth2-server-go/internal/oauth/library"
	"oauth2-server-go/pkg/er"
	"oauth2-server-go/pkg/logr"

	"github.com/gin-gonic/gin"
	"github.com/go-session/session"
	"go.uber.org/zap"
)

func UserAuthorizeHandler(c *gin.Context) {
	store, err := session.Start(c, c.Writer, c.Request)
	if err != nil {
		sessionErr := er.NewAppErr(http.StatusInternalServerError, er.UnknownError, err.Error(), err)
		_ = c.Error(sessionErr)
		return
	}

	var form url.Values
	if v, ok := store.Get("ReturnUri"); ok {
		form = v.(url.Values)
	}
	c.Request.Form = form

	store.Delete("ReturnUri")
	err = store.Save()
	if err != nil {
		storeErr := er.NewAppErr(http.StatusInternalServerError, er.UnauthorizedError, err.Error(), err)
		_ = c.Error(storeErr)
		return
	}

	err = oauthLib.Srv.HandleAuthorizeRequest(c.Writer, c.Request)
	if err != nil {
		logr.L.Error("oauth authorize error:", zap.Error(err))
		authorizeErr := er.NewAppErr(http.StatusBadRequest, er.ErrorParamInvalid, err.Error(), err)
		_ = c.Error(authorizeErr)
		return
	}
}

func LoginHandler(c *gin.Context) {
	// 取得表單內資料
	if err := c.Request.ParseForm(); err != nil {
		paramErr := er.NewAppErr(http.StatusBadRequest, er.ErrorParamInvalid, "", err)
		_ = c.Error(paramErr)
		return
	}

	account := c.PostForm("username")
	password := c.PostForm("password")

	env := api.GetEnv()
	ocr := oauthClientRepo.NewRepository(env.Orm)
	store, client, redirectUri, err := oauthLib.Validation(c, ocr)
	if err != nil {
		clientErr := er.NewAppErr(http.StatusBadRequest, er.OauthClientDataError, "", err)
		_ = c.Error(clientErr)
		return
	}

	// 驗證登入資訊
	//ur := userRepo.NewMysqlUserRepo(env.Orm)
	//_, err = oas.VerifyLogin(phone, pinCode)
	if account != "test" || password != "test" {
		pageData := apires.OauthLoginPage{
			HostIconPath: client.IconPath,
			ClientName:   client.Name,
			AccountError: true,
			RedirectUrl:  redirectUri,
			BasePath:     config.GetHtmlBasePath(),
		}
		c.HTML(http.StatusOK, "login.tmpl", pageData)
		return
	}

	store.Set("LoggedInUserID", account)
	err = store.Save()
	if err != nil {
		storeErr := er.NewAppErr(http.StatusInternalServerError, er.OauthUnknownError, err.Error(), err)
		_ = c.Error(storeErr)
		return
	}

	c.Redirect(http.StatusMovedPermanently, oauthLib.AuthUrl)
}
