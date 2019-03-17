package services

import (
	"github.com/devishot/so-go-grpc-client_project/domain"
)

type ConnectionService interface {
	GetConnection() (val domain.ConnectionValue, err error)
}
