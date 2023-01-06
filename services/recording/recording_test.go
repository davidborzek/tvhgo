package recording_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/davidborzek/tvhgo/core"
	mock_tvheadend "github.com/davidborzek/tvhgo/mock/tvheadend"
	"github.com/davidborzek/tvhgo/services/recording"
	"github.com/davidborzek/tvhgo/tvheadend"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var (
	dvrGridEntry = tvheadend.DvrGridEntry{
		Channel:     "someChannelID",
		Channelname: "someChannel",
		Create:      1234,
		Duration:    12345,
		Enabled:     true,
		StartExtra:  2,
		Stop:        2345,
		Filename:    "someFilename",
		Title: map[string]string{
			"de": "someTitle",
		},
		StopReal:    2346,
		StartReal:   2347,
		StopExtra:   1,
		DispTitle:   "someTitle",
		UUID:        "someID",
		SchedStatus: "scheduled",
	}

	idnode = tvheadend.Idnode{
		UUID: "someID",
		Params: []tvheadend.InodeParams{
			{
				ID:    "enabled",
				Value: true,
			},
			{
				ID: "title",
				Value: map[string]interface{}{
					"key": "value",
				},
			},
			{
				ID:    "disp_title",
				Value: "someTitle",
			},
			{
				ID:    "filename",
				Value: "someFilename",
			},
			{
				ID:    "channel",
				Value: "someChannelId",
			},
			{
				ID:    "channelname",
				Value: "someChannelName",
			},
			{
				ID:    "start",
				Value: float64(1234),
			},
			{
				ID:    "start_real",
				Value: float64(12345),
			},
			{
				ID:    "start_extra",
				Value: float64(1),
			},
			{
				ID:    "stop",
				Value: float64(234),
			},
			{
				ID:    "stop_real",
				Value: float64(2345),
			},
			{
				ID:    "stop_extra",
				Value: float64(2),
			},
			{
				ID:    "duration",
				Value: float64(1000),
			},
			{
				ID:    "create",
				Value: float64(3456),
			},
		},
	}

	ctx = context.TODO()
)

func mockClientExecSucceedsForGetAll(
	ctx context.Context,
	path string,
	dst interface{},
	query ...tvheadend.Query,
) (*tvheadend.Response, error) {
	res := &tvheadend.Response{
		Response: &http.Response{
			StatusCode: 200,
		},
	}

	g := dst.(*tvheadend.DvrGrid)
	g.Entries = []tvheadend.DvrGridEntry{
		dvrGridEntry,
	}
	g.Total = 20

	return res, nil
}

func mockClientExecReturnsEmptyEntriesForGet(
	ctx context.Context,
	path string,
	dst interface{},
	query ...tvheadend.Query,
) (*tvheadend.Response, error) {
	res := &tvheadend.Response{
		Response: &http.Response{
			StatusCode: 200,
		},
	}

	g := dst.(*tvheadend.IdnodeLoadResponse)
	g.Entries = []tvheadend.Idnode{}

	return res, nil
}

func mockClientExecSucceedsForGet(
	ctx context.Context,
	path string,
	dst interface{},
	query ...tvheadend.Query,
) (*tvheadend.Response, error) {
	res := &tvheadend.Response{
		Response: &http.Response{
			StatusCode: 200,
		},
	}

	g := dst.(*tvheadend.IdnodeLoadResponse)
	g.Entries = []tvheadend.Idnode{
		idnode,
	}

	return res, nil
}

func TestCreateReturnsError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_tvheadend.NewMockClient(ctrl)
	mockClient.EXPECT().
		Exec(ctx, "/api/dvr/entry/create", gomock.Any(), gomock.Any()).
		DoAndReturn(mock_tvheadend.MockClientExecReturnsError).
		Times(1)

	service := recording.New(mockClient)
	err := service.Create(ctx, core.CreateRecording{})

	assert.EqualError(t, err, "error")
}

func TestCreateRequestFailedError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_tvheadend.NewMockClient(ctrl)
	mockClient.EXPECT().
		Exec(ctx, "/api/dvr/entry/create", gomock.Any(), gomock.Any()).
		DoAndReturn(mock_tvheadend.MockClientExecReturnsErroneousHttpStatus).
		Times(1)

	service := recording.New(mockClient)
	err := service.Create(ctx, core.CreateRecording{})

	assert.Equal(t, err, recording.ErrRequestFailed)
}
func TestCreateSucceeds(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	opts := core.CreateRecording{
		Title:     "someTitle",
		ChannelID: "someChannelID",
		StartsAt:  1000,
		EndsAt:    2000,
	}

	tvhq := tvheadend.NewQuery()
	tvhq.Conf(&tvheadend.DvrCreateRecordingOpts{
		DispTitle: opts.Title,
		Channel:   opts.ChannelID,
		Start:     opts.StartsAt,
		Stop:      opts.EndsAt,
	})

	mockClient := mock_tvheadend.NewMockClient(ctrl)
	mockClient.EXPECT().
		Exec(ctx, "/api/dvr/entry/create", gomock.Any(), tvhq).
		Return(&tvheadend.Response{
			Response: &http.Response{
				StatusCode: 200,
				Body:       http.NoBody,
			}}, nil).
		Times(1)

	service := recording.New(mockClient)
	err := service.Create(ctx, opts)

	assert.Nil(t, err)
}

func TestCreateByEventReturnsError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_tvheadend.NewMockClient(ctrl)
	mockClient.EXPECT().
		Exec(ctx, "/api/dvr/entry/create_by_event", gomock.Any(), gomock.Any()).
		DoAndReturn(mock_tvheadend.MockClientExecReturnsError).
		Times(1)

	service := recording.New(mockClient)
	err := service.CreateByEvent(ctx, core.CreateRecordingByEvent{})

	assert.EqualError(t, err, "error")
}

func TestCreateByEventReturnsRequestFailedError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_tvheadend.NewMockClient(ctrl)
	mockClient.EXPECT().
		Exec(ctx, "/api/dvr/entry/create_by_event", gomock.Any(), gomock.Any()).
		DoAndReturn(mock_tvheadend.MockClientExecReturnsErroneousHttpStatus).
		Times(1)

	service := recording.New(mockClient)
	err := service.CreateByEvent(ctx, core.CreateRecordingByEvent{})

	assert.Equal(t, err, recording.ErrRequestFailed)
}

func TestCreateByEventSucceeds(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	opts := core.CreateRecordingByEvent{
		EventID:  1234,
		ConfigID: "someConfigID",
	}

	tvhq := tvheadend.NewQuery()
	tvhq.SetInt("event_id", opts.EventID)
	tvhq.Set("config_uuid", opts.ConfigID)

	mockClient := mock_tvheadend.NewMockClient(ctrl)
	mockClient.EXPECT().
		Exec(ctx, "/api/dvr/entry/create_by_event", gomock.Any(), tvhq).
		Return(&tvheadend.Response{
			Response: &http.Response{
				StatusCode: 200,
				Body:       http.NoBody,
			}}, nil).
		Times(1)

	service := recording.New(mockClient)
	err := service.CreateByEvent(ctx, opts)

	assert.Nil(t, err)
}

func TestGetAllReturnsError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_tvheadend.NewMockClient(ctrl)
	mockClient.EXPECT().
		Exec(ctx, "/api/dvr/entry/grid", gomock.Any(), gomock.Any()).
		DoAndReturn(mock_tvheadend.MockClientExecReturnsError).
		Times(1)

	service := recording.New(mockClient)
	res, err := service.GetAll(ctx, core.GetRecordingsParams{})

	assert.Nil(t, res)
	assert.EqualError(t, err, "error")
}

func TestGetAllReturnsRequestFailedError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_tvheadend.NewMockClient(ctrl)
	mockClient.EXPECT().
		Exec(ctx, "/api/dvr/entry/grid", gomock.Any(), gomock.Any()).
		DoAndReturn(mock_tvheadend.MockClientExecReturnsErroneousHttpStatus).
		Times(1)

	service := recording.New(mockClient)
	res, err := service.GetAll(ctx, core.GetRecordingsParams{})

	assert.Nil(t, res)
	assert.Equal(t, err, recording.ErrRequestFailed)
}

func TestGetAllSucceeds(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tvhq := tvheadend.NewQuery()
	tvhq.Limit(10)
	tvhq.Start(5)
	tvhq.SortKey("disp_title")
	tvhq.SortDir("asc")

	mockClient := mock_tvheadend.NewMockClient(ctrl)
	mockClient.EXPECT().
		Exec(ctx, "/api/dvr/entry/grid", gomock.Any(), tvhq).
		DoAndReturn(mockClientExecSucceedsForGetAll).
		Times(1)

	service := recording.New(mockClient)

	q := core.GetRecordingsParams{}
	q.Limit = 10
	q.Offset = 5
	q.SortDirection = "asc"
	q.SortKey = "title"

	recordings, err := service.GetAll(ctx, q)

	assert.Nil(t, err)
	assert.NotNil(t, recordings)
	assert.Len(t, recordings, 1)

	assert.Equal(t, dvrGridEntry.Channel, recordings[0].ChannelID)
	assert.Equal(t, dvrGridEntry.Channelname, recordings[0].ChannelName)
	assert.Equal(t, int64(dvrGridEntry.Create), recordings[0].CreatedAt)
	assert.Equal(t, int64(dvrGridEntry.Duration), recordings[0].Duration)
	assert.Equal(t, dvrGridEntry.Enabled, recordings[0].Enabled)
	assert.Equal(t, dvrGridEntry.StopExtra, recordings[0].EndPadding)
	assert.Equal(t, int64(dvrGridEntry.Stop), recordings[0].EndsAt)
	assert.Equal(t, dvrGridEntry.Filename, recordings[0].Filename)
	assert.Equal(t, dvrGridEntry.UUID, recordings[0].ID)
	assert.Equal(t, dvrGridEntry.Title, recordings[0].LangTitle)
	assert.Equal(t, int64(dvrGridEntry.StopReal), recordings[0].OriginalEndsAt)
	assert.Equal(t, int64(dvrGridEntry.StartReal), recordings[0].OriginalStartsAt)
	assert.Equal(t, dvrGridEntry.StartExtra, recordings[0].StartPadding)
	assert.Equal(t, int64(dvrGridEntry.Start), recordings[0].StartsAt)
	assert.Equal(t, dvrGridEntry.DispTitle, recordings[0].Title)
	assert.Equal(t, dvrGridEntry.SchedStatus, recordings[0].Status)
}

func TestGetAllMapsSortKeyCorrectly(t *testing.T) {
	keys := map[string]string{
		"channelName":      "channelname",
		"endsAt":           "stop",
		"filename":         "filename",
		"originalEndsAt":   "stop_real",
		"originalStartsAt": "start_real",
		"startsAt":         "start",
		"title":            "disp_title",
	}

	for key, mappedKey := range keys {
		t.Run(key, testGetAllMapsSortKeyCorrectlyParametrize(key, mappedKey))
	}
}

func testGetAllMapsSortKeyCorrectlyParametrize(key string, mappedKey string) func(t *testing.T) {
	return func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		tvhq := tvheadend.NewQuery()
		tvhq.SortKey(mappedKey)

		mockClient := mock_tvheadend.NewMockClient(ctrl)
		mockClient.EXPECT().
			Exec(ctx, "/api/dvr/entry/grid", gomock.Any(), tvhq).
			DoAndReturn(mockClientExecSucceedsForGetAll).
			Times(1)

		q := core.GetRecordingsParams{}
		q.SortKey = key

		service := recording.New(mockClient)
		res, err := service.GetAll(ctx, q)

		assert.NotNil(t, res)
		assert.Nil(t, err)
	}
}

func TestStopReturnsError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_tvheadend.NewMockClient(ctrl)
	mockClient.EXPECT().
		Exec(ctx, "/api/dvr/entry/stop", gomock.Any(), gomock.Any()).
		DoAndReturn(mock_tvheadend.MockClientExecReturnsError).
		Times(1)

	service := recording.New(mockClient)
	err := service.Stop(ctx, "someID")

	assert.EqualError(t, err, "error")
}

func TestStopReturnsRequestFailedError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_tvheadend.NewMockClient(ctrl)
	mockClient.EXPECT().
		Exec(ctx, "/api/dvr/entry/stop", gomock.Any(), gomock.Any()).
		DoAndReturn(mock_tvheadend.MockClientExecReturnsErroneousHttpStatus).
		Times(1)

	service := recording.New(mockClient)
	err := service.Stop(ctx, "someID")

	assert.Equal(t, err, recording.ErrRequestFailed)
}

func TestStopSucceeds(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	id := "someID"

	tvhq := tvheadend.NewQuery()
	tvhq.Set("uuid", id)

	mockClient := mock_tvheadend.NewMockClient(ctrl)
	mockClient.EXPECT().
		Exec(ctx, "/api/dvr/entry/stop", gomock.Any(), tvhq).
		Return(&tvheadend.Response{
			Response: &http.Response{
				StatusCode: 200,
				Body:       http.NoBody,
			}}, nil).
		Times(1)

	service := recording.New(mockClient)
	err := service.Stop(ctx, id)

	assert.Nil(t, err)
}

func TestCancelReturnsError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_tvheadend.NewMockClient(ctrl)
	mockClient.EXPECT().
		Exec(ctx, "/api/dvr/entry/cancel", gomock.Any(), gomock.Any()).
		DoAndReturn(mock_tvheadend.MockClientExecReturnsError).
		Times(1)

	service := recording.New(mockClient)
	err := service.Cancel(ctx, "someID")

	assert.EqualError(t, err, "error")
}

func TestCancelReturnsRequestFailedError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_tvheadend.NewMockClient(ctrl)
	mockClient.EXPECT().
		Exec(ctx, "/api/dvr/entry/cancel", gomock.Any(), gomock.Any()).
		DoAndReturn(mock_tvheadend.MockClientExecReturnsErroneousHttpStatus).
		Times(1)

	service := recording.New(mockClient)
	err := service.Cancel(ctx, "someID")

	assert.Equal(t, err, recording.ErrRequestFailed)
}

func TestCancelSucceeds(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	id := "someID"

	tvhq := tvheadend.NewQuery()
	tvhq.Set("uuid", id)

	mockClient := mock_tvheadend.NewMockClient(ctrl)
	mockClient.EXPECT().
		Exec(ctx, "/api/dvr/entry/cancel", gomock.Any(), tvhq).
		Return(&tvheadend.Response{
			Response: &http.Response{
				StatusCode: 200,
				Body:       http.NoBody,
			}}, nil).
		Times(1)

	service := recording.New(mockClient)
	err := service.Cancel(ctx, id)

	assert.Nil(t, err)
}

func TestRemoveReturnsError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_tvheadend.NewMockClient(ctrl)
	mockClient.EXPECT().
		Exec(ctx, "/api/dvr/entry/remove", gomock.Any(), gomock.Any()).
		DoAndReturn(mock_tvheadend.MockClientExecReturnsError).
		Times(1)

	service := recording.New(mockClient)
	err := service.Remove(ctx, "someID")

	assert.EqualError(t, err, "error")
}

func TestRemoveReturnsRequestFailedError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_tvheadend.NewMockClient(ctrl)
	mockClient.EXPECT().
		Exec(ctx, "/api/dvr/entry/remove", gomock.Any(), gomock.Any()).
		DoAndReturn(mock_tvheadend.MockClientExecReturnsErroneousHttpStatus).
		Times(1)

	service := recording.New(mockClient)
	err := service.Remove(ctx, "someID")

	assert.Equal(t, err, recording.ErrRequestFailed)
}

func TestRemoveSucceeds(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	id := "someID"

	tvhq := tvheadend.NewQuery()
	tvhq.Set("uuid", id)

	mockClient := mock_tvheadend.NewMockClient(ctrl)
	mockClient.EXPECT().
		Exec(ctx, "/api/dvr/entry/remove", gomock.Any(), tvhq).
		Return(&tvheadend.Response{
			Response: &http.Response{
				StatusCode: 200,
				Body:       http.NoBody,
			}}, nil).
		Times(1)

	service := recording.New(mockClient)
	err := service.Remove(ctx, id)

	assert.Nil(t, err)
}

func TestMoveFinishedReturnsError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_tvheadend.NewMockClient(ctrl)
	mockClient.EXPECT().
		Exec(ctx, "/api/dvr/entry/move/finished", gomock.Any(), gomock.Any()).
		DoAndReturn(mock_tvheadend.MockClientExecReturnsError).
		Times(1)

	service := recording.New(mockClient)
	err := service.MoveFinished(ctx, "someID")

	assert.EqualError(t, err, "error")
}

func TestMoveFinishedReturnsRequestFailedError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_tvheadend.NewMockClient(ctrl)
	mockClient.EXPECT().
		Exec(ctx, "/api/dvr/entry/move/finished", gomock.Any(), gomock.Any()).
		DoAndReturn(mock_tvheadend.MockClientExecReturnsErroneousHttpStatus).
		Times(1)

	service := recording.New(mockClient)
	err := service.MoveFinished(ctx, "someID")

	assert.Equal(t, err, recording.ErrRequestFailed)
}

func TestMoveFinishedSucceeds(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	id := "someID"

	tvhq := tvheadend.NewQuery()
	tvhq.Set("uuid", id)

	mockClient := mock_tvheadend.NewMockClient(ctrl)
	mockClient.EXPECT().
		Exec(ctx, "/api/dvr/entry/move/finished", gomock.Any(), tvhq).
		Return(&tvheadend.Response{
			Response: &http.Response{
				StatusCode: 200,
				Body:       http.NoBody,
			}}, nil).
		Times(1)

	service := recording.New(mockClient)
	err := service.MoveFinished(ctx, id)

	assert.Nil(t, err)
}

func TestMoveFailedReturnsError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_tvheadend.NewMockClient(ctrl)
	mockClient.EXPECT().
		Exec(ctx, "/api/dvr/entry/move/failed", gomock.Any(), gomock.Any()).
		DoAndReturn(mock_tvheadend.MockClientExecReturnsError).
		Times(1)

	service := recording.New(mockClient)
	err := service.MoveFailed(ctx, "someID")

	assert.EqualError(t, err, "error")
}

func TestMoveFailedReturnsRequestFailedError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_tvheadend.NewMockClient(ctrl)
	mockClient.EXPECT().
		Exec(ctx, "/api/dvr/entry/move/failed", gomock.Any(), gomock.Any()).
		DoAndReturn(mock_tvheadend.MockClientExecReturnsErroneousHttpStatus).
		Times(1)

	service := recording.New(mockClient)
	err := service.MoveFailed(ctx, "someID")

	assert.Equal(t, err, recording.ErrRequestFailed)
}

func TestMoveFailedSucceeds(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	id := "someID"

	tvhq := tvheadend.NewQuery()
	tvhq.Set("uuid", id)

	mockClient := mock_tvheadend.NewMockClient(ctrl)
	mockClient.EXPECT().
		Exec(ctx, "/api/dvr/entry/move/failed", gomock.Any(), tvhq).
		Return(&tvheadend.Response{
			Response: &http.Response{
				StatusCode: 200,
				Body:       http.NoBody,
			}}, nil).
		Times(1)

	service := recording.New(mockClient)
	err := service.MoveFailed(ctx, id)

	assert.Nil(t, err)
}

func TestUpdateRecordingReturnsErrorWhenIdnodeLoadFails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_tvheadend.NewMockClient(ctrl)
	mockClient.EXPECT().
		Exec(ctx, "/api/idnode/load", gomock.Any(), gomock.Any()).
		DoAndReturn(mock_tvheadend.MockClientExecReturnsError).
		Times(1)

	service := recording.New(mockClient)
	err := service.UpdateRecording(ctx, "someID", core.UpdateRecording{})

	assert.EqualError(t, err, "error")
}

func TestUpdateRecordingReturnsError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_tvheadend.NewMockClient(ctrl)
	mockClient.EXPECT().
		Exec(ctx, "/api/idnode/load", gomock.Any(), gomock.Any()).
		DoAndReturn(mockClientExecSucceedsForGet).
		Times(1)
	mockClient.EXPECT().
		Exec(ctx, "/api/idnode/save", gomock.Any(), gomock.Any()).
		DoAndReturn(mock_tvheadend.MockClientExecReturnsError).
		Times(1)

	service := recording.New(mockClient)
	err := service.UpdateRecording(ctx, "someID", core.UpdateRecording{})

	assert.EqualError(t, err, "error")
}

func TestUpdateRecordingReturnsRequestFailedError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_tvheadend.NewMockClient(ctrl)
	mockClient.EXPECT().
		Exec(ctx, "/api/idnode/load", gomock.Any(), gomock.Any()).
		DoAndReturn(mockClientExecSucceedsForGet).
		Times(1)
	mockClient.EXPECT().
		Exec(ctx, "/api/idnode/save", gomock.Any(), gomock.Any()).
		DoAndReturn(mock_tvheadend.MockClientExecReturnsErroneousHttpStatus).
		Times(1)

	service := recording.New(mockClient)
	err := service.UpdateRecording(ctx, "someID", core.UpdateRecording{})

	assert.Equal(t, err, recording.ErrRequestFailed)
}

func TestUpdateRecordingSucceeds(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	id := "someID"

	mockClient := mock_tvheadend.NewMockClient(ctrl)
	mockClient.EXPECT().
		Exec(ctx, "/api/idnode/load", gomock.Any(), gomock.Any()).
		DoAndReturn(mockClientExecSucceedsForGet).
		Times(1)

	mockClient.EXPECT().
		Exec(ctx, "/api/idnode/save", gomock.Any(), gomock.Any()).
		Return(&tvheadend.Response{
			Response: &http.Response{
				StatusCode: 200,
				Body:       http.NoBody,
			}}, nil).
		Times(1)

	service := recording.New(mockClient)
	err := service.UpdateRecording(ctx, id, core.UpdateRecording{})

	assert.Nil(t, err)
}

func TestGetReturnsRequestFailedError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_tvheadend.NewMockClient(ctrl)
	mockClient.EXPECT().
		Exec(ctx, "/api/idnode/load", gomock.Any(), gomock.Any()).
		DoAndReturn(mock_tvheadend.MockClientExecReturnsErroneousHttpStatus).
		Times(1)

	service := recording.New(mockClient)
	rec, err := service.Get(ctx, "someID")

	assert.Equal(t, err, recording.ErrRequestFailed)
	assert.Nil(t, rec)
}

func TestGetReturnsRecordingNotFoundError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_tvheadend.NewMockClient(ctrl)
	mockClient.EXPECT().
		Exec(ctx, "/api/idnode/load", gomock.Any(), gomock.Any()).
		DoAndReturn(mockClientExecReturnsEmptyEntriesForGet).
		Times(1)

	service := recording.New(mockClient)
	rec, err := service.Get(ctx, "someID")

	assert.Equal(t, err, core.ErrRecordingNotFound)
	assert.Nil(t, rec)
}

func TestGetSucceeds(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_tvheadend.NewMockClient(ctrl)
	mockClient.EXPECT().
		Exec(ctx, "/api/idnode/load", gomock.Any(), gomock.Any()).
		DoAndReturn(mockClientExecSucceedsForGet).
		Times(1)

	service := recording.New(mockClient)
	rec, err := service.Get(ctx, "someID")

	assert.Nil(t, err)
	assert.NotNil(t, rec)
}
