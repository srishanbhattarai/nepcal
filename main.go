package main

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/nepcal/nepcal/internal/conversion"
)

// Cheap testing.
var (
	stdout io.Writer = os.Stdout
)

func main() {
	showDate(time.Now())
}

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
