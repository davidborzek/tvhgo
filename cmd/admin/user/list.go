package user

import (
	"fmt"
	"time"

	"github.com/davidborzek/tvhgo/cmd/common"
	"github.com/davidborzek/tvhgo/core"
	"github.com/davidborzek/tvhgo/repository/user"
	"github.com/davidborzek/tvhgo/services/clock"
	"github.com/urfave/cli/v2"
)

var listCmd = &cli.Command{
	Name:   "list",
	Usage:  "List users.",
	Action: list,
}

func list(ctx *cli.Context) error {
	_, db := common.Init()

	userRepository := user.New(db, clock.NewClock())

	users, err := userRepository.Find(ctx.Context, core.UserQueryParams{})
	if err != nil {
		return err
	}

	if len(users) == 0 {
		fmt.Printf("No users found.")
		return nil
	}

	common.PrintTable(
		[]string{"ID", "Username", "Email", "Name", "Created", "Updated"},
		common.MapRows(users, func(user *core.User) []any {
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
