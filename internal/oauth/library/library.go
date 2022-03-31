package oauthLibrary

import (
	"fmt"
	"log"
	"net/http"
	"oauth2-server-go/api"
	"oauth2-server-go/config"
	"oauth2-server-go/dto/model"
	"strings"

	"github.com/go-playground/validator/v10"

	"github.com/go-oauth2/oauth2/v4/errors"

	"github.com/go-session/session"

	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/server"

	oauthXorm "oauth2-server-go/internal/oauth/xorm"

	"github.com/go-oauth2/oauth2/v4/generates"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/golang-jwt/jwt/v4"
)

var Srv *server.Server
var validate *validator.Validate

func InitOauth2() {
	validate = validator.New()
	env := api.GetEnv()

	manager := manage.NewDefaultManager()
	manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)

	store := oauthXorm.NewStore(env.Orm, "", 0, false)
	manager.MapTokenStorage(store)

	// generate jwt access token
	manager.MapAccessGenerate(generates.NewJWTAccessGenerate("", []byte(config.GetOauthSalt()), jwt.SigningMethodHS512))

	clientStore, _ := oauthXorm.NewClientStore(env.Orm, oauthXorm.WithClientStoreTableName("oauth_client"))
	manager.MapClientStorage(clientStore)

	cfg := server.NewConfig()
	cfg.AllowedGrantTypes = []oauth2.GrantType{
		oauth2.AuthorizationCode,
		oauth2.Refreshing,
	}

	Srv = server.NewServer(cfg, manager)

	Srv.SetUserAuthorizationHandler(userAuthorizeHandler)

	Srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		return
	})

	Srv.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("Response Error:", re.Error.Error())
	})

	Srv.SetAuthorizeScopeHandler(func(w http.ResponseWriter, r *http.Request) (scope string, err error) {
		if err := r.ParseForm(); err != nil {
			return "", err
		}

		reqScope := r.Form["scope"][0]
		clientId := r.Form["client_id"][0]

		client := model.OauthClient{
			Id: clientId,
		}

		has, err := env.Orm.Get(&client)
		if err != nil {
			return "", err
		}
		if !has {
			return "", fmt.Errorf("client scope not found error")
		}

		// Validate scope request
		reqScopes := strings.Split(reqScope, " ")
		for _, s := range reqScopes {
			err = validate.Var(s, fmt.Sprintf("oneof=%s", client.Scope))
			if err != nil {
				return "", err
			}
		}

		return reqScope, nil
	})
}

func userAuthorizeHandler(w http.ResponseWriter, r *http.Request) (userID string, err error) {
	store, err := session.Start(r.Context(), w, r)
	if err != nil {
		return
	}

	account, isPass := ValidateLogin(store)
	if !isPass {
		if r.Form == nil {
			_ = r.ParseForm()
		}

		store.Set("ReturnUri", r.Form)
		_ = store.Save()

		w.Header().Set("Location", LoginUrl)
		w.WriteHeader(http.StatusFound)
		return
	}

	isPass = ValidateAuthStatus(store)
	if !isPass {
		if r.Form == nil {
			_ = r.ParseForm()
		}

		store.Set("ReturnUri", r.Form)
		_ = store.Save()

		w.Header().Set("Location", LoginUrl)
		w.WriteHeader(http.StatusFound)
		return
	}

	userID = account
	store.Delete("LoggedInUserID")
	store.Delete("AuthorizationStatus")
	_ = store.Save()
	return
}
