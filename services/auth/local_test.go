package auth_test

import (
	"context"
	"errors"
	"testing"

	"github.com/davidborzek/tvhgo/core"
	mock_core "github.com/davidborzek/tvhgo/mock/core"
	"github.com/davidborzek/tvhgo/services/auth"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const (
	username     = "someUsername"
	password     = "somePassword"
	passwordHash = "$2a$12$3jHVVMgg/0R4oQOXtQWnzeVtKUw5NNrf9.aTToK6zQ9bd/sM1uV3q"
)

var (
	ctx = context.TODO()

	expectedUser = core.User{
		ID:           1234,
		Username:     username,
		PasswordHash: passwordHash,
	}
)

func TestLocalPasswordAuthenticatorLoginReturnsUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock_core.NewMockUserRepository(ctrl)
	mockRepository.EXPECT().
		FindByUsername(ctx, username).
		Return(&expectedUser, nil).
		Times(1)

	authenticator := auth.NewLocalPasswordAuthenticator(mockRepository)

	user, err := authenticator.Login(ctx, username, password)

	assert.Nil(t, err)
	assert.Equal(t, expectedUser, *user)
}

func TestLocalPasswordAuthenticatorLoginReturnsErrInvalidUsernameOrPasswordWhenUserWasNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock_core.NewMockUserRepository(ctrl)
	mockRepository.EXPECT().
		FindByUsername(ctx, username).
		Return(nil, nil).
		Times(1)

	authenticator := auth.NewLocalPasswordAuthenticator(mockRepository)

	user, err := authenticator.Login(ctx, username, password)

	assert.Nil(t, user)
	assert.Equal(t, core.ErrInvalidUsernameOrPassword, err)
}

func TestLocalPasswordAuthenticatorLoginReturnsErrUnexpectedError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock_core.NewMockUserRepository(ctrl)
	mockRepository.EXPECT().
		FindByUsername(ctx, username).
		Return(nil, errors.New("some unexpected error")).
		Times(1)

	authenticator := auth.NewLocalPasswordAuthenticator(mockRepository)

	user, err := authenticator.Login(ctx, username, password)

	assert.Nil(t, user)
	assert.Equal(t, core.ErrUnexpectedError, err)
}

func TestLocalPasswordAuthenticatorLoginReturnsErrInvalidUsernameOrPasswordForInvalidPassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock_core.NewMockUserRepository(ctrl)
	mockRepository.EXPECT().
		FindByUsername(ctx, username).
		Return(&expectedUser, nil).
		Times(1)

	authenticator := auth.NewLocalPasswordAuthenticator(mockRepository)

	user, err := authenticator.Login(ctx, username, "invalid password")

	assert.Nil(t, user)
	assert.Equal(t, core.ErrInvalidUsernameOrPassword, err)
}
