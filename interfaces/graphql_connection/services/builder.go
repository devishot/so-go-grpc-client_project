package services

import (
	conn "github.com/devishot/so-go-grpc-client_project/interfaces/graphql_connection/domain"
	"github.com/devishot/so-go-grpc-client_project/interfaces/graphql_connection/repository"
)

type Builder struct {
	ClientConnRepo  repository.ClientConnectionRepository
	ProjectConnRepo repository.ProjectConnectionRepository
}

func (b Builder) GetClientConnectionService(page conn.ArgumentsValue) ConnectionService {
	ok, _ := page.IsForward()

	if ok {
		return &ClientForwardConnectionService{
			ConnRepo: b.ClientConnRepo,
			Page:     page,
		}
	} else {
		// TODO: create ClientBackwardConnectionService
		return nil
	}
}

func (b Builder) GetProjectConnectionService(page conn.ArgumentsValue) ConnectionService {
	ok, _ := page.IsForward()

	if ok {
		return &ProjectForwardConnectionService{
			ConnRepo: b.ProjectConnRepo,
			Page:     page,
		}
	} else {
		// TODO: create ProjectBackwardConnectionService
		return nil
	}
}
