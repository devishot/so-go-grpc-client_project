package connection

import (
	"github.com/devishot/so-go-grpc-client_project/domain"
	"github.com/devishot/so-go-grpc-client_project/interface/graphql"
)

type ProjectConnectionRepository interface {
	GetLastPageItemByTimestamp(clientID domain.ID, args graphql.ProjectRepositoryPageArgs) (domain.ProjectEntity, error)
	GetFirstPageItemByTimestamp(clientID domain.ID, args graphql.ProjectRepositoryPageArgs) (domain.ProjectEntity, error)
	PaginateByTimestamp(clientID domain.ID, args graphql.ProjectRepositoryPageArgs) ([]domain.ProjectEntity, error)
}
