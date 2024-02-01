package auth

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/davidborzek/tvhgo/core"
	"github.com/rs/zerolog/log"
)

type tokenService struct {
	tokenRepository core.TokenRepository
}

func NewTokenService(tokenRepository core.TokenRepository) core.TokenService {
	return &tokenService{
		tokenRepository: tokenRepository,
	}
}

func (s *tokenService) Create(ctx context.Context, userID int64, name string) (string, error) {
	token, err := generateToken()
	if err != nil {
		log.Error().Err(err).Int64("user", userID).
			Msg("could not generate token")

		return "", core.ErrUnexpectedError
	}

	hashedToken := hashToken(token)
	tokenEntity := &core.Token{
		Name:        name,
		UserID:      userID,
		HashedToken: hashedToken,
	}

	if err := s.tokenRepository.Create(ctx, tokenEntity); err != nil {
		log.Error().Err(err).Int64("user", userID).
			Msg("could not persist token")

		return "", core.ErrUnexpectedError
	}

	return token, nil
}

func (s *tokenService) Revoke(ctx context.Context, id int64) error {
	err := s.tokenRepository.Delete(ctx, &core.Token{ID: id})
	if err != nil {
		log.Error().Err(err).Int64("token", id).
			Msg("could not revoke token")

		return core.ErrUnexpectedError
	}
	return nil
}

func (s *tokenService) Validate(
	ctx context.Context,
	token string,
) (*core.AuthContext, error) {
	hashedToken := hashToken(token)
	tokenEntity, err := s.tokenRepository.FindByToken(ctx, hashedToken)
	if err != nil {
		log.Error().Err(err).
			Msg("could not get token")

		return nil, core.ErrUnexpectedError
	}

	if tokenEntity == nil {
		return nil, core.InvalidOrExpiredTokenError{
			Reason: core.ErrTokenInvalid,
		}
	}

	authCtx := core.AuthContext{
		UserID: tokenEntity.UserID,
	}

	return &authCtx, nil

}

// Generates a random 256-bit token.
func generateToken() (string, error) {
	randBytes := make([]byte, 32)
	_, err := rand.Read(randBytes)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", randBytes), nil
}

// Hashes a token with SHA256.
func hashToken(token string) string {
	hash := sha256.Sum256([]byte(token))
	return hex.EncodeToString(hash[:])
}
