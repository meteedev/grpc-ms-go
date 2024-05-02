package grpc

import (
	"fmt"

	"log"
	"net"

	"github.com/meteedev/grpc-ms-go/payment/config"
	"github.com/meteedev/grpc-ms-go/payment/proto/go_proto"
	"google.golang.org/grpc/reflection"

	"google.golang.org/grpc"
)

func (a Adapter) Run() {
	var err error

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		log.Fatalf("failed to listen on port %d, error: %v", a.port, err)
	}

	grpcServer := grpc.NewServer()
	a.server = grpcServer
	grpc_payment.RegisterPaymentServer(grpcServer, a)
	if config.GetEnv() == "dev" {
		reflection.Register(grpcServer)
	}

	log.Printf("starting payment service on port %d ...", a.port)
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve grpc on port ")
	}
}
