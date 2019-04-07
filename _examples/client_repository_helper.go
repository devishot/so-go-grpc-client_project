package _examples

import (
	"fmt"
	"log"

	"github.com/devishot/so-go-grpc-client_project/domain"
	"github.com/devishot/so-go-grpc-client_project/infrastructure/database/postgres"
	impl "github.com/devishot/so-go-grpc-client_project/interfaces/db/repository_impl"
)

var clientTestRows *ClientTestRows

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

func getClientRepository(db *postgres.DB) (r *impl.ClientRepository) {
	r, err := impl.NewClientRepository(db)
	if err != nil {
		log.Fatalf("cannot init ClientRepository, error=%v", err)
	}

	return
}

type ClientTestRows struct {
	db   *postgres.DB
	repo *impl.ClientRepository
}

func (r *ClientTestRows) Get(n int) (items []domain.ClientEntity) {
	log.Println("inserting ClientEntity rows for testing")

	for i := 1; i <= n; i++ {
		fn, ln, cn := fmt.Sprintf("test%dFirstName", i),
			fmt.Sprintf("test%dtestLastName", i),
			fmt.Sprintf("test%dtestCompanyName", i)
		cl := domain.NewClient(fn, ln, cn)

		err := r.repo.Create(cl)
		if err != nil {
			log.Fatalf("cannot insert row for testing, error=%v entity=%v", err, cl)
		}

		items = append(items, cl)
	}

	return items
}

func (r *ClientTestRows) Release() {
	if err := r.db.Close(); err != nil {
		log.Fatalf("cannot close db for testing, error=%v", err)
	}
}

func (r *ClientTestRows) Delete(items []domain.ClientEntity) {
	for _, cl := range items {
		if err := r.repo.Delete(cl.ID); err != nil {
			log.Fatalf("cannot delete row for testing, error=%v entity=%v", err, cl)
		}
	}
}
