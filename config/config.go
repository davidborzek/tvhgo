package config

import (
	"errors"
	"os"
	"path"
	"strings"

	"github.com/caarlos0/env/v9"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

type (
	Config struct {
		Server    ServerConfig    `yaml:"server" envPrefix:"SERVER_"`
		Tvheadend TvheadendConfig `yaml:"tvheadend" envPrefix:"TVHEADEND_"`
		Auth      AuthConfig      `yaml:"auth" envPrefix:"AUTH_"`
		Database  DatabaseConfig  `yaml:"database" envPrefix:"DATABASE_"`

		LogLevel string `yaml:"log_level"`
	}
)

var (
	// Possible paths of the configuration file.
	paths = []string{
		"./",
		"/etc/tvhgo/",
	}

	// Options to load configuration from environment variables.
	envOpts = env.Options{
		Prefix: "TVHGO_",
	}
)

// existsConfig checks if either config.yml or config.yaml
// exists in the given directory and returns the full path.
func existsConfig(p string) (string, bool, error) {
	c := path.Join(p, "config.yml")
	_, err := os.Stat(c)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return "", false, err
	}

	if err == nil {
		return c, true, nil
	}

	c = path.Join(p, "config.yaml")
	_, err = os.Stat(c)
	if errors.Is(err, os.ErrNotExist) {
		return "", false, nil
	}

	if err != nil {
		return "", false, err
	}

	return c, true, nil
}

// findConfig tries to find a config file in
// various directories.
func findConfig() (string, error) {
	for _, p := range paths {
		if c, ok, err := existsConfig(p); err != nil {
			return "", err
		} else if ok {
			return c, nil
		}
	}

	return "", nil
}

// Load loads a config from a config file at a given path
// and overrides from environment variables.
func Load() (*Config, error) {
	cfgPath, err := findConfig()
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := env.ParseWithOptions(&cfg, envOpts); err != nil {
		return nil, err
	}

	if err := loadFromFile(cfgPath, &cfg); err != nil {
		return nil, err
	}

	if err := cfg.validate(); err != nil {
		return nil, err
	}

	cfg.loadDefaults()

	log.SetLevel(
		parseLogLevel(cfg.LogLevel),
	)

	return &cfg, nil
}

// Loads the config from a file if the path is not empty.
func loadFromFile(cfgPath string, cfg *Config) error {
	if cfgPath == "" {
		return nil
	}

	log.WithField("path", cfgPath).
		Info("loading config from file")

	cfgFile, err := os.ReadFile(cfgPath)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(cfgFile, &cfg)
}

func (c *Config) validate() error {
	if err := c.Tvheadend.Validate(); err != nil {
		return err
	}

	return nil
}

func (c *Config) loadDefaults() {
	c.Server.SetDefaults()
	c.Tvheadend.SetDefaults()
	c.Auth.Session.SetDefaults()
	c.Auth.TOTP.SetDefaults()
	c.Database.SetDefaults()

	if c.LogLevel == "" {
		c.LogLevel = "info"
	}
}

func parseLogLevel(level string) log.Level {
	switch strings.ToLower(level) {
	case "debug":
		return log.DebugLevel
	case "info":
		return log.InfoLevel
	case "warning":
		return log.WarnLevel
	case "error":
		return log.ErrorLevel
	case "fatal":
		return log.FatalLevel
	}

	log.WithField("level", level).
		Warn("invalid log level provided - falling back to 'info'")

	return log.InfoLevel
}
