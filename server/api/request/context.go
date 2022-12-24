package request

import (
	"context"

	"github.com/davidborzek/tvhgo/core"
	log "github.com/sirupsen/logrus"
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
	if !ok {
		log.Error("auth context could not be obtained from request context")
	}

	return auth, ok
}
