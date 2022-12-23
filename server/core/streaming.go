package core

import (
	"context"
	"net/http"
)

type (
	StreamingService interface {
		// GetChannelStream returns the raw http response the channel stream.
		GetChannelStream(ctx context.Context, channelNumber int64) (*http.Response, error)
	}
)
