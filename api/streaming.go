package api

import (
	"context"
	"errors"
	"net/http"
	"syscall"

	"github.com/davidborzek/tvhgo/api/request"
	"github.com/davidborzek/tvhgo/api/response"
	log "github.com/sirupsen/logrus"
)

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
