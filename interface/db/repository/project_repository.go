package repository

import (
	"github.com/pkg/errors"

	"github.com/devishot/so-go-grpc-client_project/domain"
	"github.com/devishot/so-go-grpc-client_project/infrastructure/database/postgres"
)

const ProjectTableName = "cp_project"

func NewProjectRepository(db *postgres.DB) (*ProjectRepository, error) {
	r := &ProjectRepository{DB: db}

	err := r.createTable()
	if err != nil {
		return nil, err
	}

	return r, nil
}

type ProjectRepository struct {
	DB *postgres.DB
}

func (r *ProjectRepository) createTable() error {
	if _, err := r.DB.Conn.Exec(ProjectCreateTable); err != nil {
		return errors.WithMessagef(err, "when: ProjectCreateTable | table: %s", ProjectTableName)
	}
	return nil
}

func (r *ProjectRepository) Get(id domain.ID) (p domain.ProjectEntity, err error) {
	err = r.DB.Conn.QueryRow(ProjectFindByID, id).
		Scan(&p.ID, &p.ClientID, &p.Timestamp, &p.Title, &p.Description)
	if err != nil {
		err = errors.WithMessagef(err, "when: ProjectFindByID | table: %s", ProjectTableName)
		return
	}

	return
}
