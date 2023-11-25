package core

import "context"

type (
	Token struct {
		ID          int64  `json:"id"`
		UserID      int64  `json:"-"`
		Name        string `json:"name"`
		HashedToken string `json:"-"`
		CreatedAt   int64  `json:"createdAt"`
		UpdatedAt   int64  `json:"updatedAt"`
	}

	// TokenRepository defines CRUD operations working with Tokens.
	TokenRepository interface {
		// FindByToken returns a Token by a the hashed token.
		FindByToken(ctx context.Context, token string) (*Token, error)

		// FindByUser returns all token for a user.
		FindByUser(ctx context.Context, userID int64) ([]*Token, error)

		// Create persists a new Token.
		Create(ctx context.Context, token *Token) error

		// Delete deletes a Token.
		Delete(ctx context.Context, token *Token) error
	}
)
