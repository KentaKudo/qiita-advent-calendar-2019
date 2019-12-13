package main

import (
	"database/sql"

	"github.com/KentaKudo/qiita-advent-calendar-2019/internal/schema"
	"github.com/pkg/errors"
)

func initDB(connString string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, errors.Wrap(err, "ping")
	}
	return db, nil
}

var _ todoManager = (*store)(nil)

type store struct {
	db *sql.DB
}

func newStore(db *sql.DB, version int) (*store, error) {
	if err := schema.Migrate(db, version); err != nil {
		return nil, errors.Wrap(err, "migrate db schema")
	}

	return &store{db: db}, nil
}

func (s *store) projectTodo(t todo) error {
	if _, err := s.db.Exec(
		`INSERT INTO todo (id, title, description) VALUES ($1, $2, $3)`,
		t.id, t.title, t.description,
	); err != nil {
		return err
	}

	return nil
}
