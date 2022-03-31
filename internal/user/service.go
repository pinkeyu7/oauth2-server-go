package user

import (
	"oauth2-server-go/dto/model"
)

type Service interface {
	Get(account string) (*model.User, error)
	Verify(account, password string) error
}
