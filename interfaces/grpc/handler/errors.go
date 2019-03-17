package handler

import (
	"errors"
)

var NotImplementedError = errors.New("gRPC handler method not implented")
