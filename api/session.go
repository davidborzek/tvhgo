package api

import (
	"net/http"

	"github.com/davidborzek/tvhgo/api/request"
	"github.com/davidborzek/tvhgo/api/response"
	"github.com/rs/zerolog/log"
)

// GetSessionsForCurrentUser godoc
//
//	@Summary	Get list of sessions for the current user
//	@Tags		sessions
//
//	@Produce	json
//	@Success	200	{array}		core.Session
//	@Failure	401	{object}	response.ErrorResponse
//	@Failure	500	{object}	response.ErrorResponse
//	@Security	JWT
//	@Router		/sessions [get]
func (s *router) GetSessionsForCurrentUser(w http.ResponseWriter, r *http.Request) {
	ctx, ok := request.GetAuthContext(r.Context())
	if !ok {
		response.InternalErrorCommon(w)
		return
	}

	sessions, err := s.sessions.FindByUser(r.Context(), ctx.UserID)
	if err != nil {
		log.Error().Err(err).Msg("failed to get sessions")

		response.InternalErrorCommon(w)
		return
	}

	response.JSON(w, sessions, 200)
}

// GetSessions godoc
//
//	@Summary	Get list of session for a user
//	@Tags		sessions
//	@Param		id	path	int	true	"User id"
//	@Produce	json
//	@Success	200	{array}		core.Session
//	@Failure	401	{object}	response.ErrorResponse
//	@Failure	500	{object}	response.ErrorResponse
//	@Security	JWT
//	@Router		/users/{id}/sessions [get]
func (s *router) GetSessions(w http.ResponseWriter, r *http.Request) {
	userID, err := request.NumericURLParam(r, "id")
	if err != nil {
		response.BadRequestf(w, "invalid value for parameter 'id'")
		return
	}

	sessions, err := s.sessions.FindByUser(r.Context(), userID)
	if err != nil {
		log.Error().Err(err).Msg("failed to get sessions")

		response.InternalErrorCommon(w)
		return
	}

	response.JSON(w, sessions, 200)
}

// DeleteSession godoc
//
//	@Summary	Revokes a session
//	@Tags		sessions
//	@Param		id	path	int	true	"Session id"
//	@Produce	json
//	@Success	204
//	@Failure	401	{object}	response.ErrorResponse
//	@Failure	500	{object}	response.ErrorResponse
//	@Security	JWT
//	@Router		/sessions/{id} [delete]
func (s *router) DeleteSession(w http.ResponseWriter, r *http.Request) {
	ctx, ok := request.GetAuthContext(r.Context())
	if !ok {
		response.InternalErrorCommon(w)
		return
	}

	sessionID, err := request.NumericURLParam(r, "id")
	if err != nil {
		response.BadRequestf(w, "invalid value for parameter 'id'")
		return
	}

	if ctx.SessionID != nil && sessionID == *ctx.SessionID {
		response.BadRequestf(w, "current session cannot be revoked")
		return
	}

	if err := s.sessionManager.Revoke(r.Context(), sessionID, ctx.UserID); err != nil {
		log.Error().Int64("id", sessionID).
			Err(err).Msg("failed to revoke session")

		response.InternalErrorCommon(w)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// DeleteUserSession godoc
//
//	@Summary	Revokes a session for a user
//	@Tags		sessions
//	@Param		id		path	int	true	"Session id"
//	@Param		userId	path	int	true	"User id"
//	@Produce	json
//	@Success	204
//	@Failure	401	{object}	response.ErrorResponse
//	@Failure	500	{object}	response.ErrorResponse
//	@Security	JWT
//	@Router		/users/{userId}/sessions/{id} [delete]
func (s *router) DeleteUserSession(w http.ResponseWriter, r *http.Request) {
	userID, err := request.NumericURLParam(r, "userId")
	if err != nil {
		response.BadRequestf(w, "invalid value for parameter 'userId'")
		return
	}

	sessionID, err := request.NumericURLParam(r, "id")
	if err != nil {
		response.BadRequestf(w, "invalid value for parameter 'id'")
		return
	}

	ctx, ok := request.GetAuthContext(r.Context())
	if !ok {
		response.InternalErrorCommon(w)
		return
	}

	if ctx.UserID == userID && ctx.SessionID != nil && sessionID == *ctx.SessionID {
		response.BadRequestf(w, "current session cannot be revoked")
		return
	}

	if err := s.sessionManager.Revoke(r.Context(), sessionID, userID); err != nil {
		log.Error().Int64("id", sessionID).
			Err(err).Msg("failed to revoke session")

		response.InternalErrorCommon(w)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
