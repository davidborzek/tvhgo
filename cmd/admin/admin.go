package admin

import (
	"github.com/davidborzek/tvhgo/cmd/admin/user"
	"github.com/urfave/cli/v2"
)

var (
	Cmd = &cli.Command{
		Name:  "admin",
		Usage: "Admin controls for the tvhgo server",
		Subcommands: []*cli.Command{
			user.Cmd,
		},
	}
)
