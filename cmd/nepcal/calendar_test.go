package main

import (
	"bytes"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var fixtures = map[string]time.Time{
	"May17":  time.Date(2018, time.May, 17, 0, 0, 0, 0, time.UTC),
	"May19":  time.Date(2018, time.May, 19, 0, 0, 0, 0, time.UTC),
	"May26":  time.Date(2018, time.May, 26, 0, 0, 0, 0, time.UTC),
	"June15": time.Date(2018, time.June, 15, 0, 0, 0, 0, time.UTC),
	"Feb13_2020": time.Date(2020, time.February, 13, 0, 0, 0, 0, time.UTC),
	"Mar21_2020": time.Date(2020, time.March, 21, 0, 0, 0, 0, time.UTC),
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
		{
			"Mar21_2020",
			fixtures["Mar21_2020"],
			`
			चैत 8, 2076
 Su Mo Tu We Th Fr Sa 
                   1
 2  3  4  5  6  7  8
 9  10 11 12 13 14 15
 16 17 18 19 20 21 22
 23 24 25 26 27 28 29
 30
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
			c := newCalendar(b)
			c.Render(test.t)
			assert.Equal(t, clean(test.expected), clean(b.String()))
		})
	}
}
