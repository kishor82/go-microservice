package server

import (
	"context"

	"github.com/hashicorp/go-hclog"

	cp "github.com/kishor82/go-microservices/currency/protos/currency"
)

type Currency struct {
	cp.UnimplementedCurrencyServer
	log hclog.Logger
}

func NewCurrency(l hclog.Logger) *Currency {
	return &Currency{log: l}
}

func (c *Currency) GetRate(ctx context.Context, rr *cp.RateRequest) (*cp.RateResponse, error) {
	c.log.Info("Hallo GetRate", "base", rr.GetBase(), "destination", rr.GetDestination())

	return &cp.RateResponse{Rate: 0.5}, nil
}
