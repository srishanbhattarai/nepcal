package main

import (
	"fmt"
	"io"
	"os"
	"time"

	dateconv "github.com/srishanbhattarai/nepcal/time"
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
		showDateBS(w, dateconv.ToBS(t), t.Weekday())
	}
}

// Convert AD date to BS date after validation.
func (nepcalCli) convADToBS(c *cli.Context) {
	if !validateArgs(c) {
		fmt.Fprintln(os.Stderr, "Please supply a valid date in the format mm-dd-yyyy. Example: `nepcal conv adtobs 08-21-1994`")
		return
	}

	mm, dd, yy, _ := parseRawDate(c.Args().First())
	adDate := toTime(yy, time.Month(mm), dd)
	if err := dateconv.CheckBoundsAD(adDate); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	bsDate := dateconv.ToBS(adDate)

	showDateBS(writer, bsDate, adDate.Weekday())
}

// Convert BS date to AD date after validation.
func (nepcalCli) convBSToAD(c *cli.Context) {
	if !validateArgs(c) {
		fmt.Fprintln(os.Stderr, "Please supply a valid date in the format mm-dd-yyyy. Example: `nepcal conv bstoad 08-18-2053`")
		return
	}

	mm, dd, yy, _ := parseRawDate(c.Args().First())
	bsdate := dateconv.NewBSDate(yy, mm, dd)
	if err := bsdate.CheckBounds(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	adyy, admm, addd := dateconv.ToAD(dateconv.NewBSDate(yy, mm, dd)).Date()
	date := toTime(adyy, time.Month(admm), addd)

	showDateAD(writer, date)
}

// Validates the arguments provided to the program.
func validateArgs(c *cli.Context) bool {
	if c.NArg() < 1 {
		return false
	}

	_, _, _, ok := parseRawDate(c.Args().First())

	return ok
}
