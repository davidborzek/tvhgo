package user

import "github.com/urfave/cli/v2"

var (
	Cmd = &cli.Command{
		Name:  "user",
		Usage: "Manage users of the tvhgo server",
		Subcommands: []*cli.Command{
			userAddCmd,
			userListCmd,
			userDeleteCmd,
			user2FACmd,
			tokenCmd,
		},
	}
)
