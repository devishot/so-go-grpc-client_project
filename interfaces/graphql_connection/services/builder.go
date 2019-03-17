package services

import (
	"github.com/devishot/so-go-grpc-client_project/domain"
	"github.com/devishot/so-go-grpc-client_project/interfaces/graphql_connection/repository"
)

type Builder struct {
	ClientConnRepo repository.ClientConnectionRepository
}

func (b Builder) GetClientConnectionService(page domain.ConnectionArgumentsValue) ConnectionService {
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
