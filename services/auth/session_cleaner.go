package auth

import (
	"context"
	"time"

	"github.com/davidborzek/tvhgo/core"
	log "github.com/sirupsen/logrus"
)

type sessionCleaner struct {
	sessions         core.SessionRepository
	clock            core.Clock
	interval         time.Duration
	inactiveLifetime time.Duration
	lifetime         time.Duration
}

func NewSessionCleaner(
	sessions core.SessionRepository,
	clock core.Clock,
	interval time.Duration,
	inactiveLifetime time.Duration,
	lifetime time.Duration,
) *sessionCleaner {
	return &sessionCleaner{
		sessions:         sessions,
		clock:            clock,
		interval:         interval,
		lifetime:         lifetime,
		inactiveLifetime: inactiveLifetime,
	}
}

func (s *sessionCleaner) Start() {
	log.WithField("interval", s.interval).
		Info("starting session cleaner")
	ticker := time.NewTicker(s.interval)

	s.RunNow()

	go func() {
		for {
			<-ticker.C
			log.Debug("running scheduled session cleanup")
			s.RunNow()
		}
	}()
}

func (s *sessionCleaner) RunNow() {
	expirationDate := s.clock.Now().Add(-s.lifetime)
	inActiveExpirationDate := s.clock.Now().Add(-s.inactiveLifetime)

	rows, err := s.sessions.DeleteExpired(
		context.Background(),
		expirationDate.Unix(),
		inActiveExpirationDate.Unix(),
	)

	if err != nil {
		log.WithError(err).Error("failed to cleanup expired sessions")
	} else if rows > 0 {
		log.WithField("sessions", rows).Debug("cleaned up expired sessions")
	}
}
