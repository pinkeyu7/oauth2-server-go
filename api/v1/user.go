package v1

import (
	"net/http"
	"oauth2-server-go/api"
	"oauth2-server-go/dto/apireq"
	"oauth2-server-go/dto/apires"
	userRepo "oauth2-server-go/internal/user/repository"
	userSrv "oauth2-server-go/internal/user/service"

	"github.com/gin-gonic/gin"
)

// GetUser
// @Summary Get User 取得使用者資料
// @Produce json
// @Accept json
// @Tags User(Oauth)
// @Security Bearer
// @Param Bearer header string true "OAuth Access Token"
// @Success 200 {object} apires.User
// @Failure 400 {object} er.AppErrorMsg "{"code":"400400","message":"Wrong parameter format or invalid"}"
// @Failure 401 {object} er.AppErrorMsg "{"code":"400401","message":"Unauthorized"}"
// @Failure 403 {object} er.AppErrorMsg "{"code":"400403","message":"Permission denied"}"
// @Failure 404 {object} er.AppErrorMsg "{"code":"400404","message":"Resource not found"}"
// @Failure 500 {object} er.AppErrorMsg "{"code":"500000","message":"Database unknown error"}"
// @Router /v1/users [get]
func GetUser(c *gin.Context) {
	req := apireq.GetUser{
		ClientId: c.GetString("aud"),
		Account:  c.GetString("sub"),
	}

	env := api.GetEnv()
	ur := userRepo.NewRepository(env.Orm)
	us := userSrv.NewService(ur)

	usr, err := us.Get(req.Account)
	if err != nil {
		_ = c.Error(err)
		return
	}
	res := apires.User{
		Id:      usr.Id,
		Account: usr.Account,
		Phone:   usr.Phone,
		Email:   usr.Email,
		Name:    usr.Name,
	}
	c.JSON(http.StatusOK, res)
}
