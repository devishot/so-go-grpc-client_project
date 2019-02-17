package domain

import (
	"time"
)

func NewProject(clientID ID, title, description string) ProjectEntity {
	return ProjectEntity{
		ID:          generateID(),
		Timestamp:   time.Now(),
		Title:       title,
		Description: description,
		ClientID:    clientID,
	}
}

type ProjectEntity struct {
	ID          ID        `db:"id"`
	Timestamp   time.Time `db:"created_at"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	ClientID    ID        `db:"client_id"`
}

func (e ProjectEntity) IsZero() bool {
	return (e.ID == "" || e.Timestamp.IsZero())
}
