package twofactorsettings

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/davidborzek/tvhgo/core"
)

type sqlRepository struct {
	db *sql.DB
}

func New(db *sql.DB) core.TwoFactorSettingsRepository {
	return &sqlRepository{
		db: db,
	}
}

func (s *sqlRepository) Find(ctx context.Context, userID int64) (*core.TwoFactorSettings, error) {
	row := s.db.QueryRowContext(ctx, queryByUserID, userID)
	settings := new(core.TwoFactorSettings)
	if err := scanRow(row, settings); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}
	return settings, nil
}

func (s *sqlRepository) Create(ctx context.Context, settings *core.TwoFactorSettings) error {
	createdAt := time.Now().Unix()

	_, err := s.db.ExecContext(ctx, stmtInsert,
		settings.UserID,
		settings.Secret,
		settings.Enabled,
		createdAt,
		createdAt,
	)

	if err != nil {
		return err
	}

	settings.UpdatedAt = createdAt
	settings.CreatedAt = createdAt
	return nil
}

func (s *sqlRepository) Delete(ctx context.Context, settings *core.TwoFactorSettings) error {
	_, err := s.db.ExecContext(ctx, stmtDelete, settings.UserID)
	return err
}

func (s *sqlRepository) UpdateEnabled(ctx context.Context, settings *core.TwoFactorSettings) error {
	res, err := s.db.ExecContext(ctx, stmtUpdateEnabled, settings.Enabled, settings.UserID)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("no rows updated")
	}

	return err
}
