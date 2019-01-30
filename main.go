package main

import (
	"log"

	"github.com/kelseyhightower/envconfig"

	"github.com/devishot/so-go-grpc-client_project/infrastructure/database/postgres"
	"github.com/devishot/so-go-grpc-client_project/infrastructure/tcp_server"
	"github.com/devishot/so-go-grpc-client_project/usecase"
)

const DBEnvPrefix string = "DATABASE"
const GRPCEnvPrefix string = "GRPC"

func main() {
	b := &usecase.Builder{}
	b.InitGRPC(loadGRPCEnv())

	b.GRPCServer.Listen()
}

func loadDatabaseEnv() (cfg postgres.Config) {
	err := envconfig.Process(DBEnvPrefix, &cfg)
	if err != nil {
		log.Fatalf("cannot load env configs for %t", DBEnvPrefix)
	}
	return
}

func loadGRPCEnv() (cfg tcp_server.Config) {
	err := envconfig.Process(GRPCEnvPrefix, &cfg)
	if err != nil {
		log.Fatalf("cannot load env configs for %t", DBEnvPrefix)
	}
	return
}
