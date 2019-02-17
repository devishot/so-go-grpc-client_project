package domain

import (
	"time"
)

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
	ID          ID        `db:"id"`
	Timestamp   time.Time `db:"created_at"`
	FirstName   string    `db:"first_name"`
	LastName    string    `db:"last_name"`
	CompanyName string    `db:"company_name"`
}

func (e ClientEntity) IsZero() bool {
	return (e.ID == "" || e.Timestamp.IsZero())
}
