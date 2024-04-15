package api

import (
	"errors"
	"net/http"

	"github.com/davidborzek/tvhgo/api/response"
	"github.com/davidborzek/tvhgo/core"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
)

// GetDVRConfigList godoc
//
//	@Summary	Get list of dvr configs
//	@Tags		dvr
//
//	@Produce	json
//	@Success	200	{array}		core.DVRConfig
//	@Failure	401	{object}	response.ErrorResponse
//	@Failure	500	{object}	response.ErrorResponse
//	@Router		/dvr/config [get]
func (s *router) GetDVRConfigList(w http.ResponseWriter, r *http.Request) {
	configs, err := s.dvrConfigService.GetAll(r.Context())
	if err != nil {
		log.Error().Err(err).Msg("failed to get channels")

		response.InternalErrorCommon(w)
		return
	}

	response.JSON(w, configs, 200)
}

// GetDVRConfig godoc
//
//	@Summary	Get a dvr configs by id
//	@Tags		dvr
//	@Param		id	path	string	true	"DVR Config ID"
//
//	@Produce	json
//	@Success	200	{object}	core.DVRConfig
//	@Failure	401	{object}	response.ErrorResponse
//	@Failure	500	{object}	response.ErrorResponse
//	@Router		/dvr/config/{id} [get]
func (s *router) GetDVRConfig(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	// Using the existing GetAll method to get all configs and then filter the wanted config,
	// because it's easier to implement and the performance impact is negligible.
	configs, err := s.dvrConfigService.GetAll(r.Context())
	if err != nil {
		log.Error().Err(err).Msg("failed to get channels")

		response.InternalErrorCommon(w)
		return
	}

	var config *core.DVRConfig
	for _, c := range configs {
		if c.ID == id {
			config = &c
			break
		}
	}

	if config == nil {
		response.NotFound(w, errors.New("dvr config not found"))
		return
	}

	response.JSON(w, config, 200)
}

// DeleteDVRConfig godoc
//
//	@Summary	Deletes a dvr config by id
//	@Tags		dvr
//	@Param		id	path	string	true	"DVR config ID"
//	@Produce	json
//	@Success	204
//	@Failure	401	{object}	response.ErrorResponse
//	@Failure	404	{object}	response.ErrorResponse
//	@Failure	500	{object}	response.ErrorResponse
//	@Router		/dvr/config/{id} [delete]
func (s *router) DeleteDVRConfig(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := s.dvrConfigService.Delete(r.Context(), id)
	if err != nil {
		if err == core.ErrDVRConfigNotFound {
			response.NotFound(w, err)
			return
		}

		log.Error().Err(err).Msg("failed to delete dvr config")

		response.InternalErrorCommon(w)
		return
	}

	response.JSON(w, nil, 204)
}
