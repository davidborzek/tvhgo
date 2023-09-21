package auth

import (
	"context"
	"errors"

	"github.com/davidborzek/tvhgo/config"
	"github.com/davidborzek/tvhgo/core"
	"github.com/pquerna/otp/totp"
)

var (
	errTwoFactorServiceUserNotFound = errors.New("two factor service: user not found")
)

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
	settings, err := s.twoFactorSettingsRepository.Find(ctx, userID)
	if err != nil {
		return "", err
	}

	if settings != nil && settings.Enabled {
		return "", core.ErrTwoFactorAuthAlreadyEnabled
	}

	user, err := s.userRepository.FindById(ctx, userID)
	if err != nil {
		return "", err
	}

	if user == nil {
		return "", errTwoFactorServiceUserNotFound
	}

	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      s.cfg.Issuer,
		AccountName: user.Username,
	})

	if err != nil {
		return "", err
	}

	twoFactorSettings := &core.TwoFactorSettings{
		UserID:  user.ID,
		Secret:  key.Secret(),
		Enabled: false,
	}

	if err := s.twoFactorSettingsRepository.Save(ctx, twoFactorSettings); err != nil {
		return "", err
	}

	return key.URL(), nil
}

func (s *twoFactorAuthService) Deactivate(ctx context.Context, userID int64, code string) error {
	settings, err := s.twoFactorSettingsRepository.Find(ctx, userID)
	if err != nil {
		return err
	}

	if settings == nil {
		return core.ErrTwoFactorAuthNotEnabled
	}

	if !totp.Validate(code, settings.Secret) {
		return core.ErrTwoFactorCodeInvalid
	}

	if err := s.twoFactorSettingsRepository.Delete(ctx, settings); err != nil {
		return err
	}

	return nil
}

func (s *twoFactorAuthService) Verify(ctx context.Context, userID int64, code *string) error {
	settings, err := s.twoFactorSettingsRepository.Find(ctx, userID)
	if err != nil {
		return err
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
		return err
	}

	if settings == nil {
		return core.ErrTwoFactorAuthSetupNotRunning
	}

	if settings.Enabled {
		return core.ErrTwoFactorAuthAlreadyEnabled
	}

	if !totp.Validate(code, settings.Secret) {
		return core.ErrTwoFactorCodeInvalid
	}

	settings.Enabled = true

	return s.twoFactorSettingsRepository.Update(ctx, settings)
}

func (s *twoFactorAuthService) GetSettings(
	ctx context.Context,
	userID int64,
) (*core.TwoFactorSettings, error) {
	settings, err := s.twoFactorSettingsRepository.Find(ctx, userID)
	if err != nil {
		return nil, err
	}

	if settings == nil {
		return &core.TwoFactorSettings{
			Enabled: false,
		}, nil
	}

	return settings, nil
}
