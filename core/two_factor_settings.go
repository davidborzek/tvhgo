package core

import "context"

type (
	// TwoFactorSettings defines the two factor settings of a user.
	TwoFactorSettings struct {
		UserID    int64  `json:"-"`
		Secret    string `json:"-"`
		Enabled   bool   `json:"enabled"`
		CreatedAt int64  `json:"-"`
		UpdatedAt int64  `json:"-"`
	}

	// TwoFactorSettingsRepository defines CRUD operations working with TwoFactorSettings.
	TwoFactorSettingsRepository interface {
		// Find returns two factor settings by a user id.
		Find(ctx context.Context, userID int64) (*TwoFactorSettings, error)

		// Create persists new two factor settings.
		Create(ctx context.Context, settings *TwoFactorSettings) error

		// Delete deletes two factor settings.
		Delete(ctx context.Context, settings *TwoFactorSettings) error

		Update(ctx context.Context, settings *TwoFactorSettings) error

		Save(ctx context.Context, settings *TwoFactorSettings) error
	}
)
