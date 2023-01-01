package server

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
	"simple-grpc-crud/config"
	"simple-grpc-crud/internal/service"
	"simple-grpc-crud/protobuf/pb"
)

type GrpcServer struct {
	grpcServer *grpc.Server
}

func NewGrpcServer(grpcServer *grpc.Server) *GrpcServer {
	return &GrpcServer{grpcServer: grpcServer}
}

func (s *GrpcServer) Start(cfg *config.Config, clientService *service.ClientService) error {
	address := getAddrString(cfg)

	lis, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	defer lis.Close()

	pb.RegisterClientGreeterServer(s.grpcServer, clientService)

	err = s.grpcServer.Serve(lis)
	if err != nil {
		return err
	}

	return nil
}

func getAddrString(cfg *config.Config) string {
	return fmt.Sprintf("%s:%s", cfg.GrpcAddr, cfg.GrpcPort)
}
