package core

import (
	"context"
	"errors"
)

var (
	ErrExpiredTokenLifetime         = errors.New("expired token lifetime")
	ErrExpiredInactiveTokenLifetime = errors.New("expired inactive token lifetime")
	ErrTokenInvalid                 = errors.New("token invalid")

	ErrInvalidUsernameOrPassword   = errors.New("invalid username or password")
	ErrConfirmationPasswordInvalid = errors.New("confirmation password is invalid")

	ErrTwoFactorRequired    = errors.New("two factor auth is required")
	ErrTwoFactorCodeInvalid = errors.New("invalid two factor code provided")

	ErrTwoFactorAuthAlreadyEnabled  = errors.New("two factor auth is already enabled")
	ErrTwoFactorAuthNotEnabled      = errors.New("two factor auth is not enabled")
	ErrTwoFactorAuthSetupNotRunning = errors.New("two factor auth setup not running")
)

type (
	// InvalidOrExpiredTokenError is returned when a token / session is invalid or expired.
	InvalidOrExpiredTokenError struct {
		// Reason is the reason why the token is invalid or expired.
		Reason error
	}

	// AuthContext represent the authenticated context for a user and a session.
	AuthContext struct {
		UserID int64

		// SessionID is the session id for authorizations via session tokens.
		SessionID *int64
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
		// Login logs in a user via login, password and optional totp code.
		Login(ctx context.Context, login string, username string, totp *string) (*User, error)

		// ConfirmPassword confirms the password of a user.
		ConfirmPassword(ctx context.Context, userID int64, password string) error
	}

	// TwoFactorAuthService defines operations to manage two factor auth for a user.
	TwoFactorAuthService interface {
		// GetSettings returns the current two factor settings for a user.
		GetSettings(ctx context.Context, userId int64) (*TwoFactorSettings, error)

		// Setup starts the setup process for two factor auth for a user.
		Setup(ctx context.Context, userId int64) (string, error)

		// Deactivate deactivates two factor auth for a user.
		Deactivate(ctx context.Context, userId int64, code string) error

		// Activate activates two factor auth for a user.
		Activate(ctx context.Context, userID int64, code string) error

		// Verify verifies a two factor code for a user.
		Verify(ctx context.Context, userId int64, code *string) error
	}

	// TokenService defines operations to manage tokens for a user.
	TokenService interface {
		// Create creates a new token for a user.
		Create(ctx context.Context, userID int64, name string) (string, error)

		// Validate validates a token.
		Validate(ctx context.Context, token string) (*AuthContext, error)

		// Revoke revokes a token.
		Revoke(ctx context.Context, id int64) error
	}
)

func (InvalidOrExpiredTokenError) Error() string {
	return "invalid or expired token"
}
