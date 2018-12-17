package _examples

import (
	"log"

	"github.com/devishot/so-go-grpc-client_project/infrastructure/database/postgres"
)

var currentDB *postgres.DB

var dbConfig = postgres.Config{
	Host:     "localhost",
	Port:     5432,
	User:     "devishot",
	Database: "so_client_project",
}

func getDatabase() *postgres.DB {
	if currentDB != nil {
		return currentDB
	}

	db, err := postgres.New(dbConfig)
	if err != nil {
		log.Fatalf("cannot connect to database, config=%v error=%v", dbConfig, err)
	}

	currentDB = db
	return currentDB
}
