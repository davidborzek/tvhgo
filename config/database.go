package config

const (
	defaultDatabasePath = "./tvhgo.db"
)

type (
	DatabaseConfig struct {
		Path string `yaml:"path" env:"PATH"`
	}
)

func (c *DatabaseConfig) SetDefaults() {
	if c.Path == "" {
		c.Path = defaultDatabasePath
	}
}
