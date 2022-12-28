package api

import (
	"net/http"

	"github.com/davidborzek/tvhgo/api/request"
	"github.com/davidborzek/tvhgo/api/response"
	"github.com/davidborzek/tvhgo/core"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func (s *router) GetChannels(w http.ResponseWriter, r *http.Request) {
	var q core.PaginationSortQueryParams
	if err := request.BindQuery(r, &q); err != nil {
		response.BadRequest(w, err)
		return
	}

	if err := q.Validate(); err != nil {
		response.BadRequest(w, err)
		return
	}

	channels, err := s.channels.GetAll(r.Context(), q)
	if err != nil {
		log.WithError(err).
			Error("failed to get channels")
		response.InternalErrorCommon(w)
		return
	}

	response.JSON(w, channels, 200)
}

func (s *router) GetChannel(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	channel, err := s.channels.Get(r.Context(), id)
	if err != nil {
		if err == core.ErrChannelNotFound {
			response.NotFound(w, err)
			return
		}

		log.WithError(err).
			Error("failed to get channel")
		response.InternalErrorCommon(w)
		return
	}

	response.JSON(w, channel, 200)
}
