package config

import (
	"errors"
	"os"
	"path"

	"github.com/caarlos0/env/v6"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

type (
	Config struct {
		Server    ServerConfig    `yaml:"server" envPrefix:"SERVER_"`
		Tvheadend TvheadendConfig `yaml:"tvheadend" envPrefix:"TVHEADEND_"`
		Auth      AuthConfig      `yaml:"auth" envPrefix:"AUTH_"`
		Database  DatabaseConfig  `yaml:"database" envPrefix:"DATABASE_"`
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

	return "", errors.New("no config file found")
}

// Load loads a config from a config file at a given path
// and overrides from environment variables.
func Load() (*Config, error) {
	cfgPath, err := findConfig()
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := env.Parse(&cfg, envOpts); err != nil {
		return nil, err
	}

	log.WithField("path", cfgPath).
		Info("loading config from file")

	cfgFile, err := os.ReadFile(cfgPath)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(cfgFile, &cfg); err != nil {
		return nil, err
	}

	if err := cfg.validate(); err != nil {
		return nil, err
	}

	cfg.loadDefaults()

	return &cfg, nil
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
	c.Database.SetDefaults()
}
