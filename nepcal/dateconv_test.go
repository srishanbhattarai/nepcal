package nepcal

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
	// TODO: Remove dummy time.Now()
	tests := []struct {
		name   string
		input  time.Time
		output Time
	}{
		{
			"case1",
			toTime(2018, 04, 01),
			Time{2074, 12, 18, time.Now()},
		},
		{
			"case2",
			toTime(1943, 04, 15),
			Time{2000, 01, 02, time.Now()},
		},
		{
			"case3",
			toTime(2018, 04, 17),
			Time{2075, 01, 04, time.Now()},
		},
		{
			"case4",
			toTime(2018, 05, 01),
			Time{2075, 01, 18, time.Now()},
		},
		{
			"case5",
			toTime(1960, 9, 16),
			Time{2017, 06, 1, time.Now()},
		},
		{
			"case6",
			toTime(2037, 9, 16),
			Time{-1, -1, -1, time.Now()},
		},
		{
			"case7",
			toTime(2019, 06, 15),
			Time{2076, 02, 32, time.Now()},
		},
		{
			"case8",
			toTime(2019, 06, 13),
			Time{2076, 02, 30, time.Now()},
		},
		{
			"case9",
			dummyNepaliTime(2019, 05, 05),
			Time{2076, 01, 22, time.Now()},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			bs := ToBS(test.input)

			// TODO: Add assertion for time.Time field.
			assert.Equal(t, test.output.year, bs.year)
			assert.Equal(t, test.output.month, bs.month)
			assert.Equal(t, test.output.day, bs.day)
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

func TestMonthString(t *testing.T) {
	tests := []struct {
		name        string
		month       Month
		expectedStr string
	}{
		{"when in range", Month(1), "बैशाख"},
		{"when not in range", Month(100), ""},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mth := test.month.String()

			assert.Equal(t, test.expectedStr, mth)
		})
	}
}

func TestWeekdayString(t *testing.T) {
	tests := []struct {
		name        string
		weekday     Weekday
		expectedStr string
	}{
		{"when in range", Weekday(0), "आइतबार"},
		{"when not in range", Weekday(7), ""},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			wd := test.weekday.String()

			assert.Equal(t, test.expectedStr, wd)
		})
	}
}

func TestBsDaysInMonthsByYear(t *testing.T) {
	tests := []struct {
		name         string
		year         int
		month        Month
		expected     int
		expectedBool bool
	}{
		{"when in range", bsLBound, Month(1), 30, true},
		{"when year not in range", bsUBound + 1, Month(1), 0, false},
		{"when year not in range", bsLBound - 1, Month(1), 0, false},
		{"when query month exceeds 12", bsUBound, Month(13), 0, false},
		{"when query month is less than 1", bsUBound, Month(-1), 0, false},
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

func TestStartingWeekdayOfMonth(t *testing.T) {
	var fixtures = map[string]time.Time{
		"May_17_2018":  time.Date(2018, time.May, 17, 0, 0, 0, 0, time.UTC),
		"May_19_2018":  time.Date(2018, time.May, 19, 0, 0, 0, 0, time.UTC),
		"May_26_2018":  time.Date(2018, time.May, 26, 0, 0, 0, 0, time.UTC),
		"June_15_2018": time.Date(2018, time.June, 15, 0, 0, 0, 0, time.UTC),
	}

	tests := []struct {
		name     string
		bsDate   Time
		expected int
	}{
		{
			"less than 7",
			ToBS(fixtures["May_17_2018"]),
			2,
		},
		{
			"less than 7",
			ToBS(fixtures["May_19_2018"]),
			2,
		},
		{
			"less than 7",
			ToBS(fixtures["June_15_2018"]),
			5,
		},
		{
			"more than 7",
			ToBS(fixtures["May_26_2018"]),
			2,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, int(test.bsDate.StartingWeekdayOfMonth()))
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
			bsDate := ToBS(test.date)
			total := bsDate.TotalDaysSpanned()

			assert.Equal(t, test.expected, total)
		})
	}
}
