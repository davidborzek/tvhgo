package user

import (
	"context"
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	actx "github.com/davidborzek/tvhgo/cmd/admin/context"
	"github.com/davidborzek/tvhgo/core"
	"github.com/davidborzek/tvhgo/repository/user"
	"github.com/urfave/cli/v2"
	"golang.org/x/crypto/bcrypt"
)

const bcryptCost = 12

var addUserCmd = &cli.Command{
	Name:   "add",
	Usage:  "Add a new user",
	Action: addUser,
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

func addUser(ctx *cli.Context) error {
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

	userRepository := user.New(actx.GetDB())

	hash, err := bcrypt.GenerateFromPassword([]byte(answers.Password), bcryptCost)
	if err != nil {
		return err
	}

	err = userRepository.Create(context.Background(), &core.User{
		Username:     answers.Username,
		Email:        answers.Email,
		DisplayName:  answers.DisplayName,
		PasswordHash: string(hash),
	})
	if err != nil {
		return err
	}

	fmt.Println("User created successfully")
	return nil
}
