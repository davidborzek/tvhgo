package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"
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
