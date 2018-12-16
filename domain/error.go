package domain

import (
	"github.com/pkg/errors"
)

var NotFoundProjectRepositoryError = errors.New("Not Found: project")
var NotFoundClientRepositoryError = errors.New("Not Found: client")
