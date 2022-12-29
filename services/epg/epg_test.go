package epg_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/davidborzek/tvhgo/core"
	mock_tvheadend "github.com/davidborzek/tvhgo/mock/tvheadend"
	"github.com/davidborzek/tvhgo/services/epg"
	"github.com/davidborzek/tvhgo/tvheadend"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var (
	epgEventGridEntry = tvheadend.EpgEventGridEntry{
		EventID: 1234,
	}

	ctx = context.TODO()
)

func mockClientReturnsEvents(
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

	g := dst.(*tvheadend.EpgEventGrid)
	g.Entries = []tvheadend.EpgEventGridEntry{
		epgEventGridEntry,
	}
	g.Total = 20

	return res, nil
}

func mockClientReturnsEmptyEntries(
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

	g := dst.(*tvheadend.EpgEventGrid)
	g.Entries = []tvheadend.EpgEventGridEntry{}
	g.Total = 0

	return res, nil
}

func TestGetEventsReturnsError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_tvheadend.NewMockClient(ctrl)
	mockClient.EXPECT().
		Exec(ctx, "/api/epg/events/grid", gomock.Any(), gomock.Any()).
		DoAndReturn(mock_tvheadend.MockClientExecReturnsError).
		Times(1)

	service := epg.New(mockClient)
	res, err := service.GetEvents(ctx, core.GetEpgQueryParams{})

	assert.Nil(t, res)
	assert.EqualError(t, err, "error")
}

func TestGetEventsReturnsRequestFailedError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_tvheadend.NewMockClient(ctrl)
	mockClient.EXPECT().
		Exec(ctx, "/api/epg/events/grid", gomock.Any(), gomock.Any()).
		DoAndReturn(mock_tvheadend.MockClientExecReturnsErroneousHttpStatus).
		Times(1)

	service := epg.New(mockClient)
	res, err := service.GetEvents(ctx, core.GetEpgQueryParams{})

	assert.Nil(t, res)
	assert.Equal(t, err, epg.ErrRequestFailed)
}

func TestGetEvents(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tvhq := tvheadend.NewQuery()
	tvhq.Limit(10)
	tvhq.Start(5)
	tvhq.SortKey("title")
	tvhq.SortDir("asc")
	tvhq.Set("fulltext", "0")
	tvhq.Set("mode", "all")

	mockClient := mock_tvheadend.NewMockClient(ctrl)
	mockClient.EXPECT().
		Exec(ctx, "/api/epg/events/grid", gomock.Any(), tvhq).
		DoAndReturn(mockClientReturnsEvents).
		Times(1)

	service := epg.New(mockClient)

	q := core.GetEpgQueryParams{}
	q.Limit = 10
	q.Offset = 5
	q.SortDirection = "asc"
	q.SortKey = "title"

	channels, err := service.GetEvents(ctx, q)

	assert.Nil(t, err)
	assert.NotNil(t, channels)
}

func TestGetEventsMapsSortKeyCorrectly(t *testing.T) {
	keys := map[string]string{
		"title":         "title",
		"subtitle":      "subtitle",
		"startsAt":      "start",
		"endsAt":        "stop",
		"channelName":   "channelName",
		"channelNumber": "channelNumber",
		"description":   "description",
	}

	for key, mappedKey := range keys {
		t.Run(key, testGetEventsMapsSortKeyCorrectlyParametrize(key, mappedKey))
	}
}

func testGetEventsMapsSortKeyCorrectlyParametrize(key string, mappedKey string) func(t *testing.T) {
	return func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		tvhq := tvheadend.NewQuery()
		tvhq.SortKey(mappedKey)
		tvhq.Set("fulltext", "0")
		tvhq.Set("mode", "all")

		mockClient := mock_tvheadend.NewMockClient(ctrl)
		mockClient.EXPECT().
			Exec(ctx, "/api/epg/events/grid", gomock.Any(), tvhq).
			DoAndReturn(mockClientReturnsEvents).
			Times(1)

		q := core.GetEpgQueryParams{}
		q.SortKey = key

		service := epg.New(mockClient)
		res, err := service.GetEvents(ctx, q)

		assert.NotNil(t, res)
		assert.Nil(t, err)
	}
}

func TestGetEventReturnsError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_tvheadend.NewMockClient(ctrl)
	mockClient.EXPECT().
		Exec(ctx, "/api/epg/events/load", gomock.Any(), gomock.Any()).
		DoAndReturn(mock_tvheadend.MockClientExecReturnsError).
		Times(1)

	service := epg.New(mockClient)
	res, err := service.GetEvent(ctx, int64(1234))

	assert.Nil(t, res)
	assert.EqualError(t, err, "error")
}

func TestGetEventReturnsRequestFailedError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_tvheadend.NewMockClient(ctrl)
	mockClient.EXPECT().
		Exec(ctx, "/api/epg/events/load", gomock.Any(), gomock.Any()).
		DoAndReturn(mock_tvheadend.MockClientExecReturnsErroneousHttpStatus).
		Times(1)

	service := epg.New(mockClient)
	res, err := service.GetEvent(ctx, int64(1234))

	assert.Nil(t, res)
	assert.Equal(t, err, epg.ErrRequestFailed)
}

func TestGetEventReturnsEpgEventNotFoundError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_tvheadend.NewMockClient(ctrl)
	mockClient.EXPECT().
		Exec(ctx, "/api/epg/events/load", gomock.Any(), gomock.Any()).
		DoAndReturn(mockClientReturnsEmptyEntries).
		Times(1)

	service := epg.New(mockClient)
	res, err := service.GetEvent(ctx, int64(1234))

	assert.Nil(t, res)
	assert.Equal(t, err, core.ErrEpgEventNotFound)
}

func TestGetEventSucceeds(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_tvheadend.NewMockClient(ctrl)
	mockClient.EXPECT().
		Exec(ctx, "/api/epg/events/load", gomock.Any(), gomock.Any()).
		DoAndReturn(mockClientReturnsEvents).
		Times(1)

	service := epg.New(mockClient)
	res, err := service.GetEvent(ctx, int64(1234))

	assert.NotNil(t, res)
	assert.Nil(t, err)
}
