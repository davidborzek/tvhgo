package config

import (
	"net"
	"strconv"
)

const (
	defaultServerPort = 8080
)

type (
	ServerConfig struct {
		Host string `yaml:"host" env:"HOST"`
		Port int    `yaml:"port" env:"PORT"`
	}
)

func (c *ServerConfig) SetDefaults() {
	if c.Port == 0 {
		c.Port = defaultServerPort
	}
}

func (c *ServerConfig) Addr() string {
	return net.JoinHostPort(
		c.Host,
		strconv.Itoa(c.Port),
	)
}
