package postgres

import (
	"database/sql"
	"errors"
	"fmt"

	// PostgreSQL library which implements database/sql interface
	_ "github.com/lib/pq"
)

type DB struct {
	Conn *sql.DB
	cfg  Config
}

func New(cfg Config) (*DB, error) {
	conn, err := sql.Open("postgres", fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%d sslmode=disable",
		cfg.User, cfg.Password, cfg.Database, cfg.Host, cfg.Port))
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
