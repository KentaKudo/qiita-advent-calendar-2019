package main

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testDB = fmt.Sprintf("test-%s", uuid.New().String())

func openDB(t *testing.T) *sql.DB {
	connStr := "postgres://root@localhost:26257/?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	require.NoError(t, err)

	_, err = db.Exec(fmt.Sprintf(`CREATE DATABASE IF NOT EXISTS "%s"; USE "%s"`, testDB, testDB))
	require.NoError(t, err)

	return db
}

func closeDB(t *testing.T, db *sql.DB) {
	require.NoError(t, db.Close())
}

func TestStore_ProjectTodo(t *testing.T) {
	t.Run("project a new todo", func(t *testing.T) {
		db := openDB(t)
		defer closeDB(t, db)

		sut, err := newStore(db, defaultSchemaVersion)
		require.NoError(t, err)

		input := todo{
			id:          uuid.New().String(),
			title:       "foo title",
			description: "foo description",
		}

		require.NoError(t, sut.projectTodo(input))

		var got todo
		require.NoError(t, db.QueryRow(
			`SELECT id, title, description FROM todo WHERE id = $1`,
			input.id,
		).Scan(&got.id, &got.title, &got.description))
		assert.Equal(t, input, got)
	})
}
