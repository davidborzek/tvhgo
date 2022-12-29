package api

import (
	"net/http"

	"github.com/davidborzek/tvhgo/api/request"
	"github.com/davidborzek/tvhgo/api/response"
	"github.com/davidborzek/tvhgo/core"
	log "github.com/sirupsen/logrus"
)

func (s *router) GetEpg(w http.ResponseWriter, r *http.Request) {
	var q core.GetEpgQueryParams
	if err := request.BindQuery(r, &q); err != nil {
		response.BadRequest(w, err)
		return
	}

	if err := q.Validate(); err != nil {
		response.BadRequest(w, err)
		return
	}

	events, err := s.epg.GetEvents(r.Context(), q)
	if err != nil {
		log.WithError(err).
			Error("failed to get epg events")
		response.InternalErrorCommon(w)
		return
	}

	response.JSON(w, events, 200)
}

func (s *router) GetEpgChannelEvents(w http.ResponseWriter, r *http.Request) {
	var q core.GetEpgChannelEventsQueryParams
	if err := request.BindQuery(r, &q); err != nil {
		response.BadRequest(w, err)
		return
	}

	if err := q.Validate(); err != nil {
		response.BadRequest(w, err)
		return
	}

	events, err := s.epg.GetChannelEvents(r.Context(), q)
	if err != nil {
		log.WithError(err).
			Error("failed to get epg events")
		response.InternalErrorCommon(w)
		return
	}

	response.JSON(w, events, 200)
}

func (s *router) GetEpgEvent(w http.ResponseWriter, r *http.Request) {
	id, err := request.NumericURLParam(r, "id")
	if err != nil {
		response.BadRequestf(w, "invalid value for parameter 'id'")
		return
	}

	event, err := s.epg.GetEvent(r.Context(), id)
	if err != nil {
		if err == core.ErrEpgEventNotFound {
			response.NotFound(w, err)
			return
		}

		log.WithError(err).
			WithField("id", id).
			Error("failed to get epg event")

		response.InternalErrorCommon(w)
		return
	}

	response.JSON(w, event, 200)
}

func (s *router) GetEpgContentTypes(w http.ResponseWriter, r *http.Request) {
	contentTypes, err := s.epg.GetContentTypes(r.Context())
	if err != nil {
		log.WithError(err).
			Error("failed to get epg content types")
		response.InternalErrorCommon(w)
		return
	}

	response.JSON(w, contentTypes, 200)
}
