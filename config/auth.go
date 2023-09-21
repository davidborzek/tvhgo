package config

import "time"

const (
	defaultSessionCookieName              = "tvhgo_session"
	defaultSessionMaximumInactiveLifeTime = 7 * 24 * time.Hour
	defaultSessionMaximumLifetime         = 30 * 24 * time.Hour
	defaultSessionTokenRotationInterval   = 30 * time.Minute
	defaultSessionCleanupInterval         = 12 * time.Hour
	defaultTOTPIssuer                     = "tvhgo"
)

type (
	SessionConfig struct {
		CookieName              string        `yaml:"cookie_name"               env:"COOKIE_NAME"`
		CookieSecure            bool          `yaml:"cookie_secure"             env:"COOKIE_SECURE"`
		MaximumInactiveLifetime time.Duration `yaml:"maximum_inactive_lifetime" env:"MAXIMUM_INACTIVE_LIFETIME"`
		MaximumLifetime         time.Duration `yaml:"maximum_lifetime"          env:"MAXIMUM_LIFETIME"`
		TokenRotationInterval   time.Duration `yaml:"token_rotation_interval"   env:"TOKEN_ROTATION_INTERVAL"`
		CleanupInterval         time.Duration `yaml:"cleanup_interval"          env:"CLEANUP_INTERVAL"`
	}

	TOTPConfig struct {
		Issuer string `yaml:"issuer" env:"ISSUER"`
	}

	AuthConfig struct {
		Session SessionConfig `yaml:"session" envPrefix:"SESSION_"`
		TOTP    TOTPConfig    `yaml:"totp"    envPrefix:"TOTP_"`
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
	if s.CleanupInterval == 0 {
		s.CleanupInterval = defaultSessionCleanupInterval
	}
}

func (c *TOTPConfig) SetDefaults() {
	if c.Issuer == "" {
		c.Issuer = defaultTOTPIssuer
	}
}
