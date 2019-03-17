package repository

import (
	"time"

	"github.com/devishot/so-go-grpc-client_project/domain"
)

//go:generate moq -out client_connection_repository__gen_mock.go . ClientConnectionRepository
type ClientConnectionRepository interface {
	GetLast() (domain.ClientEntity, error)
	GetFirst() (domain.ClientEntity, error)
	PaginateForwardByTimestamp(first int, after time.Time) ([]domain.ClientEntity, error)
	PaginateBackwardByTimestamp(last int, before time.Time) ([]domain.ClientEntity, error)
}
