package nepcal

import (
	"time"
)

// IsInRangeGregorian checks if 't' is inside Nepcal's supported range of dates.
func IsInRangeGregorian(t time.Time) bool {
	adLBound := gregorian(adLBoundY, adLBoundM, adLBoundD)
	adUBound := gregorian(adUBoundY, adUBoundM, adUBoundD)

	satisfiesLowerBound := t.Equal(adLBound) || t.After(adLBound)
	satisfiesUpperBound := t.Equal(adUBound) || t.Before(adUBound)

	return satisfiesLowerBound && satisfiesUpperBound
}

// IsInRangeBS checks if the provided date represents a B.S. that
// we have data for and can be supported for conversions to/from A.D.
func IsInRangeBS(year int, month Month, day int) bool {
	if month < Baisakh || month > Chaitra {
		return false
	}

	t := (raw{year, month, day}).String()

	// Lower bound raw date.
	bslow := raw{bsLBoundY, bsLBoundM, bsLBoundD}
	bshigh := raw{bsUBoundY, bsUBoundM, bsUBoundD}

	satisfiesLowerBound := t >= bslow.String()
	satisfiesUpperBound := t <= bshigh.String()

	return satisfiesLowerBound && satisfiesUpperBound
}

// IsInRangeYear return true if the provided bsYear is within the supported
// BS date range.
func IsInRangeYear(bsYear int) bool {
	return bsYear >= bsLBoundY && bsYear <= bsUBoundY
}

// gregorian creates a new time.Time with the basic yy/mm/dd parameters.
// Crucially, the time returned is in UTC.
func gregorian(yy, mm, dd int) time.Time {
	return time.Date(yy, time.Month(mm), dd, 0, 0, 0, 0, time.UTC)
}

// Sum the number of days in the year: 365 or 366.
func numDaysInYear(year int) int {
	// Invariant: year always in bounds.
	daysDistribution, _ := bsDaysInMonthsByYear[year]

	sum := 0
	for _, value := range daysDistribution {
		sum += value
	}

	return sum
}
