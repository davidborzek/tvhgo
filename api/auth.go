package api

import (
	"errors"
	"net/http"
	"strings"

	"github.com/davidborzek/tvhgo/api/request"
	"github.com/davidborzek/tvhgo/api/response"
	"github.com/davidborzek/tvhgo/core"
)

func (router *router) HandleAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := extractToken(r, router.cfg.Auth.Session.CookieName)

		ctx, rotatedToken, err := router.sessionManager.Validate(r.Context(), token)
		if err != nil {
			if errors.As(err, &core.InvalidOrExpiredTokenError{}) {
				response.Unauthorized(w, err)
				return
			}

			response.InternalError(w, err)
			return
		}

		if rotatedToken != nil {
			setSessionCookie(
				w,
				router.cfg.Auth.Session.CookieName,
				*rotatedToken,
				router.cfg.Auth.Session.CookieSecure,
			)
		}

		next.ServeHTTP(w, r.WithContext(
			request.WithAuthContext(r.Context(), ctx),
		))
	})
}

// Internal implementation to obtain (bearer) token from Authorization header.
func extractTokenFromHeader(r *http.Request) string {
	h := r.Header.Get("Authorization")
	return strings.ReplaceAll(h, "Bearer ", "")
}

// Internal implementation to obtain token from cookie.
func extractTokenFromCookie(r *http.Request, cookieName string) string {
	token, err := r.Cookie(cookieName)
	if err != nil {
		return ""
	}

	return token.Value
}

// Internal implementation to obtain token from Authorization header or cookie.
// Header is prioritized.
func extractToken(r *http.Request, cookieName string) string {
	if token := extractTokenFromHeader(r); token != "" {
		return token
	}

	return extractTokenFromCookie(r, cookieName)
}
