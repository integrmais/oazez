package oazes

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const DefaultContentType = "application/json"

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

func (c *Client) SetAccessToken() error {
	connectUrl := fmt.Sprintf(
		"%s/api/%s/v1/login",
		c.BaseUrl,
		c.DatabaseId,
	)

	access := &Access{
		Username: c.userName,
		Password: c.passWord,
	}

	b, err := json.Marshal(access)
	if err != nil {
		return err
	}

	req, err := http.DefaultClient.Post(connectUrl, DefaultContentType, strings.NewReader(string(b)))

	if err != nil {
		return err
	}

	defer req.Body.Close()

	if req.StatusCode != http.StatusOK {
		return errors.New(
			fmt.Sprintf("Error trying to login: %d", req.StatusCode),
		)
	}

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return err
	}

	c.Token = string(body)

	return nil
}
