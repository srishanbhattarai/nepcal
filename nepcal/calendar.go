package nepcal

import (
	"fmt"
	"io"
	"text/tabwriter"
)

// helper to generate a formatted string representation of a
// calendar for any given date.
type calendar struct {
	// state that determines how far along the iteration process
	// we have gotten, in generating the calendar for the month
	iter int

	// the time for which the calendar is being created
	when Time

	// the tabwriter to write into, only safe to use in flush()
	tw *tabwriter.Writer
}

func newCalendar(t Time) *calendar {
	return &calendar{
		iter: 1,
		when: t,
		tw:   nil,
	}
}

// Creates the calendar and flushes it onto the underlying tabwriter,
// and transitively to the io.Writer with which it was created.
func (c *calendar) flushInto(w io.Writer) {
	tw := tabwriter.NewWriter(w, 0, 0, 1, ' ', 0)
	c.tw = tw

	c.renderBSDateHeader()
	c.renderStaticDaysHeader()
	c.renderFirstRow()
	c.renderCalWithoutFirstRow()

	tw.Flush()
}

// renderFirstRow renders the first row of the calendar. The reason this needs
// to be handled separately is because there is a offset in each month which
// determines which day the month starts from - we need to tab space the 'offset' number
// of days, then start printing from the day after the offset.
func (c *calendar) renderFirstRow() {
	offset := int(c.when.StartWeekday())
	for i := 0; i < offset; i++ {
		fmt.Fprintf(c.tw, "\t")
	}

	for i := 0; i < (7 - offset); i++ {
		fmt.Fprintf(c.tw, "\t%s", c.reprValue(c.iter))
		c.next()
	}

	fmt.Fprint(c.tw, "\n")
}

// renderCalWithoutFirstRow renders the rest of the calendar without the first row.
// renderFirstRow will handle that due to special circumstances. We basically loop over
// each row and print 7 numbers until we are at the end of the month.
func (c *calendar) renderCalWithoutFirstRow() {
	daysInMonth := c.when.NumDaysInMonth()

	for c.iter <= daysInMonth {
		for i := 0; i < 7; i++ {
			if c.iter > daysInMonth {
				break
			}

			fmt.Fprintf(c.tw, "\t%s", c.reprValue(c.iter))
			c.next()
		}

		fmt.Fprint(c.tw, "\n")
	}
}

// renderStaticDaysHeader prints the static list of days for the calendar
func (c *calendar) renderStaticDaysHeader() {
	for _, v := range []string{"Su", "Mo", "Tu", "We", "Th", "Fr", "Sa"} {
		fmt.Fprintf(c.tw, "%s\t", v)
	}

	fmt.Fprint(c.tw, "\n")
}

// renderBSDateHeader prints the date corresponding to the time 't'. This will
// be the header of the calendar.
func (c *calendar) renderBSDateHeader() {
	yy, mm, dd := c.when.Date()

	fmt.Fprintf(c.tw, "\t\t%s %s, %s\n\t", mm.String(), c.reprValue(dd), c.reprValue(yy))
}

// next increments the value counter.
func (c *calendar) next() {
	c.iter++
}

func (c *calendar) reprValue(val int) string {
	return Numeral(val).String()
}
