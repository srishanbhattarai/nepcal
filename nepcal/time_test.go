package nepcal

import (
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

// effectively checks Ad->Bs tests.
func TestFromGregorian(t *testing.T) {
	tests := []struct {
		name  string
		input time.Time
		bsy   int
		bsm   Month
		bsd   int
	}{
		{
			"case1",
			gregorian(2018, 04, 01),
			2074, Chaitra, 18,
		},
		{
			"case2",
			gregorian(1943, 04, 15),
			2000, Baisakh, 02,
		},
		{
			"case3",
			gregorian(2018, 04, 17),
			2075, Baisakh, 04,
		},
		{
			"case4",
			gregorian(2018, 05, 01),
			2075, Baisakh, 18,
		},
		{
			"case5",
			gregorian(1960, 9, 16),
			2017, Ashoj, 1,
		},
		{
			"case6",
			gregorian(2039, 9, 16),
			-1, -1, -1,
		},
		{
			"case7",
			gregorian(2019, 06, 15),
			2076, Jestha, 32,
		},
		{
			"case8",
			gregorian(2019, 06, 13),
			2076, Jestha, 30,
		},
		{
			"case9",
			dummyNepaliTime(2019, 05, 05),
			2076, Baisakh, 22,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			bs, err := FromGregorian(test.input)

			assert.NoError(t, err)

			yy, mm, dd := bs.Date()
			assert.Equal(t, test.bsy, yy)
			assert.Equal(t, test.bsm, mm)
			assert.Equal(t, test.bsd, dd)
			assert.Equal(t, test.input, bs.in)
		})
	}

	t.Run("panics if date is before 1943 April 14", func(t *testing.T) {
		_, err := FromGregorian(gregorian(1943, 04, 01))
		assert.Equal(t, err, ErrOutOfBounds)
	})
}

func TestBsAdConversion(t *testing.T) {
	tests := []struct {
		name   string
		input  raw
		output time.Time
	}{
		{
			"case1",
			raw{2074, 12, 18},
			gregorian(2018, 04, 01),
		},
		{
			"case2",
			raw{2000, 01, 02},
			gregorian(1943, 04, 15),
		},
		{
			"case3",
			raw{2075, 01, 04},
			gregorian(2018, 04, 17),
		},
		{
			"case4",
			raw{2075, 01, 18},
			gregorian(2018, 05, 01),
		},
		{
			"case5",
			raw{2017, 06, 1},
			gregorian(1960, 9, 16),
		},
		{
			"case6",
			raw{2076, 02, 32},
			gregorian(2019, 06, 15),
		},
		{
			"case7",
			raw{2076, 02, 30},
			gregorian(2019, 06, 13),
		},
		{
			"case8",
			raw{2076, 01, 22},
			gregorian(2019, 05, 05),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			bs, err := Date(test.input.year, test.input.month, test.input.day)

			assert.NoError(t, err)
			assert.Equal(t, test.output, bs.in)
		})
	}

	t.Run("panics if date is before 2000 Baisakh 1", func(t *testing.T) {
		_, err := Date(1999, 04, 01)
		assert.Equal(t, err, ErrOutOfBounds)
	})
}

var fixtures = map[string]time.Time{
	"May_17_2018":  time.Date(2018, time.May, 17, 0, 0, 0, 0, time.UTC),
	"May_19_2018":  time.Date(2018, time.May, 19, 0, 0, 0, 0, time.UTC),
	"May_26_2018":  time.Date(2018, time.May, 26, 0, 0, 0, 0, time.UTC),
	"June_15_2018": time.Date(2018, time.June, 15, 0, 0, 0, 0, time.UTC),
}

func TestNumDays(t *testing.T) {
	// This is also tested through the tests on 'Month' in 'parts_test.go'
	bs, err := Date(2076, Baisakh, 8)
	assert.NoError(t, err)
	assert.Equal(t, 31, bs.NumDaysInMonth())

	assert.Equal(t, 365, bs.NumDaysInYear())
}

func TestStartWeekday(t *testing.T) {
	tests := []struct {
		name     string
		adDate   time.Time
		bsDate   Time
		expected Weekday
	}{
		{
			"less than 7",
			fixtures["May_17_2018"],
			FromGregorianUnchecked(fixtures["May_17_2018"]),
			Tuesday,
		},
		{
			"less than 7",
			fixtures["May_19_2018"],
			FromGregorianUnchecked(fixtures["May_19_2018"]),
			Tuesday,
		},
		{
			"less than 7",
			fixtures["June_15_2018"],
			FromGregorianUnchecked(fixtures["June_15_2018"]),
			Friday,
		},
		{
			"more than 7",
			fixtures["May_26_2018"],
			FromGregorianUnchecked(fixtures["May_26_2018"]),
			Tuesday,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, test.bsDate.StartWeekday())
		})
	}
}

func TestTotalDaysSpanned(t *testing.T) {
	tests := []struct {
		name     string
		date     time.Time
		expected int
	}{
		{"June_13_2018", gregorian(2018, 8, 03), 112},
		{"June_13_2019", gregorian(2019, 8, 03), 112},
		{"April_12_2020", gregorian(2020, 04, 12), 365},
		{"April_13_2020", gregorian(2020, 04, 13), 1},
		{"April_14_2020", gregorian(2020, 04, 14), 2},
		{"April_13_2021", gregorian(2021, 04, 13), 366},
		{"April_14_2021", gregorian(2021, 04, 14), 1},
		{"April_15_2021", gregorian(2021, 04, 15), 2},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			bs := FromGregorianUnchecked(test.date)
			total := bs.NumDaysSpanned()

			assert.Equal(t, test.expected, total)
		})
	}
}

func TestWeekday(t *testing.T) {
	tests := []struct {
		name     string
		date     Time
		expected Weekday
	}{
		{"Chaitra 9 2076", DateUnchecked(2076, Chaitra, 9), Sunday},
		{"Mangshir 29 2067", DateUnchecked(2067, Mangshir, 29), Wednesday},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, test.date.Weekday())
		})
	}
}

func TestAfter(t *testing.T) {
	t1 := DateUnchecked(2076, Chaitra, 9)
	t2 := DateUnchecked(2067, Mangshir, 29)

	assert.Equal(t, true, t1.After(t2))
	assert.Equal(t, false, t2.After(t1))
}
