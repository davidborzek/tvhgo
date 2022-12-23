package core

import (
	"context"
	"net/http"
)

type (
	PiconService interface {
		// GetPicon returns the raw http response of the channels picon.
		Get(ctx context.Context, id int) (*http.Response, error)
	}
)
