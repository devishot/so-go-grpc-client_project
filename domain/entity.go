package domain

import (
	"encoding/base64"
	"time"

	"github.com/satori/go.uuid"
)

type ID string
type ConnCursor string

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
	Timestamp   time.Time
	FirstName   string
	LastName    string
	CompanyName string
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

func (p ProjectEntity) getTimestampCursor() ConnCursor {
	ts := p.Timestamp.Unix()
	tsStr := string(ts)
	cursor := base64.StdEncoding.EncodeToString([]byte(tsStr))
	return ConnCursor(cursor)
}
