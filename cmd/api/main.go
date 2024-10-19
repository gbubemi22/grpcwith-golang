package main

import (
	"go.uber.org/zap"

	"vendorService/internal/server"
)

func main() {
	// Initialize zap logger
	logger, _ := zap.NewProduction()
	zap.ReplaceGlobals(logger)
	defer logger.Sync()

	// Create gRPC server
	grpcServer := server.NewServer()

	// Initialize and start the gRPC server
	srv := &server.Server{}
	srv.StartGRPCServer(grpcServer)
}
