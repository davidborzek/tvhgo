package session

import (
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
