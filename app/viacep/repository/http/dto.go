package http

import "github.com/lemuelZara/fastcep/app"

type result struct {
	CEP        string `json:"cep"`
	Logradouro string `json:"logradouro"`
	Bairro     string `json:"bairro"`
	Localidade string `json:"localidade"`
	Estado     string `json:"estado"`
}

func toAddress(r result) app.Address {
	return app.Address{
		CEP:          r.CEP,
		Street:       r.Logradouro,
		Neighborhood: r.Bairro,
		City:         r.Localidade,
		State:        r.Estado,
	}
}
