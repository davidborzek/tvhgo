package core

import (
	"context"
)

type (
	// Session defines a internal representation for a session of a user.
	Session struct {
		ID          int64
		UserId      int64
		HashedToken string
		ClientIP    string
		UserAgent   string
		CreatedAt   int64
		LastUsedAt  int64
		RotatedAt   int64
	}

	// SessionRepository defines CRUD operations for working with sessions.
	SessionRepository interface {
		// Find returns a sessions.
		Find(ctx context.Context, hashedToken string) (*Session, error)

		// Create persists a new session.
		Create(ctx context.Context, session *Session) error

		// Create persists a updated session.
		Update(ctx context.Context, session *Session) error

		// Delete deletes a session.
		Delete(ctx context.Context, id int64) error
	}
)
