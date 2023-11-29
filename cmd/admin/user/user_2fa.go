package user

import "github.com/urfave/cli/v2"

var user2FACmd = &cli.Command{
	Name:  "2fa",
	Usage: "Manage 2FA of a user",
	Subcommands: []*cli.Command{
		user2FADisableCmd,
	},
}
