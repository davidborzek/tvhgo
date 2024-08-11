package config

import (
	"fmt"
	"net/url"
)

const (
	defaultDatabasePath = "./tvhgo.db"
)

// DatabaseType represents the type of the database.
type DatabaseType string

const (
	// DatabaseTypeSqlite represents the sqlite3 database type.
	DatabaseTypeSqlite DatabaseType = "sqlite3"
	// DatabaseTypePostgres represents the postgres database type.
	DatabaseTypePostgres DatabaseType = "postgres"
)

type (
	DatabaseConfig struct {
		// Type is the type of the database.
		Type DatabaseType `yaml:"type" env:"TYPE" default:"sqlite3"`
		// Path is the path to the database file (only for sqlite3).
		Path string `yaml:"path" env:"PATH"`
		// Host is the host of the database server (only for postgres).
		Host string `yaml:"host" env:"HOST"`
		// Port is the port of the database server (only for postgres).
		Port int `yaml:"port" env:"PORT"`
		// User is the user of the database server (only for postgres).
		User string `yaml:"user" env:"USER"`
		// Password is the password of the database server (only for postgres).
		Password string `yaml:"password" env:"PASSWORD"`
		// Database is the name of the database (only for postgres).
		Database string `yaml:"database" env:"DATABASE"`
		// SSLMode is the SSL mode for the connection (only for postgres).
		SSLMode string `yaml:"ssl_mode" env:"SSL_MODE"`
	}
)

func (c *DatabaseConfig) Validate() error {
	if c.Type != DatabaseTypeSqlite && c.Type != DatabaseTypePostgres {
		return fmt.Errorf("invalid database type: %s", c.Type)
	}

	return nil
}

func (c *DatabaseConfig) SetDefaults() {
	if c.Type == "" {
		c.Type = DatabaseTypeSqlite
	}

	if c.Type == DatabaseTypePostgres {
		if c.Port == 0 {
			c.Port = 5432
		}

		if c.SSLMode == "" {
			c.SSLMode = "disable"
		}
	}

	if c.Path == "" {
		c.Path = defaultDatabasePath
	}

	if c.Host == "" {
		c.Host = "localhost"
	}

	if c.User == "" {
		c.User = "tvhgo"
	}

	if c.Database == "" {
		c.Database = "tvhgo"
	}
}

func (c *DatabaseConfig) DSN() (string, error) {
	switch c.Type {
	case DatabaseTypeSqlite:
		query := url.Values{}
		// See https://github.com/mattn/go-sqlite3#connection-string
		query.Set("_foreign_keys", "on")

		return fmt.Sprintf("file:%s?%s", c.Path, query.Encode()), nil
	case DatabaseTypePostgres:
		query := url.Values{}
		query.Set("sslmode", c.SSLMode)

		return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?%s", c.User, c.Password, c.Host, c.Port, c.Database, query.Encode()), nil
	}

	return "", fmt.Errorf("unsupported database type: %s", c.Type)
}
