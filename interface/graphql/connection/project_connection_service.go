package connection

import (
	"strconv"
	"time"

	"github.com/devishot/so-go-grpc-client_project/domain"
	conn "github.com/devishot/so-go-grpc-client_project/infrastructure/graphql_connection"
	"github.com/devishot/so-go-grpc-client_project/interface/graphql"
)

type ProjectConnectionService struct {
	ConnRepo ProjectConnectionRepository
}

func (s *ProjectConnectionService) encodeTimestampCursor(t time.Time) conn.Cursor {
	ts := t.Unix()
	str := string(ts)
	return conn.NewCursor(str)
}

func (s *ProjectConnectionService) decodeTimestampCursor(c conn.Cursor) time.Time {
	str := conn.Must(conn.FromCursor(c)).(string)

	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		panic(err)
	}

	t := time.Unix(i, 0)
	return t
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
			After: s.decodeTimestampCursor(args.After),
		}
	} else {
		pageArgs = graphql.ProjectRepositoryPageArgs{
			Last:   args.Last,
			Before: s.decodeTimestampCursor(args.Before),
		}
	}

	if forward {
		cursorProject, err = s.ConnRepo.GetLastByTimestamp(cID, pageArgs)
	} else {
		cursorProject, err = s.ConnRepo.GetFirstByTimestamp(cID, pageArgs)
	}
	if err != nil {
		return
	}

	cursor := s.encodeTimestampCursor(cursorProject.Timestamp)

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

	for _, p := range projects {
		edge := &conn.ConnectionEdgeValue{
			Cursor: s.encodeTimestampCursor(p.Timestamp),
			Node:   p,
		}
		edges = append(edges, edge)
	}

	return edges
}

func (s *ProjectConnectionService) getPageInfo(
	forward bool,
	page *graphql.ProjectConnectionPageValue,
	edges []*conn.ConnectionEdgeValue) conn.ConnectionPageInfoValue {
	if forward {
		return conn.ConnectionPageInfoValue{
			HasNextPage: edges[len(edges)-1].Cursor == page.EndCursor,
			EndCursor:   page.EndCursor,
		}
	} else {
		return conn.ConnectionPageInfoValue{
			HasPreviousPage: edges[0].Cursor == page.StartCursor,
			StartCursor:     page.StartCursor,
		}
	}
}
