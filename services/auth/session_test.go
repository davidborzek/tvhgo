package auth_test

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/davidborzek/tvhgo/core"
	mock_core "github.com/davidborzek/tvhgo/mock/core"
	"github.com/davidborzek/tvhgo/services/auth"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type eqSessionMatcher struct {
	expected *core.Session
}

func (e eqSessionMatcher) Matches(x interface{}) bool {
	actual, ok := x.(*core.Session)
	if !ok {
		return false
	}

	e.expected.HashedToken = actual.HashedToken
	return reflect.DeepEqual(e.expected, actual)
}

func (e eqSessionMatcher) String() string {
	return fmt.Sprintf("is equal to %v", e.expected)
}

func newEqSessionMatcher(expected *core.Session) gomock.Matcher {
	return eqSessionMatcher{expected}
}

const (
	userID    int64 = 1234
	clientIp        = "127.0.0.1"
	userAgent       = "Mozilla/5.0"

	sessionID int64 = 12345

	token = "someToken"
)

var (
	expectedCreatedSession = core.Session{
		UserId:    userID,
		ClientIP:  clientIp,
		UserAgent: userAgent,
	}
)

func TestSessionManagerCreateReturnsToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock_core.NewMockSessionRepository(ctrl)
	mockRepository.EXPECT().
		Create(ctx, newEqSessionMatcher(&expectedCreatedSession)).
		Return(nil).
		Times(1)

	sessionManager := auth.NewSessionManager(mockRepository, 0, 0, 0)
	token, err := sessionManager.Create(ctx, userID, clientIp, userAgent)

	assert.Nil(t, err)
	assert.NotEmpty(t, token)
}

func TestSessionManagerCreateReturnsErrUnexpectedErrWhenPersistingSessionFails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock_core.NewMockSessionRepository(ctrl)
	mockRepository.EXPECT().
		Create(ctx, newEqSessionMatcher(&expectedCreatedSession)).
		Return(errors.New("some unexpected error")).
		Times(1)

	sessionManager := auth.NewSessionManager(mockRepository, 0, 0, 0)
	token, err := sessionManager.Create(ctx, userID, clientIp, userAgent)

	assert.Equal(t, core.ErrUnexpectedError, err)
	assert.Empty(t, token)
}

func TestSessionManagerRevokeSucceeds(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock_core.NewMockSessionRepository(ctrl)
	mockRepository.EXPECT().
		Delete(ctx, sessionID).
		Return(nil).
		Times(1)

	sessionManager := auth.NewSessionManager(mockRepository, 0, 0, 0)
	err := sessionManager.Revoke(ctx, sessionID)

	assert.Nil(t, err)
}

func TestSessionManagerRevokeReturnsErrUnexpectedError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock_core.NewMockSessionRepository(ctrl)
	mockRepository.EXPECT().
		Delete(ctx, sessionID).
		Return(errors.New("some unexpected error")).
		Times(1)

	sessionManager := auth.NewSessionManager(mockRepository, 0, 0, 0)
	err := sessionManager.Revoke(ctx, sessionID)

	assert.Equal(t, core.ErrUnexpectedError, err)
}

func TestSessionManagerValidateReturnsErrUnexpectedErrorWhenFindingSessionFails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock_core.NewMockSessionRepository(ctrl)
	mockRepository.EXPECT().
		Find(ctx, gomock.Any()).
		Return(nil, errors.New("some unexpected error")).
		Times(1)

	sessionManager := auth.NewSessionManager(mockRepository, 0, 0, 0)
	authCtx, maybeRotatedToken, err := sessionManager.Validate(ctx, token)

	assert.Nil(t, authCtx)
	assert.Nil(t, maybeRotatedToken)
	assert.Equal(t, core.ErrUnexpectedError, err)
}

func TestSessionManagerValidateReturnsErrInvalidOrExpiredTokenWhenSessionWasNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock_core.NewMockSessionRepository(ctrl)
	mockRepository.EXPECT().
		Find(ctx, gomock.Any()).
		Return(nil, nil).
		Times(1)

	sessionManager := auth.NewSessionManager(mockRepository, 0, 0, 0)
	authCtx, maybeRotatedToken, err := sessionManager.Validate(ctx, token)

	assert.Nil(t, authCtx)
	assert.Nil(t, maybeRotatedToken)
	assert.Equal(t, core.ErrInvalidOrExpiredToken, err)
}

func TestSessionManagerValidateReturnsErrInvalidOrExpiredTokenWhenSessionLifetimeIsExpired(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	session := &core.Session{
		ID:          sessionID,
		UserId:      userID,
		ClientIP:    clientIp,
		UserAgent:   userAgent,
		HashedToken: "someHash",
		CreatedAt:   1,
		LastUsedAt:  1,
		RotatedAt:   1,
	}

	mockRepository := mock_core.NewMockSessionRepository(ctrl)
	mockRepository.EXPECT().
		Find(ctx, gomock.Any()).
		Return(session, nil).
		Times(1)

	sessionManager := auth.NewSessionManager(mockRepository, 1, 1, 1)
	authCtx, maybeRotatedToken, err := sessionManager.Validate(ctx, token)

	assert.Nil(t, authCtx)
	assert.Nil(t, maybeRotatedToken)
	assert.Equal(t, core.ErrInvalidOrExpiredToken, err)
}

func TestSessionManagerValidateReturnsErrInvalidOrExpiredTokenWhenSessionInactiveLifetimeIsExpired(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	createdAt := time.Now().Unix()
	lastUsedAt := time.Now().Add(-2 * time.Hour).Unix()

	session := &core.Session{
		ID:          sessionID,
		UserId:      userID,
		ClientIP:    clientIp,
		UserAgent:   userAgent,
		HashedToken: "someHash",
		CreatedAt:   createdAt,
		LastUsedAt:  lastUsedAt,
		RotatedAt:   1,
	}

	mockRepository := mock_core.NewMockSessionRepository(ctrl)
	mockRepository.EXPECT().
		Find(ctx, gomock.Any()).
		Return(session, nil).
		Times(1)

	sessionManager := auth.NewSessionManager(mockRepository, time.Hour, time.Hour, 1)
	authCtx, maybeRotatedToken, err := sessionManager.Validate(ctx, token)

	assert.Nil(t, authCtx)
	assert.Nil(t, maybeRotatedToken)
	assert.Equal(t, core.ErrInvalidOrExpiredToken, err)
}

func TestSessionManagerValidateReturnsAuthContextAndDoesNotRotateToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	now := time.Now().Unix()

	session := &core.Session{
		ID:          sessionID,
		UserId:      userID,
		ClientIP:    clientIp,
		UserAgent:   userAgent,
		HashedToken: "someHash",
		CreatedAt:   now,
		LastUsedAt:  now,
		RotatedAt:   now,
	}

	mockRepository := mock_core.NewMockSessionRepository(ctrl)
	mockRepository.EXPECT().
		Find(ctx, gomock.Any()).
		Return(session, nil).
		Times(1)

	mockRepository.EXPECT().Update(ctx, newEqSessionMatcher(session)).Return(nil).Times(1)

	sessionManager := auth.NewSessionManager(mockRepository, time.Hour, time.Hour, time.Hour)
	authCtx, maybeRotatedToken, err := sessionManager.Validate(ctx, token)

	assert.Equal(t, sessionID, authCtx.SessionID)
	assert.Equal(t, userID, authCtx.UserID)
	assert.Nil(t, maybeRotatedToken)
	assert.Nil(t, err)
}

func TestSessionManagerValidateReturnsAuthContextAndRotatesToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	now := time.Now().Unix()

	session := &core.Session{
		ID:          sessionID,
		UserId:      userID,
		ClientIP:    clientIp,
		UserAgent:   userAgent,
		HashedToken: "someHash",
		CreatedAt:   now,
		LastUsedAt:  now,
		RotatedAt:   1000,
	}

	mockRepository := mock_core.NewMockSessionRepository(ctrl)
	mockRepository.EXPECT().
		Find(ctx, gomock.Any()).
		Return(session, nil).
		Times(1)

	mockRepository.EXPECT().Update(ctx, newEqSessionMatcher(session)).Return(nil).Times(1)

	sessionManager := auth.NewSessionManager(mockRepository, time.Hour, time.Hour, time.Hour)
	authCtx, maybeRotatedToken, err := sessionManager.Validate(ctx, token)

	assert.Equal(t, sessionID, authCtx.SessionID)
	assert.Equal(t, userID, authCtx.UserID)
	assert.NotNil(t, maybeRotatedToken)
	assert.Nil(t, err)
}

func TestSessionManagerValidateReturnsErrUnexpectedErrorWhenUpdaingSessionFails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	now := time.Now().Unix()

	session := &core.Session{
		ID:          sessionID,
		UserId:      userID,
		ClientIP:    clientIp,
		UserAgent:   userAgent,
		HashedToken: "someHash",
		CreatedAt:   now,
		LastUsedAt:  now,
		RotatedAt:   now,
	}

	mockRepository := mock_core.NewMockSessionRepository(ctrl)
	mockRepository.EXPECT().
		Find(ctx, gomock.Any()).
		Return(session, nil).
		Times(1)

	mockRepository.EXPECT().
		Update(ctx, newEqSessionMatcher(session)).
		Return(errors.New("some unexpected error")).
		Times(1)

	sessionManager := auth.NewSessionManager(mockRepository, time.Hour, time.Hour, time.Hour)
	authCtx, maybeRotatedToken, err := sessionManager.Validate(ctx, token)

	assert.Nil(t, authCtx)
	assert.Nil(t, maybeRotatedToken)
	assert.Equal(t, core.ErrUnexpectedError, err)
}
