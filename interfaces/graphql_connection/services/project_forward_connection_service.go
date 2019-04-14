package services

import (
	"time"

	"github.com/devishot/so-go-grpc-client_project/domain"
	conn "github.com/devishot/so-go-grpc-client_project/interfaces/graphql_connection/domain"
	"github.com/devishot/so-go-grpc-client_project/interfaces/graphql_connection/repository"
)

type ProjectForwardConnectionService struct {
	ConnRepo repository.ProjectConnectionRepository
	Page     conn.ArgumentsValue
	Client   domain.ClientEntity

	projects []domain.ProjectEntity
}

func (s *ProjectForwardConnectionService) GetPage() (val conn.PageValue, err error) {
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

func (s *ProjectForwardConnectionService) fetchPage() (err error) {
	timestamp := time.Time{}
	if s.Page.After != "" {
		timestamp = conn.DecodeTimestampCursor(s.Page.After)
	}

	s.projects, err = s.ConnRepo.PaginateForwardByClientByTimestamp(s.Client.ID, s.Page.First, timestamp)
	if err != nil {
		return
	}

	return
}

func (s *ProjectForwardConnectionService) getEdges() []conn.EdgeValue {
	edges := make([]conn.EdgeValue, 0)

	for _, p := range s.projects {
		e := conn.EdgeValue{
			Cursor: conn.EncodeTimestampCursor(p.Timestamp),
			Node:   p,
		}
		edges = append(edges, e)
	}

	return edges
}

func (s *ProjectForwardConnectionService) getPageInfo() (pageInfo conn.PageInfoValue, err error) {
	last, err := s.ConnRepo.GetLastByClient(s.Client.ID)
	if err != nil {
		return
	}

	var hasNext bool
	if n := len(s.projects); n > 0 {
		lastInPage := s.projects[n-1]
		hasNext = lastInPage.ID != last.ID
	}

	pageInfo = conn.PageInfoValue{
		EndCursor:   conn.EncodeTimestampCursor(last.Timestamp),
		HasNextPage: hasNext,
	}
	return
}
