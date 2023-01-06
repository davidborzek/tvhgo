package core_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/davidborzek/tvhgo/conv"
	"github.com/davidborzek/tvhgo/core"
	"github.com/davidborzek/tvhgo/tvheadend"
	"github.com/stretchr/testify/assert"
)

func TestGetRecordingsParamsValidate(t *testing.T) {
	p := core.GetRecordingsParams{
		Status: "upcoming",
	}

	err := p.Validate()
	assert.Nil(t, err)
}

func TestGetRecordingsParamsValidateEmptyStatus(t *testing.T) {
	p := core.GetRecordingsParams{}

	err := p.Validate()
	assert.Nil(t, err)
}

func TestGetRecordingsParamsValidateInvalidStatus(t *testing.T) {
	p := core.GetRecordingsParams{
		Status: "invalidStatus",
	}

	err := p.Validate()
	assert.Equal(t, core.ErrGetRecordingsInvalidStatus, err)
}

func TestCreateRecordingByEventValidate(t *testing.T) {
	c := core.CreateRecordingByEvent{
		EventID:  123,
		ConfigID: "someConfigID",
	}

	err := c.Validate()
	assert.Nil(t, err)
}

func TestCreateRecordingByEventValidateReturnError(t *testing.T) {
	c := core.CreateRecordingByEvent{}

	err := c.Validate()
	assert.Equal(t, err, core.ErrCreateRecordingInvalidEventID)
}

func TestCreateRecordingValidate(t *testing.T) {
	c := core.CreateRecording{
		Title:     "someTitle",
		ChannelID: "someChannelID",
		StartsAt:  time.Now().Unix() + 10000,
		EndsAt:    time.Now().Unix() + 20000,
	}

	err := c.Validate()
	assert.Nil(t, err)
}

func TestCreateRecordingValidateReturnErrorInvalidTitle(t *testing.T) {
	c := core.CreateRecording{}

	err := c.Validate()
	assert.Equal(t, err, core.ErrRecordingInvalidTitle)
}

func TestCreateRecordingValidateReturnErrorInvalidChannelID(t *testing.T) {
	c := core.CreateRecording{
		Title: "someTitle",
	}

	err := c.Validate()
	assert.Equal(t, err, core.ErrRecordingInvalidChannelID)
}

func TestCreateRecordingValidateReturnErrorInvalidStartDate(t *testing.T) {
	c := core.CreateRecording{
		Title:     "someTitle",
		ChannelID: "someChannelID",
		StartsAt:  time.Now().Unix() - 10000,
	}

	err := c.Validate()
	assert.Equal(t, err, core.ErrRecordingInvalidStartDate)
}

func TestCreateRecordingValidateReturnErrorInvalidEndDate(t *testing.T) {
	c := core.CreateRecording{
		Title:     "someTitle",
		ChannelID: "someChannelID",
		StartsAt:  time.Now().Unix(),
		EndsAt:    time.Now().Unix() - 10000,
	}

	err := c.Validate()
	assert.Equal(t, err, core.ErrRecordingInvalidEndDate)
}

func TestUpdateRecordingValidate(t *testing.T) {
	title := "someTitle"
	startsAt := time.Now().Unix() + 10000
	endsAt := time.Now().Unix() + 20000

	c := core.UpdateRecording{
		Title:    &title,
		StartsAt: &startsAt,
		EndsAt:   &endsAt,
	}

	err := c.Validate()
	assert.Nil(t, err)
}

func TestUpdateRecordingValidateEmpty(t *testing.T) {
	c := core.UpdateRecording{}

	err := c.Validate()
	assert.Nil(t, err)
}

func TestUpdateRecordingValidateReturnsErrorInvalidTitle(t *testing.T) {
	title := ""

	c := core.UpdateRecording{
		Title: &title,
	}

	err := c.Validate()
	assert.Equal(t, core.ErrRecordingInvalidTitle, err)
}

func TestUpdateRecordingValidateReturnsErrorInvalidStartDate(t *testing.T) {
	title := "title"
	startsAt := int64(0)

	c := core.UpdateRecording{
		Title:    &title,
		StartsAt: &startsAt,
	}

	err := c.Validate()
	assert.Equal(t, core.ErrRecordingInvalidStartDate, err)
}

func TestUpdateRecordingValidateReturnsErrorInvalidEndDate(t *testing.T) {
	title := "title"
	endsAt := int64(0)

	c := core.UpdateRecording{
		Title:  &title,
		EndsAt: &endsAt,
	}

	err := c.Validate()
	assert.Equal(t, core.ErrRecordingInvalidEndDate, err)
}

func TestCreateRecordingMapToTvheadendOpts(t *testing.T) {
	c := core.CreateRecording{
		Title:        "someTitle",
		ExtraText:    "someExtraText",
		ChannelID:    "someChannelID",
		StartsAt:     time.Now().Unix(),
		EndsAt:       time.Now().Unix() + 10000,
		Comment:      "someComment",
		StartPadding: 1,
		EndPadding:   1,
		Priority:     1,
		ConfigID:     "someConfigID",
	}

	mapped := c.MapToTvheadendOpts()

	assert.Equal(t, c.Title, mapped.DispTitle)
	assert.Equal(t, c.ExtraText, mapped.DispExtratext)
	assert.Equal(t, c.ChannelID, mapped.Channel)
	assert.Equal(t, c.Comment, mapped.Comment)
	assert.Equal(t, c.ConfigID, mapped.ConfigName)
	assert.Equal(t, c.Priority, mapped.Pri)
	assert.Equal(t, c.StartsAt, mapped.Start)
	assert.Equal(t, c.StartPadding, mapped.StartExtra)
	assert.Equal(t, c.EndsAt, mapped.Stop)
	assert.Equal(t, c.EndPadding, mapped.StopExtra)
}

func TestMapToTvheadendDvrGridEntryToRecording(t *testing.T) {
	piconId := 6
	tvhEntry := tvheadend.DvrGridEntry{
		Channel:     "someChannelID",
		Channelname: "someChannelName",
		Create:      time.Now().Unix(),
		Duration:    1234,
		Enabled:     true,
		StopExtra:   1,
		Stop:        time.Now().Unix() + 10000,
		Filename:    "/some/file/name",
		UUID:        "someID",
		Title: map[string]string{
			"de": "someTitle",
		},
		StopReal:        time.Now().Unix() + 11000,
		StartReal:       time.Now().Unix() - 1000,
		StartExtra:      1,
		Start:           time.Now().Unix(),
		DispTitle:       "someTitle",
		DispSubtitle:    "someSubtitle",
		DispExtratext:   "someExtratext",
		DispDescription: "someDescription",
		SchedStatus:     "scheduled",
		Broadcast:       1234,
		ChannelIcon:     fmt.Sprintf("imagecache/%d", piconId),
	}

	recording := core.MapToTvheadendDvrGridEntryToRecording(tvhEntry)

	assert.Equal(t, tvhEntry.Channel, recording.ChannelID)
	assert.Equal(t, tvhEntry.Channelname, recording.ChannelName)
	assert.Equal(t, tvhEntry.Create, recording.CreatedAt)
	assert.Equal(t, tvhEntry.Duration, recording.Duration)
	assert.Equal(t, tvhEntry.Enabled, recording.Enabled)
	assert.Equal(t, tvhEntry.StopExtra, recording.EndPadding)
	assert.Equal(t, tvhEntry.Stop, recording.EndsAt)
	assert.Equal(t, tvhEntry.Filename, recording.Filename)
	assert.Equal(t, tvhEntry.UUID, recording.ID)
	assert.Equal(t, tvhEntry.Title, recording.LangTitle)
	assert.Equal(t, tvhEntry.StopReal, recording.OriginalEndsAt)
	assert.Equal(t, tvhEntry.StartReal, recording.OriginalStartsAt)
	assert.Equal(t, tvhEntry.StartExtra, recording.StartPadding)
	assert.Equal(t, tvhEntry.SchedStatus, recording.Status)
	assert.Equal(t, tvhEntry.DispTitle, recording.Title)
	assert.Equal(t, tvhEntry.DispSubtitle, recording.Subtitle)
	assert.Equal(t, tvhEntry.DispExtratext, recording.ExtraText)
	assert.Equal(t, tvhEntry.DispDescription, recording.Description)
	assert.Equal(t, tvhEntry.Broadcast, recording.EventID)
	assert.Equal(t, piconId, recording.PiconID)
}

func TestMapTvheadendIdnodeToRecordingFailsForUnexpectedType(t *testing.T) {
	idnode := tvheadend.Idnode{
		UUID: "someID",
		Params: []tvheadend.InodeParams{
			{
				ID:    "enabled",
				Value: "true",
			},
		},
	}

	recording, err := core.MapTvheadendIdnodeToRecording(idnode)

	assert.Nil(t, recording)
	assert.Equal(t, conv.ErrInterfaceToBool, err)
}

func TestMapTvheadendIdnodeToRecording(t *testing.T) {
	enabled := true
	title := "someTitle"
	filename := "someFilename"
	channelId := "someChannelId"
	channelName := "someChannelName"
	piconId := 6
	extraText := "someExtraText"
	description := "someDescription"
	subtitle := "someSubtitle"
	eventId := int64(1234)

	startsAt := time.Now().Unix() + 60000
	originalStartsAt := time.Now().Unix()
	startPadding := 1

	endsAt := time.Now().Unix() + 1000000
	originalEndsAt := time.Now().Unix() + 1000000
	endPadding := 0

	duration := int64(100000)
	createdAt := time.Now().Unix()

	status := "scheduled"

	idnode := tvheadend.Idnode{
		UUID: "someID",
		Params: []tvheadend.InodeParams{
			{
				ID:    "enabled",
				Value: enabled,
			},
			{
				ID: "title",
				Value: map[string]interface{}{
					"key": "value",
				},
			},
			{
				ID:    "disp_title",
				Value: title,
			},
			{
				ID:    "filename",
				Value: filename,
			},
			{
				ID:    "channel",
				Value: channelId,
			},
			{
				ID:    "channelname",
				Value: channelName,
			},
			{
				ID:    "channel_icon",
				Value: fmt.Sprintf("imagecache/%d", piconId),
			},
			{
				ID:    "start",
				Value: float64(startsAt),
			},
			{
				ID:    "start_real",
				Value: float64(originalStartsAt),
			},
			{
				ID:    "start_extra",
				Value: float64(startPadding),
			},
			{
				ID:    "stop",
				Value: float64(endsAt),
			},
			{
				ID:    "stop_real",
				Value: float64(originalEndsAt),
			},
			{
				ID:    "stop_extra",
				Value: float64(endPadding),
			},
			{
				ID:    "duration",
				Value: float64(duration),
			},
			{
				ID:    "create",
				Value: float64(createdAt),
			},
			{
				ID:    "sched_status",
				Value: status,
			},
			{
				ID:    "disp_subtitle",
				Value: subtitle,
			},
			{
				ID:    "disp_extratext",
				Value: extraText,
			},
			{
				ID:    "disp_description",
				Value: description,
			},
			{
				ID:    "broadcast",
				Value: float64(eventId),
			},
		},
	}

	recording, err := core.MapTvheadendIdnodeToRecording(idnode)

	assert.Nil(t, err)
	assert.Equal(t, enabled, recording.Enabled)
	assert.Equal(t, map[string]string{
		"key": "value",
	}, recording.LangTitle)
	assert.Equal(t, title, recording.Title)
	assert.Equal(t, filename, recording.Filename)
	assert.Equal(t, channelId, recording.ChannelID)
	assert.Equal(t, channelName, recording.ChannelName)
	assert.Equal(t, startsAt, recording.StartsAt)
	assert.Equal(t, originalStartsAt, recording.OriginalStartsAt)
	assert.Equal(t, startPadding, recording.StartPadding)
	assert.Equal(t, endsAt, recording.EndsAt)
	assert.Equal(t, originalEndsAt, recording.OriginalEndsAt)
	assert.Equal(t, endPadding, recording.EndPadding)
	assert.Equal(t, duration, recording.Duration)
	assert.Equal(t, createdAt, recording.CreatedAt)
	assert.Equal(t, status, recording.Status)
	assert.Equal(t, subtitle, recording.Subtitle)
	assert.Equal(t, extraText, recording.ExtraText)
	assert.Equal(t, description, recording.Description)
	assert.Equal(t, eventId, recording.EventID)
	assert.Equal(t, piconId, recording.PiconID)
}

func TestBuildTvheadendDvrUpdateRecordingOptsFailsForUnexpectedType(t *testing.T) {
	idnode := tvheadend.Idnode{
		UUID: "someID",
		Params: []tvheadend.InodeParams{
			{
				ID:    "enabled",
				Value: "true",
			},
		},
	}

	opts, err := core.BuildTvheadendDvrUpdateRecordingOpts(idnode, core.UpdateRecording{})

	assert.Nil(t, opts)
	assert.Equal(t, conv.ErrInterfaceToBool, err)
}

func TestBuildTvheadendDvrUpdateRecording(t *testing.T) {
	id := "someID"
	enabled := true
	title := "someTitle"
	channelId := "someChannelId"
	extraText := "someExtraText"
	comment := "someComment"
	episode := "someEpisode"
	priority := 1
	dvrConfigId := "someDvrConfigId"
	owner := "someOwner"
	creator := "someCreator"
	retention := 1
	removal := 2

	startsAt := time.Now().Unix() + 60000
	startPadding := 1

	endsAt := time.Now().Unix() + 1000000
	endPadding := 1

	idnode := tvheadend.Idnode{
		UUID: id,
		Params: []tvheadend.InodeParams{
			{
				ID:    "enabled",
				Value: enabled,
			},
			{
				ID:    "disp_title",
				Value: title,
			},
			{
				ID:    "disp_extratext",
				Value: extraText,
			},
			{
				ID:    "channel",
				Value: channelId,
			},
			{
				ID:    "start",
				Value: float64(startsAt),
			},
			{
				ID:    "stop",
				Value: float64(endsAt),
			},
			{
				ID:    "comment",
				Value: comment,
			},
			{
				ID:    "episode_disp",
				Value: episode,
			},
			{
				ID:    "start_extra",
				Value: float64(startPadding),
			},
			{
				ID:    "stop_extra",
				Value: float64(endPadding),
			},
			{
				ID:    "pri",
				Value: float64(priority),
			},
			{
				ID:    "config_name",
				Value: dvrConfigId,
			},
			{
				ID:    "owner",
				Value: owner,
			},
			{
				ID:    "creator",
				Value: creator,
			},
			{
				ID:    "removal",
				Value: float64(removal),
			},
			{
				ID:    "retention",
				Value: float64(retention),
			},
		},
	}

	// Build opts without merging and assert the values.
	notMerged, err := core.BuildTvheadendDvrUpdateRecordingOpts(idnode, core.UpdateRecording{})
	assert.Nil(t, err)

	assert.Equal(t, channelId, notMerged.Channel)
	assert.Equal(t, comment, notMerged.Comment)
	assert.Equal(t, dvrConfigId, notMerged.ConfigName)
	assert.Equal(t, creator, notMerged.Creator)
	assert.Equal(t, extraText, notMerged.DispExtratext)
	assert.Equal(t, title, notMerged.DispTitle)
	assert.Equal(t, enabled, notMerged.Enabled)
	assert.Equal(t, episode, notMerged.EpisodeDisp)
	assert.Equal(t, owner, notMerged.Owner)
	assert.Equal(t, priority, notMerged.Pri)
	assert.Equal(t, removal, notMerged.Removal)
	assert.Equal(t, retention, notMerged.Retention)
	assert.Equal(t, startsAt, notMerged.Start)
	assert.Equal(t, startPadding, notMerged.StartExtra)
	assert.Equal(t, endsAt, notMerged.Stop)
	assert.Equal(t, endPadding, notMerged.StopExtra)
	assert.Equal(t, id, notMerged.UUID)

	mergedTitle := "mergedTitle"
	mergedComment := "mergedComment"
	mergedEnabled := false
	mergedEndPadding := 2
	mergedEndsAt := int64(1234)
	mergedEpisode := "mergedEpisode"
	mergedExtratext := "mergedExtratext"
	mergedPriority := 500
	mergedStartPadding := 2
	mergedStartsAt := int64(123)

	// Build opts with merging and assert the values.
	opts := core.UpdateRecording{
		Title:        &mergedTitle,
		ExtraText:    &mergedExtratext,
		StartsAt:     &mergedStartsAt,
		EndsAt:       &mergedEndsAt,
		Comment:      &mergedComment,
		StartPadding: &mergedStartPadding,
		EndPadding:   &mergedEndPadding,
		Priority:     &mergedPriority,
		Enabled:      &mergedEnabled,
		Episode:      &mergedEpisode,
	}

	merged, err := core.BuildTvheadendDvrUpdateRecordingOpts(idnode, opts)
	assert.Nil(t, err)

	assert.Equal(t, channelId, merged.Channel)
	assert.Equal(t, mergedComment, merged.Comment)
	assert.Equal(t, dvrConfigId, merged.ConfigName)
	assert.Equal(t, creator, merged.Creator)
	assert.Equal(t, mergedExtratext, merged.DispExtratext)
	assert.Equal(t, mergedTitle, merged.DispTitle)
	assert.Equal(t, mergedEnabled, merged.Enabled)
	assert.Equal(t, mergedEpisode, merged.EpisodeDisp)
	assert.Equal(t, owner, merged.Owner)
	assert.Equal(t, mergedPriority, merged.Pri)
	assert.Equal(t, removal, merged.Removal)
	assert.Equal(t, retention, merged.Retention)
	assert.Equal(t, mergedStartsAt, merged.Start)
	assert.Equal(t, mergedStartPadding, merged.StartExtra)
	assert.Equal(t, mergedEndsAt, merged.Stop)
	assert.Equal(t, mergedEndPadding, merged.StopExtra)
	assert.Equal(t, id, merged.UUID)
}
