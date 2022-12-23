package core

import (
	"context"
)

type (
	// Channel defines a channel in tvheadend.
	Channel struct {
		ID      string `json:"id"`
		Enabled bool   `json:"enabled"`
		Name    string `json:"name"`
		Number  int    `json:"number"`
		PiconID int    `json:"piconId"`
	}

	// ChannelService provides access to channel
	// resources from the tvheadend server.
	ChannelService interface {
		// GetAll returns a list of channels.
		GetAll(ctx context.Context, params PaginationSortQueryParams) ([]*Channel, error)
	}
)
