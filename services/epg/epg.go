package epg

import (
	"context"
	"errors"
	"sort"

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

func (s *service) GetEvents(ctx context.Context, params core.GetEpgQueryParams) (*core.EpgEventsResult, error) {
	q, err := params.MapToTvheadendQuery(sortKeyMapping)

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

func (s *service) GetChannelEvents(ctx context.Context, params core.GetEpgChannelEventsQueryParams) (*core.EpgChannelEventsResult, error) {
	q, err := params.MapToTvheadendQuery(sortKeyMapping)
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

	result := buildEpgChannelEventsResult(grid, params.Offset, params.SortQueryParams)
	return &result, nil
}

func getChannelIndex(channels []*core.EpgChannel, channelId string) (int, bool) {
	for i, c := range channels {
		if c.ChannelID == channelId {
			return i, true
		}
	}
	return 0, false
}

func buildEpgChannelEventsResult(grid tvheadend.EpgEventGrid, offset int64, params core.SortQueryParams) core.EpgChannelEventsResult {
	channels := make([]*core.EpgChannel, 0)

	for _, entry := range grid.Entries {
		event := core.MapTvheadendEpgEventToEpgEvent(entry)

		index, ok := getChannelIndex(channels, event.ChannelID)
		if ok {
			channels[index].Events = append(channels[index].Events, &event)
		} else {
			channels = append(channels, &core.EpgChannel{
				ChannelID:     event.ChannelID,
				ChannelName:   event.ChannelName,
				ChannelNumber: event.ChannelNumber,
				PiconID:       event.PiconID,
				Events: []*core.EpgEvent{
					&event,
				},
			})
		}
	}

	sort.SliceStable(channels, func(i, j int) bool {
		switch params.SortKey {
		case "channelName":
			{
				if params.SortDirection == "desc" {
					return channels[i].ChannelName > channels[j].ChannelName
				}

				return channels[i].ChannelName < channels[j].ChannelName
			}
		}

		if params.SortDirection == "desc" {
			return channels[i].ChannelNumber > channels[j].ChannelNumber
		}

		return channels[i].ChannelNumber < channels[j].ChannelNumber
	})

	return core.EpgChannelEventsResult{
		Entries: channels,
		Total:   grid.Total,
		Offset:  offset,
	}
}
