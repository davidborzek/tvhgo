package api

import (
	"net/http"

	"github.com/davidborzek/tvhgo/api/request"
	"github.com/davidborzek/tvhgo/api/response"
	"github.com/davidborzek/tvhgo/core"
	"github.com/rs/zerolog/log"
)

type twoFactorAuthSetupResponse struct {
	URL string `json:"url"`
}

type twoFactorAuthSetupRequest struct {
	Password string `json:"password"`
}

type twoFactorAuthActivateRequest struct {
	Password string `json:"password"`
	Code     string `json:"code"`
}

type twoFactorAuthDeactivateRequest struct {
	Code string `json:"code"`
}

// SetupTwoFactorAuth godoc
//
//	@Summary	Starts the two factor auth setup for the current user
//	@Tags		two-factor-auth
//	@Param		body	body	twoFactorAuthSetupRequest	true	"Body"
//	@Produce	json
//	@Success	200	{object}	twoFactorAuthSetupResponse
//	@Failure	400	{object}	response.ErrorResponse
//	@Failure	401	{object}	response.ErrorResponse
//	@Failure	409	{object}	response.ErrorResponse
//	@Failure	500	{object}	response.ErrorResponse
//	@Security	JWT
//	@Router		/two-factor-auth/setup [put]
func (s *router) SetupTwoFactorAuth(w http.ResponseWriter, r *http.Request) {
	ctx, ok := request.GetAuthContext(r.Context())
	if !ok {
		response.InternalErrorCommon(w)
		return
	}

	var in twoFactorAuthSetupRequest
	if err := request.BindJSON(r, &in); err != nil {
		response.BadRequest(w, err)
		return
	}

	if err := s.passwordAuthenticator.ConfirmPassword(r.Context(), ctx.UserID, in.Password); err != nil {
		if err == core.ErrConfirmationPasswordInvalid {
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

		log.Error().Int64("id", ctx.UserID).
			Err(err).Msg("failed to setup two factor auth")

		response.InternalErrorCommon(w)
		return
	}

	response.JSON(w, twoFactorAuthSetupResponse{URL: url}, 200)
}

// ActivateTwoFactorAuth godoc
//
//	@Summary	Activates two factor auth for the current user
//	@Tags		two-factor-auth
//	@Param		body	body	twoFactorAuthActivateRequest	true	"Body"
//	@Produce	json
//	@Success	204
//	@Failure	400	{object}	response.ErrorResponse
//	@Failure	401	{object}	response.ErrorResponse
//	@Failure	409	{object}	response.ErrorResponse
//	@Failure	500	{object}	response.ErrorResponse
//	@Security	JWT
//	@Router		/two-factor-auth/activate [put]
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

	if err := s.passwordAuthenticator.ConfirmPassword(r.Context(), ctx.UserID, in.Password); err != nil {
		if err == core.ErrConfirmationPasswordInvalid {
			response.BadRequest(w, err)
		} else {
			response.InternalErrorCommon(w)
		}

		return
	}

	err := s.twoFactorService.Activate(r.Context(), ctx.UserID, in.Code)
	if err != nil {
		if err == core.ErrTwoFactorAuthAlreadyEnabled ||
			err == core.ErrTwoFactorAuthSetupNotRunning {
			response.Conflict(w, err)
			return
		}

		if err == core.ErrTwoFactorCodeInvalid {
			response.BadRequest(w, err)
			return
		}

		log.Error().Int64("id", ctx.UserID).
			Err(err).Msg("failed to activate two factor auth")

		response.InternalError(w, err)
		return
	}

	w.WriteHeader(204)
}

// DeactivateTwoFactorAuth godoc
//
//	@Summary	Deactivates two factor auth for the current user
//	@Tags		two-factor-auth
//	@Param		body	body	twoFactorAuthDeactivateRequest	true	"Body"
//	@Produce	json
//	@Success	204
//	@Failure	400	{object}	response.ErrorResponse
//	@Failure	401	{object}	response.ErrorResponse
//	@Failure	409	{object}	response.ErrorResponse
//	@Failure	500	{object}	response.ErrorResponse
//	@Security	JWT
//	@Router		/two-factor-auth/deactivate [put]
func (s *router) DeactivateTwoFactorAuth(w http.ResponseWriter, r *http.Request) {
	ctx, ok := request.GetAuthContext(r.Context())
	if !ok {
		response.InternalErrorCommon(w)
		return
	}

	var in twoFactorAuthDeactivateRequest
	if err := request.BindJSON(r, &in); err != nil {
		response.BadRequest(w, err)
		return
	}

	if err := s.twoFactorService.Deactivate(r.Context(), ctx.UserID, in.Code); err != nil {
		if err == core.ErrTwoFactorAuthNotEnabled {
			response.Conflict(w, err)
			return
		}

		if err == core.ErrTwoFactorCodeInvalid {
			response.BadRequest(w, err)
			return
		}

		log.Error().Int64("id", ctx.UserID).
			Err(err).Msg("failed to deactivate two factor auth")

		response.InternalErrorCommon(w)
		return
	}

	w.WriteHeader(204)
}

// GetTwoFactorAuthSettings godoc
//
//	@Summary	Get the two factor auth settings for the current user
//	@Tags		two-factor-auth
//
//	@Produce	json
//	@Success	200	{array}		core.TwoFactorSettings
//	@Failure	401	{object}	response.ErrorResponse
//	@Failure	500	{object}	response.ErrorResponse
//	@Security	JWT
//	@Router		/two-factor-auth [get]
func (s *router) GetTwoFactorAuthSettings(w http.ResponseWriter, r *http.Request) {
	ctx, ok := request.GetAuthContext(r.Context())
	if !ok {
		response.InternalErrorCommon(w)
		return
	}

	settings, err := s.twoFactorService.GetSettings(r.Context(), ctx.UserID)
	if err != nil {
		log.Error().Int64("id", ctx.UserID).
			Err(err).Msg("failed to get two factor auth settings")

		response.InternalErrorCommon(w)
		return
	}

	response.JSON(w, settings, 200)
}
