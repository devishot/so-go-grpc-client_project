package graphql

import (
	"time"

	"github.com/devishot/so-go-grpc-client_project/domain"
	conn "github.com/devishot/so-go-grpc-client_project/infrastructure/graphql_connection"
)

type ProjectInputValue struct {
	Title       string
	Description string
	ClientID    domain.ID
}

type ProjectConnectionPageValue struct {
	Projects    []domain.ProjectEntity
	EndCursor   conn.Cursor
	StartCursor conn.Cursor
}

type ProjectRepositoryPageArgs struct {
	First int
	After time.Time

	Last   int
	Before time.Time
}

func IsForward(args ProjectRepositoryPageArgs) bool {
	if args.First != 0 && !args.After.IsZero() {
		return true
	}
	return false
}
