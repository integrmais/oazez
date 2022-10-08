package oazes

func NewClient(baseUrl, databaseId, username, password string) *Client {
	c := &Client{
		BaseUrl:    baseUrl,
		DatabaseId: databaseId,
		userName:   username,
		passWord:   password,
	}

	c.Provider = (*ProviderService)(c)

	return c
}
