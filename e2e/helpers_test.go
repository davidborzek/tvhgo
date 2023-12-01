package e2e_test

import (
	"database/sql"
	"os"
)

func runSqlScript(db *sql.DB, path string) {
	data, err := os.ReadFile("./sql/init.sql")
	if err != nil {
		panic(err)
	}

	if _, err = db.Exec(string(data)); err != nil {
		panic(err)
	}
}
