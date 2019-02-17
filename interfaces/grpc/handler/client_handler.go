package handler

import (
	"context"
	"github.com/pkg/errors"
	"log"

	pb "github.com/devishot/so-go-grpc-client_project/interfaces/grpc/api"
)

var NotImplementedError = errors.New("gRPC not implented error")

type ClientService struct {
}

func (s *ClientService) CreateClient(ctx context.Context, cl *pb.Client) (*pb.Client, error) {
	log.Printf("createClient: client=%v", cl)
	return nil, NotImplementedError
}

func (s *ClientService) DeleteClient(ctx context.Context, req *pb.DeleteClientRequest) (*pb.Client, error) {
	log.Printf("deleteClient: request=%v", req)
	return nil, NotImplementedError
}

func (s *ClientService) GetClient(ctx context.Context, req *pb.GetClientRequest) (*pb.Client, error) {
	log.Printf("getClient: request=%v", req)
	return nil, NotImplementedError
}
