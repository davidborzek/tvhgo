package api

import (
	"net/http"

	"github.com/davidborzek/tvhgo/api/request"
	"github.com/davidborzek/tvhgo/api/response"
)

type twoFactorSetupResponse struct {
	URL string `json:"url"`
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
