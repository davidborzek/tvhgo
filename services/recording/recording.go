package recording

import (
	"context"
	"errors"
	"fmt"

	"github.com/davidborzek/tvhgo/core"
	"github.com/davidborzek/tvhgo/tvheadend"
)

type service struct {
	tvh tvheadend.Client
}

var (
	ErrRequestFailed = errors.New("recording request failed")

	// sortKeyMapping mapping of Recording model fields
	// to the tvheadend model fields used for sorting.
	sortKeyMapping = map[string]string{
		"channelName":      "channelname",
		"endsAt":           "stop",
		"filename":         "filename",
		"originalEndsAt":   "stop_real",
		"originalStartsAt": "start_real",
		"startsAt":         "start",
		"title":            "disp_title",
	}
)

func New(tvh tvheadend.Client) core.RecordingService {
	return &service{
		tvh: tvh,
	}
}

func (s *service) CreateByEvent(ctx context.Context, opts core.CreateRecordingByEvent) error {
	q := tvheadend.NewQuery()
	q.SetInt("event_id", opts.EventID)
	q.Set("config_uuid", opts.ConfigID)

	res, err := s.tvh.Exec(ctx, "/api/dvr/entry/create_by_event", nil, q)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		return ErrRequestFailed
	}

	return nil
}

func (s *service) Create(ctx context.Context, opts core.CreateRecording) error {
	q := tvheadend.NewQuery()
	conf := opts.MapToTvheadendOpts()

	if err := q.Conf(&conf); err != nil {
		return err
	}

	var created tvheadend.DvrRecordingCreated
	res, err := s.tvh.Exec(ctx, "/api/dvr/entry/create", &created, q)
	if err != nil {
		return err
	}

	if res.StatusCode >= 400 {
		return ErrRequestFailed
	}

	return nil
}

func (s *service) GetAll(
	ctx context.Context,
	params core.GetRecordingsParams,
) (*core.RecordingListResult, error) {
	q := params.PaginationSortQueryParams.MapToTvheadendQuery(sortKeyMapping)

	var url string
	if params.Status == "" {
		url = "/api/dvr/entry/grid"
	} else {
		url = fmt.Sprintf("/api/dvr/entry/grid_%s", params.Status)
	}

	var grid tvheadend.DvrGrid
	res, err := s.tvh.Exec(ctx, url, &grid, q)
	if err != nil {
		return nil, err
	}

	if res.StatusCode >= 400 {
		return nil, ErrRequestFailed
	}

	recordings := make([]*core.Recording, 0)
	for _, entry := range grid.Entries {
		r := core.MapToTvheadendDvrGridEntryToRecording(entry)
		recordings = append(recordings, &r)
	}

	result := core.RecordingListResult{
		Entries: recordings,
		Total:   grid.Total,
		Offset:  params.Offset,
	}

	return &result, nil
}

func (s *service) Stop(ctx context.Context, id string) error {
	q := tvheadend.NewQuery()
	q.Set("uuid", id)

	return s.stop(ctx, q)
}

func (s *service) BatchStop(ctx context.Context, ids []string) error {
	q := tvheadend.NewQuery()
	q.SetJSON("uuid", ids)

	return s.stop(ctx, q)
}

func (s *service) Cancel(ctx context.Context, id string) error {
	q := tvheadend.NewQuery()
	q.Set("uuid", id)

	return s.cancel(ctx, q)
}

func (s *service) BatchCancel(ctx context.Context, ids []string) error {
	q := tvheadend.NewQuery()
	q.SetJSON("uuid", ids)

	return s.cancel(ctx, q)
}

func (s *service) Remove(ctx context.Context, id string) error {
	q := tvheadend.NewQuery()
	q.Set("uuid", id)

	return s.remove(ctx, q)
}

func (s *service) BatchRemove(ctx context.Context, ids []string) error {
	q := tvheadend.NewQuery()
	q.SetJSON("uuid", ids)

	return s.remove(ctx, q)
}

func (s *service) MoveFinished(ctx context.Context, id string) error {
	q := tvheadend.NewQuery()
	q.Set("uuid", id)

	res, err := s.tvh.Exec(ctx, "/api/dvr/entry/move/finished", nil, q)
	if err != nil {
		return err
	}

	if res.StatusCode >= 400 {
		return ErrRequestFailed
	}

	return nil
}

func (s *service) MoveFailed(ctx context.Context, id string) error {
	q := tvheadend.NewQuery()
	q.Set("uuid", id)

	res, err := s.tvh.Exec(ctx, "/api/dvr/entry/move/failed", nil, q)
	if err != nil {
		return err
	}

	if res.StatusCode >= 400 {
		return ErrRequestFailed
	}

	return nil
}

func (s *service) UpdateRecording(ctx context.Context, id string, opts core.UpdateRecording) error {
	idnode, err := s.getRecordingIdnode(ctx, id)
	if err != nil {
		return err
	}

	node, err := core.BuildTvheadendDvrUpdateRecordingOpts(*idnode, opts)
	if err != nil {
		return err
	}

	q := tvheadend.NewQuery()
	if err := q.Node(node); err != nil {
		return err
	}

	res, err := s.tvh.Exec(ctx, "/api/idnode/save", nil, q)
	if err != nil {
		return err
	}

	if res.StatusCode >= 400 {
		return ErrRequestFailed
	}

	return nil
}

func (s *service) Get(ctx context.Context, id string) (*core.Recording, error) {
	idnode, err := s.getRecordingIdnode(ctx, id)
	if err != nil {
		return nil, err
	}

	recording, err := core.MapTvheadendIdnodeToRecording(*idnode)
	if err != nil {
		return nil, err
	}

	return recording, nil
}

func (s *service) getRecordingIdnode(ctx context.Context, id string) (*tvheadend.Idnode, error) {
	q := tvheadend.NewQuery()
	q.Set("uuid", id)

	var idnodeLoad tvheadend.IdnodeLoadResponse
	res, err := s.tvh.Exec(ctx, "/api/idnode/load", &idnodeLoad, q)
	if err != nil {
		return nil, err
	}

	if res.StatusCode >= 400 {
		return nil, ErrRequestFailed
	}

	if len(idnodeLoad.Entries) == 0 {
		return nil, core.ErrRecordingNotFound
	}

	return &idnodeLoad.Entries[0], nil
}

func (s *service) stop(ctx context.Context, q tvheadend.Query) error {
	res, err := s.tvh.Exec(ctx, "/api/dvr/entry/stop", nil, q)
	if err != nil {
		return err
	}

	if res.StatusCode >= 400 {
		return ErrRequestFailed
	}

	return nil
}

func (s *service) cancel(ctx context.Context, q tvheadend.Query) error {
	res, err := s.tvh.Exec(ctx, "/api/dvr/entry/cancel", nil, q)
	if err != nil {
		return err
	}

	if res.StatusCode >= 400 {
		return ErrRequestFailed
	}

	return nil
}

func (s *service) remove(ctx context.Context, q tvheadend.Query) error {
	res, err := s.tvh.Exec(ctx, "/api/dvr/entry/remove", nil, q)
	if err != nil {
		return err
	}

	if res.StatusCode >= 400 {
		return ErrRequestFailed
	}

	return nil
}
