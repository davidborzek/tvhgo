package common

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/davidborzek/tvhgo/config"
	"github.com/davidborzek/tvhgo/db"
	"github.com/urfave/cli/v2"
)

func PrintTable(headers []string, rows [][]any) {
	w := tabwriter.NewWriter(os.Stdout, 10, 0, 1, ' ', 0)

	for i, h := range headers {
		fmt.Fprintf(w, "%s", h)

		if i < len(headers) {
			fmt.Fprint(w, "\t")
		}
	}
	fmt.Fprint(w, "\n")

	for _, row := range rows {
		for i, c := range row {
			fmt.Fprintf(w, "%v", c)

			if i < len(row) {
				fmt.Fprint(w, "\t")
			}
		}

		fmt.Fprint(w, "\n")
	}

	w.Flush()
}

func MapRows[T any](values []T, mapper func(t T) []any) [][]any {
	rows := make([][]any, 0)
	for _, v := range values {
		rows = append(rows, mapper(v))
	}
	return rows
}

func Init(ctx *cli.Context) (*config.Config, *db.DB) {
	cfg, err := config.Load(ctx.String("config"))
	if err != nil {
		log.Println("failed to load config:", err)
		os.Exit(1)
	}

	dbConn, err := db.Connect(cfg.Database)
	if err != nil {
		log.Println("failed create db connection:", err)
		os.Exit(1)
	}

	return cfg, dbConn
}
