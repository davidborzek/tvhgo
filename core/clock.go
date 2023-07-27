package core

import "time"

type Clock interface {
	Now() time.Time
}
