package migration

import (
	"database/sql"
	"embed"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

//go:embed sqlite3/*.sql
var sqliteFs embed.FS

//go:embed postgres/*.sql
var postgresFs embed.FS

// Migrate runs the SQL migration provided by the sql files in the `files` directory.
func Migrate(driverType string, db *sql.DB) error {
	migrator, err := getMigrator(driverType, db)
	if err != nil {
		return err
	}

	if err := migrator.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}

func getMigrator(driverType string, db *sql.DB) (*migrate.Migrate, error) {

	var (
		driver database.Driver
		err    error
		fs     embed.FS
	)

	switch driverType {
	case "sqlite3":
		driver, err = sqlite3.WithInstance(db, &sqlite3.Config{})
		if err != nil {
			return nil, err
		}
		fs = sqliteFs
	case "postgres":
		driver, err = postgres.WithInstance(db, &postgres.Config{})
		if err != nil {
			return nil, err
		}
		fs = postgresFs
	default:
		return nil, fmt.Errorf("unsupported driver: %s", driverType)
	}

	source, err := iofs.New(fs, driverType)
	if err != nil {
		return nil, err
	}

	return migrate.NewWithInstance("iofs", source, driverType, driver)
}
