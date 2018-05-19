package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/nepcal/nepcal/internal/conversion"
)

func main() {
	dateFlag := flag.Bool("d", false, "Print only the date")
	flag.Parse()

	if *dateFlag {
		showDate(os.Stdout, time.Now())
	} else {
		renderCalendar(os.Stdout, time.Now())
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
