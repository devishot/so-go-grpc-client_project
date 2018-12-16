package _examples

import (
	"log"

	"github.com/devishot/so-go-grpc-client_project/domain"
	"github.com/devishot/so-go-grpc-client_project/domain_interface/db/repository"
	"github.com/devishot/so-go-grpc-client_project/infrastructure/database/postgres"
)

var clientTestRows *ClientTestRows

type ClientTestRows struct {
	db     *postgres.DB
	repo   domain.ClientRepository
	client domain.ClientEntity
}

func (r *ClientTestRows) Get() domain.ClientEntity {
	if !r.client.IsZero() {
		return r.client
	}

	cl := domain.NewClient("testFirstName", "testLastName", "testCompanyName")
	err := r.repo.Create(cl)
	if err != nil {
		log.Fatalf("cannot insert row for testing, entity=%v error=%v", cl, err)
	}

	r.client = cl
	return r.client
}

func (r *ClientTestRows) Release() {
	if err := r.repo.Delete(r.client.ID); err != nil {
		log.Fatalf("cannot delete row for testing, entity=%v error=%v", r.client, err)
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
