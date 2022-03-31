package repository

import (
	"encoding/json"
	"oauth2-server-go/dto/model"
	"oauth2-server-go/internal/oauth/scope"
	"time"

	"github.com/go-redis/redis"
)

type Cache struct {
	redis *redis.ClusterClient
}

func NewRedis(r *redis.ClusterClient) scope.Cache {
	return &Cache{redis: r}
}

func (c *Cache) SetOneScope(path, method string, scp *model.OauthScope) error {
	hashKey := scope.GetScopeHashKey()
	key := scope.GetScopeKey(path, method)
	jsonBytes, _ := json.Marshal(scp)

	err := c.redis.HSet(hashKey, key, jsonBytes).Err()
	if err != nil {
		return err
	}

	return err
}

func (c *Cache) FindOneScope(path, method string) (*model.OauthScope, error) {
	hashKey := scope.GetScopeHashKey()
	key := scope.GetScopeKey(path, method)
	scp := model.OauthScope{}

	res, err := c.redis.HGet(hashKey, key).Result()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(res), &scp)
	if err != nil {
		return nil, err
	}

	return &scp, nil
}

func (c *Cache) SetClientScopeList(clientId string, scopeList *model.ScopeList) error {
	key := scope.GetClientScopeListKey(clientId)
	jsonBytes, _ := json.Marshal(scopeList)

	err := c.redis.Set(key, jsonBytes, time.Hour*24).Err()
	if err != nil {
		return err
	}

	return err
}

func (c *Cache) FindClientScopeList(clientId string) (*model.ScopeList, error) {
	key := scope.GetClientScopeListKey(clientId)
	scopeList := make(model.ScopeList)

	res, err := c.redis.Get(key).Result()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(res), &scopeList)
	if err != nil {
		return nil, err
	}

	return &scopeList, nil
}
