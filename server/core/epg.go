package core

import (
	"context"
	"errors"
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
	}

	// EpgEventsResult defines a ListResult of epg events.
	EpgEventsResult = ListResult[*EpgEvent]

	// EpgContentType defines a epg content type from tvheadend.
	EpgContentType struct {
		ID   int    `json:"id"`
		Name string `json:"string"`
	}

	// GetEpgQueryParams defines query params
	// to paginate, sort and filter the epg.
	GetEpgQueryParams struct {
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
	}

	// GetEpgTimelineQueryParams defines query params
	// to paginate, sort and filter the epg timline.
	GetEpgTimelineQueryParams struct {
		PaginationSortQueryParams
	}

	// EpgService provides access to epg
	// resources from the tvheadend server.
	EpgService interface {
		// GetEvents returns a list of epg events.
		GetEvents(ctx context.Context, params GetEpgQueryParams) (*EpgEventsResult, error)

		// GetEvent returns a epg event.
		GetEvent(ctx context.Context, id int64) (*EpgEvent, error)

		// GetContentTypes returns a list of epg content types.
		GetContentTypes(ctx context.Context) ([]*EpgContentType, error)
	}
)

// MapToTvheadendQuery maps a GetEpgQueryParams model to a tvheadend
// query model.
func (p *GetEpgQueryParams) MapToTvheadendQuery(sortKeyMapping map[string]string) tvheadend.Query {
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

	return q
}

// MapTvheadendEpgEventToEpgEvent maps a epg grid event entry
// from Tvheadend to a EpgEvent model.
func MapTvheadendEpgEventToEpgEvent(src tvheadend.EpgEventGridEntry) EpgEvent {
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
	}
}

func BuildEpgEventsResult(src tvheadend.EpgEventGrid, offset int64) EpgEventsResult {
	events := make([]*EpgEvent, 0)
	for _, entry := range src.Entries {
		e := MapTvheadendEpgEventToEpgEvent(entry)
		events = append(events, &e)
	}

	result := EpgEventsResult{
		Entries: events,
		Total:   src.Total,
		Offset:  offset,
	}

	return result
}
