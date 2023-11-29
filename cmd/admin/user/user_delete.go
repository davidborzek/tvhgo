package user

import (
	"errors"
	"fmt"

	actx "github.com/davidborzek/tvhgo/cmd/admin/context"
	"github.com/davidborzek/tvhgo/repository/user"
	"github.com/davidborzek/tvhgo/services/clock"
	"github.com/urfave/cli/v2"
)

var userDeleteCmd = &cli.Command{
	Name:  "delete",
	Usage: "Deletes a user",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "username",
			Aliases:  []string{"u"},
			Usage:    "Username of the new user",
			Required: true,
		},
	},
	Action: userDeleteAction,
}

func userDeleteAction(ctx *cli.Context) error {
	userRepository := user.New(actx.GetDB(), clock.NewClock())

	user, err := userRepository.FindByUsername(ctx.Context, ctx.String("username"))
	if err != nil {
		return err
	}

	if user == nil {
		return errors.New("user not found")
	}

	if err := userRepository.Delete(ctx.Context, user); err != nil {
		return err
	}

	fmt.Println("User successfully deleted.")
	return nil
}
