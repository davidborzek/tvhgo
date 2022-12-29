package picon

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

func New(tvh tvheadend.Client) core.PiconService {
	return &service{
		tvh: tvh,
	}
}

func (s *service) Get(ctx context.Context, id int) (*http.Response, error) {
	res, err := s.tvh.Exec(ctx, fmt.Sprintf("/imagecache/%d", id), nil)
	if err != nil {
		return nil, err
	}

	return res.Response, nil
}
