package auth

import (
	"context"

	"github.com/davidborzek/tvhgo/core"
	"github.com/rs/zerolog/log"
)

func NewLocalPasswordAuthenticator(
	userRepository core.UserRepository,
	twoFactorService core.TwoFactorAuthService,
) *localPasswordAuthenticator {
	return &localPasswordAuthenticator{
		userRepository:   userRepository,
		twoFactorService: twoFactorService,
	}
}

type localPasswordAuthenticator struct {
	userRepository   core.UserRepository
	twoFactorService core.TwoFactorAuthService
}

func (s *localPasswordAuthenticator) Login(
	ctx context.Context,
	login string,
	password string,
	totp *string,
) (*core.User, error) {
	user, err := s.userRepository.FindByUsername(ctx, login)
	if err != nil {
		log.Error().Str("user", login).
			Err(err).Msg("failed to get user")

		return nil, core.ErrUnexpectedError
	}

	if user == nil {
		return nil, core.ErrInvalidUsernameOrPassword
	}

	if user.PasswordHash == "" {
		return nil, core.ErrInvalidUsernameOrPassword
	}

	if err := ComparePassword(password, user.PasswordHash); err != nil {
		return nil, core.ErrInvalidUsernameOrPassword
	}

	if err := s.twoFactorService.Verify(ctx, user.ID, totp); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *localPasswordAuthenticator) ConfirmPassword(
	ctx context.Context,
	userID int64,
	password string,
) error {
	user, err := s.userRepository.FindById(ctx, userID)
	if err != nil {
		log.Error().Int64("userId", userID).
			Err(err).Msg("failed to find user for password confirmation")

		return core.ErrUnexpectedError
	}

	if err := ComparePassword(password, user.PasswordHash); err != nil {
		return core.ErrConfirmationPasswordInvalid
	}

	return nil
}
