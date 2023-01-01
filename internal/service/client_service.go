package service

import (
	"context"
	"simple-grpc-crud/protobuf/pb"
)

var _ ClientRepository = (*ClientService)(nil)

type ClientRepository interface {
	Create(ctx context.Context, in *pb.CreateClientRequest) (*pb.CreateClientResponse, error)
	Read(ctx context.Context, in *pb.ReadClientRequest) (*pb.ReadClientResponse, error)
	Update(ctx context.Context, in *pb.UpdateClientRequest) (*pb.UpdateClientResponse, error)
	Delete(ctx context.Context, in *pb.DeleteClientRequest) (*pb.DeleteClientResponse, error)
}

type ClientService struct {
	pb.UnimplementedClientGreeterServer
	repository ClientRepository
}

func NewClientService(repository ClientRepository) *ClientService {
	return &ClientService{repository: repository}
}

func (c *ClientService) Create(ctx context.Context, in *pb.CreateClientRequest) (*pb.CreateClientResponse, error) {
	return c.repository.Create(ctx, in)
}

func (c *ClientService) Read(ctx context.Context, in *pb.ReadClientRequest) (*pb.ReadClientResponse, error) {
	return c.repository.Read(ctx, in)
}

func (c *ClientService) Update(ctx context.Context, in *pb.UpdateClientRequest) (*pb.UpdateClientResponse, error) {
	return c.repository.Update(ctx, in)
}

func (c *ClientService) Delete(ctx context.Context, in *pb.DeleteClientRequest) (*pb.DeleteClientResponse, error) {
	return c.repository.Delete(ctx, in)
}
