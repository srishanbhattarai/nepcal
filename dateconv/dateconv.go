// Package dateconv deals with conversion of A.D. dates to B.S dates as well as
// some utilities to get B.S. month names and day counts for a month.
package dateconv

import (
	"fmt"
	"time"
)

// ToBS handles conversion of an Anno Domini (A.D) date into the Nepali
// date format - Bikram Samwad (B.S).The approximate difference is
// 56 years, 8 months.
func ToBS(adDate time.Time) time.Time {
	adLBound := toTime(adLBoundY, adLBoundM, adLBoundD)
	if !adDate.After(adLBound) {
		panic("Can only work with dates after 1943 April 14.")
	}
	totalDiff := int(adDate.Sub(adLBound).Hours() / 24)

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

	return toTime(year, month, days)
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
func TotalDaysInBSYear(bsYear int) (int, error) {
	sum := 0
	days, ok := bsDaysInMonthsByYear[bsYear]

	if ok == false {
		return -1, fmt.Errorf("Year should be in between %d and %d", bsLBound, bsUBound)
	}

	for _, value := range days {
		sum = sum + value
	}

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
