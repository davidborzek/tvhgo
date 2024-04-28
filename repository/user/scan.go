package user

import (
	"database/sql"

	"github.com/davidborzek/tvhgo/core"
	"github.com/davidborzek/tvhgo/repository"
)

// Internal helper to scan a sql.Row into a user model.
func scanRow(scanner repository.Scanner, dest *core.User) error {
	var (
		twoFactor sql.NullBool
	)

	err := scanner.Scan(
		&dest.ID,
		&dest.Username,
		&dest.PasswordHash,
		&dest.Email,
		&dest.DisplayName,
		&dest.IsAdmin,
		&dest.CreatedAt,
		&dest.UpdatedAt,
		&twoFactor,
	)

	if err != nil {
		return err
	}

	if twoFactor.Valid {
		dest.TwoFactor = twoFactor.Bool
	}

	return nil
}

// Internal helper to scan sql.Rows into an array of user models.
func scanRows(rows *sql.Rows) ([]*core.User, error) {
	defer rows.Close()

	users := []*core.User{}
	for rows.Next() {
		user := new(core.User)
		if err := scanRow(rows, user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
