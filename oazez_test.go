package oazez_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	oazez "github.com/integrmais/oazez"
)

const baseUrlMock string = "http://localhost"
const databaseIdMock string = "databaseIdMock"
const userNameMock string = "username-mock"
const passWordMock string = "password-mock"

const accessTokenMock string = "base64-access-token"

func TestNewClient(t *testing.T) {
	var serverMock = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusOK)
		res.Write([]byte(accessTokenMock))
	}))

	defer func() { serverMock.Close() }()

	c := oazez.NewClient(
		serverMock.URL,
		databaseIdMock,
		userNameMock,
		passWordMock,
	)

	if c.BaseUrl != serverMock.URL {
		t.Errorf("Expected baseUrl to be %s, got %s", serverMock.URL, c.BaseUrl)
	}

	err := c.SetAccessToken()
	if err != nil {
		t.Errorf("Expected setAccessToken call sucessfully, got %s", err.Error())
	}

	if c.Token != accessTokenMock {
		t.Errorf("Expected Token to be %s, got %s", accessTokenMock, c.Token)
	}
}
