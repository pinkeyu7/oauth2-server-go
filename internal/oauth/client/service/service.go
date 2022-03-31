package service

import (
	"oauth2-server-go/dto/apireq"
	"oauth2-server-go/dto/apires"
	"oauth2-server-go/dto/model"
	"oauth2-server-go/internal/oauth/client"
)

type Service struct {
	clientRepo client.Repository
}

func NewService(ocr client.Repository) client.Service {
	return &Service{
		clientRepo: ocr,
	}
}

func (s *Service) List(req *apireq.ListOauthClient) (*apires.ListOauthClient, error) {
	return nil, nil
}

func (s *Service) Get(clientId int) (*model.OauthClient, error) {
	return nil, nil
}

func (s *Service) Add(req *apireq.AddOauthClient) error {
	return nil
}

func (s *Service) Edit(clientId int, req *apireq.EditOauthClient) error {
	return nil
}

func (s *Service) Delete(clientId int) error {
	return nil
}
