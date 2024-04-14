package user

import (
	"context"
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/davidborzek/tvhgo/cmd/common"
	"github.com/davidborzek/tvhgo/core"
	"github.com/davidborzek/tvhgo/repository/user"
	"github.com/davidborzek/tvhgo/services/auth"
	"github.com/davidborzek/tvhgo/services/clock"
	"github.com/urfave/cli/v2"
)

var addCmd = &cli.Command{
	Name:        "add",
	Usage:       "Add a new user",
	Description: "You can use the flags to create a user non-interactive.",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "username",
			Aliases: []string{"u"},
			Usage:   "Username of the new user",
		},
		&cli.StringFlag{
			Name:    "password",
			Aliases: []string{"p"},
			Usage:   "Password of the new user",
			EnvVars: []string{"TVHGO_PASSWORD"},
		},
		&cli.StringFlag{
			Name:    "email",
			Aliases: []string{"e"},
			Usage:   "Email of the new user",
		},
		&cli.StringFlag{
			Name:    "display-name",
			Aliases: []string{"n"},
			Usage:   "Display name of the new user",
		},
	},
	Action: add,
}

var qs = []*survey.Question{
	{
		Name:      "username",
		Prompt:    &survey.Input{Message: "Username:"},
		Validate:  survey.Required,
		Transform: survey.ToLower,
	},
	{
		Name:     "password",
		Prompt:   &survey.Password{Message: "Password:"},
		Validate: survey.Required,
	},
	{
		Name:      "email",
		Prompt:    &survey.Input{Message: "E-Mail:"},
		Validate:  survey.Required,
		Transform: survey.ToLower,
	},
	{
		Name:     "displayName",
		Prompt:   &survey.Input{Message: "Display name:"},
		Validate: survey.Required,
	},
}

func add(ctx *cli.Context) error {
	if ctx.IsSet("username") &&
		ctx.IsSet("password") &&
		ctx.IsSet("email") &&
		ctx.IsSet("display-name") {

		return createUser(
			ctx,
			ctx.String("username"),
			ctx.String("password"),
			ctx.String("email"),
			ctx.String("display-name"),
		)
	}

	answers := struct {
		Username    string
		Password    string
		Email       string
		DisplayName string
	}{}

	err := survey.Ask(qs, &answers, survey.WithIcons(func(is *survey.IconSet) {
		is.Question.Text = ""
	}))

	if err != nil {
		return err
	}

	return createUser(
		ctx,
		answers.Username,
		answers.Password,
		answers.Email,
		answers.DisplayName,
	)
}

func createUser(ctx *cli.Context, username, password, email, displayName string) error {
	_, db := common.Init(ctx)
	userRepository := user.New(db, clock.NewClock())

	hash, err := auth.HashPassword(password)
	if err != nil {
		return err
	}

	err = userRepository.Create(context.Background(), &core.User{
		Username:     username,
		Email:        email,
		DisplayName:  displayName,
		PasswordHash: string(hash),
	})
	if err != nil {
		return err
	}

	fmt.Println("User created successfully")
	return nil
}
