package core

import (
	"context"
	"errors"
)

var (
	ErrExpiredTokenLifetime         = errors.New("expired token lifetime")
	ErrExpiredInactiveTokenLifetime = errors.New("expired inactive token lifetime")
	ErrTokenInvalid                 = errors.New("token invalid")

	ErrInvalidUsernameOrPassword = errors.New("invalid username or password")
)

type (
	InvalidOrExpiredTokenError struct {
		Reason error
	}

	// AuthContext represent the authenticated context for a user and a session.
	AuthContext struct {
		UserID    int64
		SessionID int64
	}

	// SessionManager defines operations to manage a session of a user.
	SessionManager interface {
		// Validate validates a session and updates the last usage and rotates the token if needed.
		// The rotated token is returned as second return value when the token was rotated.
		Validate(ctx context.Context, token string) (*AuthContext, *string, error)
		// Create creates a new session for a user with a client ip and a user agent.
		Create(ctx context.Context, userId int64, clientIp string, userAgent string) (string, error)
		// Revoke revokes a specific session.
		Revoke(ctx context.Context, sessionID int64, userID int64) error
	}

	// PasswordAuthenticator defines operations to log in users via login and password.
	PasswordAuthenticator interface {
		Login(ctx context.Context, login string, username string) (*User, error)
	}
)

func (InvalidOrExpiredTokenError) Error() string {
	return "invalid or expired token"
}
