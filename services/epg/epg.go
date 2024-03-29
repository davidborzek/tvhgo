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

	// getEventsSortKeyMapping mapping of EpgEvent model fields
	// to the tvheadend model fields used for sorting.
	getEventsSortKeyMapping = map[string]string{
		"title":         "title",
		"subtitle":      "subtitle",
		"startsAt":      "start",
		"endsAt":        "stop",
		"channelName":   "channelName",
		"channelNumber": "channelNumber",
		"description":   "description",
	}

	// getEpgSortKeyMapping defines the allowed sort keys
	// for GetEpg.
	getEpgSortKeyMapping = map[string]string{
		"channelName": "channelName",
	}
)

func New(tvh tvheadend.Client) core.EpgService {
	return &service{
		tvh: tvh,
	}
}

func (s *service) GetEvents(
	ctx context.Context,
	params core.GetEpgEventsQueryParams,
) (*core.EpgEventsResult, error) {
	q, err := params.MapToTvheadendQuery(getEventsSortKeyMapping)

	if err != nil {
		return nil, err
	}

	var grid tvheadend.EpgEventGrid
	res, err := s.tvh.Exec(ctx, "/api/epg/events/grid", &grid, *q)
	if err != nil {
		return nil, err
	}

	if res.StatusCode >= 400 {
		return nil, ErrRequestFailed
	}

	result := core.BuildEpgEventsResult(grid, params.Offset)
	return &result, nil
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
	event := core.BuildEpgEvent(grid.Entries[0])

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

func (s *service) GetEpg(
	ctx context.Context,
	params core.GetEpgQueryParams,
) ([]*core.EpgChannel, error) {
	q, err := params.MapToTvheadendQuery(getEpgSortKeyMapping)
	if err != nil {
		return nil, err
	}

	q.Limit(0)

	var meta tvheadend.EpgEventGrid
	metaRes, err := s.tvh.Exec(ctx, "/api/epg/events/grid", &meta, *q)
	if err != nil {
		return nil, err
	}

	if metaRes.StatusCode >= 400 {
		return nil, ErrRequestFailed
	}

	q.Limit(meta.Total)

	var grid tvheadend.EpgEventGrid
	res, err := s.tvh.Exec(ctx, "/api/epg/events/grid", &grid, *q)
	if err != nil {
		return nil, err
	}

	if res.StatusCode >= 400 {
		return nil, ErrRequestFailed
	}

	result := core.BuildEpgResult(grid, params.SortQueryParams)
	return result, nil
}

func (s *service) GetRelatedEvents(
	ctx context.Context,
	eventId int64,
	params core.PaginationSortQueryParams,
) (*core.EpgEventsResult, error) {
	q := params.MapToTvheadendQuery(getEventsSortKeyMapping)
	q.SetInt("eventId", eventId)

	var grid tvheadend.EpgEventGrid
	res, err := s.tvh.Exec(ctx, "/api/epg/events/related", &grid, q)
	if err != nil {
		return nil, err
	}

	if res.StatusCode >= 400 {
		return nil, ErrRequestFailed
	}

	result := core.BuildEpgEventsResult(grid, params.Offset)
	return &result, nil
}
