package repository

import (
	"oauth2-server-go/dto/model"
	"oauth2-server-go/internal/oauth/scope"
	"oauth2-server-go/pkg/logr"

	"go.uber.org/zap"
	"xorm.io/xorm"
)

type Repository struct {
	orm   *xorm.EngineGroup
	cache scope.Cache
}

func NewRepository(orm *xorm.EngineGroup, c scope.Cache) scope.Repository {
	return &Repository{orm: orm, cache: c}
}

func (r *Repository) Find() ([]*model.OauthScope, error) {
	var err error
	scopes := make([]*model.OauthScope, 0)

	err = r.orm.Where(" is_disable =  ? ", 0).Find(&scopes)
	if err != nil {
		return nil, err
	}

	return scopes, nil
}

func (r *Repository) FindOne(scope *model.OauthScope) (*model.OauthScope, error) {
	// find from redis ----------------------------------------------------------------------
	scp, err := r.cache.FindOneScope(scope.Path, scope.Method)
	if err == nil {
		return scp, nil
	}

	// find from mysql ----------------------------------------------------------------------
	has, err := r.orm.Get(scope)
	if err != nil {
		return nil, err
	}

	if !has {
		return nil, nil
	}

	// set cache to redis
	err = r.cache.SetOneScope(scope.Path, scope.Method, scope)
	if err != nil {
		logr.L.Error("set scope cache error.", zap.String("error", err.Error()))
	}

	return scope, nil
}

func (r *Repository) FindScope() ([]string, error) {
	var err error
	scopes := make([]string, 0)

	err = r.orm.Table("oauth_scope").Where(" is_disable =  ? ", 0).Select("scope").Find(&scopes)
	if err != nil {
		return nil, err
	}

	return scopes, nil
}
