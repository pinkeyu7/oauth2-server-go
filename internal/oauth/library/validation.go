package library

import (
	"encoding/json"
	"errors"
	"oauth2-server-go/dto/model"
	oauthClient "oauth2-server-go/internal/oauth/client"

	"github.com/gin-gonic/gin"
	"github.com/go-session/session"
)

func Validation(c *gin.Context, ocr oauthClient.Repository) (session.Store, *model.OauthClient, string, error) {
	store, err := session.Start(c, c.Writer, c.Request)
	if err != nil {
		return nil, nil, "", err
	}

	data, ok := store.Get("ReturnUri")
	if !ok {
		return nil, nil, "", errors.New("ReturnUri not found.")
	}

	reJson := model.OauthClientRedisCache{}
	marshal, err := json.Marshal(data)
	if err != nil {
		return nil, nil, "", err
	}

	err = json.Unmarshal(marshal, &reJson)
	if err != nil {
		return nil, nil, "", err
	}

	// 取得錯誤回傳URI
	redirectUri := reJson.RedirectURI[0]
	if redirectUri == "" {
		return nil, nil, "", errors.New("redirectUri is empty.")
	}

	// 取得 client 資訊如：名稱 顯示於網頁
	clientId := reJson.ClientID[0]
	if clientId == "" {
		return nil, nil, "", errors.New("clientId is empty.")
	}

	client, err := ocr.FindOne(&model.OauthClient{Id: clientId})
	if err != nil {
		return nil, nil, "", err
	}
	if client == nil {
		return nil, nil, "", errors.New("client not found.")
	}

	return store, client, redirectUri, nil
}

func ValidateLogin(store session.Store) (string, bool) {
	account, ok := store.Get("LoggedInUserID")
	if !ok {
		return "", false
	}

	return account.(string), true
}

func ValidateAuthStatus(store session.Store) bool {
	status, ok := store.Get("AuthorizationStatus")
	if !ok {
		return false
	}

	if status.(string) != StatusAuthorized {
		return false
	}

	return true
}
