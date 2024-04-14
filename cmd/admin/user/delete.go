package user

import (
	"errors"
	"fmt"

	"github.com/davidborzek/tvhgo/cmd/common"
	"github.com/davidborzek/tvhgo/repository/user"
	"github.com/davidborzek/tvhgo/services/clock"
	"github.com/urfave/cli/v2"
)

var deleteCmd = &cli.Command{
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
	Action: delete,
}

func delete(ctx *cli.Context) error {
	_, db := common.Init(ctx)
	userRepository := user.New(db, clock.NewClock())

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
