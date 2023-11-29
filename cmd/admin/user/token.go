package user

import (
	"github.com/urfave/cli/v2"
)

var tokenCmd = &cli.Command{
	Name:  "token",
	Usage: "Manage tokens of a user",
	Subcommands: []*cli.Command{
		tokenListCmd,
		tokenGenerateCmd,
		tokenRevokeCmd,
	},
}
