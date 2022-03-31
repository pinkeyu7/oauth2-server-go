package repository

import (
	"oauth2-server-go/dto/model"
	"oauth2-server-go/internal/user"

	"xorm.io/xorm"
)

type Repository struct {
	orm *xorm.EngineGroup
}

func NewRepository(orm *xorm.EngineGroup) user.Repository {
	return &Repository{
		orm: orm,
	}
}

func (r *Repository) Insert(m *model.User) error {
	_, err := r.orm.Insert(m)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) Find(offset, limit int) ([]*model.User, error) {
	list := make([]*model.User, 0)

	err := r.orm.Limit(limit, offset).Find(&list)
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (r *Repository) FindOne(m *model.User) (*model.User, error) {
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
	count, err := r.orm.Count(&model.User{})
	if err != nil {
		return 0, err
	}

	return int(count), nil
}

func (r *Repository) Update(m *model.User) error {
	_, err := r.orm.ID(m.Id).Update(m)
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) Delete(userId int) error {
	_, err := r.orm.Delete(&model.User{Id: userId})
	if err != nil {
		return err
	}

	return nil
}
