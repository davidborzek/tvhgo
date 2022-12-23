package config

import (
	"net"
	"strconv"
)

type (
	ServerConfig struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	}
)

func (c *ServerConfig) SetDefaults() {
	if c.Port == 0 {
		c.Port = 8080
	}
}

func (c *ServerConfig) Addr() string {
	return net.JoinHostPort(
		c.Host,
		strconv.Itoa(c.Port),
	)
}
