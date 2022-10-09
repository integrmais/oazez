package oazes

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ProviderService Client

type Provider struct {
	Id           int64  `json:"codigo"`
	Sequence     int64  `json:"sequencia"`
	Name         string `json:"nome"`
	Street       string `json:"logradouro"`
	Number       string `json:"numero"`
	Neighborhood string `json:"bairro"`
	Complement   string `json:"complemento"`
	City         string `json:"city"`
	State        string `json:"uf"`
	Telephone    string `json:"telefone"`
	Mobilephone  string `json:"celular"`
	FederalTaxId int64  `json:"cpfcnpj"`
}

func (ps *ProviderService) List() ([]Provider, error) {
	url := fmt.Sprintf(
		"%s/api/%s/v1/clientes/prestadores", ps.BaseUrl, ps.DatabaseId,
	)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return []Provider{}, err
	}

	req.Header.Add("Authorization", "Bearer "+ps.Token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return []Provider{}, err
	}

	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	var providers []Provider
	err = json.Unmarshal(body, &providers)
	if err != nil {
		return []Provider{}, err
	}

	return providers, nil
}
