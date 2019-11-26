package schema

import (
	"database/sql"

	"github.com/pkg/errors"
)

// Migrate sets the schema at the requested version
func Migrate(db *sql.DB, expectedVersion int) error {
	currentVersion, err := Version(db)
	if err != nil {
		return err
	}
	switch {
	case expectedVersion == currentVersion:
		return nil
	case expectedVersion < currentVersion:
		return errors.Errorf("schema migrate: invalid request, can not migrate backwards, current: %v, expected: %v", currentVersion, expectedVersion)
	case expectedVersion > currentVersion:
		for ; expectedVersion > currentVersion; currentVersion++ {
			if _, err := db.Exec(schemas[currentVersion+1]); err != nil {
				return errors.Wrap(err, "exec migration")
			}
		}
	}
	return nil
}

// Version returns the current schema version
func Version(db *sql.DB) (v int, err error) {
	if err := db.QueryRow("SELECT id FROM schema_version WHERE md_curr = true").Scan(&v); err != nil {
		if err.Error() == `pq: relation "schema_version" does not exist` {
			return -1, nil
		}
	}
	return v, errors.Wrap(err, "select version")
}
