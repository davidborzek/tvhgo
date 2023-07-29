package auth

import (
	"context"
	"errors"
	"fmt"

	"github.com/davidborzek/tvhgo/config"
	"github.com/davidborzek/tvhgo/core"
	"github.com/pquerna/otp/totp"
)

// TODO: refactor errors
type twoFactorAuthService struct {
	userRepository              core.UserRepository
	twoFactorSettingsRepository core.TwoFactorSettingsRepository
	cfg                         *config.TOTPConfig
}

func NewTwoFactorAuthService(
	twoFactorSettingsRepository core.TwoFactorSettingsRepository,
	userRepository core.UserRepository,
	cfg *config.TOTPConfig,
) core.TwoFactorAuthService {
	return &twoFactorAuthService{
		twoFactorSettingsRepository: twoFactorSettingsRepository,
		userRepository:              userRepository,
		cfg:                         cfg,
	}
}

func (s *twoFactorAuthService) Setup(ctx context.Context, userID int64) (string, error) {
	user, err := s.userRepository.FindById(ctx, userID)
	if err != nil {
		return "", fmt.Errorf("two factor service failed to find user by id %w", err)
	}

	if user == nil {
		return "", fmt.Errorf("two factor service could not find user with id %d", userID)
	}

	existingSettings, err := s.twoFactorSettingsRepository.Find(ctx, userID)
	if err != nil {
		return "", fmt.Errorf("two factor service failed search for existing settings %w", err)
	}

	if existingSettings != nil {
		if existingSettings.Enabled {
			return "", core.ErrTwoFactorAuthAlreadyEnabled
		}

		// TODO: this or generate new secret and update database?
		return buildTotpUrl(s.cfg.Issuer, user.Username, existingSettings.Secret), nil
	}

	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      s.cfg.Issuer,
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

func (s *twoFactorAuthService) Deactivate(ctx context.Context, userID int64) error {
	settings, err := s.twoFactorSettingsRepository.Find(ctx, userID)
	if err != nil {
		return fmt.Errorf("two factor service failed search for existing settings %w", err)
	}

	if settings == nil {
		return core.ErrTwoFactorAuthNotEnabled
	}

	if err := s.twoFactorSettingsRepository.Delete(ctx, settings); err != nil {
		return fmt.Errorf("failed to delete two factor settings: %w", err)
	}

	return nil
}

func (s *twoFactorAuthService) Verify(ctx context.Context, userID int64, code *string) error {
	settings, err := s.twoFactorSettingsRepository.Find(ctx, userID)
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

func (s *twoFactorAuthService) Activate(ctx context.Context, userID int64, code string) error {
	settings, err := s.twoFactorSettingsRepository.Find(ctx, userID)
	if err != nil {
		return fmt.Errorf("two factor service failed search for existing settings %w", err)
	}

	if settings == nil {
		return errors.New("no two factor settings found")
	}

	if settings.Enabled {
		return core.ErrTwoFactorAuthAlreadyEnabled
	}

	if !totp.Validate(code, settings.Secret) {
		return core.ErrTwoFactorCodeInvalid
	}

	settings.Enabled = true

	return s.twoFactorSettingsRepository.UpdateEnabled(ctx, settings)
}

func (s *twoFactorAuthService) GetSettings(ctx context.Context, userID int64) (*core.TwoFactorSettings, error) {
	settings, err := s.twoFactorSettingsRepository.Find(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("two factor service failed find settings %w", err)
	}

	if settings == nil {
		return &core.TwoFactorSettings{
			Enabled: false,
		}, nil
	}

	return settings, nil
}

func buildTotpUrl(issuer string, username string, secret string) string {
	return fmt.Sprintf("otpauth://totp/%s:%s?algorithm=SHA1&digits=6&issuer=%s&period=30&secret=%s", issuer, username, issuer, secret)
}
