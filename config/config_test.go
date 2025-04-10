package config_test

import (
	"os"
	"testing"
	"time"

	"github.com/davidborzek/tvhgo/config"
	"github.com/stretchr/testify/assert"
)

func TestLoadRequiredConfigFromEnv(t *testing.T) {
	defer os.Clearenv()
	os.Setenv("TVHGO_TVHEADEND_HOST", "localhost")
	cfg, err := config.Load("")

	assert.Nil(t, err)
	assert.Equal(t, "localhost", cfg.Tvheadend.Host)
	assert.Equal(t, 9981, cfg.Tvheadend.Port)
	assert.Equal(t, "http", cfg.Tvheadend.Scheme)
	assert.Empty(t, cfg.Tvheadend.Username)
	assert.Empty(t, cfg.Tvheadend.Password)

	assert.Empty(t, cfg.Server.Host)
	assert.Equal(t, 8080, cfg.Server.Port)

	assert.Equal(t, "./tvhgo.db", cfg.Database.Path)
	assert.Equal(t, config.DatabaseTypeSqlite, cfg.Database.Type)
	assert.Equal(t, "tvhgo", cfg.Database.User)
	assert.Equal(t, "tvhgo", cfg.Database.Database)
	assert.Equal(t, "localhost", cfg.Database.Host)

	assert.Equal(t, "tvhgo", cfg.Auth.TOTP.Issuer)
	assert.Equal(t, "tvhgo_session", cfg.Auth.Session.CookieName)
	assert.False(t, cfg.Auth.Session.CookieSecure)
	assert.Equal(t, 7*24*time.Hour, cfg.Auth.Session.MaximumInactiveLifetime)
	assert.Equal(t, 30*24*time.Hour, cfg.Auth.Session.MaximumLifetime)
	assert.Equal(t, 30*time.Minute, cfg.Auth.Session.TokenRotationInterval)
	assert.Equal(t, 12*time.Hour, cfg.Auth.Session.CleanupInterval)

	assert.False(t, cfg.Metrics.Enabled)
	assert.Equal(t, "/metrics", cfg.Metrics.Path)
	assert.Empty(t, cfg.Metrics.Token)
	assert.Equal(t, 8081, cfg.Metrics.Port)
	assert.Empty(t, cfg.Metrics.Host)

	assert.Equal(t, "console", cfg.Log.Format)
	assert.Equal(t, "info", cfg.Log.Level)

	assert.False(t, cfg.Auth.ReverseProxy.Enabled)
	assert.Equal(t, "Remote-User", cfg.Auth.ReverseProxy.UserHeader)
	assert.Equal(t, "Remote-Email", cfg.Auth.ReverseProxy.EmailHeader)
	assert.Equal(t, "Remote-Name", cfg.Auth.ReverseProxy.NameHeader)
	assert.Empty(t, cfg.Auth.ReverseProxy.AllowedProxies)
	assert.False(t, cfg.Auth.ReverseProxy.AllowRegistration)
}

func TestLoadFailsForNoTvheadendHost(t *testing.T) {
	defer os.Clearenv()
	cfg, err := config.Load("")

	assert.EqualError(t, err, "tvheadend host is not set")
	assert.Nil(t, cfg)
}

func TestLoadFailsForWhenSamePortForServerAndMetricsIsSet(t *testing.T) {
	defer os.Clearenv()
	os.Setenv("TVHGO_TVHEADEND_HOST", "localhost")
	os.Setenv("TVHGO_SERVER_PORT", "9999")
	os.Setenv("TVHGO_METRICS_PORT", "9999")

	cfg, err := config.Load("")

	assert.EqualError(t, err, "metrics and server port cannot be the same")
	assert.Nil(t, cfg)
}

func TestLoadConfigFromEnv(t *testing.T) {
	defer os.Clearenv()
	os.Setenv("TVHGO_TVHEADEND_HOST", "localhost")
	os.Setenv("TVHGO_TVHEADEND_PORT", "1234")
	os.Setenv("TVHGO_TVHEADEND_SCHEME", "https")
	os.Setenv("TVHGO_TVHEADEND_USERNAME", "someTvheadendUsername")
	os.Setenv("TVHGO_TVHEADEND_PASSWORD", "password")

	os.Setenv("TVHGO_SERVER_HOST", "0.0.0.0")
	os.Setenv("TVHGO_SERVER_PORT", "9999")
	os.Setenv("TVHGO_SERVER_SWAGGER_UI_ENABLED", "true")

	os.Setenv("TVHGO_DATABASE_PATH", "/tmp/tvhgo.db")

	os.Setenv("TVHGO_AUTH_TOTP_ISSUER", "myTvhgo")
	os.Setenv("TVHGO_AUTH_SESSION_COOKIE_NAME", "myTvhgo_session")
	os.Setenv("TVHGO_AUTH_SESSION_COOKIE_SECURE", "true")
	os.Setenv("TVHGO_AUTH_SESSION_MAXIMUM_INACTIVE_LIFETIME", "100h")
	os.Setenv("TVHGO_AUTH_SESSION_MAXIMUM_LIFETIME", "200h")
	os.Setenv("TVHGO_AUTH_SESSION_TOKEN_ROTATION_INTERVAL", "1h")
	os.Setenv("TVHGO_AUTH_SESSION_CLEANUP_INTERVAL", "5h")

	os.Setenv("TVHGO_METRICS_ENABLED", "true")
	os.Setenv("TVHGO_METRICS_PATH", "/prometheus")
	os.Setenv("TVHGO_METRICS_TOKEN", "someMetricsToken")
	os.Setenv("TVHGO_METRICS_PORT", "8082")
	os.Setenv("TVHGO_METRICS_HOST", "0.0.0.0")

	os.Setenv("TVHGO_LOG_FORMAT", "json")
	os.Setenv("TVHGO_LOG_LEVEL", "debug")

	os.Setenv("TVHGO_AUTH_REVERSE_PROXY_ENABLED", "true")
	os.Setenv("TVHGO_AUTH_REVERSE_PROXY_USER_HEADER", "X-Remote-User")
	os.Setenv("TVHGO_AUTH_REVERSE_PROXY_EMAIL_HEADER", "X-Remote-Email")
	os.Setenv("TVHGO_AUTH_REVERSE_PROXY_NAME_HEADER", "X-Remote-Name")
	os.Setenv("TVHGO_AUTH_REVERSE_PROXY_ALLOWED_PROXIES", "127.0.0.1/24,127.0.0.1")
	os.Setenv("TVHGO_AUTH_REVERSE_PROXY_ALLOW_REGISTRATION", "true")

	cfg, err := config.Load("")

	assert.Nil(t, err)
	assert.Equal(t, "localhost", cfg.Tvheadend.Host)
	assert.Equal(t, 1234, cfg.Tvheadend.Port)
	assert.Equal(t, "https", cfg.Tvheadend.Scheme)
	assert.Equal(t, "someTvheadendUsername", cfg.Tvheadend.Username)
	assert.Equal(t, "password", cfg.Tvheadend.Password)

	enableSwaggerUI := true
	assert.Equal(t, "0.0.0.0", cfg.Server.Host)
	assert.Equal(t, 9999, cfg.Server.Port)
	assert.Equal(t, &enableSwaggerUI, cfg.Server.SwaggerUI.Enabled)

	assert.Equal(t, "/tmp/tvhgo.db", cfg.Database.Path)

	assert.Equal(t, "myTvhgo", cfg.Auth.TOTP.Issuer)
	assert.Equal(t, "myTvhgo_session", cfg.Auth.Session.CookieName)
	assert.True(t, cfg.Auth.Session.CookieSecure)
	assert.Equal(t, 100*time.Hour, cfg.Auth.Session.MaximumInactiveLifetime)
	assert.Equal(t, 200*time.Hour, cfg.Auth.Session.MaximumLifetime)
	assert.Equal(t, 1*time.Hour, cfg.Auth.Session.TokenRotationInterval)
	assert.Equal(t, 5*time.Hour, cfg.Auth.Session.CleanupInterval)

	assert.True(t, cfg.Metrics.Enabled)
	assert.Equal(t, "/prometheus", cfg.Metrics.Path)
	assert.Equal(t, "someMetricsToken", cfg.Metrics.Token)
	assert.Equal(t, 8082, cfg.Metrics.Port)
	assert.Equal(t, "0.0.0.0", cfg.Metrics.Host)

	assert.Equal(t, "json", cfg.Log.Format)
	assert.Equal(t, "debug", cfg.Log.Level)

	assert.True(t, cfg.Auth.ReverseProxy.Enabled)
	assert.Equal(t, "X-Remote-User", cfg.Auth.ReverseProxy.UserHeader)
	assert.Equal(t, "X-Remote-Email", cfg.Auth.ReverseProxy.EmailHeader)
	assert.Equal(t, "X-Remote-Name", cfg.Auth.ReverseProxy.NameHeader)
	assert.Contains(t, cfg.Auth.ReverseProxy.AllowedProxies, "127.0.0.1/24", "127.0.0.1")
	assert.True(t, cfg.Auth.ReverseProxy.AllowRegistration)
}

func TestLoadPostgresDatabaseConfigDefaults(t *testing.T) {
	defer os.Clearenv()
	// Set required env variables
	os.Setenv("TVHGO_TVHEADEND_HOST", "localhost")

	os.Setenv("TVHGO_DATABASE_TYPE", "postgres")
	cfg, err := config.Load("")

	assert.Nil(t, err)

	assert.Equal(t, config.DatabaseTypePostgres, cfg.Database.Type)
	assert.Equal(t, "tvhgo", cfg.Database.User)
	assert.Equal(t, "tvhgo", cfg.Database.Database)
	assert.Equal(t, "localhost", cfg.Database.Host)
	assert.Equal(t, 5432, cfg.Database.Port)
	assert.Empty(t, cfg.Database.Password)
	assert.Equal(t, "disable", cfg.Database.SSLMode)
}

func TestLoadPostgresDatabaseConfig(t *testing.T) {
	defer os.Clearenv()
	// Set required env variables
	os.Setenv("TVHGO_TVHEADEND_HOST", "localhost")

	os.Setenv("TVHGO_DATABASE_TYPE", "postgres")
	os.Setenv("TVHGO_DATABASE_USER", "myUser")
	os.Setenv("TVHGO_DATABASE_DATABASE", "myDatabase")
	os.Setenv("TVHGO_DATABASE_HOST", "myHost")
	os.Setenv("TVHGO_DATABASE_PORT", "1234")
	os.Setenv("TVHGO_DATABASE_PASSWORD", "myPassword")
	os.Setenv("TVHGO_DATABASE_SSL_MODE", "require")
	cfg, err := config.Load("")

	assert.Nil(t, err)

	assert.Equal(t, config.DatabaseTypePostgres, cfg.Database.Type)
	assert.Equal(t, "myUser", cfg.Database.User)
	assert.Equal(t, "myDatabase", cfg.Database.Database)
	assert.Equal(t, "myHost", cfg.Database.Host)
	assert.Equal(t, 1234, cfg.Database.Port)
	assert.Equal(t, "myPassword", cfg.Database.Password)
	assert.Equal(t, "require", cfg.Database.SSLMode)
}
