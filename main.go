package main

import (
	"fmt"
	"io"
	"os"
	"text/tabwriter"
	"time"

	"github.com/nepcal/nepcal/internal/conversion"
)

// Cheap testing.
var (
	stdout io.Writer = os.Stdout
)

func main() {
	// showDate(time.Now())
	showCal()
}

// showDate prints the current B.S. date
func showDate(t time.Time) {
	yy, mm, dd := t.Date()

	bs := conversion.ToBS(
		conversion.Epoch{
			Year:  yy,
			Month: int(mm),
			Day:   dd,
		},
	)

	fmt.Fprintf(stdout, "%s %d, %d\n", conversion.BSMonths[bs.Month], bs.Day, bs.Year)
}

// showCal prints the current calendar.
func showCal() {
	yy, mm, dd := time.Now().Date()
	w := tabwriter.NewWriter(stdout, 0, 0, 1, ' ', 0)

	bs := conversion.ToBS(
		conversion.Epoch{
			Year:  yy,
			Month: int(mm),
			Day:   dd,
		},
	)

	// Month header
	fmt.Fprintf(w, "\t\t%s %d, %d\n\t", conversion.BSMonths[bs.Month], bs.Day, bs.Year)

	// Days list
	printDay := func(day string) {
		fmt.Fprintf(w, "%s\t", day)
	}
	for _, v := range []string{"Su", "Mo", "Tu", "We", "Th", "Fr", "Sa"} {
		printDay(v)
	}
	fmt.Fprint(w, "\n")

	// Calendar
	printVal := func(v int) {
		fmt.Fprintf(w, "\t%d", v)
	}

	daysInMonth := conversion.BsDaysInMonthsByYear[bs.Year][bs.Month]
	offset := dd % 7
	val := 1
	for i := 0; i < offset-1; i++ {
		fmt.Fprintf(w, "\t")
	}
	for i := 0; i <= offset+1; i++ {
		printVal(val)
		val++
	}
	fmt.Fprint(w, "\n")

	for val < daysInMonth {
		start := daysInMonth - val
		end := start + 7

		for i := start; i < end; i++ {
			if val > daysInMonth {
				break
			}

			printVal(val)
			val++
		}
		fmt.Fprint(w, "\n")
	}
	w.Flush()
}
