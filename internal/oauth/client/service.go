package client

import (
	"oauth2-server-go/dto/apireq"
	"oauth2-server-go/dto/apires"
	"oauth2-server-go/dto/model"
)

type Service interface {
	List(req *apireq.ListOauthClient) (*apires.ListOauthClient, error)
	Get(clientId int) (*model.OauthClient, error)
	Add(req *apireq.AddOauthClient) error
	Edit(clientId int, req *apireq.EditOauthClient) error
	Delete(clientId int) error
}
