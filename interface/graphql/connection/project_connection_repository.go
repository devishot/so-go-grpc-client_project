package connection

import (
	"github.com/devishot/so-go-grpc-client_project/domain"
	"github.com/devishot/so-go-grpc-client_project/interface/graphql"
)

//go:generate moq -out project_connection_repository__gen_mock.go . ProjectConnectionRepository
type ProjectConnectionRepository interface {
	GetLastByClient(clientID domain.ID) (domain.ProjectEntity, error)
	GetFirstByClient(clientID domain.ID) (domain.ProjectEntity, error)
	PaginateByTimestamp(clientID domain.ID, args graphql.ProjectRepositoryPageArgs) ([]domain.ProjectEntity, error)
}
