package epg

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
	ErrRequestFailed = errors.New("epg request failed")

	// sortKeyMapping mapping of EpgEvent model fields
	// to the tvheadend model fields used for sorting.
	sortKeyMapping = map[string]string{
		"title":         "title",
		"subtitle":      "subtitle",
		"startsAt":      "start",
		"endsAt":        "stop",
		"channelName":   "channelName",
		"channelNumber": "channelNumber",
		"description":   "description",
	}
)

func New(tvh tvheadend.Client) core.EpgService {
	return &service{
		tvh: tvh,
	}
}

func (s *service) GetEvents(ctx context.Context, params core.GetEpgQueryParams) ([]*core.EpgEvent, error) {
	q := params.MapToTvheadendQuery(sortKeyMapping)

	var grid tvheadend.EpgEventGrid
	res, err := s.tvh.Exec(ctx, "/api/epg/events/grid", &grid, q)
	if err != nil {
		return nil, err
	}

	if res.StatusCode >= 400 {
		return nil, ErrRequestFailed
	}

	events := make([]*core.EpgEvent, 0)
	for _, entry := range grid.Entries {
		e := core.MapTvheadendEpgEventToEpgEvent(entry)
		events = append(events, &e)
	}

	return events, nil
}

func (s *service) GetEvent(ctx context.Context, id int64) (*core.EpgEvent, error) {
	q := tvheadend.NewQuery()
	q.SetInt("eventId", id)

	var grid tvheadend.EpgEventGrid
	res, err := s.tvh.Exec(ctx, "/api/epg/events/load", &grid, q)
	if err != nil {
		return nil, err
	}

	if res.StatusCode >= 400 {
		return nil, ErrRequestFailed
	}

	if len(grid.Entries) == 0 {
		return nil, core.ErrEpgEventNotFound
	}
	event := core.MapTvheadendEpgEventToEpgEvent(grid.Entries[0])

	return &event, nil
}

func (s *service) GetContentTypes(ctx context.Context) ([]*core.EpgContentType, error) {
	var list tvheadend.EpgContentTypeResponse
	res, err := s.tvh.Exec(ctx, "/api/epg/content_type/list", &list)
	if err != nil {
		return nil, err
	}

	if res.StatusCode >= 400 {
		return nil, ErrRequestFailed
	}

	contentTypes := make([]*core.EpgContentType, 0)
	for _, entry := range list.Entries {
		c := core.EpgContentType{
			ID:   entry.Key,
			Name: entry.Val,
		}

		contentTypes = append(contentTypes, &c)
	}

	return contentTypes, nil
}
