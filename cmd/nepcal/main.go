package main

import (
	"flag"
	"io"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

const versionNumber = "v1.1.0"

// Cheap testing.
var writer io.Writer = os.Stdout

func main() {
	runCli()
}

func runCli() {
	cli := bootstrapCli()

	cli.Run(os.Args)
}

func bootstrapCli() *cli.App {
	flag.Parse()
	nc := nepcalCli{}

	app := cli.NewApp()
	app.Name = "nepcal"
	app.Version = versionNumber
	app.Usage = "Calendar and conversion utilities for Nepali dates"
	app.Commands = []*cli.Command{
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
			Action:  nc.showDate(writer, time.Now()),
		},
		{
			Name:  "conv",
			Usage: "Convert AD dates to BS and vice-versa",
			Subcommands: []*cli.Command{
				{
					Name:   "adtobs",
					Usage:  "Convert AD date to BS date",
					Action: nc.convADToBS,
				},
				{
					Name:   "bstoad",
					Usage:  "Convert BS date to AD date",
					Action: nc.convBSToAD,
				},
			},
		},
	}

	return app
}
