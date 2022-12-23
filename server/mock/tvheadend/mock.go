//go:build !prod
// +build !prod

package mock_tvheadend

import (
	context "context"
	"errors"
	"net/http"

	tvheadend "github.com/davidborzek/tvhgo/tvheadend"
)

func MockClientExecReturnsError(
	ctx context.Context,
	path string,
	dst interface{},
	query ...tvheadend.Query,
) (*tvheadend.Response, error) {
	return nil, errors.New("error")
}

func MockClientExecReturnsErroneousHttpStatus(
	ctx context.Context,
	path string,
	dst interface{},
	query ...tvheadend.Query,
) (*tvheadend.Response, error) {
	res := &tvheadend.Response{
		Response: &http.Response{
			StatusCode: 500,
			Body:       http.NoBody,
		},
	}

	return res, nil
}

//go:generate mockgen -destination=mock_gen.go github.com/davidborzek/tvhgo/tvheadend Client
