package core

import (
	"context"
	"errors"
	"sort"
	"strconv"

	"github.com/davidborzek/tvhgo/tvheadend"
)

var (
	ErrEpgEventNotFound = errors.New("epg event not found")
)

type (
	// EpgEvent defines a epg event from tvheadend.
	EpgEvent struct {
		ID            int64  `json:"id"`
		AudioDesc     bool   `json:"audioDesc"`
		ChannelID     string `json:"channelId"`
		ChannelName   string `json:"channelName"`
		ChannelNumber int64  `json:"channelNumber"`
		PiconID       int    `json:"piconId"`
		Description   string `json:"description"`
		EndsAt        int64  `json:"endsAt"`
		HD            bool   `json:"hd"`
		NextEventID   int    `json:"nextEventId"`
		StartsAt      int64  `json:"startsAt"`
		Subtitle      string `json:"subtitle"`
		Subtitled     bool   `json:"subtitled"`
		Title         string `json:"title"`
		Widescreen    bool   `json:"widescreen"`
		DvrUUID       string `json:"dvrUuid,omitempty"`
		DvrState      string `json:"dvrState,omitempty"`
	}

	// EpgEventsResult defines a ListResult of epg events.
	EpgEventsResult = ListResult[*EpgEvent]

	// EpgChannel defines a channel with epg events from tvheadend.
	EpgChannel struct {
		ChannelID     string      `json:"channelId"`
		ChannelName   string      `json:"channelName"`
		ChannelNumber int64       `json:"channelNumber"`
		PiconID       int         `json:"piconId"`
		Events        []*EpgEvent `json:"events"`
	}

	// EpgContentType defines a epg content type from tvheadend.
	EpgContentType struct {
		ID   int    `json:"id"`
		Name string `json:"string"`
	}

	// GetEpgEventsQueryParams defines query params
	// to paginate, sort and filter the epg.
	GetEpgEventsQueryParams struct {
		PaginationSortQueryParams
		Title      string `schema:"title"`
		FullText   bool   `schema:"fullText"`
		Language   string `schema:"lang"`
		NowPlaying bool   `schema:"nowPlaying"`
		// Channel name or id of the channel.
		Channel     string `schema:"channel"`
		ContentType string `schema:"contentType"`
		DurationMin int64  `schema:"durationMin"`
		DurationMax int64  `schema:"durationMax"`
		StartsAt    int64  `schema:"startsAt"`
		EndsAt      int64  `schema:"endsAt"`
	}

	GetEpgQueryParams struct {
		SortQueryParams
		StartsAt int64 `schema:"startsAt"`
		EndsAt   int64 `schema:"endsAt"`
	}

	// EpgService provides access to epg
	// resources from the tvheadend server.
	EpgService interface {
		// GetEpg returns the epg (events for each channel).
		GetEpg(ctx context.Context, params GetEpgQueryParams) ([]*EpgChannel, error)

		// GetEvents returns a list of epg events.
		GetEvents(ctx context.Context, params GetEpgEventsQueryParams) (*EpgEventsResult, error)

		// GetEvent returns a epg event.
		GetEvent(ctx context.Context, id int64) (*EpgEvent, error)

		// GetRelatedEvents returns a list of epg related events for a given event.
		GetRelatedEvents(ctx context.Context, eventId int64, params PaginationSortQueryParams) (*EpgEventsResult, error)

		// GetContentTypes returns a list of epg content types.
		GetContentTypes(ctx context.Context) ([]*EpgContentType, error)
	}
)

// MapToTvheadendQuery maps a GetEpgEventsQueryParams model to a tvheadend
// query model.
func (p *GetEpgEventsQueryParams) MapToTvheadendQuery(sortKeyMapping map[string]string) (*tvheadend.Query, error) {
	q := p.PaginationSortQueryParams.MapToTvheadendQuery(sortKeyMapping)

	if p.Title != "" {
		q.Set("title", p.Title)
	}

	if p.FullText {
		q.Set("fulltext", "1")
	} else {
		q.Set("fulltext", "0")
	}

	if p.Language != "" {
		q.Set("lang", p.Language)
	}

	if p.NowPlaying {
		q.Set("mode", "now")
	} else {
		q.Set("mode", "all")
	}

	if p.Channel != "" {
		q.Set("channel", p.Channel)
	}

	if p.ContentType != "" {
		q.Set("contentType", p.ContentType)
	}

	if p.DurationMin > 0 {
		q.Set("durationMin", strconv.FormatInt(p.DurationMin, 10))
	}

	if p.DurationMax > 0 {
		q.Set("durationMax", strconv.FormatInt(p.DurationMax, 10))
	}

	filter := mapTimeRangeToTvheadendFilter(p.StartsAt, p.EndsAt)
	if len(filter) > 0 {
		if err := q.Filter(filter); err != nil {
			return nil, err
		}
	}

	return &q, nil
}

// MapToTvheadendQuery maps a GetEpgQueryParams model to a tvheadend
// query model.
func (p *GetEpgQueryParams) MapToTvheadendQuery(sortKeyMapping map[string]string) (*tvheadend.Query, error) {
	q := p.SortQueryParams.MapToTvheadendQuery(sortKeyMapping)

	filter := mapTimeRangeToTvheadendFilter(p.StartsAt, p.EndsAt)

	if len(filter) > 0 {
		if err := q.Filter(filter); err != nil {
			return nil, err
		}
	}

	return &q, nil
}

func mapTimeRangeToTvheadendFilter(startsAt int64, endsAt int64) []tvheadend.FilterQuery {
	var filter []tvheadend.FilterQuery
	if startsAt > 0 {
		filter = append(filter, tvheadend.FilterQuery{
			Field:      "start",
			Type:       "numeric",
			Value:      startsAt,
			Comparison: "gt",
		})
	}

	if endsAt > 0 {
		filter = append(filter, tvheadend.FilterQuery{
			Field:      "start",
			Type:       "numeric",
			Value:      endsAt,
			Comparison: "lt",
		})
	}

	return filter
}

// BuildEpgEvent maps a epg grid event entry
// from Tvheadend to a EpgEvent model.
func BuildEpgEvent(src tvheadend.EpgEventGridEntry) EpgEvent {
	channelNumber, _ := strconv.ParseInt(src.ChannelNumber, 10, 0)

	return EpgEvent{
		AudioDesc:     src.AudioDesc == 1,
		ChannelID:     src.ChannelUUID,
		ChannelName:   src.ChannelName,
		ChannelNumber: channelNumber,
		PiconID:       MapTvheadendIconUrlToPiconID(src.ChannelIcon),
		Description:   src.Description,
		EndsAt:        src.Stop,
		HD:            src.HD == 1,
		ID:            src.EventID,
		NextEventID:   src.NextEventID,
		StartsAt:      src.Start,
		Subtitle:      src.Subtitle,
		Subtitled:     src.Subtitled == 1,
		Title:         src.Title,
		Widescreen:    src.Widescreen == 1,
		DvrUUID:       src.DvrUUID,
		DvrState:      src.DvrState,
	}
}

// BuildEpgEventsResult builds the EpgEventsResult model for a given
// tvheadend.EpgEventGrid.
func BuildEpgEventsResult(src tvheadend.EpgEventGrid, offset int64) EpgEventsResult {
	events := make([]*EpgEvent, 0)
	for _, entry := range src.Entries {
		e := BuildEpgEvent(entry)
		events = append(events, &e)
	}

	result := EpgEventsResult{
		Entries: events,
		Total:   src.Total,
		Offset:  offset,
	}

	return result
}

// BuildEpgResult builds the epg result for a given tvheadend.EpgEventGrid
// and sorts th channels by the given SortQueryParams.
func BuildEpgResult(grid tvheadend.EpgEventGrid, params SortQueryParams) []*EpgChannel {
	channels := make([]*EpgChannel, 0)

	for _, entry := range grid.Entries {
		event := BuildEpgEvent(entry)

		index, ok := getChannelIndex(channels, event.ChannelID)
		if ok {
			channels[index].Events = append(channels[index].Events, &event)
		} else {
			channels = append(channels, &EpgChannel{
				ChannelID:     event.ChannelID,
				ChannelName:   event.ChannelName,
				ChannelNumber: event.ChannelNumber,
				PiconID:       event.PiconID,
				Events: []*EpgEvent{
					&event,
				},
			})
		}
	}

	sortEpgChannels(channels, params)
	return channels
}

// sortEpgChannels sorts an array of EpgChannel by the given
// SortQueryParams.
func sortEpgChannels(channels []*EpgChannel, params SortQueryParams) {
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
}

// A helper function to find the index of channel in an array of channels by its id.
func getChannelIndex(channels []*EpgChannel, channelId string) (int, bool) {
	for i, c := range channels {
		if c.ChannelID == channelId {
			return i, true
		}
	}
	return 0, false
}
