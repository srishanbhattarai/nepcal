package dateconv

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestToBS(t *testing.T) {
	tests := []struct {
		name   string
		input  time.Time
		output BSDate
	}{
		{
			"case1",
			toTime(2018, 04, 01),
			NewBSDate(2074, 12, 18),
		},
		{
			"case2",
			toTime(1943, 04, 15),
			NewBSDate(2000, 01, 02),
		},
		{
			"case3",
			toTime(2018, 04, 17),
			NewBSDate(2075, 01, 04),
		},
		{
			"case4",
			toTime(2018, 05, 01),
			NewBSDate(2075, 01, 18),
		},
		{
			"case5",
			toTime(1960, 9, 16),
			NewBSDate(2017, 06, 1),
		},
		{
			"case6",
			toTime(2037, 9, 16),
			NewBSDate(-1, -1, -1),
		},
		{
			"case7",
			toTime(2019, 06, 15),
			NewBSDate(2076, 02, 32),
		},
		{
			"case8",
			toTime(2019, 06, 13),
			NewBSDate(2076, 02, 30),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.output, ToBS(test.input))
		})
	}

	t.Run("panics if date is before 1943 April 14", func(t *testing.T) {
		assert.Panics(t, func() {
			ToBS(toTime(1943, 04, 01)) // april 1
		}, "Can only work with dates after 1943 April 14")
	})
}

func TestIsLeapYear(t *testing.T) {
	tests := []struct {
		name     string
		year     int
		expected bool
	}{
		{"343", 343, false},
		{"100", 100, false},
		{"1700", 1700, false},
		{"2100", 2100, false},
		{"1600", 1600, true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, isLeapYear(test.year))
		})
	}
}

func TestAdDaysInMonths(t *testing.T) {
	normalData := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	leapData := []int{31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	tests := []struct {
		name       string
		isLeapYear bool
		expected   []int
	}{
		{"leap year", true, leapData},
		{"not leap year", false, normalData},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			data := adDaysInMonths(test.isLeapYear)

			sum := func(d []int) int {
				s := 0
				for _, v := range d {
					s = s + v
				}

				return s
			}

			assert.ElementsMatch(t, test.expected, data)
			assert.Equal(t, sum(test.expected), sum(data))
		})
	}
}

func TestGetBSMonthName(t *testing.T) {
	tests := []struct {
		name         string
		month        time.Month
		expectedStr  string
		expectedBool bool
	}{
		{"when in range", time.Month(1), "बैशाख", true},
		{"when not in range", time.Month(100), "", false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mth, ok := GetBSMonthName(test.month)

			assert.Equal(t, test.expectedStr, mth)
			assert.Equal(t, test.expectedBool, ok)
		})
	}
}

func TestGetNepWeekday(t *testing.T) {
	tests := []struct {
		name         string
		weekday      time.Weekday
		expectedStr  string
		expectedBool bool
	}{
		{"when in range", time.Weekday(0), "आइतबार", true},
		{"when not in range", time.Weekday(7), "", false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mth, ok := GetNepWeekday(test.weekday)

			assert.Equal(t, test.expectedStr, mth)
			assert.Equal(t, test.expectedBool, ok)
		})
	}
}

func TestBsDaysInMonthsByYear(t *testing.T) {
	tests := []struct {
		name         string
		year         int
		month        time.Month
		expected     int
		expectedBool bool
	}{
		{"when in range", bsLBound, time.Month(1), 30, true},
		{"when year not in range", bsUBound + 1, time.Month(1), 0, false},
		{"when year not in range", bsLBound - 1, time.Month(1), 0, false},
		{"when query month exceeds 12", bsUBound, time.Month(13), 0, false},
		{"when query month is less than 1", bsUBound, time.Month(-1), 0, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			days, ok := BsDaysInMonthsByYear(test.year, test.month)

			assert.Equal(t, test.expected, days)
			assert.Equal(t, test.expectedBool, ok)
		})
	}
}

func TestTotalDaysInBSYear(t *testing.T) {
	tests := []struct {
		name     string
		year     int
		expected int
	}{
		{"returns total number of days in 2001 BS", 2001, 365},
		{"returns total number of days in 2077 BS", 2077, 366},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			total, err := TotalDaysInBSYear(test.year)

			assert.Equal(t, test.expected, total, err)
		})
	}

	t.Run("errors if BS year is out of range", func(t *testing.T) {
		_, err := TotalDaysInBSYear(3050)

		if assert.Error(t, err) {
			assert.Equal(t, fmt.Errorf("Year should be in between %d and %d", bsLBound, bsUBound), err)
		}
	})
}
