package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/nepcal/nepcal/internal/conversion"
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
	if dateFlag {
		showDate(writer, time.Now())
	} else {
		cal := newCalendar()
		cal.Render(writer, time.Now())
	}
}

// showDate prints the current B.S. date
func showDate(w io.Writer, t time.Time) {
	yy, mm, dd := t.Date()

	bs := conversion.ToBS(
		conversion.Epoch{
			Year:  yy,
			Month: int(mm),
			Day:   dd,
		},
	)

	fmt.Fprintf(w, "%s %d, %d\n", conversion.BSMonths[bs.Month], bs.Day, bs.Year)
}
