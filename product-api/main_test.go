package main

import (
	"fmt"
	"testing"

	"github.com/kishor82/go-microservices/product-api/sdk/client"
	"github.com/kishor82/go-microservices/product-api/sdk/client/products"
)

func TestOurClient(t *testing.T) {
	cfg := client.DefaultTransportConfig().WithHost("localhost:9090")
	c := client.NewHTTPClientWithConfig(nil, cfg)

	params := products.NewListProductsParams()
	currency := "USD"
	prod, err := c.Products.ListProducts(params.WithCurrency(&currency))
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%#v", prod.GetPayload()[0])
}
