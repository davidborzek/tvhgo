package config

import (
	"io"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type (
	LogConfig struct {
		Level  string `yaml:"level" env:"LEVEL"`
		Format string `yaml:"format" env:"FORMAT"`
	}
)

func (c *LogConfig) SetDefaults() {
	if c.Level == "" {
		c.Level = "info"
	}

	if c.Format == "" {
		c.Format = "console"
	}
}

func (c *LogConfig) SetupLogger() {
	zerolog.SetGlobalLevel(parseLogLevel(c.Level))

	if c.Format == "json" {
		initLogger(os.Stdout)
	} else if c.Format != "console" {
		log.Warn().Str("format", c.Format).
			Msg("invalid log format provided - falling back to 'console'")
	}
}

func parseLogLevel(level string) zerolog.Level {
	switch strings.ToLower(level) {
	case "debug":
		return zerolog.DebugLevel
	case "info":
		return zerolog.InfoLevel
	case "warning":
		return zerolog.WarnLevel
	case "error":
		return zerolog.ErrorLevel
	case "fatal":
		return zerolog.FatalLevel
	}

	log.Warn().Str("level", level).
		Msg("invalid log level provided - falling back to 'info'")

	return zerolog.InfoLevel
}

func initLogger(writer io.Writer) {
	log.Logger = zerolog.New(writer).
		With().Timestamp().Stack().Caller().Logger()
}

func InitDefaultLogger() {
	initLogger(zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
	})
}
