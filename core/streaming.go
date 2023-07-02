package core

import (
	"context"
	"net/http"
)

type (
	StreamingService interface {
		// GetChannelStream returns a raw http response of the channel stream.
		GetChannelStream(ctx context.Context, channelNumber int64, profile string) (*http.Response, error)

		// GetRecordingStream returns a raw http response of the recording stream.
		GetRecordingStream(ctx context.Context, recordingId string) (*http.Response, error)
	}
)
