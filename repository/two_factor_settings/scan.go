package twofactorsettings

import (
	"github.com/davidborzek/tvhgo/core"
	"github.com/davidborzek/tvhgo/repository"
)

// Internal helper to scan a sql.Row into a user model.
func scanRow(scanner repository.Scanner, dest *core.TwoFactorSettings) error {
	return scanner.Scan(
		&dest.UserID,
		&dest.Secret,
		&dest.CreatedAt,
		&dest.UpdatedAt,
	)
}
