package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/srishanbhattarai/nepcal/dateconv"
)

// Cheap testing.
var writer io.Writer = os.Stdout

// Flag list
var (
	dateFlag = flag.Bool("d", false, "Print only the date")
)

func init() {
	flag.Parse()
}

func main() {
	render(*dateFlag)
}

// Render decides what to show based on the flags.
func render(dateFlag bool) {
	if len(os.Args) > 3 {
		var (
			yy, _ = strconv.Atoi(os.Args[3])
			mm, _ = strconv.Atoi(os.Args[2])
			dd, _ = strconv.Atoi(os.Args[1])
		)

		showDate(writer, time.Date(yy, time.Month(mm), dd, 0, 0, 0, 0, time.UTC))
		return
	}

	if dateFlag {
		showDate(writer, time.Now())
		return
	}

	cal := newCalendar()
	cal.Render(writer, time.Now())
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
