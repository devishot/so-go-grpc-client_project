package main

import (
	"log"

	"github.com/kelseyhightower/envconfig"

	"github.com/devishot/so-go-grpc-client_project/infrastructure/database/postgres"
)

const DBEnvPrefix string = "DATABASE"

var DBConfig postgres.Config

func init() {
	loadDatabaseEnv()
}

func main() {
	db, err := postgres.New(DBConfig)
	if err != nil {
		log.Fatalf("cannot connect to database, config=%v err=%v", DBConfig, err)
	}

	err = db.Conn.Ping()
	if err != nil {
		log.Fatal("cannot ping database connection")
	}

	log.Println("Hello world")
}

func loadDatabaseEnv() {
	err := envconfig.Process(DBEnvPrefix, &DBConfig)
	if err != nil {
		log.Fatalf("cannot load env configs for %s", DBEnvPrefix)
	}
}
