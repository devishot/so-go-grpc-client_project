package _examples

import (
	"log"

	"github.com/devishot/so-go-grpc-client_project/domain"
	"github.com/devishot/so-go-grpc-client_project/infrastructure/database/postgres"
	"github.com/devishot/so-go-grpc-client_project/interfaces/db/repository"
)

var clientTestRows *ClientTestRows

type ClientTestRows struct {
	db   *postgres.DB
	repo domain.ClientRepository
	item domain.ClientEntity
}

func (r *ClientTestRows) Get() domain.ClientEntity {
	if !r.item.IsZero() {
		return r.item
	}

	log.Println("inserting ClientEntity row for testing")

	cl := domain.NewClient("testFirstName", "testLastName", "testCompanyName")

	err := r.repo.Create(cl)
	if err != nil {
		log.Fatalf("cannot insert row for testing, error=%v entity=%v", err, cl)
	}

	r.item = cl
	return r.item
}

func (r *ClientTestRows) Release() {
	if err := r.repo.Delete(r.item.ID); err != nil {
		log.Fatalf("cannot delete row for testing, error=%v entity=%v", err, r.item)
	}

	if err := r.db.Close(); err != nil {
		log.Fatalf("cannot close db for testing, error=%v", err)
	}
}

func getClientRepository(db *postgres.DB) (r *repository.ClientRepository) {
	r, err := repository.NewClientRepository(db)
	if err != nil {
		log.Fatalf("cannot init ClientRepository, error=%v", err)
	}

	return
}

func getClientTestRows() *ClientTestRows {
	if clientTestRows != nil {
		return clientTestRows
	}

	db := getDatabase()

	clientTestRows = &ClientTestRows{
		db:   db,
		repo: getClientRepository(db),
	}
	return clientTestRows
}
