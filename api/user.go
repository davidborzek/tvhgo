package api

import (
	"net/http"

	"github.com/davidborzek/tvhgo/api/request"
	"github.com/davidborzek/tvhgo/api/response"
)

// GetUser godoc
//
//	@Summary	Get the current user
//	@Tags		user
//	@Produce	json
//	@Success	200	{object}	core.User
//	@Failure	401	{object}	response.ErrorResponse
//	@Failure	500	{object}	response.ErrorResponse
//
//	@Security	JWT
//	@Router		/user [get]
func (s *router) GetUser(w http.ResponseWriter, r *http.Request) {
	ctx, ok := request.GetAuthContext(r.Context())
	if !ok {
		response.InternalErrorCommon(w)
		return
	}

	user, err := s.users.FindById(r.Context(), ctx.UserID)
	if err != nil {
		response.InternalError(w, err)
		return
	}

	response.JSON(w, user, 200)
}
