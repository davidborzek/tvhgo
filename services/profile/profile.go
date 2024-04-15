package profiles

import (
	"context"
	"errors"

	"github.com/davidborzek/tvhgo/core"
	"github.com/davidborzek/tvhgo/tvheadend"
)

var (
	ErrRequestFailed = errors.New("profile request failed")
)

type service struct {
	tvh tvheadend.Client
}

func New(tvh tvheadend.Client) core.ProfileService {
	return &service{
		tvh: tvh,
	}
}

func (s *service) GetStreamProfiles(ctx context.Context) ([]core.StreamProfile, error) {
	var list tvheadend.ListResponse[tvheadend.StreamProfile]
	res, err := s.tvh.Exec(ctx, "/api/profile/list", &list)
	if err != nil {
		return nil, err
	}

	if res.StatusCode >= 400 {
		return nil, ErrRequestFailed
	}

	profiles := make([]core.StreamProfile, 0)
	for _, entry := range list.Entries {
		profiles = append(profiles, core.NewStreamProfile(entry))
	}

	return profiles, nil
}
