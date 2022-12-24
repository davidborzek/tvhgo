package api

import (
	"net/http"

	"github.com/davidborzek/tvhgo/api/request"
	"github.com/davidborzek/tvhgo/api/response"
)

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
