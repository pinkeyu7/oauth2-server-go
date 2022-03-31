package service

import (
	"net/http"
	"oauth2-server-go/dto/model"
	"oauth2-server-go/internal/user"
	"oauth2-server-go/pkg/er"
	"oauth2-server-go/pkg/helper"
)

type Service struct {
	userRepo user.Repository
}

func NewService(ur user.Repository) user.Service {
	return &Service{
		userRepo: ur,
	}
}

func (s *Service) Get(userId int) (*model.User, error) {
	return nil, nil
}

func (s *Service) Verify(account, password string) error {
	usr, err := s.userRepo.FindOne(&model.User{Account: account})
	if err != nil {
		findErr := er.NewAppErr(http.StatusInternalServerError, er.UnknownError, "find user error.", err)
		return findErr
	}
	if usr == nil {
		notFoundErr := er.NewAppErr(http.StatusBadRequest, er.ResourceNotFoundError, "user not found.", nil)
		return notFoundErr
	}

	// user validation
	password = helper.ScryptStr(password)
	if usr.Password != password {
		notMatchErr := er.NewAppErr(http.StatusBadRequest, er.ErrorParamInvalid, "", nil)
		return notMatchErr
	}

	return nil
}
