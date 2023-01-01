package core_test

import (
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
	c := core.UpdateRecording{
		CreateRecording: core.CreateRecording{
			Title:     "someTitle",
			ChannelID: "someChannelID",
			StartsAt:  time.Now().Unix(),
			EndsAt:    time.Now().Unix() - 10000,
		},
	}

	err := c.Validate()
	assert.Nil(t, err)
}

func TestUpdateRecordingValidateReturnErrorInvalidTitle(t *testing.T) {
	c := core.UpdateRecording{}

	err := c.Validate()
	assert.Equal(t, core.ErrRecordingInvalidTitle, err)
}

func TestUpdateRecordingValidateReturnErrorInvalidChannelID(t *testing.T) {
	c := core.UpdateRecording{
		CreateRecording: core.CreateRecording{
			Title: "someTitle",
		},
	}

	err := c.Validate()
	assert.Equal(t, core.ErrRecordingInvalidChannelID, err)
}

func TestUpdateRecordingValidateReturnErrorInvalidStartDate(t *testing.T) {
	c := core.UpdateRecording{
		CreateRecording: core.CreateRecording{
			Title:     "someTitle",
			ChannelID: "someChannelID",
		},
	}

	err := c.Validate()
	assert.Equal(t, core.ErrRecordingInvalidStartDate, err)
}

func TestUpdateRecordingValidateReturnErrorInvalidEndDate(t *testing.T) {
	c := core.UpdateRecording{
		CreateRecording: core.CreateRecording{
			Title:     "someTitle",
			ChannelID: "someChannelID",
			StartsAt:  time.Now().Unix(),
		},
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

func TestUpdateRecordingMapToTvheadendOpts(t *testing.T) {
	c := core.UpdateRecording{
		CreateRecording: core.CreateRecording{
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
		},
		Enabled:       true,
		Episode:       "someEpisode",
		Owner:         "someOwner",
		Creator:       "someCreator",
		RemovalTime:   1,
		RetentionTime: 1,
	}

	id := "someId"

	mapped := c.MapToTvheadendOpts(id)

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

	assert.Equal(t, c.Creator, mapped.Creator)
	assert.Equal(t, c.Enabled, mapped.Enabled)
	assert.Equal(t, c.Episode, mapped.EpisodeDisp)
	assert.Equal(t, c.Owner, mapped.Owner)
	assert.Equal(t, c.RemovalTime, mapped.Removal)
	assert.Equal(t, c.RetentionTime, mapped.Retention)

	assert.Equal(t, id, mapped.UUID)
}

func TestMapToTvheadendDvrGridEntryToRecording(t *testing.T) {
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
		StopReal:    time.Now().Unix() + 11000,
		StartReal:   time.Now().Unix() - 1000,
		StartExtra:  1,
		Start:       time.Now().Unix(),
		DispTitle:   "someTitle",
		SchedStatus: "scheduled",
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
	assert.Equal(t, tvhEntry.DispTitle, recording.Title)
	assert.Equal(t, tvhEntry.SchedStatus, recording.Status)
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
}
