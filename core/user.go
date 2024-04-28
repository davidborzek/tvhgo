package core

import (
	"context"
	"errors"
)

var (
	ErrUsernameAlreadyExists = errors.New("username already exists")
	ErrEmailAlreadyExists    = errors.New("email already exists")
)

type (
	// User represents a user.
	User struct {
		ID       int64  `json:"id"`
		Username string `json:"username"`
		// PasswordHash hash of the users password
		PasswordHash string `json:"-"`
		Email        string `json:"email"`
		DisplayName  string `json:"displayName"`
		CreatedAt    int64  `json:"createdAt"`
		UpdatedAt    int64  `json:"updatedAt"`
	}

	UserListResult ListResult[*User]

	// UserQueryParams defines user query parameters.
	UserQueryParams struct {
		// (Optional) Limit the result.
		Limit int64
		// (Optional) Offset the result.
		// Can only be used together with Limit.
		Offset int64
	}

	UserRepository interface {
		// FindById returns a user by id.
		FindById(ctx context.Context, id int64) (*User, error)

		// FindByUsername returns a user by username.
		FindByUsername(ctx context.Context, user string) (*User, error)

		// Find returns a list of users paginated by UserQueryParams.
		Find(ctx context.Context, params UserQueryParams) (*UserListResult, error)

		// Create persists a new user.
		Create(ctx context.Context, user *User) error

		// Update persists an updated user.
		Update(ctx context.Context, user *User) error

		// Delete deletes a user.
		Delete(ctx context.Context, user *User) error
	}
)
