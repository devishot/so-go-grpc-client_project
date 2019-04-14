package services

import (
	"time"

	"github.com/devishot/so-go-grpc-client_project/domain"
	conn "github.com/devishot/so-go-grpc-client_project/interfaces/graphql_connection/domain"
	"github.com/devishot/so-go-grpc-client_project/interfaces/graphql_connection/repository"
)

type ClientForwardConnectionService struct {
	ConnRepo repository.ClientConnectionRepository
	Page     conn.ArgumentsValue

	clients []domain.ClientEntity
}

func (s *ClientForwardConnectionService) GetPage() (val conn.PageValue, err error) {
	err = s.fetchPage()
	if err != nil {
		return
	}

	pageInfo, err := s.getPageInfo()
	if err != nil {
		return
	}

	val = conn.PageValue{
		Edges:    s.getEdges(),
		PageInfo: pageInfo,
	}
	return
}

func (s *ClientForwardConnectionService) fetchPage() (err error) {
	timestamp := time.Time{}
	if s.Page.After != "" {
		timestamp = conn.DecodeTimestampCursor(s.Page.After)
	}

	s.clients, err = s.ConnRepo.PaginateForwardByTimestamp(s.Page.First, timestamp)
	if err != nil {
		return
	}

	return
}

func (s *ClientForwardConnectionService) getEdges() []conn.EdgeValue {
	edges := make([]conn.EdgeValue, 0)

	for _, p := range s.clients {
		e := conn.EdgeValue{
			Cursor: conn.EncodeTimestampCursor(p.Timestamp),
			Node:   p,
		}
		edges = append(edges, e)
	}

	return edges
}

func (s *ClientForwardConnectionService) getPageInfo() (pageInfo conn.PageInfoValue, err error) {
	last, err := s.ConnRepo.GetLast()
	if err != nil {
		return
	}

	var hasNext bool
	if n := len(s.clients); n > 0 {
		lastInPage := s.clients[n-1]
		hasNext = lastInPage.ID != last.ID
	}

	pageInfo = conn.PageInfoValue{
		EndCursor:   conn.EncodeTimestampCursor(last.Timestamp),
		HasNextPage: hasNext,
	}
	return
}
