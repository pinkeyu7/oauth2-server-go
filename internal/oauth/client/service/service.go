package service

import (
	"oauth2-server-go/dto/apireq"
	"oauth2-server-go/dto/apires"
	"oauth2-server-go/dto/model"
	oauthClient "oauth2-server-go/internal/oauth/client"
)

type Service struct {
	clientRepo oauthClient.Repository
}

func NewService(ocr oauthClient.Repository) oauthClient.Service {
	return &Service{
		clientRepo: ocr,
	}
}

func (s *Service) List(req *apireq.ListOauthClient) (*apires.ListOauthClient, error) {
	return nil, nil
}

func (s *Service) Get(contactId int) (*model.OauthClient, error) {
	return nil, nil
}

func (s *Service) Add(req *apireq.AddOauthClient) error {
	return nil
}

func (s *Service) Edit(contactId int, req *apireq.EditOauthClient) error {
	return nil
}

func (s *Service) Delete(contactId int) error {
	return nil
}
