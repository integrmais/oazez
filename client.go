package oazes

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
