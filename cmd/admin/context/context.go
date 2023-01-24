package context

import (
	"database/sql"
)

type adminContext struct {
	db *sql.DB
}

var ctx adminContext

func SetDB(db *sql.DB) {
	ctx.db = db
}

func GetDB() *sql.DB {
	if ctx.db == nil {
		panic("admin command context not initialized")
	}

	return ctx.db
}
