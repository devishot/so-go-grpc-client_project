package handler

import (
	"context"
	"log"

	"github.com/devishot/so-go-grpc-client_project/domain"
	connServices "github.com/devishot/so-go-grpc-client_project/interfaces/graphql_connection/services"
	pb "github.com/devishot/so-go-grpc-client_project/interfaces/grpc/api"
)

type ClientService struct {
	Service     domain.ClientService
	ConnBuilder connServices.Builder
}

func (s ClientService) CreateClient(ctx context.Context, req *pb.CreateClientRequest) (*pb.Client, error) {
	log.Printf("createClient: req=%v", req)

	cl, err := s.Service.Create(decodeProtoClient(req.Data))
	if err != nil {
		return nil, err
	}

	return encodeProtoClient(cl), nil
}

func (s ClientService) DeleteClient(ctx context.Context, req *pb.DeleteClientRequest) (*pb.Client, error) {
	log.Printf("deleteClient: request=%v", req)

	cl, err := s.Service.Delete(decodeID(req.Id))
	if err != nil {
		return nil, err
	}

	return encodeProtoClient(cl), nil
}

func (s ClientService) GetClient(ctx context.Context, req *pb.GetClientRequest) (*pb.Client, error) {
	log.Printf("getClient: request=%v", req)

	cl, err := s.Service.Get(decodeID(req.Id))
	if err != nil {
		return nil, err
	}

	return encodeProtoClient(cl), nil
}

func (s ClientService) GetClientConnection(ctx context.Context, req *pb.GetClientConnectionRequest) (*pb.ClientConnectionResponse, error) {
	log.Printf("getClientConnection: request=%v", req)

	connService := s.ConnBuilder.GetClientConnectionService(decodeConnRequest(req.Args))

	page, err := connService.GetPage()
	if err != nil {
		return nil, err
	}

	return encodeClientConnResponse(page), nil
}
