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
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
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

	sessionManager := auth.NewSessionManager(mockRepository, mock_core.NewMockClock(ctrl), 0, 0, 0)
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

	sessionManager := auth.NewSessionManager(mockRepository, mock_core.NewMockClock(ctrl), 0, 0, 0)
	token, err := sessionManager.Create(ctx, userID, clientIp, userAgent)

	assert.Equal(t, core.ErrUnexpectedError, err)
	assert.Empty(t, token)
}

func TestSessionManagerRevokeSucceeds(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock_core.NewMockSessionRepository(ctrl)
	mockRepository.EXPECT().
		Delete(ctx, sessionID, userID).
		Return(nil).
		Times(1)

	sessionManager := auth.NewSessionManager(mockRepository, mock_core.NewMockClock(ctrl), 0, 0, 0)
	err := sessionManager.Revoke(ctx, sessionID, userID)

	assert.Nil(t, err)
}

func TestSessionManagerRevokeReturnsErrUnexpectedError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock_core.NewMockSessionRepository(ctrl)
	mockRepository.EXPECT().
		Delete(ctx, sessionID, userID).
		Return(errors.New("some unexpected error")).
		Times(1)

	sessionManager := auth.NewSessionManager(mockRepository, mock_core.NewMockClock(ctrl), 0, 0, 0)
	err := sessionManager.Revoke(ctx, sessionID, userID)

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

	sessionManager := auth.NewSessionManager(mockRepository, mock_core.NewMockClock(ctrl), 0, 0, 0)
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

	sessionManager := auth.NewSessionManager(mockRepository, mock_core.NewMockClock(ctrl), 0, 0, 0)
	authCtx, maybeRotatedToken, err := sessionManager.Validate(ctx, token)

	assert.Nil(t, authCtx)
	assert.Nil(t, maybeRotatedToken)

	assert.ErrorIs(t, err, core.InvalidOrExpiredTokenError{
		Reason: core.ErrTokenInvalid,
	})
}

func TestSessionManagerValidateReturnsErrInvalidOrExpiredTokenWhenSessionLifetimeIsExpired(
	t *testing.T,
) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Sun Jan 01 2023 00:00:00 GMT+0000
	now := time.Unix(1672527600, 0)

	// Thu Dec 01 2022 00:00:00 GMT+0000
	createdAt := int64(1669849200)

	// 7 Days
	lifetime := 7 * 24 * time.Hour

	session := &core.Session{
		ID:          sessionID,
		UserId:      userID,
		ClientIP:    clientIp,
		UserAgent:   userAgent,
		HashedToken: "someHash",
		CreatedAt:   createdAt,
		LastUsedAt:  0,
		RotatedAt:   0,
	}

	mockRepository := mock_core.NewMockSessionRepository(ctrl)
	mockRepository.EXPECT().
		Find(ctx, gomock.Any()).
		Return(session, nil).
		Times(1)

	mockClock := mock_core.NewMockClock(ctrl)
	mockClock.EXPECT().
		Now().
		Return(now)

	sessionManager := auth.NewSessionManager(
		mockRepository,
		mockClock,
		1,
		lifetime,
		1,
	)

	authCtx, maybeRotatedToken, err := sessionManager.Validate(ctx, token)

	assert.Nil(t, authCtx)
	assert.Nil(t, maybeRotatedToken)

	assert.ErrorIs(t, err, core.InvalidOrExpiredTokenError{
		Reason: core.ErrExpiredTokenLifetime,
	})
}

func TestSessionManagerValidateReturnsErrInvalidOrExpiredTokenWhenSessionInactiveLifetimeIsExpired(
	t *testing.T,
) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Sun Jan 01 2023 00:00:00 GMT+0000
	now := time.Unix(1672527600, 0)

	// Tue Dec 06 2022 00:00:00 GMT+0000
	createdAt := int64(1670284800)

	// 30 Days
	lifetime := 30 * 24 * time.Hour

	// Tue Dec 20 2022 00:00:00 GMT+0000
	lastUsedAt := int64(1671494400)

	// 7 days
	inactiveLifetime := 7 * 24 * time.Hour

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

	mockClock := mock_core.NewMockClock(ctrl)
	mockClock.EXPECT().
		Now().
		Return(now).
		AnyTimes()

	sessionManager := auth.NewSessionManager(
		mockRepository,
		mockClock,
		inactiveLifetime,
		lifetime,
		1,
	)

	authCtx, maybeRotatedToken, err := sessionManager.Validate(ctx, token)

	assert.Nil(t, authCtx)
	assert.Nil(t, maybeRotatedToken)

	assert.ErrorIs(t, err, core.InvalidOrExpiredTokenError{
		Reason: core.ErrExpiredInactiveTokenLifetime,
	})
}

func TestSessionManagerValidateReturnsAuthContext(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Sun Jan 01 2023 00:00:00 GMT+0000
	now := time.Unix(1672527600, 0)

	// Tue Dec 06 2022 00:00:00 GMT+0000
	createdAt := int64(1670284800)

	// 30 Days
	lifetime := 30 * 24 * time.Hour

	// Sat Dec 31 2022 00:00:00 GMT+0000
	lastUsedAt := int64(1672444800)

	// 7 days
	inactiveLifetime := 7 * 24 * time.Hour

	rotatedAt := now.Unix()
	tokenRotationInterval := 1 * time.Hour

	session := &core.Session{
		ID:          sessionID,
		UserId:      userID,
		ClientIP:    clientIp,
		UserAgent:   userAgent,
		HashedToken: "someHash",
		CreatedAt:   createdAt,
		LastUsedAt:  lastUsedAt,
		RotatedAt:   rotatedAt,
	}

	mockRepository := mock_core.NewMockSessionRepository(ctrl)
	mockRepository.EXPECT().
		Find(ctx, gomock.Any()).
		Return(session, nil).
		Times(1)

	mockRepository.EXPECT().
		Update(ctx, newEqSessionMatcher(session)).
		Return(nil).Times(1)

	mockClock := mock_core.NewMockClock(ctrl)
	mockClock.EXPECT().
		Now().
		Return(now).
		AnyTimes()

	sessionManager := auth.NewSessionManager(
		mockRepository,
		mockClock,
		inactiveLifetime,
		lifetime,
		tokenRotationInterval,
	)

	authCtx, maybeRotatedToken, err := sessionManager.Validate(ctx, token)

	assert.Equal(t, sessionID, *authCtx.SessionID)
	assert.Equal(t, userID, authCtx.UserID)
	assert.Nil(t, maybeRotatedToken)
	assert.Nil(t, err)
}

func TestSessionManagerValidateReturnsAuthContextAndRotatesToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Sun Jan 01 2023 00:00:00 GMT+0000
	now := time.Unix(1672527600, 0)

	// Sat Dec 31 2022 00:00:00 GMT+0000
	rotatedAt := int64(1672444800)

	tokenRotationInterval := 1 * time.Hour

	session := &core.Session{
		ID:          sessionID,
		UserId:      userID,
		ClientIP:    clientIp,
		UserAgent:   userAgent,
		HashedToken: "someHash",
		CreatedAt:   now.Unix(),
		LastUsedAt:  now.Unix(),
		RotatedAt:   rotatedAt,
	}

	mockRepository := mock_core.NewMockSessionRepository(ctrl)
	mockRepository.EXPECT().
		Find(ctx, gomock.Any()).
		Return(session, nil).
		Times(1)

	mockRepository.EXPECT().Update(ctx, newEqSessionMatcher(session)).Return(nil).Times(1)

	mockClock := mock_core.NewMockClock(ctrl)
	mockClock.EXPECT().
		Now().
		Return(now).
		AnyTimes()

	sessionManager := auth.NewSessionManager(
		mockRepository,
		mockClock,
		time.Hour,
		time.Hour,
		tokenRotationInterval,
	)

	authCtx, maybeRotatedToken, err := sessionManager.Validate(ctx, token)

	assert.Equal(t, sessionID, *authCtx.SessionID)
	assert.Equal(t, userID, authCtx.UserID)
	assert.NotNil(t, maybeRotatedToken)
	assert.Nil(t, err)
}

func TestSessionManagerValidateReturnsErrUnexpectedErrorWhenUpdaingSessionFails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	now := time.Now()

	session := &core.Session{
		ID:          sessionID,
		UserId:      userID,
		ClientIP:    clientIp,
		UserAgent:   userAgent,
		HashedToken: "someHash",
		CreatedAt:   now.Unix(),
		LastUsedAt:  now.Unix(),
		RotatedAt:   now.Unix(),
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

	mockClock := mock_core.NewMockClock(ctrl)
	mockClock.EXPECT().
		Now().
		Return(now).
		AnyTimes()

	sessionManager := auth.NewSessionManager(
		mockRepository,
		mockClock,
		time.Hour,
		time.Hour,
		time.Hour,
	)
	authCtx, maybeRotatedToken, err := sessionManager.Validate(ctx, token)

	assert.Nil(t, authCtx)
	assert.Nil(t, maybeRotatedToken)
	assert.Equal(t, core.ErrUnexpectedError, err)
}
