package grpc

import (
	"fmt"
	"net"

	"github.com/meteedev/grpc-ms-go/order/config"
	"github.com/meteedev/grpc-ms-go/order/internal/ports"
	"github.com/meteedev/grpc-ms-go/proto/go_proto"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Adapter struct {
	api    ports.APIPort
	port   int
	server *grpc.Server
	grpc_order.UnimplementedOrderServer
}

func NewAdapter(api ports.APIPort, port int) *Adapter {
	return &Adapter{api: api, port: port}
}

func (a Adapter) Run() {
	var err error

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		log.Fatalf("failed to listen on port %d, error: %v", a.port, err)
	}

	grpcServer := grpc.NewServer()
	a.server = grpcServer
	grpc_order.RegisterOrderServer(grpcServer, a)
	if config.GetEnv() == "dev" {
		reflection.Register(grpcServer)
	}

	log.Printf("starting order service on port %d ...", a.port)
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve grpc on port %d", a.port)
	}
}

func (a Adapter) Stop() {
	a.server.Stop()
}
