package conversion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToBS(t *testing.T) {
	tests := []struct {
		name   string
		input  epoch
		output epoch
	}{
		{
			"case1",
			epoch{2018, 04, 01},
			epoch{2074, 12, 18},
		},
		{
			"case2",
			epoch{1943, 04, 15},
			epoch{2000, 01, 02},
		},
		{
			"case3",
			epoch{2018, 04, 17},
			epoch{2075, 01, 04},
		},
		{
			"case4",
			epoch{2018, 05, 01},
			epoch{2075, 01, 18},
		},
		{
			"case5",
			epoch{1960, 9, 16},
			epoch{2017, 06, 1},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.output, toBS(test.input))
		})
	}
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
