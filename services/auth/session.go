package auth

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/davidborzek/tvhgo/core"
	log "github.com/sirupsen/logrus"
)

type sessionManager struct {
	inactiveLifetime      time.Duration
	lifetime              time.Duration
	tokenRotationInterval time.Duration

	sessionRepository core.SessionRepository
}

func NewSessionManager(
	sessionRepository core.SessionRepository,
	inactiveLifetime time.Duration,
	lifetime time.Duration,
	tokenRotationInterval time.Duration,
) core.SessionManager {
	return &sessionManager{
		sessionRepository:     sessionRepository,
		inactiveLifetime:      inactiveLifetime,
		lifetime:              lifetime,
		tokenRotationInterval: tokenRotationInterval,
	}
}

func (s *sessionManager) Create(
	ctx context.Context,
	userId int64,
	clientIp string,
	userAgent string,
) (string, error) {
	token, err := generateToken()
	if err != nil {
		log.WithError(err).
			Error("could not generate session token")

		return "", core.ErrUnexpectedError
	}

	hashedToken := hashToken(token)
	session := &core.Session{
		UserId:      userId,
		HashedToken: hashedToken,
		ClientIP:    clientIp,
		UserAgent:   userAgent,
	}

	if err := s.sessionRepository.Create(ctx, session); err != nil {
		log.WithError(err).
			WithField("user", userId).
			Error("could not persist session")

		return "", core.ErrUnexpectedError
	}

	return token, nil
}

func (s *sessionManager) Revoke(ctx context.Context, sessionId int64) error {
	err := s.sessionRepository.Delete(ctx, sessionId)
	if err != nil {
		log.WithError(err).
			WithField("session", sessionId).
			Error("could not delete session")

		return core.ErrUnexpectedError
	}
	return nil
}

func (s *sessionManager) Validate(ctx context.Context, token string) (*core.AuthContext, *string, error) {
	hashedToken := hashToken(token)

	session, err := s.sessionRepository.Find(ctx, hashedToken)
	if err != nil {
		log.WithError(err).
			WithField("session", session.ID).
			Error("could not get session")

		return nil, nil, core.ErrUnexpectedError
	}

	if session == nil {
		return nil, nil, core.ErrInvalidOrExpiredToken
	}

	if isExpired(session.CreatedAt, s.lifetime) || isExpired(session.LastUsedAt, s.inactiveLifetime) {
		return nil, nil, core.ErrInvalidOrExpiredToken
	}

	rotatedToken, err := s.rotateToken(session)
	if err != nil {
		log.WithError(err).
			Error("could not rotate token")
		return nil, nil, err
	}

	session.LastUsedAt = time.Now().Unix()

	if err := s.sessionRepository.Update(ctx, session); err != nil {
		log.WithError(err).
			WithField("session", session.ID).
			Error("could not update session")

		return nil, nil, core.ErrUnexpectedError
	}

	authCtx := core.AuthContext{
		SessionID: session.ID,
		UserID:    session.UserId,
	}

	return &authCtx, rotatedToken, nil
}

// rotateToken rotates the token if necessary.
func (s *sessionManager) rotateToken(session *core.Session) (rotatedToken *string, err error) {
	if isExpired(session.RotatedAt, s.tokenRotationInterval) {
		token, err := generateToken()
		if err != nil {
			return nil, core.ErrUnexpectedError
		}

		session.HashedToken = hashToken(token)
		session.RotatedAt = time.Now().Unix()

		rotatedToken = &token
	}

	return
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

// isExpired checks if a creation date has expired for a given lifetime.
func isExpired(creation int64, lifetime time.Duration) bool {
	return time.Unix(creation, 0).Add(lifetime).Before(time.Now())
}
