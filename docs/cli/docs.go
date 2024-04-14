package main

import (
	"os"

	"github.com/davidborzek/tvhgo/cmd"
	"github.com/rs/zerolog/log"
)

func main() {
	md, err := cmd.App.ToMarkdown()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to generate cli markdown docs")
	}

	if err := os.WriteFile("./docs/cli/docs.md", []byte(md), 0655); err != nil {
		log.Fatal().Err(err).Msg("failed to generate cli markdown docs")
	}
}
