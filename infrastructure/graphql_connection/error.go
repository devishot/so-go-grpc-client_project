package graphql_connection

import (
	"errors"
)

var IncorrectConnectionArgsError = errors.New("none of args pairs 'first-after' or 'last-before' exist")
