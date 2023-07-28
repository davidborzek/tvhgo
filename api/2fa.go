package api

import (
	"net/http"

	"github.com/davidborzek/tvhgo/api/request"
	"github.com/davidborzek/tvhgo/api/response"
)

type twoFactorSetupResponse struct {
	URL string `json:"url"`
}

type twoFactorAuthEnableRequest struct {
	Code string `json:"code"`
}

func (s *router) Setup2FA(w http.ResponseWriter, r *http.Request) {
	ctx, ok := request.GetAuthContext(r.Context())
	if !ok {
		response.InternalErrorCommon(w)
		return
	}

	url, err := s.twoFactorService.Setup(r.Context(), ctx.UserID)
	if err != nil {
		response.InternalErrorCommon(w)
		return
	}

	response.JSON(w, twoFactorSetupResponse{URL: url}, 200)
}

func (s *router) Enable2FA(w http.ResponseWriter, r *http.Request) {
	ctx, ok := request.GetAuthContext(r.Context())
	if !ok {
		response.InternalErrorCommon(w)
		return
	}

	var in twoFactorAuthEnableRequest
	if err := request.BindJSON(r, &in); err != nil {
		response.BadRequest(w, err)
		return
	}

	err := s.twoFactorService.Enable(r.Context(), ctx.UserID, in.Code)
	if err != nil {
		response.InternalError(w, err)
		return
	}

	w.WriteHeader(204)
}

func (s *router) Deactivate2FA(w http.ResponseWriter, r *http.Request) {
	ctx, ok := request.GetAuthContext(r.Context())
	if !ok {
		response.InternalErrorCommon(w)
		return
	}

	if err := s.twoFactorService.Deactivate(r.Context(), ctx.UserID); err != nil {
		response.InternalErrorCommon(w)
		return
	}

	w.WriteHeader(204)
}
