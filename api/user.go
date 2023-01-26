package api

import (
	"net/http"

	"github.com/davidborzek/tvhgo/api/request"
	"github.com/davidborzek/tvhgo/api/response"
	"github.com/davidborzek/tvhgo/services/auth"
)

type userUpdate struct {
	Username    *string `json:"username"`
	Email       *string `json:"email"`
	DisplayName *string `json:"displayName"`
	Password    *string `json:"password"`
}

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

// UpdateUser godoc
//
//	@Summary	Updates the current user
//	@Tags		user
//	@Accept		json
//	@Param		body	body	userUpdate	true	"Body"
//	@Produce	json
//	@Success	200	{object}	core.User
//	@Failure	401	{object}	response.ErrorResponse
//	@Failure	500	{object}	response.ErrorResponse
//
//	@Security	JWT
//	@Router		/user [patch]
func (s *router) UpdateUser(w http.ResponseWriter, r *http.Request) {
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

	in := new(userUpdate)
	if err := request.BindJSON(r, in); err != nil {
		response.BadRequest(w, err)
		return
	}

	if in.Username != nil {
		user.Username = *in.Username
	}

	if in.DisplayName != nil {
		user.DisplayName = *in.DisplayName
	}

	if in.Email != nil {
		user.Email = *in.Email
	}

	if in.Password != nil {
		hash, err := auth.HashPassword(*in.Password)
		if err != nil {
			response.InternalError(w, err)
			return
		}

		user.PasswordHash = hash
	}

	err = s.users.Update(r.Context(), user)
	if err != nil {
		response.InternalError(w, err)
		return
	}

	response.JSON(w, user, 200)
}
