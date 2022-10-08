package oazes_test

import (
	"testing"

	oazez "github.com/integrmais/oazes"
)

const baseUrlMock string = "http://localhost"
const databaseIdMock string = "databaseIdMock"
const userNameMock string = "username-mock"
const passWordMock string = "password-mock"

func TestNewClient(t *testing.T) {
	c := oazez.NewClient(
		baseUrlMock,
		databaseIdMock,
		userNameMock,
		passWordMock,
	)

	if c.BaseUrl != baseUrlMock {
		t.Errorf("Expected baseUrl to be %s, got %s", baseUrlMock, c.BaseUrl)
	}
}
