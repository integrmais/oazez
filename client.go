package oazez

type Client struct {
	BaseUrl string

	// Database Connection Id
	DatabaseId string

	// Username and Password
	userName string
	passWord string

	Token string

	// Providers
	Provider *ProviderService
}

type Access struct {
	Username string `json:"Usuario"`
	Password string `json:"Senha"`
}
