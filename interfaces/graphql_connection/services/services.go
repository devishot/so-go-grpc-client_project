package services

import (
	conn "github.com/devishot/so-go-grpc-client_project/interfaces/graphql_connection/domain"
)

type ConnectionService interface {
	GetPage() (val conn.PageValue, err error)
}
