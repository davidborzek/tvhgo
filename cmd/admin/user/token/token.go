package token

import (
	"github.com/urfave/cli/v2"
)

var Cmd = &cli.Command{
	Name:  "token",
	Usage: "Manage tokens of a user",
	Subcommands: []*cli.Command{
		listCmd,
		generateCmd,
		revokeCmd,
	},
}
