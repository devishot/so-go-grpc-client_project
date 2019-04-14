package domain

import (
	"errors"
)

var IncorrectConnectionArgsError = errors.New("none of fields 'first' or 'last' exist")
