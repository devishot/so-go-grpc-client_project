package domain

import (
	"github.com/devishot/so-go-grpc-client_project/domain"
)

type ArgumentsValue struct {
	// forward
	First int
	After domain.Cursor
	// backward
	Last   int
	Before domain.Cursor
}

func (args ArgumentsValue) IsForward() (forward bool, err error) {
	if args.First != 0 && args.After != "" {
		forward = true
	} else if args.Last != 0 && args.Before != "" {
		forward = false
	} else {
		err = IncorrectConnectionArgsError
	}
	return forward, err
}

type PageInfoValue struct {
	// forward
	HasNextPage bool
	EndCursor   domain.Cursor
	// backward
	HasPreviousPage bool
	StartCursor     domain.Cursor
}

type EdgeValue struct {
	Cursor domain.Cursor
	Node   interface{}
}

type PageValue struct {
	Edges    []EdgeValue
	PageInfo PageInfoValue
}
