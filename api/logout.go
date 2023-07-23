package api

import (
	"net/http"

	"github.com/davidborzek/tvhgo/api/request"
	"github.com/davidborzek/tvhgo/api/response"
)

// Internal implementation to delete the session cookie.
func deleteSessionCookie(w http.ResponseWriter, cookieName string) {
	c := &http.Cookie{
		Name:     cookieName,
		Value:    "",
		Path:     "/api",
		MaxAge:   0,
		HttpOnly: true,
	}

	http.SetCookie(w, c)
}

func (s *router) Logout(w http.ResponseWriter, r *http.Request) {
	ctx, ok := request.GetAuthContext(r.Context())
	if !ok {
		response.InternalErrorCommon(w)
		return
	}

	err := s.sessionManager.Revoke(
		r.Context(), ctx.SessionID, ctx.UserID,
	)

	if err != nil {
		response.InternalError(w, err)
		return
	}

	deleteSessionCookie(w, s.cfg.Auth.Session.CookieName)
	w.WriteHeader(200)
}
