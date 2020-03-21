package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	dateconv "github.com/srishanbhattarai/nepcal/time"
	"github.com/urfave/cli"
)

const versionNumber = "0.4.0"

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
	app.Commands = []cli.Command{
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
			Subcommands: []cli.Command{
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

// Parse user input raw date into valid dd, mm, yy format. The last parameter is a boolean indicating if
// the date is valid or not.
func parseRawDate(rawDate string) (int, int, int, bool) {
	dateParts := strings.Split(rawDate, "-")
	if len(dateParts) != 3 {
		return -1, -1, -1, false
	}

	mm, err := strconv.Atoi(dateParts[0])
	if err != nil {
		return -1, -1, -1, false
	}

	dd, err := strconv.Atoi(dateParts[1])
	if err != nil {
		return -1, -1, -1, false
	}

	yy, err := strconv.Atoi(dateParts[2])
	if err != nil {
		return -1, -1, -1, false
	}

	if dd < 1 || dd > 31 || mm < 1 || mm > 12 || len(dateParts[2]) != 4 {
		return -1, -1, -1, false
	}

	return mm, dd, yy, true
}

func showDateAD(w io.Writer, t time.Time) {
	adyy, _, addd := t.Date()

	month := t.Month().String()
	weekday := t.Weekday()

	fmt.Fprintf(w, "%s %d, %d %s\n", month, addd, adyy, weekday)
}

// showDate prints the current B.S. date
func showDateBS(w io.Writer, bs dateconv.BSDate, wd time.Weekday) {
	bsyy, bsmm, bsdd := bs.Date()

	month, monthOk := dateconv.GetBSMonthName(time.Month(bsmm))
	weekday, weekdayOk := dateconv.GetNepWeekday(wd)

	if monthOk && weekdayOk {
		fmt.Fprintf(w, "%s %d, %d %s\n", month, bsdd, bsyy, weekday)
	}
}

// toTime creates a new time.Time with the basic yy/mm/dd parameters.
func toTime(yy int, mm time.Month, dd int) time.Time {
	return time.Date(yy, mm, dd, 0, 0, 0, 0, time.UTC)
}
