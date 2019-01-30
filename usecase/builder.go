package usecase

import (
	"log"

	"github.com/devishot/so-go-grpc-client_project/domain_interface/grpc"
	"github.com/devishot/so-go-grpc-client_project/infrastructure/database/postgres"
	"github.com/devishot/so-go-grpc-client_project/infrastructure/tcp_server"
)

type Builder struct {
	GRPCServer *grpc.Server
	DB         *postgres.DB
}

func (b *Builder) InitGRPC(cfg tcp_server.Config) {
	tcpServer := &tcp_server.TCPServer{Cfg: cfg}
	grpcServer := &grpc.Server{Listener: tcpServer.Listen()}
	grpcServer.Init()

	b.GRPCServer = grpcServer
}

func (b *Builder) InitDB(cfg postgres.Config) {
	db, err := postgres.New(cfg)
	if err != nil {
		log.Fatalf("cannot connect to database, config=%v err=%v", cfg, err)
	}

	err = db.Conn.Ping()
	if err != nil {
		log.Fatal("cannot ping database connection")
	}

	b.DB = db
}
