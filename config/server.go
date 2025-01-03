package config

import (
	"net"
	"strconv"
)

const (
	defaultServerPort = 8080
)

type (
	SwaggerUIConfig struct {
		Enabled *bool `yaml:"enabled" env:"ENABLED"`
	}

	ServerConfig struct {
		Host      string          `yaml:"host" env:"HOST"`
		Port      int             `yaml:"port" env:"PORT"`
		SwaggerUI SwaggerUIConfig `yaml:"swagger_ui" env:"SWAGGER_UI"`
	}
)

func (c *ServerConfig) SetDefaults() {
	if c.Port == 0 {
		c.Port = defaultServerPort
	}

	if c.SwaggerUI.Enabled == nil {
		v := true
		c.SwaggerUI.Enabled = &v
	}
}

func (c *ServerConfig) Addr() string {
	return net.JoinHostPort(
		c.Host,
		strconv.Itoa(c.Port),
	)
}
