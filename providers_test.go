package oazes_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/integrmais/oazes"
)

func TestListProviders(t *testing.T) {
	var serverMock = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusOK)
		res.Write([]byte(`[
			{
				"codigo": 10001,
				"sequencia": 0,
				"nome": "Prestador da Silva",
				"logradouro": "Rua Domingos Leme",
				"numero": "1",
				"bairro": "Bairro",
				"complemento": "",
				"cidade": "Ariquemes",
				"uf": "RO",
				"telefone": "",
				"celular": "",
				"cpfcnpj": 123000000456
			},
			{
				"codigo": 10002,
				"sequencia": 0,
				"nome": "Prestador dos Santos",
				"logradouro": "",
				"numero": "2",
				"bairro": "Bairro",
				"complemento": "",
				"cidade": "Ariquemes",
				"uf": "RO",
				"telefone": "",
				"celular": "6599XXXX99",
				"cpfcnpj": 123000000456
			},
			{
				"codigo": 10003,
				"sequencia": 0,
				"nome": "Prestador da Shopee",
				"logradouro": "",
				"numero": "2",
				"bairro": "Victor Konder",
				"complemento": "",
				"cidade": "Blumenau",
				"uf": "SC",
				"telefone": "489XXXXXXXX",
				"celular": "6599XXXX99",
				"cpfcnpj": 23000000456
			}
		]`))
	}))

	defer func() { serverMock.Close() }()

	c := oazes.NewClient(serverMock.URL, databaseIdMock, userNameMock, passWordMock)

	list, err := c.Provider.List()

	if err != nil {
		t.Errorf("Error when list providers, got %s", err.Error())
	}

	if len(list) != 3 {
		t.Errorf("Expected 3 providers, got %d", len(list))
	}

	if list[0].Id != 10001 {
		t.Errorf("Expected 10001 providers, got %d", list[0].Id)
	}
}
