package main

import (
	"log"
	"net"

	"github.com/devishot/so-go-grpc-client_project/domain"

	"github.com/devishot/so-go-grpc-client_project/interfaces/db/repository_impl"
	conn_services "github.com/devishot/so-go-grpc-client_project/interfaces/graphql_connection/services"
	"github.com/devishot/so-go-grpc-client_project/interfaces/grpc"
	"github.com/devishot/so-go-grpc-client_project/interfaces/grpc/handler"

	"github.com/devishot/so-go-grpc-client_project/infrastructure/database/postgres"
	"github.com/devishot/so-go-grpc-client_project/infrastructure/tcp_server"
)

func NewBuilder() *Builder {
	return &Builder{}
}

type Builder struct {
	TCPServer  net.Listener
	GRPCServer *grpc.Server
	DB         *postgres.DB

	ClientRepo  repository_impl.ClientRepository
	ProjectRepo repository_impl.ProjectRepository

	ConnectionBuilder conn_services.Builder

	ClientService  domain.ClientService
	ProjectService domain.ProjectService

	ClientHandler  handler.ClientService
	ProjectHandler handler.ClientProjectService
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

func (b *Builder) InitRepositories() {
	var err error

	b.ClientRepo, err = repository_impl.NewClientRepository(b.DB)
	if err != nil {
		log.Fatal(err)
	}

	b.ProjectRepo, err = repository_impl.NewProjectRepository(b.DB)
	if err != nil {
		log.Fatal(err)
	}

	b.ConnectionBuilder = conn_services.Builder{
		ClientConnRepo:  b.ClientRepo,
		ProjectConnRepo: b.ProjectRepo,
	}
}

func (b *Builder) InitServices() {
	b.ClientService = domain.ClientService{Repo: b.ClientRepo}
	b.ProjectService = domain.ProjectService{Repo: b.ProjectRepo}
}

func (b *Builder) InitHandlers() {
	b.ClientHandler = handler.ClientService{
		Service:     b.ClientService,
		ConnBuilder: b.ConnectionBuilder,
	}
	b.ProjectHandler = handler.ClientProjectService{
		Service:     b.ProjectService,
		ConnBuilder: b.ConnectionBuilder,
	}
}

func (b *Builder) InitGRPC(cfg tcp_server.Config) {
	tcp := &tcp_server.TCPServer{Cfg: cfg}
	b.TCPServer = tcp.Listen()

	grpcServer := &grpc.Server{
		Listener:             b.TCPServer,
		ClientHandler:        b.ClientHandler,
		ClientProjectHandler: b.ProjectHandler,
	}
	grpcServer.Init()

	b.GRPCServer = grpcServer
}
