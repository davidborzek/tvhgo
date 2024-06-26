package token

import (
	"errors"
	"fmt"

	"github.com/davidborzek/tvhgo/cmd/common"
	"github.com/davidborzek/tvhgo/repository/token"
	"github.com/davidborzek/tvhgo/repository/user"
	"github.com/davidborzek/tvhgo/services/auth"
	"github.com/davidborzek/tvhgo/services/clock"
	"github.com/urfave/cli/v2"
)

var generateCmd = &cli.Command{
	Name:  "generate",
	Usage: "Generate a new token",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "username",
			Aliases:  []string{"u"},
			Usage:    "Username of the new user",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "name",
			Aliases:  []string{"n"},
			Usage:    "Name of the token,",
			Required: true,
		},
	},
	Action: generate,
}

func generate(ctx *cli.Context) error {
	_, db := common.Init(ctx)
	userRepository := user.New(db, clock.NewClock())

	user, err := userRepository.FindByUsername(ctx.Context, ctx.String("username"))
	if err != nil {
		return err
	}

	if user == nil {
		return errors.New("user not found")
	}

	tokenRepository := token.New(db)
	tokenService := auth.NewTokenService(tokenRepository)

	tokenValue, err := tokenService.Create(ctx.Context, user.ID, ctx.String("name"))
	if err != nil {
		return err
	}

	fmt.Printf("Token generated: %s\n", tokenValue)

	return nil
}
