package client

import "oauth2-server-go/dto/model"

type Repository interface {
	Insert(m *model.OauthClient) error
	Find(offset, limit int) ([]*model.OauthClient, error)
	FindOne(m *model.OauthClient) (*model.OauthClient, error)
	Count() (int, error)
	Update(m *model.OauthClient) error
	Delete(clientId string) error
}
