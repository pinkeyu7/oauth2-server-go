package apireq

type GetUser struct {
	ClientId string `json:"client_id"`
	Account  string `json:"account"`
}
