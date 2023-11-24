package api

import (
	"net/http"

	"github.com/davidborzek/tvhgo/api/request"
	"github.com/davidborzek/tvhgo/api/response"
	log "github.com/sirupsen/logrus"
)

type tokenResponse struct {
	Token string `json:"token"`
}

type createTokenRequest struct {
	Name string `json:"name"`
}

// GetSessions godoc
//
//	@Summary	Get list of tokens for the current user
//	@Tags		tokens
//
//	@Produce	json
//	@Success	200	{array}		core.Token
//	@Failure	401	{object}	response.ErrorResponse
//	@Failure	500	{object}	response.ErrorResponse
//	@Security	JWT
//	@Router		/tokens [get]
func (s *router) GetTokens(w http.ResponseWriter, r *http.Request) {
	ctx, ok := request.GetAuthContext(r.Context())
	if !ok {
		response.InternalErrorCommon(w)
		return
	}

	tokens, err := s.tokens.FindByUser(r.Context(), ctx.UserID)
	if err != nil {
		log.WithError(err).
			Error("failed to get tokens")

		response.InternalErrorCommon(w)
		return
	}

	response.JSON(w, tokens, 200)
}

// CreateToken godoc
//
//	@Summary	Creates an api token
//	@Tags		tokens
//	@Param		body	body	createTokenRequest	true	"Body"
//	@Produce	json
//	@Success	200	{object}	tokenResponse
//	@Failure	401	{object}	response.ErrorResponse
//	@Failure	403	{object}	response.ErrorResponse
//	@Failure	500	{object}	response.ErrorResponse
//	@Security	JWT
//	@Router		/tokens [post]
func (s *router) CreateToken(w http.ResponseWriter, r *http.Request) {
	ctx, ok := request.GetAuthContext(r.Context())
	if !ok {
		response.InternalErrorCommon(w)
		return
	}

	if ctx.SessionID == nil {
		response.Forbiddenf(w, "a token can only be created via a web session")
		return
	}

	var in createTokenRequest
	if err := request.BindJSON(r, &in); err != nil {
		response.BadRequest(w, err)
		return
	}

	if in.Name == "" {
		response.BadRequestf(w, "invalid name")
		return
	}

	token, err := s.tokenService.Create(r.Context(), ctx.UserID, in.Name)
	if err != nil {
		log.WithError(err).
			Error("failed to create token")

		response.InternalErrorCommon(w)
		return
	}

	response.JSON(w, &tokenResponse{Token: token}, 200)
}

// DeleteSession godoc
//
//	@Summary	Revokes a token
//	@Tags		tokens
//	@Param		id	path	int	true	"Token ID"
//	@Produce	json
//	@Success	204
//	@Failure	401	{object}	response.ErrorResponse
//	@Failure	403	{object}	response.ErrorResponse
//	@Failure	500	{object}	response.ErrorResponse
//	@Security	JWT
//	@Router		/tokens/{id} [delete]
func (s *router) DeleteToken(w http.ResponseWriter, r *http.Request) {
	ctx, ok := request.GetAuthContext(r.Context())
	if !ok {
		response.InternalErrorCommon(w)
		return
	}

	if ctx.SessionID == nil {
		response.Forbiddenf(w, "a token can only be deleted via a web session")
		return
	}

	tokenID, err := request.NumericURLParam(r, "id")
	if err != nil {
		response.BadRequestf(w, "invalid value for parameter 'id'")
		return
	}

	if err := s.tokenService.Revoke(r.Context(), tokenID); err != nil {
		log.WithError(err).
			Error("failed to delete token")
		response.InternalErrorCommon(w)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
