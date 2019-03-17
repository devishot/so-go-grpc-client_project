package _examples

import (
	"fmt"
	"log"

	"github.com/devishot/so-go-grpc-client_project/domain"
	"github.com/devishot/so-go-grpc-client_project/infrastructure/database/postgres"
	"github.com/devishot/so-go-grpc-client_project/interfaces/db/repository_impl"
)

var projectTestRows *ProjectTestRows

type ProjectTestRows struct {
	db    *postgres.DB
	repo  *repository_impl.ProjectRepository
	items []domain.ProjectEntity

	clientTR *ClientTestRows
}

func (r *ProjectTestRows) Get() domain.ProjectEntity {
	if len(r.items) == 0 {
		r.prepare()
	}

	return r.items[0]
}

func (r *ProjectTestRows) GetLast() domain.ProjectEntity {
	if len(r.items) == 0 {
		r.prepare()
	}

	return r.items[len(r.items)-1]
}

func (r *ProjectTestRows) prepare() {
	log.Println("inserting ProjectEntity rows for testing")

	cID := r.getClientID()

	for i := 0; i < 2; i++ {
		p := domain.NewProject(cID, fmt.Sprintf("test project-%d", i+1), "test")

		err := r.repo.Create(p)
		if err != nil {
			log.Fatalf("cannot insert row for testing, error=%v entity=%v", err, p)
		}

		r.items = append(r.items, p)
	}
}

func (r *ProjectTestRows) getClientID() domain.ID {
	if r.clientTR == nil {
		r.clientTR = getClientTestRows()
	}

	return r.clientTR.Get().ID
}

func (r *ProjectTestRows) Release() {
	for _, p := range r.items {
		if err := r.repo.Delete(p.ID); err != nil {
			log.Fatalf("cannot delete row for testing, error=%v entity=%v", err, p)
		}
	}

	r.clientTR.Release()
}

func getProjectRepository(db *postgres.DB) (r *repository_impl.ProjectRepository) {
	r, err := repository_impl.NewProjectRepository(db)
	if err != nil {
		log.Fatalf("cannot init ProjectRepository, error=%v", err)
	}

	return
}

func getProjectTestRows() *ProjectTestRows {
	if projectTestRows != nil {
		return projectTestRows
	}

	db := getDatabase()

	projectTestRows = &ProjectTestRows{
		db:   db,
		repo: getProjectRepository(db),
	}
	return projectTestRows
}
