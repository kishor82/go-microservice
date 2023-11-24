package server

import (
	"context"

	"github.com/hashicorp/go-hclog"

	"github.com/kishor82/go-microservices/currency/data"
	cp "github.com/kishor82/go-microservices/currency/protos/currency"
)

type Currency struct {
	cp.UnimplementedCurrencyServer
	rates *data.ExchangeRates
	log   hclog.Logger
}

func NewCurrency(r *data.ExchangeRates, l hclog.Logger) *Currency {
	return &Currency{rates: r, log: l}
}

func (c *Currency) GetRate(ctx context.Context, rr *cp.RateRequest) (*cp.RateResponse, error) {
	c.log.Info("Hallo GetRate", "base", rr.GetBase(), "destination", rr.GetDestination())

	rate, err := c.rates.GetRate(rr.GetBase().String(), rr.Destination.String())
	if err != nil {
		return nil, err
	}

	return &cp.RateResponse{Rate: (rate)}, nil
}
