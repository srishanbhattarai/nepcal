package main

import (
	"fmt"
	"time"

	"github.com/nepcal/nepcal/internal/conversion"
)

func main() {
	now := time.Now()
	yy, mm, dd := now.Date()

	showDate(yy, int(mm), dd)
}

func showDate(yy, mm, dd int) {
	bs := conversion.ToBS(conversion.Epoch{yy, int(mm), dd})
	fmt.Printf("%s %d, %d\n", conversion.BSMonths[bs.Month], bs.Day, bs.Year)
}
