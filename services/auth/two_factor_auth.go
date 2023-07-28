package auth

import (
	"context"
	"errors"
	"fmt"

	"github.com/davidborzek/tvhgo/core"
	"github.com/pquerna/otp/totp"
)

const (
	twoFactorIssuer = "tvhgo"
)

type twoFactorAuthService struct {
	userRepository              core.UserRepository
	twoFactorSettingsRepository core.TwoFactorSettingsRepository
}

func NewTwoFactorAuthService(
	twoFactorSettingsRepository core.TwoFactorSettingsRepository,
	userRepository core.UserRepository,
) core.TwoFactorAuthService {
	return &twoFactorAuthService{
		twoFactorSettingsRepository: twoFactorSettingsRepository,
		userRepository:              userRepository,
	}
}

func (s *twoFactorAuthService) Setup(ctx context.Context, userId int64) (string, error) {
	// TODO: maybe Password Check?

	existingSettings, err := s.twoFactorSettingsRepository.Find(ctx, userId)
	if err != nil {
		return "", fmt.Errorf("two factor service failed search for existing settings %w", err)
	}

	if existingSettings != nil {
		return "", fmt.Errorf("two factor auth already enabled")
	}

	user, err := s.userRepository.FindById(ctx, userId)
	if err != nil {
		return "", fmt.Errorf("two factor service failed to find user by id %w", err)
	}

	if user == nil {
		return "", fmt.Errorf("two factor service could not find user with id %d", userId)
	}

	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      twoFactorIssuer,
		AccountName: user.Username,
	})

	if err != nil {
		return "", fmt.Errorf("two factor service failed to generate totp secret: %w", err)
	}

	twoFactorSettings := &core.TwoFactorSettings{
		UserID:  user.ID,
		Secret:  key.Secret(),
		Enabled: false,
	}

	if err := s.twoFactorSettingsRepository.Create(ctx, twoFactorSettings); err != nil {
		return "", fmt.Errorf("two factor service failed persist two factor settings: %w", err)
	}

	return key.URL(), nil
}

func (s *twoFactorAuthService) Deactivate(ctx context.Context, userId int64) error {
	// TODO: Password Check

	settings, err := s.twoFactorSettingsRepository.Find(ctx, userId)
	if err != nil {
		return fmt.Errorf("two factor service failed search for existing settings %w", err)
	}

	if settings == nil {
		return fmt.Errorf("two factor auth not enabled")
	}

	if err := s.twoFactorSettingsRepository.Delete(ctx, settings); err != nil {
		return fmt.Errorf("failed to delete two factor settings: %w", err)
	}

	return nil
}

func (s *twoFactorAuthService) Verify(ctx context.Context, userId int64, code *string) error {
	settings, err := s.twoFactorSettingsRepository.Find(ctx, userId)
	if err != nil {
		return fmt.Errorf("two factor service failed search for existing settings %w", err)
	}

	if settings == nil || !settings.Enabled {
		return nil
	}

	if code == nil {
		return core.ErrTwoFactorRequired
	}

	if totp.Validate(*code, settings.Secret) {
		return nil
	}

	return core.ErrTwoFactorCodeInvalid
}

func (s *twoFactorAuthService) Enable(ctx context.Context, userID int64, code string) error {
	settings, err := s.twoFactorSettingsRepository.Find(ctx, userID)
	if err != nil {
		return fmt.Errorf("two factor service failed search for existing settings %w", err)
	}

	if settings == nil {
		return errors.New("no two factor settings found")
	}

	if settings.Enabled {
		return errors.New("two factor auth already enabled")
	}

	if !totp.Validate(code, settings.Secret) {
		return core.ErrTwoFactorCodeInvalid
	}

	settings.Enabled = true

	// TODO: Password Check
	return s.twoFactorSettingsRepository.UpdateEnabled(ctx, settings)
}
