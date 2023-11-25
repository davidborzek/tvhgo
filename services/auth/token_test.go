package auth_test

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/davidborzek/tvhgo/core"
	mock_core "github.com/davidborzek/tvhgo/mock/core"
	"github.com/davidborzek/tvhgo/services/auth"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

const (
	tokenID   = 1
	tokenName = "My Token"
	tokenVal  = "b29f4163d2427ee4c91ed5c993576948666b1832960fbdb57d86965238619716"
	tokenHash = "256475849de6aecfb0434623bde0b6f16a89722ffd86994b6094101ac044762b"
)

var (
	expectedCreatedToken = core.Token{
		UserID: userID,
		Name:   tokenName,
	}
)

func TestTokenServiceCreateReturnsToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock_core.NewMockTokenRepository(ctrl)
	mockRepository.EXPECT().
		Create(ctx, newEqTokenMatcher(&expectedCreatedToken)).
		Return(nil).
		Times(1)

	tokenService := auth.NewTokenService(mockRepository)

	token, err := tokenService.Create(ctx, userID, tokenName)

	assert.Nil(t, err)
	assert.NotEmpty(t, token)
}

func TestTokenServiceCreateErrUnexpectedErrorWhenPersistingFails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock_core.NewMockTokenRepository(ctrl)
	mockRepository.EXPECT().
		Create(ctx, newEqTokenMatcher(&expectedCreatedToken)).
		Return(errors.New("some unexpected error")).
		Times(1)

	tokenService := auth.NewTokenService(mockRepository)

	token, err := tokenService.Create(ctx, userID, tokenName)

	assert.Equal(t, core.ErrUnexpectedError, err)
	assert.Empty(t, token)
}

func TestTokenServiceRevokeSucceeds(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock_core.NewMockTokenRepository(ctrl)
	mockRepository.EXPECT().
		Delete(ctx, &core.Token{ID: tokenID}).
		Return(nil).
		Times(1)

	tokenService := auth.NewTokenService(mockRepository)

	err := tokenService.Revoke(ctx, tokenID)

	assert.Nil(t, err)
}

func TestTokenServiceRevokeFails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock_core.NewMockTokenRepository(ctrl)
	mockRepository.EXPECT().
		Delete(ctx, &core.Token{ID: tokenID}).
		Return(errors.New("some unexpected error")).
		Times(1)

	tokenService := auth.NewTokenService(mockRepository)

	err := tokenService.Revoke(ctx, tokenID)
	assert.Equal(t, core.ErrUnexpectedError, err)
}

func TestTokenServiceValidateReturnsAuthContext(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tokenEntity := &core.Token{
		ID:          tokenID,
		UserID:      userID,
		Name:        tokenName,
		HashedToken: tokenHash,
	}

	mockRepository := mock_core.NewMockTokenRepository(ctrl)
	mockRepository.EXPECT().
		FindByToken(ctx, tokenHash).
		Return(tokenEntity, nil).
		Times(1)

	tokenService := auth.NewTokenService(mockRepository)

	authCtx, err := tokenService.Validate(ctx, tokenVal)

	assert.Equal(t, authCtx.UserID, userID)
	assert.Nil(t, authCtx.SessionID)
	assert.Nil(t, err)
}

func TestTokenServiceValidateReturnsErrUnexpectedError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock_core.NewMockTokenRepository(ctrl)
	mockRepository.EXPECT().
		FindByToken(ctx, tokenHash).
		Return(nil, errors.New("some unexpected error")).
		Times(1)

	tokenService := auth.NewTokenService(mockRepository)

	authCtx, err := tokenService.Validate(ctx, tokenVal)

	assert.Nil(t, authCtx)
	assert.Equal(t, err, core.ErrUnexpectedError)
}

func TestTokenServiceValidateReturnsInvalidOrExpiredTokenError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock_core.NewMockTokenRepository(ctrl)
	mockRepository.EXPECT().
		FindByToken(ctx, tokenHash).
		Return(nil, nil).
		Times(1)

	tokenService := auth.NewTokenService(mockRepository)

	authCtx, err := tokenService.Validate(ctx, tokenVal)

	assert.Nil(t, authCtx)
	assert.Equal(t, err, core.InvalidOrExpiredTokenError{
		Reason: core.ErrTokenInvalid,
	})
}

type eqTokenMatcher struct {
	expected *core.Token
}

func (e eqTokenMatcher) Matches(x interface{}) bool {
	actual, ok := x.(*core.Token)
	if !ok {
		return false
	}

	e.expected.HashedToken = actual.HashedToken
	return reflect.DeepEqual(e.expected, actual)
}

func (e eqTokenMatcher) String() string {
	return fmt.Sprintf("is equal to %v", e.expected)
}

func newEqTokenMatcher(expected *core.Token) gomock.Matcher {
	return eqTokenMatcher{expected}
}
