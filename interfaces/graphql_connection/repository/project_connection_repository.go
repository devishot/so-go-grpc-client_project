package repository

import (
	"time"

	"github.com/devishot/so-go-grpc-client_project/domain"
)

//go:generate moq -out project_connection_repository__gen_mock.go . ProjectConnectionRepository
type ProjectConnectionRepository interface {
	GetLastByClient(clientID domain.ID) (domain.ProjectEntity, error)
	GetFirstByClient(clientID domain.ID) (domain.ProjectEntity, error)
	PaginateForwardByClientByTimestamp(clientID domain.ID, first int, after time.Time) ([]domain.ProjectEntity, error)
	PaginateBackwardByClientByTimestamp(clientID domain.ID, last int, before time.Time) ([]domain.ProjectEntity, error)
}
