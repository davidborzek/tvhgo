package core

import (
	"context"

	"github.com/davidborzek/tvhgo/tvheadend"
)

type (
	StreamProfile struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}

	ProfileService interface {
		GetStreamProfiles(ctx context.Context) ([]StreamProfile, error)
	}
)

func NewStreamProfile(profile tvheadend.StreamProfile) StreamProfile {
	return StreamProfile{
		ID:   profile.Key,
		Name: profile.Val,
	}
}
