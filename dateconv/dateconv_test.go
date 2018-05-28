package dateconv

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestToBS(t *testing.T) {
	tests := []struct {
		name   string
		input  time.Time
		output time.Time
	}{
		{
			"case1",
			toTime(2018, 04, 01),
			toTime(2074, 12, 18),
		},
		{
			"case2",
			toTime(1943, 04, 15),
			toTime(2000, 01, 02),
		},
		{
			"case3",
			toTime(2018, 04, 17),
			toTime(2075, 01, 04),
		},
		{
			"case4",
			toTime(2018, 05, 01),
			toTime(2075, 01, 18),
		},
		{
			"case5",
			toTime(1960, 9, 16),
			toTime(2017, 06, 1),
		},
		{
			"case6",
			toTime(2037, 9, 16),
			toTime(-1, -1, -1),
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
