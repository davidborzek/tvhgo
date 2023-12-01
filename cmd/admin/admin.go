package admin

import (
	"github.com/davidborzek/tvhgo/cmd/admin/user"
	"github.com/sirupsen/logrus"
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
	logrus.SetLevel(logrus.FatalLevel)
	return nil
}
