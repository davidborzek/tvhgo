package api

import (
	"context"
	"errors"
	"net/http"
	"syscall"

	"github.com/davidborzek/tvhgo/api/request"
	"github.com/davidborzek/tvhgo/api/response"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

// StreamChannel godoc
//
//	@Summary	Stream a channel by channel number
//	@Tags		channels
//	@Param		number	path	string	true	"Channel number"
//	@Produce	video/*
//	@Produce	json
//	@Success	200
//	@Failure	401	{object}	response.ErrorResponse
//	@Security	JWT
//	@Router		/channels/{number}/stream [get]
func (s *router) StreamChannel(w http.ResponseWriter, r *http.Request) {
	number, err := request.NumericURLParam(r, "number")

	if err != nil {
		response.BadRequestf(w, "invalid value for parameter 'number'")
		return
	}

	res, err := s.streaming.GetChannelStream(context.Background(), number)
	if err != nil {
		log.WithError(err).
			WithField("channel", number).
			Error("failed to get channel stream")
		response.InternalErrorCommon(w)
		return
	}

	if _, err := response.CopyResponse(w, res); err != nil {
		if errors.Is(err, syscall.EPIPE) {
			return
		}

		log.WithError(err).
			WithField("channel", number).
			Error("unexpected error occurred during streaming the channel")
	}
}

// StreamRecording godoc
//
//	@Summary	Stream a recording
//	@Tags		recordings
//	@Param		id	path	string	true	"Recording id"
//	@Produce	video/*
//	@Produce	json
//	@Success	200
//	@Failure	401	{object}	response.ErrorResponse
//	@Security	JWT
//	@Router		/recordings/{id}/stream [get]
func (s *router) StreamRecording(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	res, err := s.streaming.GetRecordingStream(context.Background(), id)
	if err != nil {
		log.WithError(err).
			WithField("id", id).
			Error("failed to get recording stream")
		response.InternalErrorCommon(w)
		return
	}

	if _, err := response.CopyResponse(w, res); err != nil {
		if errors.Is(err, syscall.EPIPE) {
			return
		}

		log.WithError(err).
			WithField("recording", id).
			Error("unexpected error occurred during streaming the recording")
	}
}
