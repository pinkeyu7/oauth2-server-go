package oauthClient

import (
	"oauth2-server-go/dto/apireq"
	"oauth2-server-go/dto/apires"
	"oauth2-server-go/dto/model"
)

type Service interface {
	List(req *apireq.ListOauthClient) (*apires.ListOauthClient, error)
	Get(contactId int) (*model.OauthClient, error)
	Add(req *apireq.AddOauthClient) error
	Edit(contactId int, req *apireq.EditOauthClient) error
	Delete(contactId int) error
}
