package apires

import (
	"oauth2-server-go/dto/model"

	"gopkg.in/guregu/null.v4"
)

type ListOauthClient struct {
	List        []*model.OauthClient `json:"list"`
	Total       int                  `json:"total"`
	CurrentPage int                  `json:"current_page"`
	PerPage     int                  `json:"per_page"`
	NextPage    null.Int             `json:"next_page" swaggertype:"string"`
}
