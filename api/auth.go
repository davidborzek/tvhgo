package api

import (
	"errors"
	"net/http"
	"strings"

	"github.com/davidborzek/tvhgo/api/request"
	"github.com/davidborzek/tvhgo/api/response"
	"github.com/davidborzek/tvhgo/core"
)

type (
	authResponse struct {
		UserID      int64  `json:"userId"`
		SessionID   *int64 `json:"sessionId,omitempty"`
		ForwardAuth bool   `json:"forwardAuth"`
	}
)

func (router *router) GetAuthInfo(w http.ResponseWriter, r *http.Request) {
	ctx, ok := request.GetAuthContext(r.Context())
	if !ok {
		response.InternalErrorCommon(w)
	}

	out := authResponse{
		UserID:      ctx.UserID,
		SessionID:   ctx.SessionID,
		ForwardAuth: ctx.ForwardAuth,
	}

	response.JSON(w, out, 200)
}

func (router *router) HandleAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if router.cfg.Auth.ReverseProxy.Enabled && router.handleForwardAuth(r, w, next) {
			return
		}

		headerToken := extractTokenFromHeader(r)
		if headerToken != "" {
			router.handleTokenAuthorization(r, w, next, headerToken)
			return
		}

		cookieToken := extractTokenFromCookie(r, router.cfg.Auth.Session.CookieName)
		if cookieToken == "" {
			response.Unauthorized(w, core.ErrTokenInvalid)
			return
		}

		router.handleSessionTokenAuthorization(r, w, next, cookieToken)
	})
}

func (router *router) handleSessionTokenAuthorization(
	r *http.Request,
	w http.ResponseWriter,
	next http.Handler,
	token string,
) {
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
}

func (router *router) handleTokenAuthorization(
	r *http.Request,
	w http.ResponseWriter,
	next http.Handler,
	token string,
) {
	ctx, err := router.tokenService.Validate(r.Context(), token)
	if err != nil {
		if errors.As(err, &core.InvalidOrExpiredTokenError{}) {
			response.Unauthorized(w, err)
			return
		}

		response.InternalError(w, err)
		return
	}

	next.ServeHTTP(w, r.WithContext(
		request.WithAuthContext(r.Context(), ctx),
	))
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
