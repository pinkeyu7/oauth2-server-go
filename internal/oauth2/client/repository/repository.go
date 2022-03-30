package repository

import (
	"oauth2-server-go/dto/model"
	oauthClient "oauth2-server-go/internal/oauth2/client"

	"xorm.io/xorm"
)

type Repository struct {
	orm *xorm.EngineGroup
}

func NewRepository(orm *xorm.EngineGroup) oauthClient.Repository {
	return &Repository{
		orm: orm,
	}
}

func (r *Repository) Insert(m *model.OauthClient) error {
	return nil
}

func (r *Repository) Find(offset, limit int) ([]*model.OauthClient, error) {
	return nil, nil
}

func (r *Repository) FindOne(m *model.OauthClient) (*model.OauthClient, error) {
	return nil, nil
}

func (r *Repository) Count() (int, error) {
	return 0, nil
}

func (r *Repository) Update(m *model.OauthClient) error {
	return nil
}

func (r *Repository) Delete(clientId int) error {
	return nil
}
