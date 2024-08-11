//go:build !prod
// +build !prod

package testdb

import (
	"fmt"

	"github.com/davidborzek/tvhgo/config"
	"github.com/davidborzek/tvhgo/db"
)

// Setup configures a in-memory sqlite3 test database.
func Setup() (*db.DB, error) {
	return db.Connect(config.DatabaseConfig{
		Type: config.DatabaseTypeSqlite,
		Path: ":memory:",
	})
}

// Close closes the database connection.
func Close(db *db.DB) {
	db.Close()
}

// TruncateTables truncates the given tables.
func TruncateTables(db *db.DB, tables ...string) error {
	for _, table := range tables {
		_, err := db.Exec(fmt.Sprintf("DELETE FROM %s", table))
		if err != nil {
			return err
		}
	}
	return nil
}
