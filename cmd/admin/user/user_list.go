package user

import (
	"fmt"
	"time"

	"github.com/davidborzek/tvhgo/cmd"
	actx "github.com/davidborzek/tvhgo/cmd/admin/context"
	"github.com/davidborzek/tvhgo/core"
	"github.com/davidborzek/tvhgo/repository/user"
	"github.com/davidborzek/tvhgo/services/clock"
	"github.com/urfave/cli/v2"
)

var userListCmd = &cli.Command{
	Name:   "list",
	Usage:  "List users.",
	Action: userListAction,
}

func userListAction(ctx *cli.Context) error {
	userRepository := user.New(actx.GetDB(), clock.NewClock())

	users, err := userRepository.Find(ctx.Context, core.UserQueryParams{})
	if err != nil {
		return err
	}

	if len(users) == 0 {
		fmt.Printf("No users found.")
		return nil
	}

	cmd.PrintTable(
		[]string{"ID", "Username", "Email", "Name", "Created", "Updated"},
		cmd.MapRows(users, func(user *core.User) []any {
			return []any{
				user.ID,
				user.Username,
				user.Email,
				user.DisplayName,
				time.Unix(user.CreatedAt, 0).Format(time.RFC822),
				time.Unix(user.UpdatedAt, 0).Format(time.RFC822),
			}
		}),
	)

	return nil
}
