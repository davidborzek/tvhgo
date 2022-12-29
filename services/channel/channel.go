package channel

import (
	"context"
	"errors"

	"github.com/davidborzek/tvhgo/core"
	"github.com/davidborzek/tvhgo/tvheadend"
)

type service struct {
	tvh tvheadend.Client
}

var (
	ErrRequestFailed = errors.New("channel request failed")

	// sortKeyMapping mapping of Channel model fields
	// to the tvheadend model fields used for sorting.
	sortKeyMapping = map[string]string{
		"name":   "name",
		"number": "number",
	}
)

func New(tvh tvheadend.Client) core.ChannelService {
	return &service{
		tvh: tvh,
	}
}

func (s *service) GetAll(ctx context.Context, params core.PaginationSortQueryParams) ([]*core.Channel, error) {
	q := params.MapToTvheadendQuery(sortKeyMapping)

	var grid tvheadend.ChannelGrid
	res, err := s.tvh.Exec(ctx, "/api/channel/grid", &grid, q)
	if err != nil {
		return nil, err
	}

	if res.StatusCode >= 400 {
		return nil, ErrRequestFailed
	}

	channels := make([]*core.Channel, 0)
	for _, entry := range grid.Entries {
		c := &core.Channel{
			ID:      entry.UUID,
			Name:    entry.Name,
			Enabled: entry.Enabled,
			Number:  entry.Number,
			PiconID: core.MapTvheadendIconUrlToPiconID(entry.IconPublicURL),
		}

		channels = append(channels, c)
	}

	return channels, nil
}

func (s *service) Get(ctx context.Context, id string) (*core.Channel, error) {
	q := tvheadend.NewQuery()
	q.Set("uuid", id)

	var idnodeLoad tvheadend.IdnodeLoadResponse
	res, err := s.tvh.Exec(ctx, "/api/idnode/load", &idnodeLoad, q)
	if err != nil {
		return nil, err
	}

	if res.StatusCode >= 400 {
		return nil, ErrRequestFailed
	}

	if len(idnodeLoad.Entries) == 0 {
		return nil, core.ErrChannelNotFound
	}

	channel, err := core.MapTvheadendIdnodeToChannel(idnodeLoad.Entries[0])
	if err != nil {
		return nil, err
	}

	return channel, nil
}
