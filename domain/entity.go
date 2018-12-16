package domain

import (
	"time"

	"github.com/satori/go.uuid"
)

type ID string

func generateID() ID {
	id := uuid.Must(uuid.NewV4())
	idStr := id.String()
	return ID(idStr)
}

func NewClient(fName, lName, cName string) ClientEntity {
	return ClientEntity{
		ID:          generateID(),
		Timestamp:   time.Now(),
		FirstName:   fName,
		LastName:    lName,
		CompanyName: cName,
	}
}

type ClientEntity struct {
	ID          ID
	Timestamp   time.Time `db:"created_at"`
	FirstName   string    `db:"first_name"`
	LastName    string    `db:"last_name"`
	CompanyName string    `db:"company_name"`
}

func (e ClientEntity) IsZero() bool {
	return (e.ID == "" || e.Timestamp.IsZero())
}

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
	ID          ID
	Timestamp   time.Time
	Title       string
	Description string
	ClientID    ID
}
