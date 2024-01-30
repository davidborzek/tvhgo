package auth_test

import (
	"context"
	"errors"
	"testing"

	"github.com/davidborzek/tvhgo/core"
	mock_core "github.com/davidborzek/tvhgo/mock/core"
	"github.com/davidborzek/tvhgo/services/auth"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
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

	totp := "123456"

	mockRepository := mock_core.NewMockUserRepository(ctrl)
	mockRepository.EXPECT().
		FindByUsername(ctx, username).
		Return(&expectedUser, nil).
		Times(1)

	mockTwoFactorAuthService := mock_core.NewMockTwoFactorAuthService(ctrl)
	mockTwoFactorAuthService.EXPECT().
		Verify(ctx, expectedUser.ID, &totp).
		Return(nil).
		Times(1)

	authenticator := auth.NewLocalPasswordAuthenticator(mockRepository, mockTwoFactorAuthService)

	user, err := authenticator.Login(ctx, username, password, &totp)

	assert.Nil(t, err)
	assert.Equal(t, expectedUser, *user)
}

func TestLocalPasswordAuthenticatorLoginReturnsErrWhenTwoFactorVerifyFails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	totp := "123456"

	mockRepository := mock_core.NewMockUserRepository(ctrl)
	mockRepository.EXPECT().
		FindByUsername(ctx, username).
		Return(&expectedUser, nil).
		Times(1)

	mockTwoFactorAuthService := mock_core.NewMockTwoFactorAuthService(ctrl)
	mockTwoFactorAuthService.EXPECT().
		Verify(ctx, expectedUser.ID, &totp).
		Return(errors.New("some error")).
		Times(1)

	authenticator := auth.NewLocalPasswordAuthenticator(mockRepository, mockTwoFactorAuthService)

	user, err := authenticator.Login(ctx, username, password, &totp)

	assert.EqualError(t, err, "some error")
	assert.Nil(t, user)
}

func TestLocalPasswordAuthenticatorLoginReturnsErrInvalidUsernameOrPasswordWhenUserWasNotFound(
	t *testing.T,
) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock_core.NewMockUserRepository(ctrl)
	mockRepository.EXPECT().
		FindByUsername(ctx, username).
		Return(nil, nil).
		Times(1)

	mockTwoFactorAuthService := mock_core.NewMockTwoFactorAuthService(ctrl)
	authenticator := auth.NewLocalPasswordAuthenticator(mockRepository, mockTwoFactorAuthService)

	user, err := authenticator.Login(ctx, username, password, nil)

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

	mockTwoFactorAuthService := mock_core.NewMockTwoFactorAuthService(ctrl)
	authenticator := auth.NewLocalPasswordAuthenticator(mockRepository, mockTwoFactorAuthService)

	user, err := authenticator.Login(ctx, username, password, nil)

	assert.Nil(t, user)
	assert.Equal(t, core.ErrUnexpectedError, err)
}

func TestLocalPasswordAuthenticatorLoginReturnsErrInvalidUsernameOrPasswordForInvalidPassword(
	t *testing.T,
) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock_core.NewMockUserRepository(ctrl)
	mockRepository.EXPECT().
		FindByUsername(ctx, username).
		Return(&expectedUser, nil).
		Times(1)

	mockTwoFactorAuthService := mock_core.NewMockTwoFactorAuthService(ctrl)

	authenticator := auth.NewLocalPasswordAuthenticator(mockRepository, mockTwoFactorAuthService)

	user, err := authenticator.Login(ctx, username, "invalid password", nil)

	assert.Nil(t, user)
	assert.Equal(t, core.ErrInvalidUsernameOrPassword, err)
}

func TestLocalPasswordAuthenticatorConfirmPasswordSucceeds(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock_core.NewMockUserRepository(ctrl)
	mockRepository.EXPECT().
		FindById(ctx, expectedUser.ID).
		Return(&expectedUser, nil).
		Times(1)

	mockTwoFactorAuthService := mock_core.NewMockTwoFactorAuthService(ctrl)

	authenticator := auth.NewLocalPasswordAuthenticator(mockRepository, mockTwoFactorAuthService)

	err := authenticator.ConfirmPassword(ctx, expectedUser.ID, password)

	assert.Nil(t, err)
}

func TestLocalPasswordAuthenticatorConfirmPasswordReturnsErrConfirmationPasswordInvalid(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock_core.NewMockUserRepository(ctrl)
	mockRepository.EXPECT().
		FindById(ctx, expectedUser.ID).
		Return(&expectedUser, nil).
		Times(1)

	mockTwoFactorAuthService := mock_core.NewMockTwoFactorAuthService(ctrl)

	authenticator := auth.NewLocalPasswordAuthenticator(mockRepository, mockTwoFactorAuthService)

	err := authenticator.ConfirmPassword(ctx, expectedUser.ID, "invalid")

	assert.Equal(t, core.ErrConfirmationPasswordInvalid, err)
}

func TestLocalPasswordAuthenticatorConfirmPasswordReturnsErrUnexpectedError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock_core.NewMockUserRepository(ctrl)
	mockRepository.EXPECT().
		FindById(ctx, expectedUser.ID).
		Return(nil, errors.New("some unexpected error")).
		Times(1)

	mockTwoFactorAuthService := mock_core.NewMockTwoFactorAuthService(ctrl)

	authenticator := auth.NewLocalPasswordAuthenticator(mockRepository, mockTwoFactorAuthService)

	err := authenticator.ConfirmPassword(ctx, expectedUser.ID, password)

	assert.Equal(t, core.ErrUnexpectedError, err)
}
