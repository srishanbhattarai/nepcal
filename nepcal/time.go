package nepcal

import (
	"time"
)

// A Month represents an integer Bikram Sambat month.
type Month int

// String returns the plain text Nepali version of the provided month.
func (m Month) String() string {
	return nepMonths[m]
}

// A Weekday represents an integer Bikram Sambat weekday.
type Weekday int

// String returns the plain text Nepali version of the provided day.
func (wd Weekday) String() string {
	return nepWeekdays[wd]
}

// A Time represents a Bikram Sambat date.
// The Time struct tries to maintain a similar API to `time.Time`.
// Additional information about a particular `Time` can be computed with methods
// that are available on this struct. A new `Time` instance can be created from
// an Anno Domini date using a UNIX timestamp, via the `nepcal.Unix` function.
type Time struct {
	year  int
	month Month
	day   int
	time  time.Time
}

// Now returns the current B.S. time.
func Now() Time {
	return ToBS(time.Now())
}

// Date returns the yy, mm, dd values in which t occurs.
func (t Time) Date() (int, Month, int) {
	return t.Year(), t.Month(), t.Day()
}

// Year returns the year in which t occurs.
func (t Time) Year() int {
	return t.year
}

// Month returns the month of the year specified by t.
func (t Time) Month() Month {
	return t.month
}

// Day returns the day of the month specified by t.
func (t Time) Day() int {
	return t.day
}

// Weekday returns the weekday of the week specified by t.
func (t Time) Weekday() int {
	panic("Not implemented")
}

// DaysInMonth returns the total number of days in the month for this date.
// The invariant here is that 't.month' is always a valid month.
func (t Time) DaysInMonth() (int, bool) {
	return BsDaysInMonthsByYear(t.year, t.month)
}

// StartingWeekdayOfMonth calculates the offset at the beginning of the month.
// Given an AD date we calculate the diff in days from the BS date to the start
// of the month in BS. We subtract that from the AD date, and get the weekday.
func (t Time) StartingWeekdayOfMonth() Weekday {
	dayDiff := (t.Day() % 7) - 1
	adWithoutbsDiffDays := t.time.AddDate(0, 0, -dayDiff)
	d := adWithoutbsDiffDays.Weekday()

	// Since Weekday is an iota and not an iota + 1 we can avoid
	// subtracting 1 from the return value.
	return Weekday(int(d))
}

// TotalDaysSpanned returns the total number of days spanned in the year
// specified by t, inclusive of the current day.
func (t Time) TotalDaysSpanned() int {
	days := bsDaysInMonthsByYear[t.Year()]

	sum := t.Day()
	for i := 0; i < int(t.Month())-1; i++ {
		sum += days[i]
	}

	return sum
}
