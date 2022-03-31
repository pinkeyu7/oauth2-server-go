package scope

import (
	"fmt"
	"oauth2-server-go/dto/model"
)

type Repository interface {
	Find() ([]*model.OauthScope, error)
	FindOne(scope *model.OauthScope) (*model.OauthScope, error)
	FindScope() ([]string, error)
}

type Cache interface {
	SetOneScope(path, method string, scope *model.OauthScope) error
	FindOneScope(path, method string) (*model.OauthScope, error)
	SetClientScopeList(clientId string, scopeList *model.ScopeList) error
	FindClientScopeList(clientId string) (*model.ScopeList, error)
}

func GetScopeHashKey() string {
	return "oauth_scope"
}

func GetScopeKey(path, method string) string {
	return fmt.Sprintf("%s:%s", path, method)
}

func GetClientScopeListKey(clientId string) string {
	return fmt.Sprintf("oauth_client:%s:scope_list", clientId)
}
