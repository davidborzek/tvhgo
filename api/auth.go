package api

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"strings"

	"github.com/davidborzek/tvhgo/api/request"
	"github.com/davidborzek/tvhgo/api/response"
	"github.com/davidborzek/tvhgo/core"
	"github.com/rs/zerolog/log"
)

func (router *router) HandleAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if router.cfg.Auth.ReverseProxy.Enabled {
			router.handleReverseProxyAuthentication(r, w, next)
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

// handleReverseProxyAuthentication authenticates the user based on the remote address and the
// user header. The remote address must be contained in the list of allowed proxies and the user
// header must contain a valid username.
func (router *router) handleReverseProxyAuthentication(
	r *http.Request,
	w http.ResponseWriter,
	next http.Handler,
) {
	if !isIPAllowed(r.RemoteAddr, router.cfg.Auth.ReverseProxy.AllowedProxies) {
		log.Debug().
			Str("remote_addr", r.RemoteAddr).
			Interface("allowed_proxies", router.cfg.Auth.ReverseProxy.AllowedProxies).
			Msg("[reverse proxy auth] remote address not allowed")
		response.Unauthorized(w, core.ErrTokenInvalid)
		return
	}

	remoteUser := r.Header.Get(router.cfg.Auth.ReverseProxy.UserHeader)
	if remoteUser == "" {
		response.Unauthorized(w, core.ErrTokenInvalid)
		return
	}

	user, err := router.users.FindByUsername(r.Context(), remoteUser)
	if err != nil {
		log.Error().Err(err).Str("remote_user", remoteUser).
			Msg("[reverse proxy auth] failed to find user")

		response.InternalError(w, err)
		return
	}

	if user == nil {
		if !router.cfg.Auth.ReverseProxy.AllowRegistration {
			log.Debug().Str("remote_user", remoteUser).
				Msg("[reverse proxy auth] remote user not found and registrations are disabled")

			response.Unauthorized(w, core.ErrTokenInvalid)
			return
		}

		displayName := r.Header.Get(router.cfg.Auth.ReverseProxy.NameHeader)
		if displayName == "" {
			displayName = remoteUser
		}

		email := r.Header.Get(router.cfg.Auth.ReverseProxy.EmailHeader)
		if email == "" {
			email = fmt.Sprintf("%s@tvhgo.local", remoteUser)
		}

		log.Debug().Str("remote_user", remoteUser).
			Str("display_name", displayName).
			Str("email", email).
			Msg("[reverse proxy auth] creating new user")

		user = &core.User{
			Username:    remoteUser,
			DisplayName: displayName,
			Email:       email,
		}

		if err := router.users.Create(r.Context(), user); err != nil {
			log.Error().Err(err).Str("remote_user", remoteUser).
				Msg("[reverse proxy auth] failed to create user")

			response.InternalError(w, err)
			return
		}
	}

	ctx := &core.AuthContext{
		UserID: user.ID,
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

// isIPAllowed checks if the remote address is contained in the list of allowed networks.
// The list of allowed networks can be either IP addresses or CIDR notation.
func isIPAllowed(addr string, allowedNetworks []string) bool {
	if addr == "" || len(allowedNetworks) == 0 {
		return false
	}

	if net.ParseIP(addr) == nil {
		addr, _, _ = net.SplitHostPort(addr)
	}

	if addr == "" {
		return false
	}

	ip, _, err := net.ParseCIDR(fmt.Sprintf("%s/32", addr))
	if err != nil {
		log.Debug().Err(err).Str("addr", addr).
			Msg("[reverse proxy auth] failed to parse remote addr")

		return false
	}

	for _, allowed := range allowedNetworks {
		if allowed == addr {
			return true
		}

		if _, ipNet, err := net.ParseCIDR(allowed); err == nil && ipNet.Contains(ip) {
			return true
		} else if err != nil {
			log.Debug().Err(err).Str("proxy", allowed).
				Msg("[reverse proxy auth] failed to parse allowed proxy")
		}
	}

	return false
}
