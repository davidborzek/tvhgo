package main

import (
	"fmt"
	"os"

	"github.com/davidborzek/tvhgo/cmd/admin"
	"github.com/davidborzek/tvhgo/cmd/server"
	"github.com/urfave/cli/v2"
)

//	@title			tvhgo
//	@version		1.0
//	@description	tvhgo REST API documentation.

//	@BasePath	/api

func main() {
	app := cli.App{
		Name:  "tvhgo",
		Usage: "Modern and secure api and web interface for Tvheadend",
		Commands: []*cli.Command{
			server.Cmd,
			admin.Cmd,
		},
		DefaultCommand: "server",
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err.Error())
	}
}
