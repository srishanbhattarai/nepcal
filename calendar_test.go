package main

import (
	"bytes"
	"strings"
	"testing"
	"time"

	"github.com/srishanbhattarai/nepcal/dateconv"
	"github.com/stretchr/testify/assert"
)

var fixtures = map[string]time.Time{
	"May17":  time.Date(2018, time.May, 17, 0, 0, 0, 0, time.UTC),
	"May19":  time.Date(2018, time.May, 19, 0, 0, 0, 0, time.UTC),
	"May26":  time.Date(2018, time.May, 26, 0, 0, 0, 0, time.UTC),
	"June15": time.Date(2018, time.June, 15, 0, 0, 0, 0, time.UTC),
}

// Test the 'calculateSkew' function.
func TestCalculateSkew(t *testing.T) {
	tests := []struct {
		name     string
		adDate   time.Time
		bsDate   time.Time
		expected int
	}{
		{
			"less than 7",
			fixtures["May17"],
			dateconv.ToBS(fixtures["May17"]),
			2,
		},
		{
			"less than 7",
			fixtures["May19"],
			dateconv.ToBS(fixtures["May19"]),
			2,
		},
		{
			"less than 7",
			fixtures["June15"],
			dateconv.ToBS(fixtures["June15"]),
			5,
		},
		{
			"more than 7",
			fixtures["May26"],
			dateconv.ToBS(fixtures["May26"]),
			2,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := newCalendar()
			assert.Equal(t, test.expected, c.calculateSkew(test.adDate, test.bsDate))
		})
	}
}

// Test the 'renderCalendar' function
func TestRenderCalendar(t *testing.T) {
	tests := []struct {
		name     string
		t        time.Time
		expected string
	}{
		{
			"May",
			fixtures["May17"],
			`
			जेठ 3, 2075
 Su Mo Tu We Th Fr Sa 
       1  2  3  4  5
 6  7  8  9  10 11 12
 13 14 15 16 17 18 19
 20 21 22 23 24 25 26
 27 28 29 30 31
		`,
		},
	}

	clean := func(op string) string {
		woNewLines := strings.Trim(op, "\n")
		woSpaces := strings.TrimSpace(woNewLines)

		return woSpaces
	}

	for _, test := range tests {
		b := bytes.NewBuffer([]byte(""))

		t.Run(test.name, func(t *testing.T) {
			b.Reset()
			c := newCalendar()
			c.Render(b, test.t)
			assert.Equal(t, clean(test.expected), clean(b.String()))
		})
	}
}
