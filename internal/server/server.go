package server

import (
	"fmt"
	"net"
	"os"
	"strconv"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	// pb "vendorService/proto"  // Uncomment after adding proto import
	// "vendorService/internal/service"  // Uncomment after adding the service implementation
)

type Server struct {
	port int
}

// NewServer creates a new gRPC server instance
func NewServer() *grpc.Server {
	grpcServer := grpc.NewServer() // Create the gRPC server instance
	return grpcServer
}

// StartGRPCServer starts the gRPC server on the specified port
func (s *Server) StartGRPCServer(grpcServer *grpc.Server) {
	port, _ := strconv.Atoi(os.Getenv("GRPC_PORT"))
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		zap.L().Fatal("Failed to listen", zap.Error(err))
	}

	// Register your gRPC services here (assuming `VendorService` as an example)
	// pb.RegisterVendorServiceServer(grpcServer, &service.VendorService{})

	zap.L().Info("Starting gRPC server", zap.Int("port", port))
	if err := grpcServer.Serve(lis); err != nil {
		zap.L().Fatal("Failed to serve", zap.Error(err))
	}
}
