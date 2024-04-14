package main

import (
	"fmt"
	"os"

	"github.com/davidborzek/tvhgo/cmd"
)

//	@title			tvhgo
//	@version		1.0
//	@description	tvhgo REST API documentation.

//	@BasePath	/api

func main() {
	if err := cmd.App.Run(os.Args); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
