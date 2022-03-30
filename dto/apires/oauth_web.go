package apires

type OauthLoginPage struct {
	HostIconPath string `json:"host_icon_path"`
	ClientName   string `json:"client_name"`
	AccountError bool   `json:"account_error"`
	RedirectUrl  string `json:"redirect_url"`
	BasePath     string `json:"base_path"`
}

type OauthAuthPage struct {
	HostIconPath   string   `json:"host_icon_path"`
	ClientIconPath string   `json:"client_icon_path"`
	ClientName     string   `json:"client_name"`
	RedirectUrl    string   `json:"redirect_url"`
	Scopes         []*Scope `json:"scopes"`
	UserName       string   `json:"user_name"`
	BasePath       string   `json:"base_path"`
}

type Scope struct {
	Name   string `json:"name"`
	Detail string `json:"detail"`
}
