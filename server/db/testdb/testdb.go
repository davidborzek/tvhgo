//go:build !prod
// +build !prod

package testdb

import (
	"database/sql"
	"fmt"

	"github.com/davidborzek/tvhgo/db"
)

// Setup configures a in-memory sqlite3 test database.
func Setup() (*sql.DB, error) {
	return db.Connect(":memory:")
}

// Close closes the database connection.
func Close(db *sql.DB) {
	db.Close()
}

// TruncateTables truncates the given tables.
func TruncateTables(db *sql.DB, tables ...string) error {
	for _, table := range tables {
		_, err := db.Exec(fmt.Sprintf("DELETE FROM %s", table))
		if err != nil {
			return err
		}
	}
	return nil
}
