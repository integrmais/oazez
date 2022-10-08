package oazes

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ProviderService Client

type Provider struct {
	Id           string `json:"codigo"`
	Sequence     string `json:"sequencia"`
	Name         string `json:"nome"`
	Street       string `json:"logradouro"`
	Number       string `json:"numero"`
	Neighborhood string `json:"bairro"`
	Complement   string `json:"complemento"`
	City         string `json:"city"`
	State        string `json:"uf"`
	Telephone    string `json:"telefone"`
	Mobilephone  string `json:"celular"`
	FederalTaxId string `json:"cpfcnpj"`
}

func (ps *ProviderService) List() ([]Provider, error) {
	url := fmt.Sprintf(
		"%s/api/%s/v1/clientes/prestadores", ps.BaseUrl, ps.DatabaseId,
	)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return []Provider{}, err
	}

	req.Header.Add("Authorization", "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJmaWxpYWwiOiIxIiwicGVyZmlsIjoiVVNVIiwidXNlciI6IjM3IiwiQ29ubmVjdGlvbklkIjoiMjY1MiIsIkNvbm5lY3Rpb25JZEFsaWFzIjoiU0FOQ0FSRCIsInN1YnNlcSI6IiIsInN1YnRpcG8iOiJVUyIsInN1Ym5vbWUiOiJJbnRlZ3JNYWlzIiwic3VidmVuZGVkb3IiOiIiLCJjb2RpZ29tYXN0ZXIiOiIiLCJzdWIiOiIzNyIsImlzcyI6IkludGVnck1haXMiLCJob21vbG9nYWNhbyI6ZmFsc2UsImV4cCI6MTY2NTIzNDYwMH0.Y1o8DX5HKdM5QcHy7d9ENoS3N2Y9MsQDuHOW4Bzz12E")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return []Provider{}, err
	}

	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	fmt.Println(url, string(body))

	var providers []Provider
	err = json.Unmarshal(body, &providers)
	if err != nil {
		return []Provider{}, err
	}

	return providers, nil
}
