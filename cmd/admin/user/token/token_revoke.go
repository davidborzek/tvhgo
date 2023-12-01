package token

import (
	"fmt"

	"github.com/davidborzek/tvhgo/cmd"
	"github.com/davidborzek/tvhgo/repository/token"
	"github.com/davidborzek/tvhgo/services/auth"
	"github.com/urfave/cli/v2"
)

var revokeCmd = &cli.Command{
	Name:  "revoke",
	Usage: "Revokes a token",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "id",
			Usage:    "ID of the token,",
			Required: true,
		},
	},
	Action: revoke,
}

func revoke(ctx *cli.Context) error {
	_, db := cmd.Init()
	tokenRepository := token.New(db)
	tokenService := auth.NewTokenService(tokenRepository)

	if err := tokenService.Revoke(ctx.Context, ctx.Int64("id")); err != nil {
		return err
	}

	fmt.Println("Token successfully revoked.")

	return nil
}
