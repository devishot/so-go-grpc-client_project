package repository_impl

import (
	"database/sql"
	"github.com/pkg/errors"

	"github.com/devishot/so-go-grpc-client_project/domain"
	"github.com/devishot/so-go-grpc-client_project/infrastructure/database"
	"github.com/devishot/so-go-grpc-client_project/infrastructure/database/postgres"
	q "github.com/devishot/so-go-grpc-client_project/interfaces/db/query"
)

func NewClientRepository(db *postgres.DB) (r *ClientRepository, err error) {
	r = &ClientRepository{DB: db}

	err = r.createTable()
	if err != nil {
		return
	}

	return
}

type ClientRepository struct {
	DB *postgres.DB
}

func (r *ClientRepository) createTable() error {
	if _, err := r.DB.Conn.Exec(q.ClientCreateTable); err != nil {
		return errors.WithMessagef(err, "when: ClientCreateTable | table: %s", q.ClientTableName)
	}

	return nil
}

func (r *ClientRepository) Get(id domain.ID) (cl domain.ClientEntity, err error) {
	err = r.DB.Conn.QueryRow(q.ClientFindRowByID, id).
		Scan(&cl.ID, &cl.Timestamp, &cl.FirstName, &cl.LastName, &cl.CompanyName)

	switch err {
	case sql.ErrNoRows:
		err = domain.NotFoundClientRepositoryError
	default:
		err = errors.WithMessagef(err, "when: ClientFindRowByID | table: %s", q.ClientTableName)
	}

	return
}

func (r *ClientRepository) Delete(id domain.ID) error {
	if _, err := r.DB.Conn.Exec(q.ClientDeleteRowByID, id); err != nil {
		return errors.WithMessagef(err, "when: ClientDeleteRowByID | table: %s", q.ClientTableName)
	}
	return nil
}

func (r *ClientRepository) Create(entity domain.ClientEntity) error {
	values, err := database.ExtractValuesFromTaggedStruct(entity, q.ClientTableColumns)
	if err != nil {
		return err
	}

	if _, err := r.DB.Conn.Exec(q.ClientInsertRow, values...); err != nil {
		return errors.WithMessagef(err, "when: ClientInsertRow | table: %s", q.ClientTableName)
	}
	return nil
}
