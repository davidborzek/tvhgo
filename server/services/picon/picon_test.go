package picon_test

import (
	"context"
	"net/http"
	"testing"

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
