package server

import (
	"context"
	"io"
	"time"

	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/kishor82/go-microservices/currency/data"
	cp "github.com/kishor82/go-microservices/currency/protos/currency"
)

type Currency struct {
	cp.UnimplementedCurrencyServer
	rates         *data.ExchangeRates
	log           hclog.Logger
	subscriptions map[cp.Currency_SubscribeRatesServer][]*cp.RateRequest
}

func NewCurrency(r *data.ExchangeRates, l hclog.Logger) *Currency {
	c := &Currency{
		rates:         r,
		log:           l,
		subscriptions: make(map[cp.Currency_SubscribeRatesServer][]*cp.RateRequest),
	}
	go c.handleUpdates()
	return c
}

func (c *Currency) handleUpdates() {
	ru := c.rates.MonitorRates(5 * time.Second)

	for range ru {
		c.log.Info("Got Updated rates")
		// loop over subscribed clients
		for k, v := range c.subscriptions {
			// loop over subscribed rates
			for _, rr := range v {
				r, err := c.rates.GetRate(rr.GetBase().String(), rr.GetDestination().String())
				if err != nil {
					c.log.Error(
						"Unable to get update rate",
						"base",
						rr.GetBase().String(),
						"destination",
						rr.GetDestination().String(),
					)
				}
				err = k.Send(
					&cp.StreamingRateResponse{
						Message: &cp.StreamingRateResponse_RateResponse{
							RateResponse: &cp.RateResponse{
								Base:        rr.Base,
								Destination: rr.Destination,
								Rate:        r,
							},
						},
					},
				)
				if err != nil {
					c.log.Error(
						"Unable to send updated reate",
						"base",
						rr.GetBase().String(),
						"destination",
						rr.GetDestination().String(),
					)
				}
			}
		}
	}
}

func (c *Currency) GetRate(ctx context.Context, rr *cp.RateRequest) (*cp.RateResponse, error) {
	c.log.Info("GetRate", "base", rr.GetBase(), "destination", rr.GetDestination())

	if rr.Base == rr.Destination {
		// create the grpc error and return to the client
		err := status.Newf(
			codes.InvalidArgument,
			"Base currency %s can not be same as destination currency %s",
			rr.Base.String(),
			rr.Destination.String(),
		)

		err, wde := err.WithDetails(rr)
		if wde != nil {
			return nil, wde
		}

		return nil, err.Err()
	}

	rate, err := c.rates.GetRate(rr.GetBase().String(), rr.GetDestination().String())
	if err != nil {
		return nil, err
	}

	return &cp.RateResponse{Base: rr.Base, Destination: rr.Destination, Rate: rate}, nil
}

// SubscribeRates implements the gRPC bidirection streaming method for the server
func (c *Currency) SubscribeRates(src cp.Currency_SubscribeRatesServer) error {
	// handle client messages
	for {
		rr, err := src.Recv() // Recv is a blocking method which returns client data
		// io.EOF signals that the client has closed the connection
		if err == io.EOF {
			c.log.Info("Client has closed connection")
			break
		}
		// any other error means the transport between the server and client is unavailable
		if err != nil {
			c.log.Error("Unable to read from client", "error", err)
			break
		}

		c.log.Info("Handle client request", "request", rr)
		rrs, ok := c.subscriptions[src]
		if !ok {
			rrs = []*cp.RateRequest{}
		}

		// check that subscription does not exists
		var validationError *status.Status
		for _, v := range rrs {
			if v.Base == rr.Base && v.Destination == rr.Destination {
				// subscription exists return errors
				// NOTE: here bymistake i created newVariable with := which override the aboce declared validationError due to scope :(
				// INFO:  In Go, if you declare a variable with the same name inside a block (such as the for loop in your case), it will create a new variable with that name, shadowing any variables with the same name in the outer scope.
				// to prevent this in future i added "analyses": {"shadow": true} setting in gopls lsp config.
				validationError = status.Newf(
					codes.AlreadyExists,
					"Unable to subscribe for currency as subscription already exist",
				)
				// add the original request as metadata
				validationError, err = validationError.WithDetails(rr)
				if err != nil {
					c.log.Error("Unable to add metadata to error", "error", err)
					break
				}
				break
			}
		}
		if validationError != nil {
			err := src.Send(
				&cp.StreamingRateResponse{
					Message: &cp.StreamingRateResponse_Error{
						Error: validationError.Proto(),
					},
				},
			)
			if err != nil {
				c.log.Error(
					"Unable to send error",
					"base",
					rr.GetBase().String(),
					"destination",
					rr.GetDestination().String(),
				)
			}
			continue
		}

		rrs = append(rrs, rr)
		c.subscriptions[src] = rrs
	}
	return nil
}
