package main

import (
	"net"
	"os"

	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/kishor82/go-microservices/currency/protos/currency"
	"github.com/kishor82/go-microservices/currency/server"
)

func main() {
	log := hclog.Default()

	gs := grpc.NewServer()

	cp := server.NewCurrency(log)

	currency.RegisterCurrencyServer(gs, cp)
	reflection.Register(gs)

	l, err := net.Listen("tcp", ":9092")
	if err != nil {
		log.Error("Unable to Listen", "error", err)
		os.Exit(1)
	}

	gs.Serve(l)
}
