package config

import "errors"

const (
	defaultMetricsPath = "/metrics"
)

type (
	MetricsConfig struct {
		Enabled bool   `yaml:"enabled" env:"ENABLED"`
		Path    string `yaml:"path" env:"PATH"`
		Token   string `yaml:"token" env:"TOKEN"`
	}
)

func (c *MetricsConfig) SetDefaults() {
	if c.Path == "" {
		c.Path = defaultMetricsPath
	}
}

func (c *MetricsConfig) Validate() error {
	if !c.Enabled {
		return nil
	}

	if c.Token == "" {
		return errors.New("metrics token is not set")
	}

	return nil
}
