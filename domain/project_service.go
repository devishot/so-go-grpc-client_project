package domain

import (
	"errors"
)

var IncorrectConnectionArgsError = errors.New("args should not be empty for 'first-after' or 'last-before'")

type ProjectService struct {
	Repo ProjectRepository
}

func (s *ProjectService) Create(in ProjectInputValue) error {
	p := NewProject(in.ClientID, in.Title, in.Description)

	return s.Repo.Create(p)
}

func (s *ProjectService) Delete(id ID) error {
	return s.Repo.Delete(id)
}

func (s *ProjectService) Connection(cID ID, args ConnectionArgumentsValue) (*ConnectionValue, error) {
	page, err := s.Repo.Connection(cID, args)
	if err != nil {
		return nil, err
	}

	forward, err := s.isForward(args)
	if err != nil {
		return nil, err
	}

	edges := s.getConnectionEdges(page.Projects)
	pageInfo := s.getConnectionPageInfo(forward, page, edges)

	conn := &ConnectionValue{
		Edges:    edges,
		PageInfo: pageInfo,
	}
	return conn, err
}

func (s *ProjectService) getConnectionEdges(projects []ProjectEntity) []*ConnectionEdgeValue {
	edges := make([]*ConnectionEdgeValue, len(projects))

	for _, p := range projects {
		edge := &ConnectionEdgeValue{
			Cursor: p.getTimestampCursor(),
			Node:   p,
		}
		edges = append(edges, edge)
	}

	return edges
}

func (s *ProjectService) getConnectionPageInfo(
	forward bool,
	page *ProjectConnectionPageValue,
	edges []*ConnectionEdgeValue) ConnectionPageInfoValue {
	if forward {
		return ConnectionPageInfoValue{
			HasNextPage: edges[len(edges)-1].Cursor == page.EndCursor,
			EndCursor:   page.EndCursor,
		}
	} else {
		return ConnectionPageInfoValue{
			HasPreviousPage: edges[0].Cursor == page.StartCursor,
			StartCursor:     page.StartCursor,
		}
	}
}

func (s *ProjectService) isForward(args ConnectionArgumentsValue) (forward bool, err error) {
	if args.First != 0 && args.After != "" {
		forward = true
	} else if args.Last != 0 && args.Before != "" {
		forward = false
	} else {
		err = IncorrectConnectionArgsError
	}
	return forward, err
}
