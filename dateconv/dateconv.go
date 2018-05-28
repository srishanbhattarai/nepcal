// Package dateconv deals with conversion of A.D. dates to B.S dates as well as
// some utilities to get B.S. month names and day counts for a month.
package dateconv

import (
	"time"
)

// The list of months in the B.S. system.
const (
	baisakh = iota + 1
	jestha
	ashar
	shrawan
	bhadra
	ashoj
	kartik
	mangshir
	poush
	magh
	falgun
	chaitra
)

// bsMonths is a map to get each month's name in the Nepali language.
var bsMonths = map[int]string{
	baisakh:  "बैशाख",
	jestha:   "जेठ",
	ashar:    "असार",
	shrawan:  "सावन",
	bhadra:   "भदौ",
	ashoj:    "असोज",
	kartik:   "कार्तिक",
	mangshir: "मंसिर",
	poush:    "पौष",
	magh:     "माघ",
	falgun:   "फागुन",
	chaitra:  "चैत",
}

// Lower and Upper bounds for AD and BS years along with diffs for
// month and days
const (
	adLBoundY = 1943
	adLBoundM = int(time.April)
	adLBoundD = 14

	bsLBound = 2000
	bsUBound = 2090
)

// bsDaysInMonthsByYear is a map of each BS year from BSLBound to BSUBound with a slice
// of 12 ints indicating the number of days in each month.
var bsDaysInMonthsByYear = map[int][]int{
	bsLBound: []int{30, 32, 31, 32, 31, 30, 30, 30, 29, 30, 29, 31},
	2001:     []int{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
	2002:     []int{31, 31, 32, 32, 31, 30, 30, 29, 30, 29, 30, 30},
	2003:     []int{31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31},
	2004:     []int{30, 32, 31, 32, 31, 30, 30, 30, 29, 30, 29, 31},
	2005:     []int{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
	2006:     []int{31, 31, 32, 32, 31, 30, 30, 29, 30, 29, 30, 30},
	2007:     []int{31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31},
	2008:     []int{31, 31, 31, 32, 31, 31, 29, 30, 30, 29, 29, 31},
	2009:     []int{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
	2010:     []int{31, 31, 32, 32, 31, 30, 30, 29, 30, 29, 30, 30},
	2011:     []int{31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31},
	2012:     []int{31, 31, 31, 32, 31, 31, 29, 30, 30, 29, 30, 30},
	2013:     []int{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
	2014:     []int{31, 31, 32, 32, 31, 30, 30, 29, 30, 29, 30, 30},
	2015:     []int{31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31},
	2016:     []int{31, 31, 31, 32, 31, 31, 29, 30, 30, 29, 30, 30},
	2017:     []int{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
	2018:     []int{31, 32, 31, 32, 31, 30, 30, 29, 30, 29, 30, 30},
	2019:     []int{31, 32, 31, 32, 31, 30, 30, 30, 29, 30, 29, 31},
	2020:     []int{31, 31, 31, 32, 31, 31, 30, 29, 30, 29, 30, 30},
	2021:     []int{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
	2022:     []int{31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 30},
	2023:     []int{31, 32, 31, 32, 31, 30, 30, 30, 29, 30, 29, 31},
	2024:     []int{31, 31, 31, 32, 31, 31, 30, 29, 30, 29, 30, 30},
	2025:     []int{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
	2026:     []int{31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31},
	2027:     []int{30, 32, 31, 32, 31, 30, 30, 30, 29, 30, 29, 31},
	2028:     []int{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
	2029:     []int{31, 31, 32, 31, 32, 30, 30, 29, 30, 29, 30, 30},
	2030:     []int{31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31},
	2031:     []int{30, 32, 31, 32, 31, 30, 30, 30, 29, 30, 29, 31},
	2032:     []int{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
	2033:     []int{31, 31, 32, 32, 31, 30, 30, 29, 30, 29, 30, 30},
	2034:     []int{31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31},
	2035:     []int{30, 32, 31, 32, 31, 31, 29, 30, 30, 29, 29, 31},
	2036:     []int{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
	2037:     []int{31, 31, 32, 32, 31, 30, 30, 29, 30, 29, 30, 30},
	2038:     []int{31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31},
	2039:     []int{31, 31, 31, 32, 31, 31, 29, 30, 30, 29, 30, 30},
	2040:     []int{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
	2041:     []int{31, 31, 32, 32, 31, 30, 30, 29, 30, 29, 30, 30},
	2042:     []int{31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31},
	2043:     []int{31, 31, 31, 32, 31, 31, 29, 30, 30, 29, 30, 30},
	2044:     []int{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
	2045:     []int{31, 32, 31, 32, 31, 30, 30, 29, 30, 29, 30, 30},
	2046:     []int{31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31},
	2047:     []int{31, 31, 31, 32, 31, 31, 30, 29, 30, 29, 30, 30},
	2048:     []int{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
	2049:     []int{31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 30},
	2050:     []int{31, 32, 31, 32, 31, 30, 30, 30, 29, 30, 29, 31},
	2051:     []int{31, 31, 31, 32, 31, 31, 30, 29, 30, 29, 30, 30},
	2052:     []int{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
	2053:     []int{31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 30},
	2054:     []int{31, 32, 31, 32, 31, 30, 30, 30, 29, 30, 29, 31},
	2055:     []int{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
	2056:     []int{31, 31, 32, 31, 32, 30, 30, 29, 30, 29, 30, 30},
	2057:     []int{31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31},
	2058:     []int{30, 32, 31, 32, 31, 30, 30, 30, 29, 30, 29, 31},
	2059:     []int{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
	2060:     []int{31, 31, 32, 32, 31, 30, 30, 29, 30, 29, 30, 30},
	2061:     []int{31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31},
	2062:     []int{30, 32, 31, 32, 31, 31, 29, 30, 29, 30, 29, 31},
	2063:     []int{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
	2064:     []int{31, 31, 32, 32, 31, 30, 30, 29, 30, 29, 30, 30},
	2065:     []int{31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31},
	2066:     []int{31, 31, 31, 32, 31, 31, 29, 30, 30, 29, 29, 31},
	2067:     []int{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
	2068:     []int{31, 31, 32, 32, 31, 30, 30, 29, 30, 29, 30, 30},
	2069:     []int{31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31},
	2070:     []int{31, 31, 31, 32, 31, 31, 29, 30, 30, 29, 30, 30},
	2071:     []int{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
	2072:     []int{31, 32, 31, 32, 31, 30, 30, 29, 30, 29, 30, 30},
	2073:     []int{31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31},
	2074:     []int{31, 31, 31, 32, 31, 31, 30, 29, 30, 29, 30, 30},
	2075:     []int{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
	2076:     []int{31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 30},
	2077:     []int{31, 32, 31, 32, 31, 30, 30, 30, 29, 30, 29, 31},
	2078:     []int{31, 31, 31, 32, 31, 31, 30, 29, 30, 29, 30, 30},
	2079:     []int{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
	2080:     []int{31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 30},
	2081:     []int{31, 31, 32, 32, 31, 30, 30, 30, 29, 30, 30, 30},
	2082:     []int{30, 32, 31, 32, 31, 30, 30, 30, 29, 30, 30, 30},
	2083:     []int{31, 31, 32, 31, 31, 30, 30, 30, 29, 30, 30, 30},
	2084:     []int{31, 31, 32, 31, 31, 30, 30, 30, 29, 30, 30, 30},
	2085:     []int{31, 32, 31, 32, 30, 31, 30, 30, 29, 30, 30, 30},
	2086:     []int{30, 32, 31, 32, 31, 30, 30, 30, 29, 30, 30, 30},
	2087:     []int{31, 31, 32, 31, 31, 31, 30, 30, 29, 30, 30, 30},
	2088:     []int{30, 31, 32, 32, 30, 31, 30, 30, 29, 30, 30, 30},
	2089:     []int{30, 32, 31, 32, 31, 30, 30, 30, 29, 30, 30, 30},
	bsUBound: []int{30, 32, 31, 32, 31, 30, 30, 30, 29, 30, 30, 30},
}

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
