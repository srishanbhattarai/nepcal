package main

import (
	"fmt"
	"io"
	"time"

	"github.com/urfave/cli"
)

// nepcalCli is a struct to hold all the CLI behavior.
// It is a small wrapper around the urfave/cli package to keep things clean.
type nepcalCli struct{}

// Shows the calendar for the current day.
func (nepcalCli) showCalendar(c *cli.Context) {
	cal := newCalendar(writer)
	cal.Render(time.Now())
}

// Shows the date for the provided time. Returns a cli 'action'.
func (nepcalCli) showDate(w io.Writer, t time.Time) func(c *cli.Context) {
	return func(c *cli.Context) {
		showDate(w, t)
	}
}

// Convert AD date to BS date after validation.
func (nepcalCli) convADToBS(c *cli.Context) {
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
}

// Not supported yet.
func (nepcalCli) convBSToAD(c *cli.Context) {
	fmt.Fprintln(writer, "Unfortunately BS to AD isn't supported at this time. :(")
}
