package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/micro/cli"
	"github.com/srishanbhattarai/nepcal/dateconv"
)

const versionNumber = "0.3.2"

// Cheap testing.
var writer io.Writer = os.Stdout

func init() {
	flag.Parse()
}

func main() {
	runCli()
}

func runCli() {
	cli := bootstrapCli()

	err := cli.Run(os.Args)
	if err != nil {
		fmt.Printf("Something went wrong: %s\n", err.Error())
	}
}

func bootstrapCli() *cli.App {
	app := cli.NewApp()
	app.Name = "nepcal"
	app.Version = versionNumber
	app.Usage = "Calendar and conversion utilities for Nepali dates"
	app.Commands = []cli.Command{
		{
			Name:    "cal",
			Aliases: []string{"c"},
			Usage:   "Show calendar for the month",
			Action: func(c *cli.Context) {
				cal := newCalendar()
				cal.Render(writer, time.Now())
			},
		},
		{
			Name:    "date",
			Aliases: []string{"d"},
			Usage:   "Show today's date",
			Action: func(c *cli.Context) {
				showDate(writer, time.Now())
			},
		},
		{
			Name:  "conv",
			Usage: "Convert AD dates to BS and vice-versa",
			Subcommands: []cli.Command{
				{
					Name:  "adtobs",
					Usage: "Convert AD date to BS date",
					Action: func(c *cli.Context) {
						areArgsValid := func() bool {
							if c.NArg() < 1 {
								return false
							}

							_, _, _, ok := parseRawDate(c.Args().First())
							if !ok {
								return false
							}

							return true
						}()

						if !areArgsValid {
							fmt.Println("Please supply a valid date in the format mm-dd-yyyy. Example: `nepcal conv adtobs 08-21-1994`")
							return
						}

						mm, dd, yy, _ := parseRawDate(c.Args().First())
						showDate(writer, time.Date(yy, time.Month(mm), dd, 0, 0, 0, 0, time.UTC))
					},
				},
				{
					Name:  "bstoad",
					Usage: "Convert BS date to AD date",
					Action: func(c *cli.Context) {
						fmt.Println("Unfortunately BS to AD isn't supported at this time. :(")
					},
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

	fmt.Println("dd, mm, yy", dd, mm, yy)

	if dd < 1 || dd > 31 || mm < 1 || mm > 12 || len(dateParts[2]) != 4 {
		return -1, -1, -1, false
	}

	return mm, dd, yy, true
}

// showDate prints the current B.S. date
func showDate(w io.Writer, t time.Time) {
	yy, mm, dd := t.Date()

	bsyy, bsmm, bsdd := dateconv.ToBS(toTime(yy, mm, dd)).Date()
	month, monthOk := dateconv.GetBSMonthName(bsmm)
	weekday, weekdayOk := dateconv.GetNepWeekday(t.Weekday())

	if monthOk && weekdayOk {
		fmt.Fprintf(w, "%s %d, %d %s\n", month, bsdd, bsyy, weekday)
	}
}

// toTime creates a new time.Time with the basic yy/mm/dd parameters.
func toTime(yy int, mm time.Month, dd int) time.Time {
	return time.Date(yy, mm, dd, 0, 0, 0, 0, time.UTC)
}
