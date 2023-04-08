package api

import (
	"net/http"

	"github.com/davidborzek/tvhgo/api/request"
	"github.com/davidborzek/tvhgo/api/response"
	"github.com/davidborzek/tvhgo/core"
	log "github.com/sirupsen/logrus"
)

// GetEpgEvents godoc
//
//	@Summary	Get epg events
//	@Tags		epg
//	@Param		limit		query	int		false	"Limit"
//	@Param		offset		query	int		false	"Offset"
//	@Param		sort_key	query	string	false	"Sort key"
//	@Param		sort_dir	query	string	false	"Sort direction"
//	@Param		title		query	string	false	"Title"
//	@Param		fullText	query	bool	false	"Enable full test search"
//	@Param		lang		query	string	false	"Language"
//	@Param		nowPlaying	query	bool	false	"Now playing"
//	@Param		channel		query	string	false	"Channel name or channel id"
//	@Param		contentType	query	string	false	"Content type"
//	@Param		durationMin	query	int64	false	"Minimum Duration"
//	@Param		durationMax	query	int64	false	"Maximum Duration"
//	@Param		startsAt	query	int64	false	"Start timestamp"
//	@Param		endsAt		query	int64	false	"End timestamp"
//
//	@Produce	json
//	@Success	200	{array}		core.EpgEvent
//	@Failure	401	{object}	response.ErrorResponse
//	@Failure	500	{object}	response.ErrorResponse
//
//	@Security	JWT
//	@Router		/epg/events [get]
func (s *router) GetEpgEvents(w http.ResponseWriter, r *http.Request) {
	var q core.GetEpgEventsQueryParams
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

// GetEpg godoc
//
//	@Summary	Get epg
//	@Tags		epg
//	@Param		sort_key	query	string	false	"Sort key"
//	@Param		sort_dir	query	string	false	"Sort direction"
//	@Param		startsAt	query	int64	false	"Start timestamp"
//	@Param		endsAt		query	int64	false	"End timestamp"
//
//	@Produce	json
//	@Success	200	{object}	core.ListResult[core.EpgChannel]
//	@Failure	401	{object}	response.ErrorResponse
//	@Failure	500	{object}	response.ErrorResponse
//
//	@Security	JWT
//	@Router		/epg [get]
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

	events, err := s.epg.GetEpg(r.Context(), q)
	if err != nil {
		log.WithError(err).
			Error("failed to get epg events")
		response.InternalErrorCommon(w)
		return
	}

	response.JSON(w, events, 200)
}

// GetEpgEvent godoc
//
//	@Summary	Get a epg event by id
//	@Tags		epg
//	@Param		id	path	string	true	"Event id"
//
//	@Produce	json
//	@Success	200	{object}	core.EpgEvent
//	@Failure	401	{object}	response.ErrorResponse
//	@Failure	404	{object}	response.ErrorResponse
//	@Failure	500	{object}	response.ErrorResponse
//
//	@Security	JWT
//	@Router		/epg/events/{id} [get]
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

// GetRelatedEpgEvents godoc
//
//	@Summary	Get related epg events
//	@Tags		epg
//	@Param		id			path	string	true	"Event id"
//	@Param		limit		query	int		false	"Limit"
//	@Param		offset		query	int		false	"Offset"
//	@Param		sort_key	query	string	false	"Sort key"
//	@Param		sort_dir	query	string	false	"Sort direction"
//
//	@Produce	json
//	@Success	200	{array}		core.EpgEvent
//	@Failure	401	{object}	response.ErrorResponse
//	@Failure	500	{object}	response.ErrorResponse
//
//	@Security	JWT
//	@Router		/epg/events/{id}/related [get]
func (s *router) GetRelatedEpgEvents(w http.ResponseWriter, r *http.Request) {
	id, err := request.NumericURLParam(r, "id")
	if err != nil {
		response.BadRequestf(w, "invalid value for parameter 'id'")
		return
	}

	var q core.PaginationSortQueryParams
	if err := request.BindQuery(r, &q); err != nil {
		response.BadRequest(w, err)
		return
	}

	if err := q.Validate(); err != nil {
		response.BadRequest(w, err)
		return
	}

	events, err := s.epg.GetRelatedEvents(r.Context(), id, q)
	if err != nil {
		log.WithError(err).
			Error("failed to get epg events")
		response.InternalErrorCommon(w)
		return
	}

	response.JSON(w, events, 200)
}

// GetEpgContentTypes godoc
//
//	@Summary	Get epg content types
//	@Tags		epg
//
//	@Produce	json
//	@Success	200	{array}		core.EpgContentType
//	@Failure	401	{object}	response.ErrorResponse
//	@Failure	500	{object}	response.ErrorResponse
//
//	@Security	JWT
//	@Router		/epg/content-types [get]
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
