package cmd

import (
	"github.com/davidborzek/tvhgo/cmd/admin"
	"github.com/davidborzek/tvhgo/cmd/server"
	"github.com/urfave/cli/v2"
)

var App = cli.App{
	Name:  "tvhgo",
	Usage: "Modern and secure api and web interface for Tvheadend",
	Commands: []*cli.Command{
		server.Cmd,
		admin.Cmd,
	},
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "config",
			Aliases: []string{"c"},
			Usage:   "Path to the configuration file",
			EnvVars: []string{"TVHGO_CONFIG"},
		},
	},
	DefaultCommand: "server",
}
