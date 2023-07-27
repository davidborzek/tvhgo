package clock

import (
	"time"

	"github.com/davidborzek/tvhgo/core"
)

func NewClock() core.Clock {
	return &clockImpl{}
}

type clockImpl struct{}

func (clockImpl) Now() time.Time {
	return time.Now()
}
