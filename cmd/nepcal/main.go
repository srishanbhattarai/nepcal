package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"time"

	"github.com/urfave/cli/v2"
)

const version = "v1.1.0"

// Cheap testing.
var globalWriter io.Writer = os.Stdout

func main() {
	runCli()
}

func runCli() {
	cli := bootstrapCli()

	err := cli.Run(os.Args)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func bootstrapCli() *cli.App {
	flag.Parse()
	nc := nepcalCli{}

	app := &cli.App{
		Name:            "nepcal",
		Version:         version,
		Usage:           "Calendar and conversion utilities for Nepali dates",
		HideVersion:     true,
		HideHelpCommand: true,
		Authors: []*cli.Author{
			{
				Name:  "Srishan Bhattarai",
				Email: "srishanbhattarai@gmail.com",
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "cal",
				Aliases: []string{"c"},
				Usage:   "Show calendar for the month",
				Action:  nc.showCalendar,
			},
			{
				Name:    "date",
				Aliases: []string{"d"},
				Usage:   "Show today's date",
				Action:  nc.showDate(globalWriter, time.Now()),
			},
			{
				Name:  "conv",
				Usage: "Convert AD dates to BS and vice-versa",
				Subcommands: []*cli.Command{
					{
						Name:   "tobs",
						Usage:  "Convert AD date to BS date",
						Action: nc.convADToBS,
					},
					{
						Name:   "toad",
						Usage:  "Convert BS date to AD date",
						Action: nc.convBSToAD,
					},
				},
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	return app
}
