package grpc

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/devishot/so-go-grpc-client_project/interfaces/grpc/api"
	"github.com/devishot/so-go-grpc-client_project/interfaces/grpc/handler"
)

type Server struct {
	Listener   net.Listener
	gRPCServer *grpc.Server
}

func (s *Server) Init() {
	s.gRPCServer = grpc.NewServer()

	pb.RegisterClientServiceServer(s.gRPCServer, &handler.ClientService{})
}

func (s *Server) Listen() {
	if err := s.gRPCServer.Serve(s.Listener); err != nil {
		log.Fatalf("grpc Serve: error=%v", err)
		return
	}
}
