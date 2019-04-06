package domain

import (
	"errors"
)

var IncorrectConnectionArgsError = errors.New("none of field pairs 'first-after' or 'last-before' exist")
var IncorrectPageInfoError = errors.New("none of field pairs 'hasNext-EndCursor' or 'hasPrevious-StartCursor' exist")
