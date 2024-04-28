package api

import (
	"fmt"
	"net/http"

	"github.com/davidborzek/tvhgo/api/request"
	"github.com/davidborzek/tvhgo/api/response"
	"github.com/davidborzek/tvhgo/core"
	"github.com/davidborzek/tvhgo/services/auth"
)

type createUser struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	DisplayName string `json:"displayName"`
}

type userUpdate struct {
	Username    *string `json:"username"`
	Email       *string `json:"email"`
	DisplayName *string `json:"displayName"`
}

type userUpdatePassword struct {
	CurrentPassword string `json:"currentPassword"`
	Password        string `json:"password"`
}

// GetCurrentUser godoc
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
func (s *router) GetCurrentUser(w http.ResponseWriter, r *http.Request) {
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
//	@Failure	400	{object}	response.ErrorResponse
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

	err = s.users.Update(r.Context(), user)
	if err == core.ErrEmailAlreadyExists || err == core.ErrUsernameAlreadyExists {
		response.BadRequest(w, err)
		return
	}

	if err != nil {
		response.InternalError(w, err)
		return
	}

	response.JSON(w, user, 200)
}

// UpdateUserPassword godoc
//
//	@Summary	Updates the password of the current user
//	@Tags		user
//	@Accept		json
//	@Param		body	body	userUpdatePassword	true	"Body"
//	@Produce	json
//	@Success	204
//	@Failure	400	{object}	response.ErrorResponse
//	@Failure	401	{object}	response.ErrorResponse
//	@Failure	500	{object}	response.ErrorResponse
//
//	@Security	JWT
//	@Router		/user/password [patch]
func (s *router) UpdateUserPassword(w http.ResponseWriter, r *http.Request) {
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

	in := new(userUpdatePassword)
	if err := request.BindJSON(r, in); err != nil {
		response.BadRequest(w, err)
		return
	}

	if err := auth.ComparePassword(in.CurrentPassword, user.PasswordHash); err != nil {
		response.BadRequestf(w, "current password is invalid")
		return
	}

	hash, err := auth.HashPassword(in.Password)
	if err != nil {
		response.InternalError(w, err)
		return
	}

	user.PasswordHash = hash

	err = s.users.Update(r.Context(), user)
	if err != nil {
		response.InternalError(w, err)
		return
	}

	response.JSON(w, user, 200)
}

// GetUsers godoc
//
//	@Summary	Get a list of users
//	@Tags		user
//	@Produce	json
//	@Success	200	{array}		core.UserListResult
//	@Failure	401	{object}	response.ErrorResponse
//	@Failure	500	{object}	response.ErrorResponse
//
//	@Router		/users [get]
func (s *router) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := s.users.Find(r.Context(), core.UserQueryParams{})
	if err != nil {
		response.InternalError(w, err)
		return
	}

	response.JSON(w, users, 200)
}

// CreateUser godoc
//
//	@Summary	Creates a new user
//	@Tags		user
//	@Accept		json
//	@Param		body	body	createUser	true	"Body"
//	@Produce	json
//	@Success	201	{object}	core.User
//	@Failure	400	{object}	response.ErrorResponse
//	@Failure	401	{object}	response.ErrorResponse
//	@Failure	500	{object}	response.ErrorResponse
//
//	@Router		/users [post]
func (s *router) CreateUser(w http.ResponseWriter, r *http.Request) {
	in := new(createUser)
	if err := request.BindJSON(r, in); err != nil {
		response.BadRequest(w, err)
		return
	}

	hash, err := auth.HashPassword(in.Password)
	if err != nil {
		response.InternalError(w, err)
		return
	}

	user := &core.User{
		Username:     in.Username,
		Email:        in.Email,
		DisplayName:  in.DisplayName,
		PasswordHash: hash,
	}

	err = s.users.Create(r.Context(), user)
	if err == core.ErrEmailAlreadyExists || err == core.ErrUsernameAlreadyExists {
		response.BadRequest(w, err)
		return
	}

	if err != nil {
		response.InternalError(w, err)
		return
	}

	response.JSON(w, user, 201)
}

// DeleteUser godoc
//
//	@Summary	Deletes a user
//	@Tags		user
//	@Param		id	path	string	true	"User ID"
//	@Produce	json
//	@Success	204
//	@Failure	400	{object}	response.ErrorResponse
//	@Failure	401	{object}	response.ErrorResponse
//	@Failure	404	{object}	response.ErrorResponse
//	@Failure	500	{object}	response.ErrorResponse
//
//	@Router		/users/{id} [delete]
func (s *router) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := request.NumericURLParam(r, "id")
	if err != nil {
		response.BadRequest(w, err)
		return
	}

	ctx, ok := request.GetAuthContext(r.Context())
	if !ok {
		response.InternalErrorCommon(w)
		return
	}

	if ctx.UserID == id {
		response.BadRequest(w, fmt.Errorf("current user cannot be deleted"))
		return
	}

	user, err := s.users.FindById(r.Context(), id)
	if err != nil {
		response.InternalError(w, err)
		return
	}

	if user == nil {
		response.NotFound(w, fmt.Errorf("user not found"))
		return
	}

	err = s.users.Delete(r.Context(), user)
	if err != nil {
		response.InternalError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GetUsers godoc
//
//	@Summary	Get a user by ID
//	@Tags		user
//	@Param		id	path	string	true	"User ID"
//	@Produce	json
//	@Success	200	{object}	core.User
//	@Failure	400	{object}	response.ErrorResponse
//	@Failure	401	{object}	response.ErrorResponse
//	@Failure	404	{object}	response.ErrorResponse
//	@Failure	500	{object}	response.ErrorResponse
//
//	@Router		/users/{id} [get]
func (s *router) GetUser(w http.ResponseWriter, r *http.Request) {
	id, err := request.NumericURLParam(r, "id")
	if err != nil {
		response.BadRequest(w, err)
		return
	}

	user, err := s.users.FindById(r.Context(), id)
	if err != nil {
		response.InternalError(w, err)
		return
	}

	if user == nil {
		response.NotFound(w, fmt.Errorf("user not found"))
		return
	}

	response.JSON(w, user, 200)
}
