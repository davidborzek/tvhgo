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
	DefaultCommand: "server",
}
