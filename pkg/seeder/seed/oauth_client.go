package seed

import (
	"oauth2-server-go/dto/model"

	"xorm.io/xorm"
)

func CreateOauthClient(engine *xorm.Engine, id, secret, domain, data string) error {
	con := model.OauthClient{
		Id:     id,
		Secret: secret,
		Domain: domain,
		Data:   data,
	}

	_, err := engine.Insert(&con)
	return err
}

func AllOauthClient() []Seed {
	return []Seed{
		{
			Name: "Create Oauth Client",
			Run: func(engine *xorm.Engine) error {
				err := CreateOauthClient(engine, "test_oauth_client_id", "test_oauth_client_secret", "http:localhost:9094", "")
				if err != nil {
					return err
				}
				return nil
			},
		},
	}
}
