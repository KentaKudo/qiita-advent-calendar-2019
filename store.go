package main

import (
	"database/sql"

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
