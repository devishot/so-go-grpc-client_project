package grpc

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/devishot/so-go-grpc-client_project/interfaces/grpc/api"
)

type Server struct {
	Listener             net.Listener
	ClientHandler        pb.ClientServiceServer
	ClientProjectHandler pb.ClientProjectServiceServer

	gRPCServer *grpc.Server
}

func (s *Server) Init() {
	s.gRPCServer = grpc.NewServer()

	pb.RegisterClientServiceServer(s.gRPCServer, s.ClientHandler)
	pb.RegisterClientProjectServiceServer(s.gRPCServer, s.ClientProjectHandler)
}

func (s *Server) Listen() {
	if err := s.gRPCServer.Serve(s.Listener); err != nil {
		log.Fatalf("grpc Serve: error=%v", err)
		return
	}
}

func (s *Server) Shutdown(ctx context.Context) {
	<-ctx.Done()
	s.gRPCServer.Stop()
}
