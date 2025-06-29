package streaming

import (
	"context"
	"fmt"
	"net/http"

	"github.com/davidborzek/tvhgo/core"
	"github.com/davidborzek/tvhgo/tvheadend"
	"github.com/rs/zerolog/log"
)

type service struct {
	tvh tvheadend.Client
}

func New(tvh tvheadend.Client) core.StreamingService {
	return &service{
		tvh: tvh,
	}
}

func (s *service) GetChannelStream(
	ctx context.Context,
	channelNumber int64,
	profile string,
) (*http.Response, error) {
	q := tvheadend.NewQuery()

	if profile != "" {
		q.Set("profile", profile)
	}

	// Log the URL being called for debugging
	url := fmt.Sprintf("/stream/channelnumber/%d", channelNumber)
	log.Info().
		Int64("channelNumber", channelNumber).
		Str("profile", profile).
		Str("url", url).
		Msg("streaming channel from tvheadend")

	res, err := s.tvh.Exec(ctx, url, nil, q)
	if err != nil {
		return nil, err
	}

	if res.StatusCode >= 400 {
		return nil, fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}

	return res.Response, nil
}

func (s *service) GetRecordingStream(
	ctx context.Context,
	recordingId string,
) (*http.Response, error) {
	res, err := s.tvh.Exec(ctx, fmt.Sprintf("/dvrfile/%s", recordingId), nil)
	if err != nil {
		return nil, err
	}

	if res.StatusCode >= 400 {
		return nil, fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}

	return res.Response, nil
}
