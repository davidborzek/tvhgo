package picon

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/davidborzek/tvhgo/core"
	"github.com/davidborzek/tvhgo/tvheadend"
)

var (
	ErrRequestFailed = errors.New("picon request failed")
)

type service struct {
	tvh tvheadend.Client
}

func New(tvh tvheadend.Client) core.PiconService {
	return &service{
		tvh: tvh,
	}
}

func (s *service) Get(ctx context.Context, id int) ([]byte, error) {
	res, err := s.tvh.Exec(ctx, fmt.Sprintf("/imagecache/%d", id), nil)
	if err != nil {
		return nil, err
	}

	if res.StatusCode == 404 {
		return nil, core.ErrPiconNotFound
	}

	if res.StatusCode >= 400 {
		return nil, ErrRequestFailed
	}

	picon, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return picon, nil
}
