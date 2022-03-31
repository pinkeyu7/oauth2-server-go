package oauthLibrary

import (
	"oauth2-server-go/dto/model"
	"oauth2-server-go/pkg/er"
	"strings"
)

func CheckScope(scopeList *model.ScopeList, scope string) (bool, error) {
	scpList := *scopeList
	// 檢查權限的欄位格式
	_, category, item, err := ParseScope(scope)
	if err != nil {
		return false, err
	}

	// 檢查 Level 1 授權
	if scpList[category] == nil {
		return false, nil
	}
	if scpList[category].IsAuth {
		return true, nil
	}

	// 檢查 Level 2 授權
	if scpList[category].Items[item] == nil {
		return false, nil
	}
	if scpList[category].Items[item].IsAuth {
		return true, nil
	}

	return false, nil
}

func GenerateClientScopeList(scopeList *model.ScopeList, clientScopeStr string) (*model.ScopeList, error) {
	scpList := *scopeList
	clientScopes := strings.Split(clientScopeStr, " ")
	for _, scope := range clientScopes {
		// 檢查權限的欄位格式
		level, category, item, err := ParseScope(scope)
		if err != nil {
			return nil, err
		}

		switch level {
		case 1:
			// 設定 Level 1 授權
			if scpList[category] == nil {
				continue
			}
			if scpList[category] != nil {
				scpList[category].IsAuth = true
			}
		case 2:
			// 設定 Level 2 授權
			if scpList[category] == nil {
				continue
			}
			if scpList[category].Items[item] == nil {
				continue
			}
			if scpList[category].Items[item] != nil {
				scpList[category].Items[item].IsAuth = true
			}
		}
	}

	return scopeList, nil
}

func GenerateScopeList(scopes []string) (*model.ScopeList, error) {
	scopeList := make(model.ScopeList, 0)

	for _, scope := range scopes {
		// 檢查權限的欄位格式
		level, category, item, err := ParseScope(scope)
		if err != nil {
			return nil, err
		}
		if level != 2 {
			notMatchErr := er.NewAppErr(400, er.ErrorParamInvalid, "scope level error.", nil)
			return nil, notMatchErr
		}

		if scopeList[category] == nil {
			scopeList[category] = &model.ScopeCategory{
				Name:   category,
				Items:  make(map[string]*model.ScopeItem, 0),
				IsAuth: false,
			}
		}

		if scopeList[category].Items[item] == nil {
			scopeList[category].Items[item] = &model.ScopeItem{
				Name:   item,
				IsAuth: false,
			}
		}
	}

	return &scopeList, nil
}

func ParseScope(scope string) (int, string, string, error) {
	var level int
	var category string
	var item string

	if len(scope) == 0 {
		emptyErr := er.NewAppErr(400, er.ErrorParamInvalid, "scope empty error.", nil)
		return 0, "", "", emptyErr
	}

	arr := strings.Split(scope, ".")
	level = len(arr)

	switch len(arr) {
	case 1:
		category = arr[0]
	case 2:
		category = arr[0]
		item = arr[1]
	default:
		formatErr := er.NewAppErr(400, er.ErrorParamInvalid, "scope format error.", nil)
		return 0, "", "", formatErr
	}
	return level, category, item, nil
}
