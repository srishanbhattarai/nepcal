package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/srishanbhattarai/nepcal/nepcal"
	"github.com/urfave/cli/v2"
)

// nepcalCli is a struct to hold all the CLI behavior.
// It is a small wrapper around the urfave/cli package to keep things clean.
type nepcalCli struct{}

// Shows the calendar for the current day.
func (nepcalCli) showCalendar(c *cli.Context) error {
	// get the calendar representation
	calReader := nepcal.CalendarNow()

	// stream into globalwriter
	io.Copy(globalWriter, calReader)

	return nil
}

// Shows the date for the provided time. Returns a cli 'action'.
func (nepcalCli) showDate(w io.Writer, t time.Time) func(c *cli.Context) error {
	return func(c *cli.Context) error {
		// This will stop working in year bsUBoundY + 1 (:
		bs := nepcal.FromGregorianUnchecked(t)

		fmt.Fprintln(w, bs.String())

		return nil
	}
}

// Convert AD date to BS date after validation.
func (nepcalCli) convADToBS(c *cli.Context) error {
	if !validateArgs(c) {
		fmt.Fprintln(os.Stderr, "Please supply a valid date in the format mm-dd-yyyy. Example: `nepcal conv tobs 08-21-1994`")

		return cli.Exit("", 1)
	}

	mm, dd, yy, _ := parseRawDate(c.Args().First())

	ad := gregorian(yy, mm, dd)
	bs, err := nepcal.FromGregorian(ad)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Please supply a date after 04/14/1943.")

		return cli.Exit("", 1)
	}

	fmt.Fprintln(globalWriter, bs.String())

	return nil
}

// Convert BS date to AD date after validation.
func (nepcalCli) convBSToAD(c *cli.Context) error {
	if !validateArgs(c) {
		fmt.Fprintln(os.Stderr, "Please supply a valid date in the format mm-dd-yyyy. Example: `nepcal conv toad 08-18-2053`")

		return cli.Exit("", 1)
	}

	mm, dd, yy, _ := parseRawDate(c.Args().First())

	d, err := nepcal.Date(yy, nepcal.Month(mm), dd)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Please ensure the date is between 1/1/2000 and 12/30/2095")

		return cli.Exit("", 1)
	}

	printGregorian(globalWriter, d.Gregorian())

	return nil
}

// Validates the arguments provided to the program.
func validateArgs(c *cli.Context) bool {
	if c.NArg() < 1 {
		return false
	}

	_, _, _, ok := parseRawDate(c.Args().First())

	return ok
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

func printGregorian(w io.Writer, t time.Time) {
	adyy, _, addd := t.Date()
	month := t.Month()
	weekday := t.Weekday()

	fmt.Fprintf(w, "%s %d, %d %s\n", month, addd, adyy, weekday)
}

// gregorian creates a new time.Time with the basic yy/mm/dd parameters.
// Crucially, the time returned is in UTC.
func gregorian(yy, mm, dd int) time.Time {
	return time.Date(yy, time.Month(mm), dd, 0, 0, 0, 0, time.UTC)
}
