package config

import "time"

const (
	defaultSessionCookieName              = "tvhgo_session"
	defaultSessionMaximumInactiveLifeTime = 7 * 24 * time.Hour
	defaultSessionMaximumLifetime         = 30 * 24 * time.Hour
	defaultSessionTokenRotationInterval   = 30 * time.Minute
)

type (
	SessionConfig struct {
		CookieName              string        `yaml:"cookie_name"`
		MaximumInactiveLifetime time.Duration `yaml:"maximum_inactive_lifetime"`
		MaximumLifetime         time.Duration `yaml:"maximum_lifetime"`
		TokenRotationInterval   time.Duration `yaml:"token_rotation_interval"`
	}

	AuthConfig struct {
		Session SessionConfig `yaml:"session"`
	}
)

func (s *SessionConfig) SetDefaults() {
	if s.CookieName == "" {
		s.CookieName = defaultSessionCookieName
	}
	if s.MaximumInactiveLifetime == 0 {
		s.MaximumInactiveLifetime = defaultSessionMaximumInactiveLifeTime
	}
	if s.MaximumLifetime == 0 {
		s.MaximumLifetime = defaultSessionMaximumLifetime
	}
	if s.TokenRotationInterval == 0 {
		s.TokenRotationInterval = defaultSessionTokenRotationInterval
	}
}
