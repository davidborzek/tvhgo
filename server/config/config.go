package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type (
	Config struct {
		Server    ServerConfig    `yaml:"server"`
		Tvheadend TvheadendConfig `yaml:"tvheadend"`
		Auth      AuthConfig      `yaml:"auth"`
	}
)

// Load loads a config from a given path.
func Load(path string) (*Config, error) {
	cfgFile, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
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
}
