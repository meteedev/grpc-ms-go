package main

import (
	"log"

	"github.com/meteedev/grpc-ms-go/payment/internal/adapters/grpc"
	"github.com/meteedev/grpc-ms-go/payment/internal/adapters/db"
	"github.com/meteedev/grpc-ms-go/payment/config"
	"github.com/meteedev/grpc-ms-go/payment/internal/application/core/api"
)

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}

	application := api.NewApplication(dbAdapter)
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()
}
