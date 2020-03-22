package time

import (
	"fmt"
	"time"
)

// IsInRangeGregorian checks if 't' is after 04/14/1943.
func IsInRangeGregorian(t time.Time) bool {
	adLBound := gregorian(adLBoundY, adLBoundM, adLBoundD)

	return t.After(adLBound)
}

// IsInRangeBS checks if the provided date represents a B.S. date after
// 04/14/1943 and before 30/12/2090 which is the supported date range.
func IsInRangeBS(year int, month Month, day int) bool {
	// Lower bound raw date.
	bslow := raw{bsLBoundY, bsLBoundM, bsLBoundD}

	// Input raw date.
	inraw := raw{year, month, day}

	return after(inraw, bslow)
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

// Check if 't' is after 'u'.
func after(t raw, u raw) bool {
	// Comparing their string representations is an easy way to do this
	// as we do not deal with sub-day precisions.
	tstr := fmt.Sprintf("%d-%02d-%02d", t.year, t.month, t.day)
	ustr := fmt.Sprintf("%d-%02d-%02d", u.year, u.month, u.day)

	return tstr > ustr
}
