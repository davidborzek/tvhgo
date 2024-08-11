package twofactorsettings

import (
	"context"
	"database/sql"
	"time"

	"github.com/davidborzek/tvhgo/core"
	"github.com/davidborzek/tvhgo/db"
)

type sqlRepository struct {
	db *db.DB
}

func New(db *db.DB) core.TwoFactorSettingsRepository {
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

func (s *sqlRepository) Update(ctx context.Context, settings *core.TwoFactorSettings) error {
	updatedAt := time.Now().Unix()

	_, err := s.db.ExecContext(ctx, stmtUpdate,
		settings.Secret,
		settings.Enabled,
		updatedAt,
		settings.UserID,
	)

	if err != nil {
		return err
	}

	settings.UpdatedAt = updatedAt
	return nil
}

func (s *sqlRepository) Save(ctx context.Context, settings *core.TwoFactorSettings) error {
	e, err := s.Find(ctx, settings.UserID)
	if err != nil {
		return err
	}

	if e == nil {
		return s.Create(ctx, settings)
	}

	return s.Update(ctx, settings)
}
