package config

import (
	"errors"
	"net"
	"net/url"
	"strconv"
)

type (
	TvheadendConfig struct {
		Scheme   string `yaml:"scheme"`
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	}
)

func (c *TvheadendConfig) Validate() error {
	if c.Host == "" {
		return errors.New("tvheadend host is not set")
	}

	return nil
}

func (c *TvheadendConfig) SetDefaults() {
	if c.Scheme == "" {
		c.Scheme = "http"
	}
	if c.Port == 0 {
		c.Port = 9981
	}
}

func (c *TvheadendConfig) URL() string {
	tvhUrl := url.URL{
		Scheme: c.Scheme,
		Host: net.JoinHostPort(
			c.Host,
			strconv.Itoa(c.Port),
		),
	}
	return tvhUrl.String()
}
