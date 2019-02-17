package infrastructure

import (
	"log"
	"net"

	"github.com/devishot/so-go-grpc-client_project/infrastructure/database/postgres"
	"github.com/devishot/so-go-grpc-client_project/infrastructure/tcp_server"
	"github.com/devishot/so-go-grpc-client_project/interfaces/grpc"
)

type Builder struct {
	GRPCServer *grpc.Server
	TCPServer  net.Listener
	DB         *postgres.DB
}

func (b *Builder) InitGRPC(cfg tcp_server.Config) {
	tcp := &tcp_server.TCPServer{Cfg: cfg}
	b.TCPServer = tcp.Listen()

	grpcServer := &grpc.Server{Listener: b.TCPServer}
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
