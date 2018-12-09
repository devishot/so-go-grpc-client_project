package repository

import (
	"database/sql"
	"github.com/pkg/errors"

	"github.com/devishot/so-go-grpc-client_project/domain"
	"github.com/devishot/so-go-grpc-client_project/infrastructure/database"
	"github.com/devishot/so-go-grpc-client_project/infrastructure/database/postgres"
)

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
	err = r.DB.Conn.QueryRow(ProjectFindRowByID, id).
		Scan(&p.ID, &p.ClientID, &p.Timestamp, &p.Title, &p.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			err = domain.NotFoundProjectRepositoryError
		} else {
			err = errors.WithMessagef(err, "when: ProjectFindByID | table: %s", ProjectTableName)
		}
	}

	return
}

func (r *ProjectRepository) Delete(id domain.ID) error {
	if _, err := r.DB.Conn.Exec(ProjectDeleteRowByID, id); err != nil {
		return errors.WithMessagef(err, "when: ProjectDeleteByID | table: %s", ProjectTableName)
	}
	return nil
}

func (r *ProjectRepository) Create(entity domain.ProjectEntity) error {
	values, err := database.GetValuesInOrder(entity, ProjectFields)
	if err != nil {
		return err
	}

	if _, err := r.DB.Conn.Exec(ProjectInsertRow, values); err != nil {
		return errors.WithMessagef(err, "when: ProjectCreateRow | table: %s", ProjectTableName)
	}
	return nil
}

func (r *ProjectRepository) GetByClient(clientID domain.ID) (projects []domain.ProjectEntity, err error) {
	rows, err := r.DB.Conn.Query(ProjectFindRowsByClientID, clientID)
	if err != nil {
		return
	}
	defer database.Must(rows.Close())

	for rows.Next() {
		p := domain.ProjectEntity{}

		err := rows.Scan(&p.ID, &p.ClientID, &p.Timestamp, &p.Title, &p.Description)
		if err != nil {
			return
		}

		projects = append(projects, p)
	}

	err = rows.Err()
	if err != nil {
		return
	}

	return
}
