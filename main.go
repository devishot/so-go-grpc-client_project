package main

import (
	"log"

	"github.com/kelseyhightower/envconfig"

	"github.com/devishot/so-go-grpc-client_project/infrastructure/database/postgres"
	"github.com/devishot/so-go-grpc-client_project/infrastructure/tcp_server"
)

const DBEnvPrefix string = "DATABASE"
const GRPCEnvPrefix string = "GRPC"

func main() {
	b := NewBuilder()
	b.InitDB(loadDatabaseEnv())
	b.InitRepositories()
	b.InitServices()
	b.InitHandlers()
	b.InitGRPC(loadGRPCEnv())

	b.GRPCServer.Listen()
}

func loadDatabaseEnv() (cfg postgres.Config) {
	err := envconfig.Process(DBEnvPrefix, &cfg)
	if err != nil || cfg.IsZero() {
		log.Fatalf("cannot load env configs for %s", DBEnvPrefix)
	}
	return
}

func loadGRPCEnv() (cfg tcp_server.Config) {
	err := envconfig.Process(GRPCEnvPrefix, &cfg)
	if err != nil || cfg.IsZero() {
		log.Fatalf("cannot load env configs for %s", DBEnvPrefix)
	}

	return
}
