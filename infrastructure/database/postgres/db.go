package postgres

import (
	"errors"
	"fmt"

	"database/sql"
	_ "github.com/lib/pq"
)

type DB struct {
	Conn *sql.DB
	cfg  Config
}

func New(cfg Config) (*DB, error) {
	conn, err := sql.Open("postgres", fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database))
	if err != nil {
		return nil, err
	}

	if err := conn.Ping(); err != nil {
		return nil, err
	}

	return &DB{Conn: conn, cfg: cfg}, nil
}

func (db *DB) Close() error {
	if db.Conn == nil {
		return errors.New("no connection")
	}

	return db.Conn.Close()
}
