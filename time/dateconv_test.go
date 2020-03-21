package time

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// dummyNepaliTime returns a time between 00:00 to 05:45
// which when converted to UTC is the previous day
func dummyNepaliTime(yy int, mm int, dd int) time.Time {
	loc, _ := time.LoadLocation("Asia/Kathmandu")

	hour := rand.Intn(5)
	minute := rand.Intn(45)

	return time.Date(yy, time.Month(mm), dd, hour, minute, 0, 0, loc)
}

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
		{
			"case9",
			dummyNepaliTime(2019, 05, 05),
			NewBSDate(2076, 01, 22),
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

func TestToAD(t *testing.T) {
	tests := []struct {
		name   string
		input  BSDate
		output time.Time
	}{
		{
			"case1",
			NewBSDate(2074, 12, 18),
			toTime(2018, 04, 01),
		},
		{
			"case2",
			NewBSDate(2000, 01, 02),
			toTime(1943, 04, 15),
		},
		{
			"case3",
			NewBSDate(2075, 01, 04),
			toTime(2018, 04, 17),
		},
		{
			"case4",
			NewBSDate(2075, 01, 18),
			toTime(2018, 05, 01),
		},
		{
			"case5",
			NewBSDate(2017, 06, 1),
			toTime(1960, 9, 16),
		},
		{
			"case6",
			NewBSDate(2076, 02, 32),
			toTime(2019, 06, 15),
		},
		{
			"case7",
			NewBSDate(2076, 02, 30),
			toTime(2019, 06, 13),
		},
		{
			"case8",
			NewBSDate(2076, 01, 22),
			toTime(2019, 05, 05),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.output, ToAD(test.input))
		})
	}

	t.Run("panics if date is before 2000 Baisakh 1", func(t *testing.T) {
		assert.Panics(t, func() {
			ToAD(NewBSDate(1999, 04, 01)) // april 1
		}, "Can only work with dates after 2000 Baisakh 1")
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
		{"when in range", bsLBoundY, time.Month(1), 30, true},
		{"when year not in range", bsUBoundY + 1, time.Month(1), 0, false},
		{"when year not in range", bsLBoundY - 1, time.Month(1), 0, false},
		{"when query month exceeds 12", bsUBoundY, time.Month(13), 0, false},
		{"when query month is less than 1", bsUBoundY, time.Month(-1), 0, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			days, ok := BsDaysInMonthsByYear(test.year, test.month)

			assert.Equal(t, test.expected, days)
			assert.Equal(t, test.expectedBool, ok)
		})
	}
}

func TestBSAfter(t *testing.T) {
	tests := []struct {
		name     string
		dateA    BSDate
		dateB    BSDate
		expected bool
	}{
		{"date A after date B", NewBSDate(2050, 10, 10), NewBSDate(2020, 11, 11), true},
		{"date A before date B", NewBSDate(2020, 10, 10), NewBSDate(2050, 11, 11), false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			isAfter := test.dateA.After(test.dateB)

			assert.Equal(t, test.expected, isAfter)
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
			assert.Equal(t, fmt.Errorf("Year should be in between %d and %d", bsLBoundY, bsUBoundY), err)
		}
	})
}

func TestMonthStartsAtDay(t *testing.T) {
	var fixtures = map[string]time.Time{
		"May_17_2018":  time.Date(2018, time.May, 17, 0, 0, 0, 0, time.UTC),
		"May_19_2018":  time.Date(2018, time.May, 19, 0, 0, 0, 0, time.UTC),
		"May_26_2018":  time.Date(2018, time.May, 26, 0, 0, 0, 0, time.UTC),
		"June_15_2018": time.Date(2018, time.June, 15, 0, 0, 0, 0, time.UTC),
	}

	tests := []struct {
		name     string
		adDate   time.Time
		bsDate   BSDate
		expected int
	}{
		{
			"less than 7",
			fixtures["May_17_2018"],
			ToBS(fixtures["May_17_2018"]),
			2,
		},
		{
			"less than 7",
			fixtures["May_19_2018"],
			ToBS(fixtures["May_19_2018"]),
			2,
		},
		{
			"less than 7",
			fixtures["June_15_2018"],
			ToBS(fixtures["June_15_2018"]),
			5,
		},
		{
			"more than 7",
			fixtures["May_26_2018"],
			ToBS(fixtures["May_26_2018"]),
			2,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, test.bsDate.MonthStartsAtDay(test.adDate))
		})
	}
}

func TestTotalDaysSpannedUntilDate(t *testing.T) {
	tests := []struct {
		name     string
		date     time.Time
		expected int
	}{
		{"June_13_2018", toTime(2018, 8, 03), 112},
		{"June_13_2019", toTime(2019, 8, 03), 112},
		{"April_12_2020", toTime(2020, 04, 12), 365},
		{"April_13_2020", toTime(2020, 04, 13), 1},
		{"April_14_2020", toTime(2020, 04, 14), 2},
		{"April_13_2021", toTime(2021, 04, 13), 366},
		{"April_14_2021", toTime(2021, 04, 14), 1},
		{"April_15_2021", toTime(2021, 04, 15), 2},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			total, err := totalDaysSpannedUntilDate(test.date)

			assert.Equal(t, test.expected, total, err)
		})
	}

	t.Run("errors if BS year is out of range", func(t *testing.T) {
		date := toTime(3050, 06, 15)
		_, err := totalDaysSpannedUntilDate(date)

		if assert.Error(t, err) {
			assert.Equal(t, fmt.Errorf("Year should be in between %d and %d", bsLBoundY, bsUBoundY), err)
		}
	})
}

func TestCheckBoundsBS(t *testing.T) {
	tests := []struct {
		name     string
		date     BSDate
		expected error
	}{
		{"within bounds", NewBSDate(2053, 12, 30), nil},
		{"out of bounds", NewBSDate(1999, 1, 1), fmt.Errorf("Error: can only work with dates after 2000 Baisakh 1")},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.date.CheckBounds()

			assert.Equal(t, test.expected, err)
		})
	}
}

func TestCheckBoundsAD(t *testing.T) {
	tests := []struct {
		name     string
		date     time.Time
		expected error
	}{
		{"within bounds", toTime(2053, 12, 30), nil},
		{"out of bounds", toTime(1910, 1, 1), fmt.Errorf("Error: can only work with dates after 1943 April 14")},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := CheckBoundsAD(test.date)

			assert.Equal(t, test.expected, err)
		})
	}
}
