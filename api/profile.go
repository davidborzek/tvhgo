package api

import (
	"net/http"

	"github.com/davidborzek/tvhgo/api/response"
	"github.com/rs/zerolog/log"
)

// GetStreamProfiles godoc
//
//	@Summary	Get list of stream profiles
//	@Tags		profiles
//
//	@Produce	json
//	@Success	200	{array}		core.StreamProfile
//	@Failure	401	{object}	response.ErrorResponse
//	@Failure	500	{object}	response.ErrorResponse
//	@Router		/profiles/stream [get]
func (s *router) GetStreamProfiles(w http.ResponseWriter, r *http.Request) {
	profiles, err := s.profileService.GetStreamProfiles(r.Context())
	if err != nil {
		log.Error().Err(err).Msg("failed to get channels")

		response.InternalErrorCommon(w)
		return
	}

	response.JSON(w, profiles, 200)
}
