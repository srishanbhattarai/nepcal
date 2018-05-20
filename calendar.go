package main

import (
	"fmt"
	"io"
	"text/tabwriter"
	"time"

	"github.com/srishanbhattarai/nepcal/dateconv"
)

// calendar struct represents the state required to render
// the B.S. calendar
type calendar struct {
	val int
}

// newCalendar returns a new instance of calendar with the initial value of 1.
func newCalendar() *calendar {
	return &calendar{1}
}

// Render prints the BS calendar for the given time.Time.
// For printing formatted/aligned output, we use a tabwriter from the
// standard library. It doesn't support ANSI escapes so we cant have
// color/other enhancements to the output.(https://github.com/nepcal/nepcal/issues/4)
func (c *calendar) Render(parentWriter io.Writer, t time.Time) {
	w := tabwriter.NewWriter(parentWriter, 0, 0, 1, ' ', 0)
	ad := toEpoch(t)
	bs := dateconv.ToBS(ad)

	c.renderBSDateHeader(w, bs)
	c.renderStaticDaysHeader(w)
	c.renderFirstRow(w, ad, bs)
	c.renderCalWithoutFirstRow(w, ad, bs)

	w.Flush()
}

// renderFirstRow renders the first row of the calendar. The reason this needs
// to be handled separately is because there is a skew in each month which
// determines which day the month starts from - we need to tab space the 'skew' number
// of days, then start printing from the day after the skew.
func (c *calendar) renderFirstRow(w io.Writer, ad, bs dateconv.Epoch) {
	offset := c.calculateSkew(ad, bs)
	for i := 0; i < offset; i++ {
		fmt.Fprintf(w, "\t")
	}

	for i := 0; i < (7 - offset); i++ {
		fmt.Fprintf(w, "\t%d", c.val)
		c.val++
	}

	fmt.Fprint(w, "\n")
}

// renderCalWithoutFirstRow renders the rest of the calendar without the first row.
// renderFirstRow will handle that due to special circumstances. We basically loop over
// each row and print 7 numbers until we are at the end of the month.
func (c *calendar) renderCalWithoutFirstRow(w io.Writer, ad, bs dateconv.Epoch) {
	daysInMonth := dateconv.BsDaysInMonthsByYear[bs.Year][bs.Month-1]

	for c.val < daysInMonth {
		start := daysInMonth - c.val
		end := start + 7

		for i := start; i < end; i++ {
			if c.val > daysInMonth {
				break
			}

			fmt.Fprintf(w, "\t%d", c.val)
			c.val++
		}

		fmt.Fprint(w, "\n")
	}
}

// renderStaticDaysHeader prints the static list of days for the calendar
func (c *calendar) renderStaticDaysHeader(w io.Writer) {
	for _, v := range []string{"Su", "Mo", "Tu", "We", "Th", "Fr", "Sa"} {
		fmt.Fprintf(w, "%s\t", v)
	}

	fmt.Fprint(w, "\n")
}

// renderBSDateHeader prints the date corresponding to the epoch.
func (c *calendar) renderBSDateHeader(w io.Writer, e dateconv.Epoch) {
	fmt.Fprintf(w, "\t\t%s %d, %d\n\t", dateconv.BSMonths[e.Month], e.Day, e.Year)
}

// calculateSkew calculates the offset at the beginning of the month. Given an AD and
// BS date, we calculate the diff in days from the BS date to the start of the month in BS.
// We subtract that from the AD date, and get the weekday.
// For example, a skew of 2 means the month starts from Tuesday.
func (c *calendar) calculateSkew(ad, bs dateconv.Epoch) int {
	adDate := fromEpoch(ad)
	dayDiff := (bs.Day % 7) - 1
	adWithoutbsDiffDays := adDate.AddDate(0, 0, -dayDiff)
	d := adWithoutbsDiffDays.Weekday()

	// Since time.Weekday is an iota and not an iota + 1 we can avoid
	// subtracting 1 from the return value.
	return int(d)
}

// fromEpoch creates a time.Time type from an Epoch
func fromEpoch(e dateconv.Epoch) time.Time {
	return time.Date(e.Year, time.Month(e.Month), e.Day, 0, 0, 0, 0, time.UTC)
}

// toEpoch creates a dateconv.Epoch from a time.Time
func toEpoch(t time.Time) dateconv.Epoch {
	yy, mm, dd := t.Date()

	return dateconv.Epoch{
		Year:  yy,
		Month: int(mm),
		Day:   dd,
	}
}
