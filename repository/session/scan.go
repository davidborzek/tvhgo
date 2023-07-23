package session

import (
	"database/sql"

	"github.com/davidborzek/tvhgo/core"
	"github.com/davidborzek/tvhgo/repository"
)

// Internal helper to scan a sql.Row into a session model.
func scanRow(scanner repository.Scanner, dest *core.Session) error {
	return scanner.Scan(
		&dest.ID,
		&dest.UserId,
		&dest.HashedToken,
		&dest.ClientIP,
		&dest.UserAgent,
		&dest.CreatedAt,
		&dest.LastUsedAt,
		&dest.RotatedAt,
	)
}

// Internal helper to scan sql.Rows into an array of user models.
func scanRows(rows *sql.Rows) ([]*core.Session, error) {
	defer rows.Close()

	sessions := []*core.Session{}
	for rows.Next() {
		session := new(core.Session)
		if err := scanRow(rows, session); err != nil {
			return nil, err
		}
		sessions = append(sessions, session)
	}
	return sessions, nil
}
