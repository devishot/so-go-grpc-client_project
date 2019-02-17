package domain

import (
	"github.com/satori/go.uuid"
)

func generateID() ID {
	id := uuid.Must(uuid.NewV4())
	idStr := id.String()
	return ID(idStr)
}
