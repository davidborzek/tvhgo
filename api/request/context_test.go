package request_test

import (
	"context"
	"testing"

	"github.com/davidborzek/tvhgo/api/request"
	"github.com/davidborzek/tvhgo/core"
	"github.com/stretchr/testify/assert"
)

func TestWithAuthContextEntrichesAContextWithAuthContext(t *testing.T) {
	ctx := context.Background()
	expectedAuthCtx := core.AuthContext{
		UserID:    1234,
		SessionID: 2345,
	}

	newCtx := request.WithAuthContext(ctx, &expectedAuthCtx)

	authCtx, ok := request.GetAuthContext(newCtx)

	assert.Equal(t, expectedAuthCtx, *authCtx)
	assert.True(t, ok)
}

func TestGetAuthContextReturnsFalseWhenAuthContextCouldNotBeObtained(t *testing.T) {
	ctx := context.Background()
	authCtx, ok := request.GetAuthContext(ctx)

	assert.Nil(t, authCtx)
	assert.False(t, ok)
}
