package user

import (
	"errors"
	"fmt"

	actx "github.com/davidborzek/tvhgo/cmd/admin/context"
	"github.com/davidborzek/tvhgo/core"
	twofactorsettings "github.com/davidborzek/tvhgo/repository/two_factor_settings"
	"github.com/davidborzek/tvhgo/repository/user"
	"github.com/davidborzek/tvhgo/services/clock"
	"github.com/urfave/cli/v2"
)

var user2FADisableCmd = &cli.Command{
	Name:  "disable",
	Usage: "Disable 2FA for a user.",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "username",
			Aliases:  []string{"u"},
			Usage:    "Username of the new user",
			Required: true,
		},
	},
	Action: user2FADisableAction,
}

func user2FADisableAction(ctx *cli.Context) error {
	userRepository := user.New(actx.GetDB(), clock.NewClock())

	user, err := userRepository.FindByUsername(ctx.Context, ctx.String("username"))
	if err != nil {
		return err
	}

	if user == nil {
		return errors.New("user not found")
	}

	twoFactorSettingsRepository := twofactorsettings.New(actx.GetDB())

	err = twoFactorSettingsRepository.Delete(ctx.Context, &core.TwoFactorSettings{
		UserID: user.ID,
	})

	if err != nil {
		return err
	}

	fmt.Println("Two factor auth successfully disabled.")
	return nil
}
