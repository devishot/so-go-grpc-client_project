package connection

import (
	"github.com/devishot/so-go-grpc-client_project/domain"
	"github.com/devishot/so-go-grpc-client_project/domain_interface/graphql"
	conn "github.com/devishot/so-go-grpc-client_project/infrastructure/graphql_connection"
)

type ProjectConnectionService struct {
	ConnRepo ProjectConnectionRepository
}

func (s *ProjectConnectionService) Connection(cID domain.ID, args conn.ConnectionArgumentsValue) (
	*conn.ConnectionValue, error) {
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

	val := &conn.ConnectionValue{
		Edges:    edges,
		PageInfo: pageInfo,
	}
	return val, nil
}

func (s *ProjectConnectionService) fetchPage(cID domain.ID, args conn.ConnectionArgumentsValue) (
	val *graphql.ProjectConnectionPageValue, err error) {
	var cursorProject domain.ProjectEntity
	var pageArgs graphql.ProjectRepositoryPageArgs

	forward, err := args.IsForward()
	if err != nil {
		return
	}

	if forward {
		pageArgs = graphql.ProjectRepositoryPageArgs{
			First: args.First,
			After: decodeTimestampCursor(args.After),
		}
	} else {
		pageArgs = graphql.ProjectRepositoryPageArgs{
			Last:   args.Last,
			Before: decodeTimestampCursor(args.Before),
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

	cursor := encodeTimestampCursor(cursorProject.Timestamp)

	projects, err := s.ConnRepo.PaginateByTimestamp(cID, pageArgs)
	if err != nil {
		return
	}

	if forward {
		val = &graphql.ProjectConnectionPageValue{
			Projects:  projects,
			EndCursor: cursor,
		}
	} else {
		val = &graphql.ProjectConnectionPageValue{
			Projects:    projects,
			StartCursor: cursor,
		}
	}

	return
}

func (s *ProjectConnectionService) getEdges(projects []domain.ProjectEntity) []*conn.ConnectionEdgeValue {
	edges := make([]*conn.ConnectionEdgeValue, len(projects))

	for i, p := range projects {
		edge := &conn.ConnectionEdgeValue{
			Cursor: encodeTimestampCursor(p.Timestamp),
			Node:   p,
		}
		edges[i] = edge
	}

	return edges
}

func (s *ProjectConnectionService) getPageInfo(forward bool, page *graphql.ProjectConnectionPageValue, edges []*conn.ConnectionEdgeValue) conn.ConnectionPageInfoValue {
	if forward {
		return conn.ConnectionPageInfoValue{
			HasNextPage: edges[len(edges)-1].Cursor != page.EndCursor,
			EndCursor:   page.EndCursor,
		}
	} else {
		return conn.ConnectionPageInfoValue{
			HasPreviousPage: edges[0].Cursor != page.StartCursor,
			StartCursor:     page.StartCursor,
		}
	}
}
