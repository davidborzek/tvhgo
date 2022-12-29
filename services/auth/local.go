package auth

import (
	"context"

	"github.com/davidborzek/tvhgo/core"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func NewLocalPasswordAuthenticator(userRepository core.UserRepository) *localPasswordAuthenticator {
	return &localPasswordAuthenticator{
		userRepository: userRepository,
	}
}

type localPasswordAuthenticator struct {
	userRepository core.UserRepository
}

func (s *localPasswordAuthenticator) Login(ctx context.Context, login string, password string) (*core.User, error) {
	user, err := s.userRepository.FindByUsername(ctx, login)
	if err != nil {
		log.WithError(err).
			WithField("user", login).
			Error("cloud not get user")

		return nil, core.ErrUnexpectedError
	}

	if user == nil {
		return nil, core.ErrInvalidUsernameOrPassword
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return nil, core.ErrInvalidUsernameOrPassword
	}

	return user, nil
}
