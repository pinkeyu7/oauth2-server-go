package scope

type Service interface {
	GetScope(path, method string) (string, error)
	VerifyScope(clientId, scope string) (bool, error)
}
