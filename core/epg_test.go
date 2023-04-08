package core_test

import (
	"encoding/json"
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
		ChannelIcon:   "imagecache/7",
	}

	tvhEventSameChannel = tvheadend.EpgEventGridEntry{
		AudioDesc:     1,
		ChannelUUID:   "someChannelID",
		ChannelName:   "someChannel",
		ChannelNumber: "1",
		Description:   "someDescription",
		Stop:          1234,
		HD:            1,
		EventID:       2,
		NextEventID:   2,
		Start:         123,
		Subtitle:      "someSubtitle",
		Subtitled:     1,
		Title:         "someOtherTitle",
		Widescreen:    1,
		ChannelIcon:   "imagecache/7",
	}

	tvhEventSecondChannel = tvheadend.EpgEventGridEntry{
		AudioDesc:     1,
		ChannelUUID:   "someOtherChannelID",
		ChannelName:   "1someOtherChannel",
		ChannelNumber: "2",
		Description:   "someOtherDescription",
		Stop:          1234,
		HD:            1,
		EventID:       3,
		NextEventID:   2,
		Start:         123,
		Subtitle:      "someSubtitle",
		Subtitled:     1,
		Title:         "someTitle",
		Widescreen:    1,
		ChannelIcon:   "imagecache/8",
	}

	tvhFilter = []tvheadend.FilterQuery{
		{
			Field:      "start",
			Type:       "numeric",
			Value:      20,
			Comparison: "gt",
		},
		{
			Field:      "start",
			Type:       "numeric",
			Value:      40,
			Comparison: "lt",
		},
	}
)

func TestGetEpgEventsQueryParamsMapToTvheadendQuery(t *testing.T) {
	q := core.GetEpgEventsQueryParams{
		Title:       "someTitle",
		FullText:    true,
		Language:    "de",
		NowPlaying:  true,
		Channel:     "someChannel",
		ContentType: "someContentType",
		DurationMin: 123,
		DurationMax: 1234,
		StartsAt:    20,
		EndsAt:      40,
	}

	m, err := q.MapToTvheadendQuery(map[string]string{})

	assert.Nil(t, err)

	assert.Equal(t, q.Title, m.Get("title"))
	assert.Equal(t, "1", m.Get("fulltext"))
	assert.Equal(t, "de", m.Get("lang"))
	assert.Equal(t, "now", m.Get("mode"))
	assert.Equal(t, q.Channel, m.Get("channel"))
	assert.Equal(t, q.ContentType, m.Get("contentType"))
	assert.Equal(t, "123", m.Get("durationMin"))
	assert.Equal(t, "1234", m.Get("durationMax"))

	filterRaw, _ := json.Marshal(&tvhFilter)
	assert.Equal(t, string(filterRaw), m.Get("filter"))
}

func TestGetEpgEventsQueryParamsMapToTvheadendQueryFalsyBoolean(t *testing.T) {
	q := core.GetEpgEventsQueryParams{
		FullText:   false,
		NowPlaying: false,
	}

	m, err := q.MapToTvheadendQuery(map[string]string{})

	assert.Nil(t, err)
	assert.Equal(t, "0", m.Get("fulltext"))
	assert.Equal(t, "all", m.Get("mode"))
}

func TestBuildEpgEvent(t *testing.T) {
	event := core.BuildEpgEvent(tvhEvent)

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

func TestBuildEpgEventFalsyBoolean(t *testing.T) {
	tvhEvent := tvheadend.EpgEventGridEntry{}

	event := core.BuildEpgEvent(tvhEvent)

	assert.False(t, event.Subtitled)
	assert.False(t, event.AudioDesc)
	assert.False(t, event.Widescreen)
	assert.False(t, event.HD)
}

func TestBuildEpgEventSetChannelNumberToZeroWhenFailure(t *testing.T) {
	tvhEvent := tvheadend.EpgEventGridEntry{
		ChannelNumber: "noNumber",
	}
	event := core.BuildEpgEvent(tvhEvent)
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

func TestBuildEpgResultDefaultSort(t *testing.T) {
	tvhGrid := tvheadend.EpgEventGrid{
		Entries: []tvheadend.EpgEventGridEntry{
			// This one has channel number equals to 2
			tvhEventSecondChannel,
			// These ones have channel number equals to 1
			tvhEvent,
			tvhEventSameChannel,
		},
	}

	result := core.BuildEpgResult(tvhGrid, core.SortQueryParams{})

	// Verify the first channel in the result with channel number 1.
	// Should have been sorted as first entry.
	assert.Len(t, result, 2)
	assert.Equal(t, result[0].ChannelID, tvhEvent.ChannelUUID)
	assert.Equal(t, result[0].ChannelName, tvhEvent.ChannelName)
	assert.Equal(t, result[0].ChannelNumber, int64(1))
	assert.Equal(t, result[0].PiconID, 7)
	assert.Len(t, result[0].Events, 2)

	// Verify the second channel in the result with channel number 2.
	assert.Equal(t, result[1].ChannelID, tvhEventSecondChannel.ChannelUUID)
	assert.Len(t, result[1].Events, 1)
}

func TestBuildEpgResultSortChannelName(t *testing.T) {
	tvhGrid := tvheadend.EpgEventGrid{
		Entries: []tvheadend.EpgEventGridEntry{
			tvhEvent,
			tvhEventSameChannel,
			tvhEventSecondChannel,
		},
	}

	result := core.BuildEpgResult(tvhGrid, core.SortQueryParams{
		SortKey: "channelName",
	})

	// Verify the channel sorted by channel name is first in the result.
	assert.Len(t, result, 2)
	assert.Equal(t, result[0].ChannelID, tvhEventSecondChannel.ChannelUUID)
	assert.Equal(t, result[0].ChannelName, tvhEventSecondChannel.ChannelName)
	assert.Equal(t, result[0].ChannelNumber, int64(2))
	assert.Equal(t, result[0].PiconID, 8)
	assert.Len(t, result[0].Events, 1)

	assert.Equal(t, result[1].ChannelID, tvhEvent.ChannelUUID)
	assert.Len(t, result[1].Events, 2)
}
