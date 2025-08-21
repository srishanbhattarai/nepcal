// Package nepcal is the root of all functionality for B.S. dates.
//
// This package is similar to the `time.Time` package in the standard
// library which works with Gregorian dates.
//
// Usage
//
// Anywhere a Gregorian date is expected, the standard `time.Time` struct
// is required. To avoid confusion, anywhere the func names, comments etc.
// do _not_ explicitly mention "gregorian", it refers to the B.S. values.
// For example, the `IsInRangeYear` function is for B.S. dates whereas
// IsInRangeGregorian is explicitly for Gregorian dates.
//
// This package only works for a fixed range of dates, which can be checked using the
// IsInRangeBS or IsInRangeGregorian methods. As an invariant, various methods will
// require that the dates they work with to be within this range.
// In general, this is only relevant when _constructing_ the dates.
package nepcal

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"time"
)

// ErrOutOfBounds is the error returned for out of bounds Gregorian dates.
var ErrOutOfBounds = errors.New("Provided date out of bounds; consult function/method documentation")

// A Time struct represents a single Bikram Sambat date. An instance of this struct is
// the primary way to interact with most functionality.
// It can be created in two ways:
//	 1. If you have a Gregorian date, and want a B.S. date, use the "FromGregorian" method.
//	 2. If you have B.S. date, and want to access additional functionality on it, or convert it to
//		Gregorian, then use the "Date" method.
type Time struct {
	// The inner Gregorian date that this Time corresponds to, in UTC.
	in time.Time

	// BS date specific information.
	year  int
	month Month
	day   int
}

// A 'raw' date is an internally used struct to represent yy/mm/dd triples.
type raw struct {
	year  int
	month Month
	day   int
}

// String satisfies the stringer interface for 'raw'. The string format returned
// is intended to be used to establish order between dates.
func (r raw) String() string {
	return fmt.Sprintf("%d-%02d-%02d", r.year, r.month, r.day)
}

// Now returns the nepcal.Time struct corresponding to the current time.
// This method uses FromGregorianUnchecked - read the documentation
// for that method to understand limitations.
func Now() Time {
	return FromGregorianUnchecked(time.Now())
}

// CalendarNow returns a Reader that returns a formatted calendar representation of the current date.
// This method uses FromGregorianUnchecked - read the documentation for that method to understand limitations.
func CalendarNow() io.Reader {
	t := Now()

	return t.Calendar()
}

// FromGregorian constructs a Bikram Sambat date from the provided Gregorian date.
// This function returns an error if the date is out of the supported date range,
// as defined in the 'IsInRangeGregorian' function.
func FromGregorian(t time.Time) (Time, error) {
	if !IsInRangeGregorian(t) {
		return Time{}, ErrOutOfBounds
	}

	return FromGregorianUnchecked(t), nil
}

// FromGregorianUnchecked performs the same function as FromGregorianDate
// minus any bounds checking using the IsInRangeGregorian function.
// If you are _absolutely_ sure that your gregorian date will is within
// the supported date range of this package, then you can use this unchecked constructor.
//
// If you violate this invariant, the correctness of other methods that use the returned value is undefined.
//
// In summary - for all times 't' such that IsInRangeGregorian(t) == true, this function is safe to use.
// An example of where this is useful is when you are constructing from today's date.
func FromGregorianUnchecked(t time.Time) Time {
	return fromGregorian(t)
}

// Date constructs a B.S. date using raw parts "year, month, date". As with the,
// "From_" constructors, the specified B.S date must be in the supported range as
// specified by the IsInRangeBS function.
//
// Unlike Go's time package, Nepcal does not auto normalize dates, therefore if a date was passed in that
// would have wrapped over, an error is returned.
func Date(year int, month Month, day int) (Time, error) {
	if month < Baisakh || month > Chaitra {
		return Time{}, fmt.Errorf("Month values can only be between 1 and 12, but a value of %d was specified", int(month))
	}

	if !IsInRangeBS(year, month, day) {
		return Time{}, ErrOutOfBounds
	}

	// Check that the provided day actually exists in the month. While Go's time package
	// performs auto date normalization (i.e. February 30 is normalized to April 2nd),
	// Nepcal does not do this, since this is most likely a user error.
	daysInMonth := bsDaysInMonthsByYear[year][int(month)-1]
	if day > daysInMonth {
		return Time{}, fmt.Errorf(fmt.Sprintf("The month of %s has only %d days in the year %d, but the provided date specifies a value of %d. Nepcal does not perform auto date normalization.", month, daysInMonth, year, day))
	}

	inraw := raw{year, month, day}
	return fromRaw(inraw), nil
}

// DateUnchecked is the unchecked range variant of 'Date' similar to FromGregorianUnchecked. The
// same invariants apply - if the date is not in the valid range as indicated by IsInRangeBS,
// then the correctness of using the returned value is undefined.
func DateUnchecked(year int, month Month, day int) Time {
	inraw := raw{year, month, day}

	return fromRaw(inraw)
}

// Gregorian returns the A.D. equivalent of this date. If this struct was initially created
// from a gregorian date, then it returns the same input date. Otherwise, if it was created from a raw
// B.S. date using the "Date" method, then it returns the A.D. representation of that date.
// Note that the "Date" method already does the conversion during creation, so this method
// is free of any computation in either of the two cases.
func (t Time) Gregorian() time.Time {
	return t.in
}

// Date returns the yy, mm, dd values represented by the Time. To get individual values, use the
// respective 'Year', 'Month' or 'Day' methods.
func (t Time) Date() (int, Month, int) {
	return t.Year(), t.Month(), t.Day()
}

// Year returns the B.S. year value for this date.
func (t Time) Year() int {
	return t.year
}

// Month returns the B.S. month value for this date.
func (t Time) Month() Month {
	return t.month
}

// Day returns the B.S. day of the month for this date.
func (t Time) Day() int {
	return t.day
}

// Weekday returns the B.S. weekday for this date.
func (t Time) Weekday() Weekday {
	return Weekday(t.in.Weekday())
}

// StartWeekday returns the Weekday at which this date's month starts on.
func (t Time) StartWeekday() Weekday {
	dayDiff := (t.day % 7) - 1
	adWithoutbsDiffDays := t.in.AddDate(0, 0, -dayDiff)
	d := adWithoutbsDiffDays.Weekday()

	return Weekday(int(d))
}

// NumDaysInMonth returns the number of days in the month for this B.S. date.
// Each month has a different number of days, and this also differs each year.
func (t Time) NumDaysInMonth() int {
	return t.month.numDaysUnchecked(t.year)
}

// NumDaysInYear returns the total number of days in this year. Practically, this
// will always be 365 or 366.
func (t Time) NumDaysInYear() int {
	return numDaysInYear(t.year)
}

// NumDaysSpanned returns the number of days spanned in the current year for
// this date.
func (t Time) NumDaysSpanned() int {
	// Invariant: year always in bounds.
	days, _ := bsDaysInMonthsByYear[t.year]

	sum := t.day
	for i := 0; i < int(t.month)-1; i++ {
		sum += days[i]
	}

	return sum
}

// Calendar returns an io.Reader which can be used to read the calendar
// representation of this date.
func (t Time) Calendar() io.Reader {
	buf := bytes.NewBuffer([]byte(""))

	newCalendar(t).flushInto(buf)

	return buf
}

// After reports whether the Time t, is after u.
func (t Time) After(u Time) bool {
	return t.toRaw().String() > u.toRaw().String()
}

// String satisfies the stringer interface.
func (t Time) String() string {
	return fmt.Sprintf("%s %s, %s %s", t.Month(), Numeral(t.Day()), Numeral(t.Year()), t.Weekday())
}

// Internal method to generate raw dates from valid B.S. dates.
func (t Time) toRaw() raw {
	return raw{t.year, t.month, t.day}
}
