package core

import (
	"context"
	"errors"
	"time"

	"github.com/davidborzek/tvhgo/conv"
	"github.com/davidborzek/tvhgo/tvheadend"
)

var (
	ErrRecordingNotFound = errors.New("recording not found")

	ErrGetRecordingsInvalidStatus = errors.New("get recording status invalid")

	ErrCreateRecordingInvalidEventID = errors.New("recording event id invalid")

	ErrRecordingInvalidTitle     = errors.New("recording invalid title")
	ErrRecordingInvalidChannelID = errors.New("recording invalid channel id")
	ErrRecordingInvalidStartDate = errors.New("recording invalid start date")
	ErrRecordingInvalidEndDate   = errors.New("recording invalid end date")
	ErrRecordingInvalidEventID   = errors.New("recording invalid event id")
)

type (
	// Recording defines a dvr entry from tvheadend.
	// TODO: extend necessary fields.
	Recording struct {
		ChannelID   string            `json:"channelId"`
		ChannelName string            `json:"channelName"`
		CreatedAt   int64             `json:"createdAt"`
		Duration    int64             `json:"duration"`
		Enabled     bool              `json:"enabled"`
		Filename    string            `json:"filename"`
		ID          string            `json:"id"`
		LangTitle   map[string]string `json:"langTitle"`
		Title       string            `json:"title"`
		// OriginalStartsAt time stamp of the original start date
		// without StartPadding.
		OriginalStartsAt int64 `json:"originalStartsAt"`
		// OriginalEndsAt time stamp of the original end date
		// without theEndPadding.
		OriginalEndsAt int64 `json:"originalEndsAt"`
		// StartAt start date of the recording as unix timestamp.
		StartsAt int64 `json:"startsAt"`
		// EndsAt end date of the recording as unix timestamp.
		EndsAt int64 `json:"endsAt"`
		// StartPadding optional padding in minutes to record
		// before the recording starts.
		StartPadding int `json:"startPadding"`
		// EndPadding optional padding in minutes to record
		// after the recording ends.
		EndPadding int    `json:"endPadding"`
		Status     string `json:"status"`
	}

	// GetRecordingsParams defines query params
	// to paginate, sort and filter the recordings.
	GetRecordingsParams struct {
		PaginationSortQueryParams
		// upcoming, finished, failed, removed
		Status string `schema:"status"`
	}

	// CreateRecordingByEvent defines options
	// to create a recording by an epg event.
	CreateRecordingByEvent struct {
		EventID  int64  `json:"eventId"`
		ConfigID string `json:"configId"`
	}

	// CreateRecording recording defines options to
	// manually create a recording timer.
	CreateRecording struct {
		// Title title of the recording.
		Title string `json:"title"`
		// ExtraText optional extraText of the recording.
		ExtraText string `json:"extraText"`
		// ChannelID the channel id for the recording.
		ChannelID string `json:"channelId"`
		// StartAt start date of the recording as unix timestamp.
		StartsAt int64 `json:"startsAt"`
		// EndsAt end date of the recording as unix timestamp.
		EndsAt int64 `json:"endsAt"`
		// Comment optional comment of the recording.
		Comment string `json:"comment"`
		// StartPadding optional padding in minutes to record
		// before the recording starts.
		StartPadding int `json:"startPadding"`
		// EndPadding optional padding in minutes to record
		// after the recording ends.
		EndPadding int `json:"endPadding"`
		// Priority priority of the recording.
		Priority int `json:"priority"`
		// ConfigID configuration id of the dvr config.
		ConfigID string `json:"configId"`
	}

	// UpdateRecording recording defines options to
	// update a recording timer.
	UpdateRecording struct {
		CreateRecording
		// Enabled enabled status of the recording.
		Enabled bool `json:"enabled"`
		// Episode episode of the recording.
		Episode string `json:"episode"`
		// Owner owner of the recording.
		Owner string `json:"owner"`
		// Creator creator of the recording.
		Creator string `json:"creator"`
		// RemovalTime removal time in days of the recording.
		RemovalTime int `json:"removalTime"`
		// RetentionTime log retention time in days of the recording.
		RetentionTime int `json:"retentionTime"`
	}

	// RecordingService provides access to recording
	// resources from the tvheadend server.
	RecordingService interface {
		// CreateByEvent creates a new recording by an epg event.
		CreateByEvent(ctx context.Context, opts CreateRecordingByEvent) error
		// CreateByEvent creates a new recording by an epg event.
		Create(ctx context.Context, opts CreateRecording) error
		// GetAll returns a list of recordings.
		GetAll(ctx context.Context, params GetRecordingsParams) ([]*Recording, error)
		// Get returns a recording by its id.
		Get(ctx context.Context, id string) (*Recording, error)
		// Stop gracefully stops a running recording.
		Stop(ctx context.Context, id string) error
		// Cancel deletes a upcoming recording or aborts a running recording.
		Cancel(ctx context.Context, id string) error
		// Remove removes a finished recording from disk.
		Remove(ctx context.Context, id string) error
		// MoveFinished moves a recording to finished status.
		MoveFinished(ctx context.Context, id string) error
		// MoveFailed moves a recording to failed status.
		MoveFailed(ctx context.Context, id string) error
		// UpdateRecording updates a recording.
		UpdateRecording(ctx context.Context, id string, opts UpdateRecording) error
	}
)

// Validate validates the minimum requirements of GetRecordingsParams.
func (o *GetRecordingsParams) Validate() error {
	if o.Status != "" && o.Status != "upcoming" && o.Status != "finished" && o.Status != "failed" && o.Status != "remove" {
		return ErrGetRecordingsInvalidStatus
	}
	return nil
}

// Validate validates the minimum requirements of CreateRecordingByEvent.
func (o *CreateRecordingByEvent) Validate() error {
	if o.EventID == 0 {
		return ErrCreateRecordingInvalidEventID
	}
	return nil
}

// Validate validates the minimum requirements of CreateRecording.
func (r *CreateRecording) Validate() error {
	switch {
	case r.Title == "":
		return ErrRecordingInvalidTitle
	case r.ChannelID == "":
		return ErrRecordingInvalidChannelID
	case r.StartsAt < time.Now().Unix():
		return ErrRecordingInvalidStartDate
	case r.EndsAt < time.Now().Unix():
		return ErrRecordingInvalidEndDate
	}
	return nil
}

// Validate validates the minimum requirements of UpdateRecording.
func (r *UpdateRecording) Validate() error {
	switch {
	case r.Title == "":
		return ErrRecordingInvalidTitle
	case r.ChannelID == "":
		return ErrRecordingInvalidChannelID
	case r.StartsAt == 0:
		return ErrRecordingInvalidStartDate
	case r.EndsAt == 0:
		return ErrRecordingInvalidEndDate
	}
	return nil
}

// MapToTvheadendOpts maps CreateRecording to tvheadend.DvrCreateRecordingOpts.
func (c *CreateRecording) MapToTvheadendOpts() tvheadend.DvrCreateRecordingOpts {
	return tvheadend.DvrCreateRecordingOpts{
		DispTitle:     c.Title,
		DispExtratext: c.ExtraText,
		Channel:       c.ChannelID,
		Start:         c.StartsAt,
		Stop:          c.EndsAt,
		Comment:       c.Comment,
		StartExtra:    c.StartPadding,
		StopExtra:     c.EndPadding,
		Pri:           c.Priority,
		ConfigName:    c.ConfigID,
	}
}

// MapToTvheadendOpts maps UpdateRecording to tvheadend.DvrUpdateRecordingOpts.
func (u *UpdateRecording) MapToTvheadendOpts(id string) tvheadend.DvrUpdateRecordingOpts {
	return tvheadend.DvrUpdateRecordingOpts{
		Channel:       u.ChannelID,
		Comment:       u.Comment,
		ConfigName:    u.ConfigID,
		Creator:       u.Creator,
		DispExtratext: u.ExtraText,
		DispTitle:     u.Title,
		Enabled:       u.Enabled,
		EpisodeDisp:   u.Episode,
		Owner:         u.Owner,
		Pri:           u.Priority,
		Removal:       u.RemovalTime,
		Retention:     u.RetentionTime,
		Start:         u.StartsAt,
		StartExtra:    u.StartPadding,
		Stop:          u.EndsAt,
		StopExtra:     u.EndPadding,
		UUID:          id,
	}
}

// MapToTvheadendDvrGridEntryToRecording maps a tvheadend.DvrGridEntry to a Recording.
func MapToTvheadendDvrGridEntryToRecording(entry tvheadend.DvrGridEntry) Recording {
	return Recording{
		ChannelID:        entry.Channel,
		ChannelName:      entry.Channelname,
		CreatedAt:        entry.Create,
		Duration:         entry.Duration,
		Enabled:          entry.Enabled,
		EndPadding:       entry.StopExtra,
		EndsAt:           entry.Stop,
		Filename:         entry.Filename,
		ID:               entry.UUID,
		LangTitle:        entry.Title,
		OriginalEndsAt:   entry.StopReal,
		OriginalStartsAt: entry.StartReal,
		StartPadding:     entry.StartExtra,
		StartsAt:         entry.Start,
		Title:            entry.DispTitle,
		Status:           entry.SchedStatus,
	}
}

// MapTvheadendIdnodeToRecording maps a tvheadend.Idnode to a Recording.
func MapTvheadendIdnodeToRecording(idnode tvheadend.Idnode) (*Recording, error) {
	r := Recording{
		ID: idnode.UUID,
	}

	for _, p := range idnode.Params {
		var err error

		switch p.ID {
		case "enabled":
			r.Enabled, err = conv.InterfaceToBool(p.Value)
		case "title":
			r.LangTitle, err = conv.InterfaceToStringMap(p.Value)
		case "disp_title":
			r.Title, err = conv.InterfaceToString(p.Value)
		case "filename":
			r.Filename = p.Value.(string)
		case "channel":
			r.ChannelID = p.Value.(string)
		case "channelname":
			r.ChannelName = p.Value.(string)
		case "start":
			r.StartsAt, err = conv.InterfaceToInt64(p.Value)
		case "start_real":
			r.OriginalStartsAt, err = conv.InterfaceToInt64(p.Value)
		case "start_extra":
			r.StartPadding, err = conv.InterfaceToInt(p.Value)
		case "stop":
			r.EndsAt, err = conv.InterfaceToInt64(p.Value)
		case "stop_real":
			r.OriginalEndsAt, err = conv.InterfaceToInt64(p.Value)
		case "stop_extra":
			r.EndPadding, err = conv.InterfaceToInt(p.Value)
		case "duration":
			r.Duration, err = conv.InterfaceToInt64(p.Value)
		case "create":
			r.CreatedAt, err = conv.InterfaceToInt64(p.Value)
		case "sched_status":
			r.Status, err = conv.InterfaceToString(p.Value)
		}

		if err != nil {
			return nil, err
		}
	}

	return &r, nil
}
