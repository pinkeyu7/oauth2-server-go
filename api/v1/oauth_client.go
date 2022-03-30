package v1

import (
	"net/http"
	"oauth2-server-go/api"
	"oauth2-server-go/dto/apireq"
	"oauth2-server-go/pkg/er"
	"oauth2-server-go/pkg/valider"

	oauthClientRepo "oauth2-server-go/internal/oauth2/client/repository"
	oauthClientSrv "oauth2-server-go/internal/oauth2/client/service"

	"github.com/gin-gonic/gin"
)

// AddOauthClient
// @Summary Add Oauth Client 新增 Oauth Client
// @Produce json
// @Accept json
// @Tags Oauth Client
// @Security Bearer
// @Param Bearer header string true "JWT Token"
// @Param Body body apireq.AddOauthClient true "Request Add Oauth Client"
// @Success 200 {string} string "{}"
// @Failure 400 {object} er.AppErrorMsg "{"code":"400400","message":"Wrong parameter format or invalid"}"
// @Failure 401 {object} er.AppErrorMsg "{"code":"400401","message":"Unauthorized"}"
// @Failure 403 {object} er.AppErrorMsg "{"code":"400403","message":"Permission denied"}"
// @Failure 404 {object} er.AppErrorMsg "{"code":"400404","message":"Resource not found"}"
// @Failure 500 {object} er.AppErrorMsg "{"code":"500000","message":"Database unknown error"}"
// @Router /v1/oauth/clients [post]
func AddOauthClient(c *gin.Context) {
	req := apireq.AddOauthClient{}
	err := c.BindJSON(&req)
	if err != nil {
		paramErr := er.NewAppErr(http.StatusBadRequest, er.ErrorParamInvalid, err.Error(), err)
		_ = c.Error(paramErr)
		return
	}

	// 參數驗證
	err = valider.Validate.Struct(req)
	if err != nil {
		paramErr := er.NewAppErr(http.StatusBadRequest, er.ErrorParamInvalid, err.Error(), err)
		_ = c.Error(paramErr)
		return
	}

	env := api.GetEnv()
	ocr := oauthClientRepo.NewRepository(env.Orm)
	ocs := oauthClientSrv.NewService(ocr)
	err = ocs.Add(&req)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{})
}
