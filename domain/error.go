package domain

import (
	"github.com/pkg/errors"
)

var NotFoundProjectRepositoryError = errors.New("Not Found: project")
var NotFoundClientRepositoryError = errors.New("Not Found: client")

var IncorrectConnectionArgsError = errors.New("none of field pairs 'first-after' or 'last-before' exist")
var IncorrectPageInfoError = errors.New("none of field pairs 'hasNext-EndCursor' or 'hasPrevious-StartCursor' exist")
