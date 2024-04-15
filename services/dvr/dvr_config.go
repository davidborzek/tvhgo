package dvr

import (
	"context"
	"errors"

	"github.com/davidborzek/tvhgo/core"
	"github.com/davidborzek/tvhgo/tvheadend"
)

var (
	ErrRequestFailed = errors.New("dvr config request failed")
)

type service struct {
	tvh tvheadend.Client
}

func New(tvh tvheadend.Client) core.DVRConfigService {
	return &service{
		tvh: tvh,
	}
}

func (s *service) GetAll(ctx context.Context) ([]core.DVRConfig, error) {
	q := tvheadend.NewQuery()
	q.SortKey("name")
	q.SortDir("asc")

	var grid tvheadend.ListResponse[tvheadend.DVRConfig]
	res, err := s.tvh.Exec(ctx, "/api/dvr/config/grid", &grid, q)
	if err != nil {
		return nil, err
	}

	if res.StatusCode >= 400 {
		return nil, ErrRequestFailed
	}

	configs := make([]core.DVRConfig, 0)
	for _, entry := range grid.Entries {
		configs = append(configs, core.NewDVRConfig(entry))
	}

	return configs, nil
}

func (s *service) Delete(ctx context.Context, id string) error {
	ok, err := s.isDvrConfig(ctx, []string{id})
	if err != nil {
		return err
	}

	if !ok {
		return core.ErrDVRConfigNotFound
	}

	q := tvheadend.NewQuery()
	q.Set("uuid", id)

	res, err := s.tvh.Exec(ctx, "/api/idnode/delete", nil, q)
	if err != nil {
		return err
	}

	if res.StatusCode >= 400 {
		return ErrRequestFailed
	}

	return nil
}

func (s *service) isDvrConfig(ctx context.Context, ids []string) (bool, error) {
	configs, err := s.GetAll(ctx)
	if err != nil {
		return false, err
	}

	for _, id := range ids {
		if !contains(configs, id) {
			return false, nil
		}
	}

	return true, nil
}

func contains(arr []core.DVRConfig, val string) bool {
	for _, a := range arr {
		if a.ID == val {
			return !a.Original
		}
	}

	return false
}
