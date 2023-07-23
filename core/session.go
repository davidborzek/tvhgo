package core

import (
	"context"
)

type (
	// Session defines a representation for a session of a user.
	Session struct {
		ID          int64  `json:"id"`
		UserId      int64  `json:"userId"`
		HashedToken string `json:"-"`
		ClientIP    string `json:"clientIp"`
		UserAgent   string `json:"userAgent"`
		CreatedAt   int64  `json:"createdAt"`
		LastUsedAt  int64  `json:"lastUsedAt"`
		RotatedAt   int64  `json:"-"`
	}

	// SessionRepository defines CRUD operations for working with sessions.
	SessionRepository interface {
		// Find returns a sessions.
		Find(ctx context.Context, hashedToken string) (*Session, error)

		// FindByUser returns a list of sessions for a user.
		FindByUser(ctx context.Context, userID int64) ([]*Session, error)

		// Create persists a new session.
		Create(ctx context.Context, session *Session) error

		// Create persists a updated session.
		Update(ctx context.Context, session *Session) error

		// Delete deletes a session.
		Delete(ctx context.Context, sessionID int64, userID int64) error

		// DeleteExpired deletes all expired sessions.
		DeleteExpired(ctx context.Context, expirationDate int64, inactiveExpirationDate int64) (int64, error)
	}
)
