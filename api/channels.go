package api

import (
	"net/http"

	"github.com/davidborzek/tvhgo/api/request"
	"github.com/davidborzek/tvhgo/api/response"
	"github.com/davidborzek/tvhgo/core"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

// GetChannels godoc
//
//	@Summary	Get list of channels
//	@Tags		channels
//
//	@Param		limit		query	int		false	"Limit"
//	@Param		offset		query	int		false	"Offset"
//	@Param		sort_key	query	string	false	"Sort key"
//	@Param		sort_dir	query	string	false	"Sort direction"
//
//	@Produce	json
//	@Success	200	{array}		core.Channel
//	@Failure	400	{object}	response.ErrorResponse
//	@Failure	401	{object}	response.ErrorResponse
//	@Failure	500	{object}	response.ErrorResponse
//	@Security	JWT
//	@Router		/channels [get]
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

// GetChannel godoc
//
//	@Summary	Get a channel by id
//	@Tags		channels
//	@Param		id	path	string	true	"Channel id"
//
//	@Produce	json
//	@Success	200	{object}	core.Channel
//	@Failure	404	{object}	response.ErrorResponse
//	@Failure	401	{object}	response.ErrorResponse
//	@Failure	500	{object}	response.ErrorResponse
//
//	@Security	JWT
//	@Router		/channels/{id} [get]
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
