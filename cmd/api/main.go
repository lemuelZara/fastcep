package main

import (
	"context"
	"fmt"
	"time"

	"github.com/lemuelZara/fastcep/app"
	brasilapi "github.com/lemuelZara/fastcep/app/brasilapi/repository/http"
	viacep "github.com/lemuelZara/fastcep/app/viacep/repository/http"
)

const cep string = "15610000"

type Result struct {
	Address  app.Address
	Service  string
	Duration time.Duration
	Error    error
}

func main() {
	ctx := context.Background()

	results := make(chan Result)

	b := brasilapi.NewClient()
	vc := viacep.NewClient()

	go func() {
		now := time.Now()
		a, err := b.GetByCEP(ctx, cep)
		results <- Result{a, "brasilapi", time.Since(now), fmt.Errorf("brasilapi error: %w", err)}

	}()
	go func() {
		now := time.Now()
		a, err := vc.GetByCEP(ctx, cep)
		results <- Result{a, "viacep", time.Since(now), fmt.Errorf("viacep error: %w", err)}

	}()

	select {
	case r := <-results:
		if r.Error != nil {
			fmt.Printf(r.Error.Error())
			return
		}

		outService := fmt.Sprintf("%s%s%s", "\033[32;1m", r.Service, "\033[0m")
		outDuration := fmt.Sprintf("%s%s%s", "\033[36m", r.Duration, "\033[0m")
		fmt.Printf("Success from %s: %v (Duration: %v)\n", outService, r.Address, outDuration)
	case <-ctx.Done():
		fmt.Printf("Context timed out! %v\n", ctx.Err())
	}
}
