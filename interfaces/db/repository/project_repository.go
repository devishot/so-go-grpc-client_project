package repository

import (
	"database/sql"
	"github.com/pkg/errors"

	"github.com/devishot/so-go-grpc-client_project/domain"
	"github.com/devishot/so-go-grpc-client_project/infrastructure/database"
	"github.com/devishot/so-go-grpc-client_project/infrastructure/database/postgres"
	"github.com/devishot/so-go-grpc-client_project/interfaces/graphql_connection"
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
	values, err := database.ExtractValuesFromTaggedStruct(entity, ProjectTableColumns)
	if err != nil {
		return err
	}

	if _, err := r.DB.Conn.Exec(ProjectInsertRow, values...); err != nil {
		return errors.WithMessagef(err, "when: ProjectInsertRow | table: %s", ProjectTableName)
	}
	return nil
}

func (r *ProjectRepository) GetByClient(clientID domain.ID) (projects []domain.ProjectEntity, err error) {
	rows, err := r.DB.Conn.Query(ProjectFindRowsByClientID, clientID)
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

func (r *ProjectRepository) GetLastByClient(cID domain.ID) (p domain.ProjectEntity, err error) {
	err = r.DB.Conn.QueryRow(ProjectFindLastRowByClientID, cID).
		Scan(&p.ID, &p.ClientID, &p.Timestamp, &p.Title, &p.Description)

	switch err {
	case sql.ErrNoRows:
		err = domain.NotFoundProjectRepositoryError
	default:
		err = errors.WithMessagef(err, "when: ProjectFindLastRowByClientID | table: %s", ProjectTableName)
	}

	return
}

func (r *ProjectRepository) GetFirstByClient(cID domain.ID) (
	p domain.ProjectEntity, err error) {
	err = r.DB.Conn.QueryRow(ProjectFindFirstRowByClientID, cID).
		Scan(&p.ID, &p.ClientID, &p.Timestamp, &p.Title, &p.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			err = domain.NotFoundProjectRepositoryError
		} else {
			err = errors.WithMessagef(err, "when: ProjectFindFirstRowByClientID | table: %s", ProjectTableName)
		}
	}

	return
}

func (r *ProjectRepository) PaginateByTimestamp(cID domain.ID, args graphql_connection.ProjectRepositoryPageArgs) (
	projects []domain.ProjectEntity, err error) {
	var rows *sql.Rows

	if args.IsForward() {
		rows, err = r.DB.Conn.Query(ProjectFindRowsForForwardPage, cID, args.After, args.First)
	} else {
		rows, err = r.DB.Conn.Query(ProjectFindRowsForBackwardPage, cID, args.Before, args.Last)
	}

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
