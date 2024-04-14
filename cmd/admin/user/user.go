package user

import (
	"github.com/davidborzek/tvhgo/cmd/admin/user/token"
	"github.com/davidborzek/tvhgo/cmd/admin/user/twofa"
	"github.com/urfave/cli/v2"
)

var (
	Cmd = &cli.Command{
		Name:  "user",
		Usage: "Manage users of the tvhgo server",
		Subcommands: []*cli.Command{
			addCmd,
			listCmd,
			deleteCmd,
			twofa.Cmd,
			token.Cmd,
		},
	}
)
