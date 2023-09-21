package streaming

import (
	"context"
	"fmt"
	"net/http"

	"github.com/davidborzek/tvhgo/core"
	"github.com/davidborzek/tvhgo/tvheadend"
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

	res, err := s.tvh.Exec(ctx, fmt.Sprintf("/stream/channelnumber/%d", channelNumber), nil, q)
	if err != nil {
		return nil, err
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

	return res.Response, nil
}
