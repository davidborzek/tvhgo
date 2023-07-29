package api

import (
	"net/http"

	"github.com/davidborzek/tvhgo/api/request"
	"github.com/davidborzek/tvhgo/api/response"
	"github.com/davidborzek/tvhgo/core"
	log "github.com/sirupsen/logrus"
)

type twoFactorAuthSetupResponse struct {
	URL string `json:"url"`
}

type twoFactorAuthBaseRequest struct {
	Password string `json:"password"`
}

type twoFactorAuthActivateRequest struct {
	twoFactorAuthBaseRequest
	Code string `json:"code"`
}

func (s *router) SetupTwoFactorAuth(w http.ResponseWriter, r *http.Request) {
	ctx, ok := request.GetAuthContext(r.Context())
	if !ok {
		response.InternalErrorCommon(w)
		return
	}

	var in twoFactorAuthBaseRequest
	if err := request.BindJSON(r, &in); err != nil {
		response.BadRequest(w, err)
		return
	}

	if err := s.confirmPassword(r.Context(), ctx.UserID, in.Password); err != nil {
		if err == errTwoFactorConfirmationPasswordInvalid {
			response.BadRequest(w, err)
		} else {
			response.InternalErrorCommon(w)
		}

		return
	}

	url, err := s.twoFactorService.Setup(r.Context(), ctx.UserID)
	if err != nil {
		if err == core.ErrTwoFactorAuthAlreadyEnabled {
			response.Conflict(w, err)
			return
		}

		log.WithError(err).
			WithField("userId", ctx.UserID).
			Error("failed to setup two factor auth")

		response.InternalErrorCommon(w)
		return
	}

	response.JSON(w, twoFactorAuthSetupResponse{URL: url}, 200)
}

func (s *router) ActivateTwoFactorAuth(w http.ResponseWriter, r *http.Request) {
	ctx, ok := request.GetAuthContext(r.Context())
	if !ok {
		response.InternalErrorCommon(w)
		return
	}

	var in twoFactorAuthActivateRequest
	if err := request.BindJSON(r, &in); err != nil {
		response.BadRequest(w, err)
		return
	}

	if err := s.confirmPassword(r.Context(), ctx.UserID, in.Password); err != nil {
		if err == errTwoFactorConfirmationPasswordInvalid {
			response.BadRequest(w, err)
		} else {
			response.InternalErrorCommon(w)
		}

		return
	}

	err := s.twoFactorService.Activate(r.Context(), ctx.UserID, in.Code)
	if err != nil {
		if err == core.ErrTwoFactorAuthAlreadyEnabled {
			response.Conflict(w, err)
			return
		}

		if err == core.ErrTwoFactorCodeInvalid {
			response.BadRequest(w, err)
			return
		}

		log.WithError(err).
			WithField("userId", ctx.UserID).
			Error("failed to activate two factor auth")

		response.InternalError(w, err)
		return
	}

	w.WriteHeader(204)
}

func (s *router) DeactivateTwoFactorAuth(w http.ResponseWriter, r *http.Request) {
	ctx, ok := request.GetAuthContext(r.Context())
	if !ok {
		response.InternalErrorCommon(w)
		return
	}

	var in twoFactorAuthBaseRequest
	if err := request.BindJSON(r, &in); err != nil {
		response.BadRequest(w, err)
		return
	}

	if err := s.confirmPassword(r.Context(), ctx.UserID, in.Password); err != nil {
		if err == errTwoFactorConfirmationPasswordInvalid {
			response.BadRequest(w, err)
		} else {
			response.InternalErrorCommon(w)
		}

		return
	}

	if err := s.twoFactorService.Deactivate(r.Context(), ctx.UserID); err != nil {
		if err == core.ErrTwoFactorAuthNotEnabled {
			response.Conflict(w, err)
			return
		}

		log.WithError(err).
			WithField("userId", ctx.UserID).
			Error("failed to deactivate two factor auth")

		response.InternalErrorCommon(w)
		return
	}

	w.WriteHeader(204)
}

func (s *router) GetTwoFactorAuthSettings(w http.ResponseWriter, r *http.Request) {
	ctx, ok := request.GetAuthContext(r.Context())
	if !ok {
		response.InternalErrorCommon(w)
		return
	}

	settings, err := s.twoFactorService.GetSettings(r.Context(), ctx.UserID)
	if err != nil {
		log.WithError(err).
			WithField("userId", ctx.UserID).
			Error("failed to get two factor auth settings")

		response.InternalErrorCommon(w)
		return
	}

	response.JSON(w, settings, 200)
}
