package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strings"

	"gopkg.in/urfave/cli.v2"
)

var (
	infile  = "-"
	outfile = "-"
	tabs    bool
)

func main() {
	app := cli.App{
		Name:     "sortcsv",
		Usage:    "Sort CSV file",
		Commands: buildCommands(),
	}
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "error running app: %v\n", err)
	}
}

func buildCommands() []*cli.Command {
	return []*cli.Command{
		sortCmd(),
	}
}

func sortCmd() *cli.Command {
	return &cli.Command{
		Name:    "sort",
		Aliases: []string{"s"},
		Action:  sortcsv,
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "infile", Aliases: []string{"in", "i"}, Usage: "Defaults to stdin (-) if not specified.", Destination: &infile, DefaultText: infile},
			&cli.StringFlag{Name: "outfile", Aliases: []string{"out", "o"}, Usage: "Defaults to stdout (-) if not specified.", Destination: &outfile, DefaultText: outfile},
			&cli.StringSliceFlag{Name: "sortby", Aliases: []string{"s"}, Usage: "Columns to sort by, repeat for sub-sort"},
			&cli.BoolFlag{Name: "tab", Aliases: []string{"t"}, Usage: "Input is Tab-Delimited", Destination: &tabs},
		},
	}
}

func sortcsv(ctx *cli.Context) error {
	var err error

	in := os.Stdin
	switch infile {
	case "", "-":
	default:
		if in, err = os.Open(infile); err != nil {
			return err
		}
		defer in.Close()
	}

	reader := csv.NewReader(in)
	if tabs {
		reader.Comma = '\t'
	}
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	sortKeys := ctx.StringSlice("sortby")
	order := sortOrder(sortKeys, records[0])
	sortRecords := func(i, j int) bool {
		for _, ix := range order {
			if records[i+1][ix] == records[j+1][ix] {
				continue
			}
			return records[i+1][ix] < records[j+1][ix]
		}
		return false
	}

	sort.Slice(records[1:], sortRecords)

	out := os.Stdout
	switch outfile {
	case "", "-":
	default:
		if out, err = os.Create(outfile); err != nil {
			return err
		}
		defer out.Close()
	}

	writer := csv.NewWriter(out)
	if tabs {
		writer.Comma = '\t'
	}
	return writer.WriteAll(records)
}

func sortOrder(keys []string, header []string) []int {
	ret := []int{}
	for _, lhs := range keys {
		k := key(lhs)
		for ix, rhs := range header {
			if k == key(rhs) {
				ret = append(ret, ix)
			}
		}
	}
	if len(ret) == 0 {
		ret = []int{0}
	}
	return ret
}

func key(name string) string {
	reduce := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			return r + 32
		case r >= 'a' && r <= 'z':
			return r
		}
		return -1
	}
	return strings.Map(reduce, name)
}
