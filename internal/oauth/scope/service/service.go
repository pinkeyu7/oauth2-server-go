package service

import (
	"net/http"
	"oauth2-server-go/dto/model"
	"oauth2-server-go/internal/oauth/client"
	"oauth2-server-go/internal/oauth/library"
	"oauth2-server-go/internal/oauth/scope"
	"oauth2-server-go/pkg/er"
	"oauth2-server-go/pkg/logr"

	"go.uber.org/zap"
)

type Service struct {
	scopeRepo  scope.Repository
	scopeCache scope.Cache
	clientRepo client.Repository
}

func NewService(osr scope.Repository, osc scope.Cache, ocr client.Repository) scope.Service {
	return &Service{
		scopeRepo:  osr,
		scopeCache: osc,
		clientRepo: ocr,
	}
}

func (s *Service) GetScope(path, method string) (string, error) {
	scp, err := s.scopeRepo.FindOne(&model.OauthScope{Path: path, Method: method})
	if err != nil {
		findErr := er.NewAppErr(http.StatusInternalServerError, er.UnknownError, "find scope error.", err)
		return "", findErr
	}
	if scp == nil {
		notFoundErr := er.NewAppErr(http.StatusBadRequest, er.ResourceNotFoundError, "scope not found.", nil)
		return "", notFoundErr
	}

	return scp.Scope, nil
}

func (s *Service) VerifyScope(clientId, scope string) (bool, error) {
	// 從快取取得授權清單，若是無法取得，則重新建立授權清單並儲存回快取
	scopeList, err := s.scopeCache.FindClientScopeList(clientId)
	if err != nil {
		// 取得所有 API 列表
		apis, err := s.scopeRepo.FindScope()
		if err != nil {
			findErr := er.NewAppErr(http.StatusInternalServerError, er.UnknownError, "find scope error.", err)
			return false, findErr
		}

		// 取得 client app 資訊
		clt, err := s.clientRepo.FindOne(&model.OauthClient{Id: clientId})
		if err != nil {
			findErr := er.NewAppErr(http.StatusInternalServerError, er.UnknownError, "find client error.", err)
			return false, findErr
		}

		// 從 API 列表建立授權清單
		scopeList, err = library.GenerateScopeList(apis)
		if err != nil {
			return false, err
		}

		// 依據 client app 的 scope 設定授權清單
		scopeList, err = library.GenerateClientScopeList(scopeList, clt.Scope)
		if err != nil {
			return false, err
		}

		// 將建立完成的授權清單儲存回快取
		err = s.scopeCache.SetClientScopeList(clientId, scopeList)
		if err != nil {
			logr.L.Error("set scope list cache error.", zap.String("error", err.Error()))
		}
	}

	// 檢查 API 是否在授權清單內
	isPass, err := library.CheckScope(scopeList, scope)
	if err != nil {
		return false, err
	}

	return isPass, nil
}
