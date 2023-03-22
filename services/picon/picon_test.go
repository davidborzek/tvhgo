package picon_test

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/davidborzek/tvhgo/core"
	mock_tvheadend "github.com/davidborzek/tvhgo/mock/tvheadend"
	"github.com/davidborzek/tvhgo/services/picon"
	"github.com/davidborzek/tvhgo/tvheadend"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
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
			Body: io.NopCloser(
				bytes.NewBufferString(""),
			),
		},
	}
	return res, nil
}

func mockClientExecFailsWith404(
	ctx context.Context,
	path string,
	dst interface{},
	query ...tvheadend.Query,
) (*tvheadend.Response, error) {
	res := &tvheadend.Response{
		Response: &http.Response{
			StatusCode: 404,
		},
	}
	return res, nil
}

func mockClientExecFailsWith501(
	ctx context.Context,
	path string,
	dst interface{},
	query ...tvheadend.Query,
) (*tvheadend.Response, error) {
	res := &tvheadend.Response{
		Response: &http.Response{
			StatusCode: 501,
		},
	}
	return res, nil
}

func TestGetReturnsError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.TODO()

	mockClient := mock_tvheadend.NewMockClient(ctrl)
	mockClient.EXPECT().
		Exec(ctx, "/imagecache/2", gomock.Any(), gomock.Any()).
		DoAndReturn(mock_tvheadend.MockClientExecReturnsError).
		Times(1)

	service := picon.New(mockClient)
	res, err := service.Get(ctx, 2)

	assert.Nil(t, res)
	assert.EqualError(t, err, "error")
}

func TestReturnsErrPiconNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.TODO()

	mockClient := mock_tvheadend.NewMockClient(ctrl)
	mockClient.EXPECT().
		Exec(ctx, "/imagecache/2", gomock.Any(), gomock.Any()).
		DoAndReturn(mockClientExecFailsWith404).
		Times(1)

	service := picon.New(mockClient)
	res, err := service.Get(ctx, 2)

	assert.Nil(t, res)
	assert.Equal(t, err, core.ErrPiconNotFound)
}

func TestReturnsErrRequestFailed(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.TODO()

	mockClient := mock_tvheadend.NewMockClient(ctrl)
	mockClient.EXPECT().
		Exec(ctx, "/imagecache/2", gomock.Any(), gomock.Any()).
		DoAndReturn(mockClientExecFailsWith501).
		Times(1)

	service := picon.New(mockClient)
	res, err := service.Get(ctx, 2)

	assert.Nil(t, res)
	assert.Equal(t, err, picon.ErrRequestFailed)
}

func TestGetSucceeds(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.TODO()

	mockClient := mock_tvheadend.NewMockClient(ctrl)
	mockClient.EXPECT().
		Exec(ctx, "/imagecache/2", gomock.Any(), gomock.Any()).
		DoAndReturn(mockClientExecSucceeds).
		Times(1)

	service := picon.New(mockClient)
	res, err := service.Get(ctx, 2)

	assert.NotNil(t, res)
	assert.Nil(t, err)
}
