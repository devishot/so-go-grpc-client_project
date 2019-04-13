package repository_impl

import (
	"database/sql"
	"time"

	"github.com/pkg/errors"

	"github.com/devishot/so-go-grpc-client_project/domain"
	"github.com/devishot/so-go-grpc-client_project/infrastructure/database"
	"github.com/devishot/so-go-grpc-client_project/infrastructure/database/postgres"
	q "github.com/devishot/so-go-grpc-client_project/interfaces/db/query"
)

func NewProjectRepository(db *postgres.DB) (repo ProjectRepository, err error) {
	repo = ProjectRepository{DB: db}

	err = repo.createTable()
	if err != nil {
		return
	}

	return
}

type ProjectRepository struct {
	DB *postgres.DB
}

func (r ProjectRepository) createTable() error {
	if _, err := r.DB.Conn.Exec(q.ProjectCreateTable); err != nil {
		return errors.WithMessagef(err, "when: ProjectCreateTable | table: %s", q.ProjectTableName)
	}
	return nil
}

func (r ProjectRepository) Get(id domain.ID) (p domain.ProjectEntity, err error) {
	err = r.DB.Conn.QueryRow(q.ProjectFindRowByID, id).
		Scan(&p.ID, &p.ClientID, &p.Timestamp, &p.Title, &p.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			err = domain.NotFoundProjectRepositoryError
		} else {
			err = errors.WithMessagef(err, "when: ProjectFindByID | table: %s", q.ProjectTableName)
		}
	}

	return
}

func (r ProjectRepository) Delete(id domain.ID) error {
	if _, err := r.DB.Conn.Exec(q.ProjectDeleteRowByID, id); err != nil {
		return errors.WithMessagef(err, "when: ProjectDeleteByID | table: %s", q.ProjectTableName)
	}
	return nil
}

func (r ProjectRepository) Create(entity domain.ProjectEntity) error {
	values, err := database.ExtractValuesFromTaggedStruct(entity, q.ProjectTableColumns)
	if err != nil {
		return err
	}

	if _, err := r.DB.Conn.Exec(q.ProjectInsertRow, values...); err != nil {
		return errors.WithMessagef(err, "when: ProjectInsertRow | table: %s", q.ProjectTableName)
	}
	return nil
}

func (r ProjectRepository) GetByClient(clientID domain.ID) (projects []domain.ProjectEntity, err error) {
	rows, err := r.DB.Conn.Query(q.ProjectFindRowsByClientID, clientID)
	if err != nil {
		return
	}
	defer database.MustClose(rows)

	for rows.Next() {
		p := domain.ProjectEntity{}

		err = rows.Scan(&p.ID, &p.ClientID, &p.Timestamp, &p.Title, &p.Description)
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

func (r ProjectRepository) GetLastByClient(cID domain.ID) (p domain.ProjectEntity, err error) {
	err = r.DB.Conn.QueryRow(q.ProjectFindLastRowByClientID, cID).
		Scan(&p.ID, &p.ClientID, &p.Timestamp, &p.Title, &p.Description)

	switch err {
	case sql.ErrNoRows:
		err = domain.NotFoundProjectRepositoryError
	default:
		err = errors.WithMessagef(err, "when: ProjectFindLastRowByClientID | table: %s", q.ProjectTableName)
	}

	return
}

func (r ProjectRepository) GetFirstByClient(cID domain.ID) (
	p domain.ProjectEntity, err error) {
	err = r.DB.Conn.QueryRow(q.ProjectFindFirstRowByClientID, cID).
		Scan(&p.ID, &p.ClientID, &p.Timestamp, &p.Title, &p.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			err = domain.NotFoundProjectRepositoryError
		} else {
			err = errors.WithMessagef(err, "when: ProjectFindFirstRowByClientID | table: %s", q.ProjectTableName)
		}
	}

	return
}

func (r ProjectRepository) PaginateForwardByClientByTimestamp(cID domain.ID, first int, after time.Time) (
	projects []domain.ProjectEntity, err error) {
	var rows *sql.Rows

	rows, err = r.DB.Conn.Query(q.ProjectFindRowsForForwardPage, cID, after, first)
	if err != nil {
		return
	}

	defer database.MustClose(rows)

	for rows.Next() {
		p := domain.ProjectEntity{}

		err = rows.Scan(&p.ID, &p.ClientID, &p.Timestamp, &p.Title, &p.Description)
		if err != nil {
			return
		}

		projects = append(projects, p)
	}

	err = rows.Err()
	return
}

func (r ProjectRepository) PaginateBackwardByClientByTimestamp(clientID domain.ID, last int, before time.Time) (
	projects []domain.ProjectEntity, err error) {
	var rows *sql.Rows

	rows, err = r.DB.Conn.Query(q.ProjectFindRowsForBackwardPage, clientID, before, last)
	if err != nil {
		return
	}

	defer database.MustClose(rows)

	for rows.Next() {
		p := domain.ProjectEntity{}

		err = rows.Scan(&p.ID, &p.ClientID, &p.Timestamp, &p.Title, &p.Description)
		if err != nil {
			return
		}

		projects = append(projects, p)
	}

	err = rows.Err()
	return
}
