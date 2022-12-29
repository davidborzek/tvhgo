package api

import (
	"net/http"
	"time"

	"github.com/davidborzek/tvhgo/api/request"
	"github.com/davidborzek/tvhgo/api/response"
	"github.com/davidborzek/tvhgo/core"
)

var (
	// Max cookie age
	maxAge = int((365 * 24 * time.Hour).Seconds())
)

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type loginResponse struct {
	Token string `json:"token"`
}

// Internal implementation to set the session cookie.
func setSessionCookie(w http.ResponseWriter, cookieName string, token string, secure bool) {
	c := &http.Cookie{
		Name:     cookieName,
		Value:    token,
		Path:     "/api",
		MaxAge:   maxAge,
		Secure:   secure,
		HttpOnly: true,
	}

	http.SetCookie(w, c)
}

func (s *router) Login(w http.ResponseWriter, r *http.Request) {
	var in loginRequest
	if err := request.BindJSON(r, &in); err != nil {
		response.BadRequest(w, err)
		return
	}

	user, err := s.passwordAuthenticator.Login(
		r.Context(),
		in.Username,
		in.Password,
	)

	if err != nil {
		if err == core.ErrInvalidUsernameOrPassword {
			response.Unauthorized(w, err)
			return
		}

		response.InternalError(w, err)
		return
	}

	token, err := s.sessions.Create(
		r.Context(),
		user.ID,
		request.RemoteAddr(r),
		r.UserAgent(),
	)

	if err != nil {
		response.InternalError(w, err)
		return
	}

	setSessionCookie(
		w,
		s.cfg.Auth.Session.CookieName,
		token,
		s.cfg.Auth.Session.CookieSecure,
	)
	response.JSON(w, loginResponse{Token: token}, 200)
}
