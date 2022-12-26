package core_test

import (
	"testing"

	"github.com/davidborzek/tvhgo/core"
	"github.com/davidborzek/tvhgo/tvheadend"
	"github.com/stretchr/testify/assert"
)

var (
	tvhEvent = tvheadend.EpgEventGridEntry{
		AudioDesc:     1,
		ChannelUUID:   "someChannelID",
		ChannelName:   "someChannel",
		ChannelNumber: "1",
		Description:   "someDescription",
		Stop:          1234,
		HD:            1,
		EventID:       1,
		NextEventID:   2,
		Start:         123,
		Subtitle:      "someSubtitle",
		Subtitled:     1,
		Title:         "someTitle",
		Widescreen:    1,
	}
)

func TestGetEpgQueryParamsMapToTvheadendQuery(t *testing.T) {
	q := core.GetEpgQueryParams{
		Title:       "someTitle",
		FullText:    true,
		Language:    "de",
		NowPlaying:  true,
		Channel:     "someChannel",
		ContentType: "someContentType",
		DurationMin: 123,
		DurationMax: 1234,
	}

	m := q.MapToTvheadendQuery(map[string]string{})

	assert.Equal(t, q.Title, m.Get("title"))
	assert.Equal(t, "1", m.Get("fulltext"))
	assert.Equal(t, "de", m.Get("lang"))
	assert.Equal(t, "now", m.Get("mode"))
	assert.Equal(t, q.Channel, m.Get("channel"))
	assert.Equal(t, q.ContentType, m.Get("contentType"))
	assert.Equal(t, "123", m.Get("durationMin"))
	assert.Equal(t, "1234", m.Get("durationMax"))
}

func TestGetEpgQueryParamsMapToTvheadendQueryFalsyBoolean(t *testing.T) {
	q := core.GetEpgQueryParams{
		FullText:   false,
		NowPlaying: false,
	}

	m := q.MapToTvheadendQuery(map[string]string{})

	assert.Equal(t, "0", m.Get("fulltext"))
	assert.Equal(t, "all", m.Get("mode"))
}

func TestMapTvheadendEpgEventToEpgEvent(t *testing.T) {
	event := core.MapTvheadendEpgEventToEpgEvent(tvhEvent)

	assert.True(t, event.Subtitled)
	assert.True(t, event.AudioDesc)
	assert.True(t, event.Widescreen)
	assert.True(t, event.HD)

	assert.Equal(t, tvhEvent.ChannelUUID, event.ChannelID)
	assert.Equal(t, tvhEvent.ChannelName, event.ChannelName)
	assert.Equal(t, int64(1), event.ChannelNumber)
	assert.Equal(t, tvhEvent.Description, event.Description)
	assert.Equal(t, tvhEvent.Stop, event.EndsAt)
	assert.Equal(t, tvhEvent.Start, event.StartsAt)
	assert.Equal(t, tvhEvent.EventID, event.ID)
	assert.Equal(t, tvhEvent.NextEventID, event.NextEventID)
	assert.Equal(t, tvhEvent.Subtitle, event.Subtitle)
	assert.Equal(t, tvhEvent.Title, event.Title)
}

func TestMapTvheadendEpgEventToEpgEventFalsyBoolean(t *testing.T) {
	tvhEvent := tvheadend.EpgEventGridEntry{}

	event := core.MapTvheadendEpgEventToEpgEvent(tvhEvent)

	assert.False(t, event.Subtitled)
	assert.False(t, event.AudioDesc)
	assert.False(t, event.Widescreen)
	assert.False(t, event.HD)
}

func TestMapTvheadendEpgEventToEpgEventSetChannelNumberToZeroWhenFailure(t *testing.T) {
	tvhEvent := tvheadend.EpgEventGridEntry{
		ChannelNumber: "noNumber",
	}
	event := core.MapTvheadendEpgEventToEpgEvent(tvhEvent)
	assert.Equal(t, int64(0), event.ChannelNumber)
}

func TestBuildEpgEventsResult(t *testing.T) {
	total := int64(20)
	offset := int64(10)

	tvhGrid := tvheadend.EpgEventGrid{
		Entries: []tvheadend.EpgEventGridEntry{
			tvhEvent,
		},
		Total: total,
	}

	result := core.BuildEpgEventsResult(
		tvhGrid, offset,
	)

	assert.Len(t, result.Entries, 1)
	assert.Equal(t, offset, result.Offset)
	assert.Equal(t, total, result.Total)
}
