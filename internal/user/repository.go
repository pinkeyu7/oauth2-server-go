package user

import "oauth2-server-go/dto/model"

type Repository interface {
	Insert(m *model.User) error
	Find(offset, limit int) ([]*model.User, error)
	FindOne(m *model.User) (*model.User, error)
	Count() (int, error)
	Update(m *model.User) error
	Delete(userId int) error
}
