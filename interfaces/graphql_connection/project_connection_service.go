package graphql_connection

import (
	"github.com/devishot/so-go-grpc-client_project/domain"
)

type ProjectConnectionService struct {
	ConnRepo ProjectConnectionRepository
}

func (s *ProjectConnectionService) Connection(cID domain.ID, args domain.ConnectionArgumentsValue) (
	*domain.ConnectionValue, error) {
	page, err := s.fetchPage(cID, args)
	if err != nil {
		return nil, err
	}

	forward, err := args.IsForward()
	if err != nil {
		return nil, err
	}

	edges := s.getEdges(page.Projects)
	pageInfo := s.getPageInfo(forward, page, edges)

	val := &domain.ConnectionValue{
		Edges:    edges,
		PageInfo: pageInfo,
	}
	return val, nil
}

func (s *ProjectConnectionService) fetchPage(cID domain.ID, args domain.ConnectionArgumentsValue) (
	val *domain.ConnectionArgumentsValue, err error) {
	var cursorProject domain.ProjectEntity
	var pageArgs ProjectRepositoryPageArgs

	forward, err := args.IsForward()
	if err != nil {
		return
	}

	if forward {
		pageArgs = domain.ConnectionArgumentsValue{
			First: args.First,
			After: domain.DecodeTimestampCursor(args.After),
		}
	} else {
		pageArgs = domain.ConnectionArgumentsValue{
			Last:   args.Last,
			Before: domain.DecodeTimestampCursor(args.Before),
		}
	}

	if forward {
		cursorProject, err = s.ConnRepo.GetLastByClient(cID)
	} else {
		cursorProject, err = s.ConnRepo.GetFirstByClient(cID)
	}
	if err != nil {
		return
	}

	cursor := domain.EncodeTimestampCursor(cursorProject.Timestamp)

	projects, err := s.ConnRepo.PaginateByTimestamp(cID, pageArgs)
	if err != nil {
		return
	}

	if forward {
		val = &domain.ConnectionArgumentsValue{
			Projects:  projects,
			EndCursor: cursor,
		}
	} else {
		val = &domain.ConnectionArgumentsValue{
			Projects:    projects,
			StartCursor: cursor,
		}
	}

	return
}

func (s *ProjectConnectionService) getEdges(projects []domain.ProjectEntity) []*domain.ConnectionEdgeValue {
	edges := make([]*domain.ConnectionEdgeValue, len(projects))

	for i, p := range projects {
		edge := &domain.ConnectionEdgeValue{
			Cursor: domain.EncodeTimestampCursor(p.Timestamp),
			Node:   p,
		}
		edges[i] = edge
	}

	return edges
}

func (s *ProjectConnectionService) getPageInfo(forward bool, page *domain.ConnectionArgumentsValue, edges []*domain.ConnectionEdgeValue) domain.ConnectionPageInfoValue {
	if forward {
		return domain.ConnectionPageInfoValue{
			HasNextPage: edges[len(edges)-1].Cursor != page.EndCursor,
			EndCursor:   page.EndCursor,
		}
	} else {
		return domain.ConnectionPageInfoValue{
			HasPreviousPage: edges[0].Cursor != page.StartCursor,
			StartCursor:     page.StartCursor,
		}
	}
}
