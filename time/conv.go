package time

import "time"

// fromGregorian constructs a valid Bikram Sambat date from an in-bounds
// Gregorian date.
//
// The conversion process in 5 year old speak:
// The idea is that the Gregorian time 't' is subtracted from the lowest
// bound time possible (adLBound* values) and this difference is obtained in
// the unit of 'number of days'. Since the lowest AD time possible corresponds
// to the lowest BS time possible, the number of days elapsed can be used as a
// relative value in both date systems.
//
// For example, 100 days after adLBound is 100 days after bsLBound.
// Since we know the number of days in each month for each year in BS dates
// (which is very inconsistently distributed for non-obvious reasons),
// the 100 days from adLBound is mapped to a certain month and date by
// incremental subtraction.  In this case, the first 3 months in bsLBoundY
// have 30, 32, and 31 days respectively. This adds up to 93, which is < 100,
// and the first 4 months add up to 125 (month 4 being 32). This implies that
// day 100 is month 4, day 7 (100 - 93 = 7).
func fromGregorian(t time.Time) Time {
	// Lower bound gregorian date
	glow := gregorian(adLBoundY, adLBoundM, adLBoundD)

	// Convert incoming gregorian date to UTC
	gy, gm, gd := t.Date()
	g := gregorian(gy, int(gm), gd)

	// "g - glow" to get relative number of days elapsed.
	daysElapsed := int(g.Sub(glow).Hours() / 24)

	// find the BS date according to the reasoning above, distributing the
	// daysElapsed along the data grid.
	year, month, days := func() (int, int, int) {
		for i := bsLBoundY; i < bsUBoundY; i++ {
			for j := 0; j < 12; j++ {
				days := bsDaysInMonthsByYear[i][j]

				if days <= daysElapsed {
					daysElapsed = daysElapsed - days
					continue
				}

				return i, j + 1, daysElapsed + 1
			}
		}

		return -1, -1, -1
	}()

	return Time{t, year, Month(month), days}
}

// Constructs a valid B.S. time from a 'raw' B.S. time.
//
// Typically the Time struct holds a gregorian `time.Time` and the B.S. date
// which the gregorian time corresponds to. The function 'fromGregorian' takes
// in the gregorian time.Time and finds out the the B.S. date values, to create
// the eventual struct.
// This 'fromRaw' function does the inverse - it takes a 'raw' date which is
// a (y,m,d) triple and figures out the time.Time gregorian value.
//
// The above description effectively means this function does a BS to Gregorian
// conversion during the struct creation. Any later requests to generate the
// Gregorian equivalent of a BS date is effectively free.
//
// The calculations are the inverse of what happens in fromGregorian, starting
// with a count of zero and adding the days until the difference is reached.
func fromRaw(r raw) Time {
	// Lower bound gregorian date
	glow := gregorian(adLBoundY, adLBoundM, adLBoundD)

	// Partially construct the struct; allowing us to call methods on it for some
	// of the computations below.
	t := Time{
		year:  r.year,
		month: r.month,
		day:   r.day,
	}

	year, month, day := func() (int, time.Month, int) {
		totalDiff := 0

		// Count the number of days in the years
		for i := bsLBoundY; i < r.year; i++ {
			totalDiff += numDaysInYear(i)
		}

		// Count the number of days in the months
		for i := 0; i < int(r.month)-1; i++ {
			totalDiff += bsDaysInMonthsByYear[r.year][i]
		}

		// Add the leftover days
		totalDiff += r.day - 1

		return glow.AddDate(0, 0, totalDiff).Date()
	}()

	g := gregorian(year, int(month), day)

	// Set the inner Gregorian date.
	t.in = g

	return t
}
