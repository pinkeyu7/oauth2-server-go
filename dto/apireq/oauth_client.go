package apireq

type ListOauthClient struct {
	AccountId int `form:"account_id" validate:"required"`
	Page      int `form:"page" validate:"required,numeric"`
	PerPage   int `form:"per_page" validate:"required,numeric"`
}

type GetOauthClient struct {
	AccountId int `form:"account_id" validate:"required"`
}

type AddOauthClient struct {
	AccountId int    `json:"account_id" validate:"required"`
	Id        string `json:"id" validate:"required,max=255"`
	Secret    string `json:"secret" validate:"required,max=255"`
	Domain    string `json:"domain" validate:"required,max=255"`
}

type EditOauthClient struct {
	AccountId int    `json:"account_id" validate:"required"`
	Secret    string `json:"secret" validate:"required,max=255"`
	Domain    string `json:"domain" validate:"required,max=255"`
}

type DeleteOauthClient struct {
	AccountId int `json:"account_id" validate:"required"`
}
