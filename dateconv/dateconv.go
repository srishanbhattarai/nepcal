// Package dateconv deals with conversion of A.D. dates to B.S dates as well as
// some utilities to get B.S. month names and day counts for a month.
package dateconv

import (
	"fmt"
	"time"
)

// BSDate represents a Bikram Sambat date.
type BSDate struct {
	year  int
	month int
	days  int
}

// Date returns the yy, mm, dd values represented by the BS date.
func (b BSDate) Date() (int, int, int) {
	return b.year, b.month, b.days
}

// DaysInMonth returns the total number of days in the month for this date.
// The invariant here is that 'b.month' is always a valid month.
func (b BSDate) DaysInMonth() (int, bool) {
	return BsDaysInMonthsByYear(b.year, time.Month(b.month))
}

// MonthStartsAtDay calculates the offset at the beginning of the month. Given an AD date
// we calculate the diff in days from the BS date to the start of the month in BS.
// We subtract that from the AD date, and get the weekday.
// For example, a value of 2 means the month starts from Tuesday.
func (b BSDate) MonthStartsAtDay(ad time.Time) int {
	dayDiff := (b.days % 7) - 1
	adWithoutbsDiffDays := ad.AddDate(0, 0, -dayDiff)
	d := adWithoutbsDiffDays.Weekday()

	// Since time.Weekday is an iota and not an iota + 1 we can avoid
	// subtracting 1 from the return value.
	return int(d)
}

// newBSDate is a constructor for a new Bikram Sambat date.
func newBSDate(yy, mm, dd int) BSDate {
	return BSDate{yy, mm, dd}
}

// ToBS handles conversion of an Anno Domini (A.D) date into the Nepali
// date format - Bikram Sambat (B.S).The approximate difference is
// 56 years, 8 months.
func ToBS(adDate time.Time) BSDate {
	adLBound := toTime(adLBoundY, adLBoundM, adLBoundD)

	// Convert incoming date to UTC
	adYear, adMonth, adDay := adDate.Date()
	adDateUTC := toTime(adYear, int(adMonth), adDay)

	if !adDate.After(adLBound) {
		panic("Can only work with dates after 1943 April 14.")
	}

	totalDiff := int(adDateUTC.Sub(adLBound).Hours() / 24)

	// Redistribute the diff along the BS data grid
	year, month, days := func() (int, int, int) {
		for i := bsLBound; i < bsUBound; i++ {
			for j := 0; j < 12; j++ {
				days := bsDaysInMonthsByYear[i][j]

				if days <= totalDiff {
					totalDiff = totalDiff - days
					continue
				}

				return i, j + 1, totalDiff + 1
			}
		}

		return -1, -1, -1
	}()

	return newBSDate(year, month, days)
}

// GetBSMonthName returns the B.S. month name from the time.Month type.
// Example: GetBSMonthName(1) === बैशाख
func GetBSMonthName(bsMonth time.Month) (string, bool) {
	mth, ok := bsMonths[int(bsMonth)]

	return mth, ok
}

// GetNepWeekday returns Nepali weekday from the time.Time type.
// Example: getNepWeekday(0) === आइतबार
func GetNepWeekday(weekday time.Weekday) (string, bool) {
	nepWeekday, ok := nepWeekdays[int(weekday)]

	return nepWeekday, ok
}

// BsDaysInMonthsByYear returns the number of days in the month 'mm'
// in the year 'yy'. Note that it is assumed that months start from 1
// the caller does not have to subtract by one when calling the function.
// yy must be between 2000 and 2090
// mm must be between 1 and 12.
func BsDaysInMonthsByYear(yy int, mm time.Month) (int, bool) {
	months, ok := bsDaysInMonthsByYear[yy]
	if !ok {
		return 0, ok
	}

	query := int(mm) - 1

	if query > 11 || query < 0 {
		return 0, false
	}

	return months[query], true
}

// TotalDaysInBSYear returns total number of days in a particular BS year.
func TotalDaysInBSYear(year int) (int, error) {
	days, ok := bsDaysInMonthsByYear[year]

	if !ok {
		return -1, fmt.Errorf("Year should be in between %d and %d", bsLBound, bsUBound)
	}

	sum := 0

	for _, value := range days {
		sum += value
	}

	return sum, nil
}

// now returns the current time
var now = func() time.Time {
	return time.Now()
}

// TotalDaysSpanned returns the total number of days spanned in the
// current year inclusive of the current day.
func TotalDaysSpanned() (int, error) {
	bsDate := ToBS(now())
	days, ok := bsDaysInMonthsByYear[bsDate.year]

	if !ok {
		return -1, fmt.Errorf("Year should be in between %d and %d", bsLBound, bsUBound)
	}

	sum := 0

	for i := 0; i < bsDate.month; i++ {
		sum += days[i]
	}

	sum += bsDate.days

	return sum, nil
}

// toTime creates a new time.Time with the basic yy/mm/dd parameters.
func toTime(yy, mm, dd int) time.Time {
	return time.Date(yy, time.Month(mm), dd, 0, 0, 0, 0, time.UTC)
}

// isLeapYear returns if the passed in year is a leap year.
func isLeapYear(year int) bool {
	if year%4 != 0 {
		return false
	}

	if year%100 == 0 && year%400 == 0 {
		return true
	}

	return false
}

// adDaysInMonths is the number of days in each month in a year which is only dependent on
// the leap year status. This function is the equivalent of the bsDaysInMonthsByYear map for
// AD dates.
func adDaysInMonths(isLeapYear bool) []int {
	normalData := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	leapData := []int{31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	if isLeapYear {
		return leapData
	}

	return normalData
}
