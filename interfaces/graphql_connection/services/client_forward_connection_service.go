package services

import (
	"github.com/devishot/so-go-grpc-client_project/domain"
	"github.com/devishot/so-go-grpc-client_project/interfaces/graphql_connection/repository"
)

type ClientForwardConnectionService struct {
	ConnRepo repository.ClientConnectionRepository
	Page     domain.ConnectionArgumentsValue

	clients []domain.ClientEntity
}

func (s *ClientForwardConnectionService) GetConnection() (val domain.ConnectionValue, err error) {
	err = s.fetchPage()
	if err != nil {
		return
	}

	pageInfo, err := s.getPageInfo()
	if err != nil {
		return
	}

	val = domain.ConnectionValue{
		Edges:    s.getEdges(),
		PageInfo: pageInfo,
	}
	return
}

func (s *ClientForwardConnectionService) fetchPage() (err error) {
	s.clients, err = s.ConnRepo.PaginateForwardByTimestamp(s.Page.First, domain.DecodeTimestampCursor(s.Page.After))
	if err != nil {
		return
	}

	return
}

func (s *ClientForwardConnectionService) getEdges() []domain.ConnectionEdgeValue {
	edges := make([]domain.ConnectionEdgeValue, 0)

	for _, p := range s.clients {
		e := domain.ConnectionEdgeValue{
			Cursor: domain.EncodeTimestampCursor(p.Timestamp),
			Node:   p,
		}
		edges = append(edges, e)
	}

	return edges
}

func (s *ClientForwardConnectionService) getPageInfo() (pageInfo domain.ConnectionPageInfoValue, err error) {
	last, err := s.ConnRepo.GetLast()
	if err != nil {
		return
	}

	var hasNext bool
	if n := len(s.clients); n > 0 {
		lastInPage := s.clients[n-1]
		hasNext = lastInPage.ID != last.ID
	}

	pageInfo = domain.ConnectionPageInfoValue{
		EndCursor:   domain.EncodeTimestampCursor(last.Timestamp),
		HasNextPage: hasNext,
	}
	return
}
