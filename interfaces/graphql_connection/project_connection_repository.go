package graphql_connection

import (
	"github.com/devishot/so-go-grpc-client_project/domain"
)

//go:generate moq -out project_connection_repository__gen_mock.go . ProjectConnectionRepository
type ProjectConnectionRepository interface {
	GetLastByClient(clientID domain.ID) (domain.ProjectEntity, error)
	GetFirstByClient(clientID domain.ID) (domain.ProjectEntity, error)
	PaginateByTimestamp(clientID domain.ID, args domain.ConnectionArgumentsValue) ([]domain.ProjectEntity, error)
}
