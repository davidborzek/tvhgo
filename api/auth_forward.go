package api

import (
	"fmt"
	"net"
	"net/http"

	"github.com/davidborzek/tvhgo/api/request"
	"github.com/davidborzek/tvhgo/core"
	"github.com/rs/zerolog/log"
)

// handleForwardAuth authenticates the user based on the remote address and the
// user header. The remote address must be contained in the list of allowed proxies and the user
// header must contain a valid username.
func (router *router) handleForwardAuth(
	r *http.Request,
	w http.ResponseWriter,
	next http.Handler,
) bool {
	if !isIPAllowed(r.RemoteAddr, router.cfg.Auth.ReverseProxy.AllowedProxies) {
		log.Debug().
			Str("remote_addr", r.RemoteAddr).
			Interface("allowed_proxies", router.cfg.Auth.ReverseProxy.AllowedProxies).
			Msg("[reverse proxy auth] remote address not allowed")
		return false
	}

	remoteUser := r.Header.Get(router.cfg.Auth.ReverseProxy.UserHeader)
	if remoteUser == "" {
		return false
	}

	user, err := router.users.FindByUsername(r.Context(), remoteUser)
	if err != nil {
		log.Error().Err(err).Str("remote_user", remoteUser).
			Msg("[reverse proxy auth] failed to find user")

		return false
	}

	if user == nil {
		if !router.cfg.Auth.ReverseProxy.AllowRegistration {
			log.Debug().Str("remote_user", remoteUser).
				Msg("[reverse proxy auth] remote user not found and registrations are disabled")

			return false
		}

		user, err = router.createForwardAuthUser(r, remoteUser)
		if err != nil {
			log.Error().Err(err).Str("remote_user", remoteUser).
				Msg("[reverse proxy auth] failed to create user")

			return false
		}
	}

	ctx := &core.AuthContext{
		UserID:      user.ID,
		ForwardAuth: true,
	}

	next.ServeHTTP(w, r.WithContext(
		request.WithAuthContext(r.Context(), ctx),
	))
	return true
}

// createForwardAuthUser creates a new user based on the remote user header.
func (router *router) createForwardAuthUser(r *http.Request, remoteUser string) (*core.User, error) {
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

	user := &core.User{
		Username:    remoteUser,
		DisplayName: displayName,
		Email:       email,
	}

	if err := router.users.Create(r.Context(), user); err != nil {
		return nil, err
	}

	return user, nil
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
