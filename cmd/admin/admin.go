package admin

import (
	"github.com/davidborzek/tvhgo/cmd/admin/user"
	"github.com/rs/zerolog"
	"github.com/urfave/cli/v2"
)

var (
	Cmd = &cli.Command{
		Name:  "admin",
		Usage: "Admin controls for the tvhgo server",
		Subcommands: []*cli.Command{
			user.Cmd,
		},
		Before: before,
	}
)

func before(_c *cli.Context) error {
	zerolog.SetGlobalLevel(zerolog.Disabled)
	return nil
}
