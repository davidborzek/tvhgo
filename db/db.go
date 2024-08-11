package db

import (
	"database/sql"

	"github.com/davidborzek/tvhgo/config"
	"github.com/davidborzek/tvhgo/db/migration"
)

type DB struct {
	*sql.DB
	Type config.DatabaseType
}

// Connect connects to a database and migrates schema.
func Connect(cfg config.DatabaseConfig) (*DB, error) {
	driver := string(cfg.Type)
	dsn, err := cfg.DSN()
	if err != nil {
		return nil, err
	}

	pool, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, err
	}

	if err := migration.Migrate(driver, pool); err != nil {
		return nil, err
	}

	return &DB{
		DB:   pool,
		Type: cfg.Type,
	}, nil
}
