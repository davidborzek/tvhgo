package config

import (
	"net"
	"strconv"
)

const (
	defaultMetricsPath = "/metrics"
	defaultMetricsPort = 8081
)

type (
	MetricsConfig struct {
		Enabled bool   `yaml:"enabled" env:"ENABLED"`
		Path    string `yaml:"path" env:"PATH"`
		Port    int    `yaml:"port" env:"PORT"`
		Host    string `yaml:"host" env:"HOST"`
		Token   string `yaml:"token" env:"TOKEN"`
	}
)

func (c *MetricsConfig) SetDefaults() {
	if c.Path == "" {
		c.Path = defaultMetricsPath
	}

	if c.Port == 0 {
		c.Port = defaultMetricsPort
	}
}

func (c *MetricsConfig) Addr() string {
	return net.JoinHostPort(
		c.Host,
		strconv.Itoa(c.Port),
	)
}
