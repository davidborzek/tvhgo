package auth

import (
	"context"
	"fmt"

	"github.com/davidborzek/tvhgo/core"
	"github.com/pquerna/otp/totp"
)

const (
	twoFactorIssuer = "tvhgo"
)

type twoFactorService struct {
	userRepository              core.UserRepository
	twoFactorSettingsRepository core.TwoFactorSettingsRepository
}

func NewTwoFactorService(
	twoFactorSettingsRepository core.TwoFactorSettingsRepository,
	userRepository core.UserRepository,
) core.TwoFactorService {
	return &twoFactorService{
		twoFactorSettingsRepository: twoFactorSettingsRepository,
		userRepository:              userRepository,
	}
}

func (s *twoFactorService) Setup(ctx context.Context, userId int64) (string, error) {
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
		UserID: user.ID,
		Secret: key.Secret(),
	}

	if err := s.twoFactorSettingsRepository.Create(ctx, twoFactorSettings); err != nil {
		return "", fmt.Errorf("two factor service failed persist two factor settings: %w", err)
	}

	return key.URL(), nil
}
