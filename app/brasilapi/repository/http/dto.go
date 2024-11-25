package http

import "github.com/lemuelZara/fastcep/app"

type result struct {
	CEP          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
}

func toAddress(r result) app.Address {
	return app.Address{
		CEP:          r.CEP,
		Street:       r.Street,
		Neighborhood: r.Neighborhood,
		City:         r.City,
		State:        r.State,
	}
}
