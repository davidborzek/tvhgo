package request

import (
	"context"

	"github.com/davidborzek/tvhgo/core"
)

type key int

const (
	sessionKey key = iota
)

// WithAuthContext returns a copy of parent ctx which includes the auth context.
func WithAuthContext(parent context.Context, auth *core.AuthContext) context.Context {
	return context.WithValue(parent, sessionKey, auth)
}

// SessionFrom returns the auth context for the given context.
func GetAuthContext(ctx context.Context) (*core.AuthContext, bool) {
	auth, ok := ctx.Value(sessionKey).(*core.AuthContext)
	return auth, ok
}
