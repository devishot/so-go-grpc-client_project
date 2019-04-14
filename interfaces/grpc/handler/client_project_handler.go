package handler

import (
	"context"
	"log"

	"github.com/devishot/so-go-grpc-client_project/domain"
	connServices "github.com/devishot/so-go-grpc-client_project/interfaces/graphql_connection/services"
	pb "github.com/devishot/so-go-grpc-client_project/interfaces/grpc/api"
)

type ClientProjectService struct {
	Service     domain.ProjectService
	ConnBuilder connServices.Builder
}

func (s ClientProjectService) CreateClientProject(ctx context.Context, req *pb.CreateClientProjectRequest) (*pb.ClientProject, error) {
	log.Printf("createClient: req=%v", req)

	p, err := s.Service.Create(decodeProtoClientProject(req.Data))
	if err != nil {
		return nil, err
	}

	return encodeProtoClientProject(p), nil
}

func (s ClientProjectService) DeleteClientProject(ctx context.Context, req *pb.DeleteClientProjectRequest) (*pb.ClientProject, error) {
	log.Printf("deleteClient: request=%v", req)

	cl, err := s.Service.Delete(decodeID(req.Id))
	if err != nil {
		return nil, err
	}

	return encodeProtoClientProject(cl), nil
}

func (s ClientProjectService) GetClientProject(ctx context.Context, req *pb.GetClientProjectRequest) (*pb.ClientProject, error) {
	log.Printf("getClient: request=%v", req)

	cl, err := s.Service.Get(decodeID(req.Id))
	if err != nil {
		return nil, err
	}

	return encodeProtoClientProject(cl), nil
}

func (s ClientProjectService) GetClientProjectConnection(ctx context.Context, req *pb.GetClientProjectConnectionRequest) (*pb.ClientProjectConnectionResponse, error) {
	log.Printf("getClientConnection: request=%v", req)

	connReq := decodeConnRequest(req.Args)
	if _, err := connReq.IsForward(); err != nil {
		return nil, err
	}

	connService := s.ConnBuilder.GetProjectConnectionService(connReq)
	page, err := connService.GetPage()
	if err != nil {
		return nil, err
	}

	return encodeClientProjectConnResponse(page), nil
}
