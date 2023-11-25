package channel_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/davidborzek/tvhgo/core"
	mock_tvheadend "github.com/davidborzek/tvhgo/mock/tvheadend"
	"github.com/davidborzek/tvhgo/services/channel"
	"github.com/davidborzek/tvhgo/tvheadend"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

var (
	channelGridEntry = tvheadend.ChannelGridEntry{
		UUID:          "someId",
		Enabled:       true,
		Autoname:      true,
		Name:          "someName",
		Number:        1,
		Icon:          "someIcon",
		IconPublicURL: "imagecache/3",
		EpgAuto:       true,
	}

	ctx = context.TODO()
)

func mockClientExecSucceeds(
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

	g := dst.(*tvheadend.ChannelGrid)
	g.Entries = []tvheadend.ChannelGridEntry{
		channelGridEntry,
	}
	g.Total = 20

	return res, nil
}

func TestGetAllReturnsError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_tvheadend.NewMockClient(ctrl)
	mockClient.EXPECT().
		Exec(ctx, "/api/channel/grid", gomock.Any(), gomock.Any()).
		DoAndReturn(mock_tvheadend.MockClientExecReturnsError).
		Times(1)

	service := channel.New(mockClient)
	res, err := service.GetAll(ctx, core.PaginationSortQueryParams{})

	assert.Nil(t, res)
	assert.EqualError(t, err, "error")
}

func TestGetAllReturnsRequestFailedError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_tvheadend.NewMockClient(ctrl)
	mockClient.EXPECT().
		Exec(ctx, "/api/channel/grid", gomock.Any(), gomock.Any()).
		DoAndReturn(mock_tvheadend.MockClientExecReturnsErroneousHttpStatus).
		Times(1)

	service := channel.New(mockClient)
	res, err := service.GetAll(ctx, core.PaginationSortQueryParams{})

	assert.Nil(t, res)
	assert.Equal(t, err, channel.ErrRequestFailed)
}

func TestGetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tvhq := tvheadend.NewQuery()
	tvhq.Limit(10)
	tvhq.Start(5)
	tvhq.SortKey("name")
	tvhq.SortDir("asc")

	mockClient := mock_tvheadend.NewMockClient(ctrl)
	mockClient.EXPECT().
		Exec(ctx, "/api/channel/grid", gomock.Any(), tvhq).
		DoAndReturn(mockClientExecSucceeds).
		Times(1)

	service := channel.New(mockClient)

	q := core.PaginationSortQueryParams{}
	q.Limit = 10
	q.Offset = 5
	q.SortDirection = "asc"
	q.SortKey = "name"

	channels, err := service.GetAll(ctx, q)

	assert.Nil(t, err)
	assert.NotNil(t, channels)

	assert.Equal(t, channelGridEntry.UUID, channels[0].ID)
	assert.Equal(t, channelGridEntry.Name, channels[0].Name)
	assert.Equal(t, channelGridEntry.Enabled, channels[0].Enabled)
	assert.Equal(t, channelGridEntry.Number, channels[0].Number)
	assert.Equal(t, 3, channels[0].PiconID)
}
