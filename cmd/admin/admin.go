package admin

import (
	"github.com/davidborzek/tvhgo/cmd/admin/context"
	"github.com/davidborzek/tvhgo/cmd/admin/user"
	"github.com/davidborzek/tvhgo/config"
	"github.com/davidborzek/tvhgo/db"
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

func before(_ *cli.Context) error {
	logrus.SetLevel(logrus.FatalLevel)

	cfg, err := config.Load()
	if err != nil {
		return err
	}

	dbConn, err := db.Connect(cfg.Database.Path)
	if err != nil {
		return err
	}

	context.SetDB(dbConn)

	return nil
}
