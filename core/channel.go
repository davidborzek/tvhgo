package core

import (
	"context"
	"errors"
	"strconv"
	"strings"

	"github.com/davidborzek/tvhgo/conv"
	"github.com/davidborzek/tvhgo/tvheadend"
)

var (
	ErrChannelNotFound = errors.New("channel not found")
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

	GetChannelsParams struct {
		PaginationSortQueryParams
		// Name of a channel.
		Name string `schema:"name"`
	}

	// ChannelService provides access to channel
	// resources from the tvheadend server.
	ChannelService interface {
		// GetAll returns a list of channels.
		// TODO: return ListResult with limit, offset, ...
		GetAll(ctx context.Context, params GetChannelsParams) ([]*Channel, error)
		// Get returns a channel by id.
		Get(ctx context.Context, id string) (*Channel, error)
	}
)

func (p *GetChannelsParams) MapToTvheadendQuery(
	sortKeyMapping map[string]string,
) (*tvheadend.Query, error) {
	q := p.PaginationSortQueryParams.MapToTvheadendQuery(sortKeyMapping)

	if p.Name != "" {
		err := q.Filter([]tvheadend.FilterQuery{
			{
				Field:      "name",
				Type:       "string",
				Value:      p.Name,
				Comparison: "eq",
			},
		})

		if err != nil {
			return nil, err
		}
	}

	return &q, nil
}

func MapTvheadendIconUrlToPiconID(iconUrl string) int {
	split := strings.Split(iconUrl, "/")

	var piconID int
	if len(split) == 2 {
		piconID, _ = strconv.Atoi(split[1])
	}

	return piconID
}

// MapTvheadendIdnodeToChannel maps a tvheadend.Idnode to a Channel.
func MapTvheadendIdnodeToChannel(idnode tvheadend.Idnode) (*Channel, error) {
	r := Channel{
		ID: idnode.UUID,
	}

	for _, p := range idnode.Params {
		var err error

		switch p.ID {
		case "enabled":
			r.Enabled, err = conv.InterfaceToBool(p.Value)
		case "name":
			r.Name, err = conv.InterfaceToString(p.Value)
		case "number":
			r.Number, err = conv.InterfaceToInt(p.Value)
		case "icon_public_url":
			var value string
			value, err = conv.InterfaceToString(p.Value)
			r.PiconID = MapTvheadendIconUrlToPiconID(value)
		}

		if err != nil {
			return nil, err
		}
	}

	return &r, nil
}
