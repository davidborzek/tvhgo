package core_test

import (
	"testing"

	"github.com/davidborzek/tvhgo/conv"
	"github.com/davidborzek/tvhgo/core"
	"github.com/davidborzek/tvhgo/tvheadend"
	"github.com/stretchr/testify/assert"
)

func TestMapTvheadendIconUrlToPiconID(t *testing.T) {
	piconId := core.MapTvheadendIconUrlToPiconID("imagecache/223")
	assert.Equal(t, 223, piconId)
}

func TestMapTvheadendIconUrlToPiconIDReturnsZeroOnError(t *testing.T) {
	piconId := core.MapTvheadendIconUrlToPiconID("erroneousUrl")
	assert.Equal(t, 0, piconId)
}

func TestMapTvheadendIdnodeToChannel(t *testing.T) {
	enabled := true
	name := "someName"
	number := 123
	iconUrl := "imagecache/223"

	idnode := tvheadend.Idnode{
		UUID: "someID",
		Params: []tvheadend.InodeParams{
			{
				ID:    "enabled",
				Value: enabled,
			},
			{
				ID:    "name",
				Value: name,
			},
			{
				ID:    "number",
				Value: float64(number),
			},
			{
				ID:    "icon_public_url",
				Value: iconUrl,
			},
		},
	}

	channel, err := core.MapTvheadendIdnodeToChannel(idnode)

	assert.Nil(t, err)
	assert.Equal(t, enabled, channel.Enabled)
	assert.Equal(t, name, channel.Name)
	assert.Equal(t, number, channel.Number)
	assert.Equal(t, 223, channel.PiconID)
}

func TestMapTvheadendIdnodeToChannelFailsForUnexpectedType(t *testing.T) {
	idnode := tvheadend.Idnode{
		UUID: "someID",
		Params: []tvheadend.InodeParams{
			{
				ID:    "enabled",
				Value: "true",
			},
		},
	}

	channel, err := core.MapTvheadendIdnodeToChannel(idnode)

	assert.Nil(t, channel)
	assert.Equal(t, conv.ErrInterfaceToBool, err)
}
