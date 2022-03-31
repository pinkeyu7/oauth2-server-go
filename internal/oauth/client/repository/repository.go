package repository

import (
	"oauth2-server-go/dto/model"
	"oauth2-server-go/internal/oauth/client"

	"xorm.io/xorm"
)

type Repository struct {
	orm *xorm.EngineGroup
}

func NewRepository(orm *xorm.EngineGroup) client.Repository {
	return &Repository{
		orm: orm,
	}
}

func (r *Repository) Insert(m *model.OauthClient) error {
	_, err := r.orm.Insert(m)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) Find(offset, limit int) ([]*model.OauthClient, error) {
	list := make([]*model.OauthClient, 0)

	err := r.orm.Limit(limit, offset).Find(&list)
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (r *Repository) FindOne(m *model.OauthClient) (*model.OauthClient, error) {
	has, err := r.orm.Get(m)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}

	return m, nil
}

func (r *Repository) Count() (int, error) {
	count, err := r.orm.Count(&model.OauthClient{})
	if err != nil {
		return 0, err
	}

	return int(count), nil
}

func (r *Repository) Update(m *model.OauthClient) error {
	_, err := r.orm.Where("id = ? ", m.Id).Cols("sys_account_id", "name", "secret", "domain", "scope", "icon_path", "data").Update(m)
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) Delete(cliendId string) error {
	_, err := r.orm.Delete(&model.OauthClient{Id: cliendId})
	if err != nil {
		return err
	}

	return nil
}
