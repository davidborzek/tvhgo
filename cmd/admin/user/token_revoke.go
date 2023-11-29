package user

import (
	"fmt"

	actx "github.com/davidborzek/tvhgo/cmd/admin/context"
	"github.com/davidborzek/tvhgo/repository/token"
	"github.com/davidborzek/tvhgo/services/auth"
	"github.com/urfave/cli/v2"
)

var tokenRevokeCmd = &cli.Command{
	Name:  "revoke",
	Usage: "Revokes a token",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "id",
			Usage:    "ID of the token,",
			Required: true,
		},
	},
	Action: tokenRevokeAction,
}

func tokenRevokeAction(ctx *cli.Context) error {
	tokenRepository := token.New(actx.GetDB())
	tokenService := auth.NewTokenService(tokenRepository)

	if err := tokenService.Revoke(ctx.Context, ctx.Int64("id")); err != nil {
		return err
	}

	fmt.Println("Token successfully revoked.")

	return nil
}
