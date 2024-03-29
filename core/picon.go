package core

import (
	"context"
	"errors"
	"io"
)

var (
	ErrPiconNotFound = errors.New("picon not found")
)

type (
	PiconService interface {
		// GetPicon returns the picon of a channel.
		Get(ctx context.Context, id int) (io.Reader, error)
	}
)
