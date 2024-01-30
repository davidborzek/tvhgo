package twofa

import "github.com/urfave/cli/v2"

var Cmd = &cli.Command{
	Name:  "2fa",
	Usage: "Manage 2FA of a user",
	Subcommands: []*cli.Command{
		disableCmd,
	},
}
