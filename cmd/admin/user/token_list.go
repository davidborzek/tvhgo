package user

import (
	"errors"
	"fmt"
	"time"

	"github.com/davidborzek/tvhgo/cmd"
	actx "github.com/davidborzek/tvhgo/cmd/admin/context"
	"github.com/davidborzek/tvhgo/core"
	"github.com/davidborzek/tvhgo/repository/token"
	"github.com/davidborzek/tvhgo/repository/user"
	"github.com/davidborzek/tvhgo/services/clock"
	"github.com/urfave/cli/v2"
)

var tokenListCmd = &cli.Command{
	Name:  "list",
	Usage: "List tokens of a user",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "username",
			Aliases:  []string{"u"},
			Usage:    "Username of the new user",
			Required: true,
		},
	},
	Action: tokenListAction,
}

func tokenListAction(ctx *cli.Context) error {
	userRepository := user.New(actx.GetDB(), clock.NewClock())

	user, err := userRepository.FindByUsername(ctx.Context, ctx.String("username"))
	if err != nil {
		return err
	}

	if user == nil {
		return errors.New("user not found")
	}

	tokenRepository := token.New(actx.GetDB())
	tokens, err := tokenRepository.FindByUser(ctx.Context, user.ID)
	if err != nil {
		return err
	}

	if len(tokens) == 0 {
		fmt.Printf("No tokens found.")
		return nil
	}

	cmd.PrintTable(
		[]string{"ID", "Name", "Created"},
		cmd.MapRows(tokens, func(token *core.Token) []any {
			return []any{
				token.ID,
				token.Name,
				time.Unix(token.CreatedAt, 0).Format(time.RFC822),
			}
		}),
	)

	return nil
}
