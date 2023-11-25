package token

import (
	"database/sql"

	"github.com/davidborzek/tvhgo/core"
	"github.com/davidborzek/tvhgo/repository"
)

// Internal helper to scan a sql.Row into a token model.
func scanRow(scanner repository.Scanner, dest *core.Token) error {
	return scanner.Scan(
		&dest.ID,
		&dest.UserID,
		&dest.Name,
		&dest.HashedToken,
		&dest.CreatedAt,
		&dest.UpdatedAt,
	)
}

// Internal helper to scan sql.Rows into an array of user models.
func scanRows(rows *sql.Rows) ([]*core.Token, error) {
	defer rows.Close()

	tokens := []*core.Token{}
	for rows.Next() {
		token := new(core.Token)
		if err := scanRow(rows, token); err != nil {
			return nil, err
		}
		tokens = append(tokens, token)
	}
	return tokens, nil
}
