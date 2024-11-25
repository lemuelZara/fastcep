package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/lemuelZara/fastcep/app"
)

type Client struct {
	executor http.Client
}

func NewClient() Client {
	executor := http.Client{}

	return Client{executor}
}

func (c Client) GetByCEP(ctx context.Context, cep string) (app.Address, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	path := fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s ", cep)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, path, nil)
	if err != nil {
		return app.Address{}, err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := c.executor.Do(req)
	if err != nil {
		return app.Address{}, fmt.Errorf(`failed on getting cep: %w`, err)
	}
	defer res.Body.Close()

	return parseResponse(res)
}

func parseResponse(resp *http.Response) (app.Address, error) {
	var r result

	err := json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		return app.Address{}, fmt.Errorf("parse address failed: %w", err)
	}

	return toAddress(r), nil
}
