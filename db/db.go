package db

import (
	"database/sql"

	"github.com/davidborzek/tvhgo/db/migration"
)

// Connect connects to a database and migrates schema.
func Connect(dsn string) (*sql.DB, error) {
	// TODO: make connection work with other sql drivers.

	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}

	// Enable foreign keys in sqlite3 (TODO: only do this for sqlite3 driver)
	if _, err := db.Exec("PRAGMA foreign_keys = ON"); err != nil {
		return nil, err
	}

	if err := migration.Migrate(db); err != nil {
		return nil, err
	}

	return db, nil
}
