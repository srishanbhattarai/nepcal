package main

import (
	"bytes"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var fixtures = map[string]time.Time{
	"May17":      time.Date(2018, time.May, 17, 0, 0, 0, 0, time.UTC),
	"May19":      time.Date(2018, time.May, 19, 0, 0, 0, 0, time.UTC),
	"May26":      time.Date(2018, time.May, 26, 0, 0, 0, 0, time.UTC),
	"June15":     time.Date(2018, time.June, 15, 0, 0, 0, 0, time.UTC),
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
			जेठ ३, २०७५
 Su Mo Tu We Th Fr Sa 
       १  २  ३  ४  ५
 ६  ७  ८  ९  १० ११ १२
 १३ १४ १५ १६ १७ १८ १९
 २० २१ २२ २३ २४ २५ २६
 २७ २८ २९ ३० ३१
		`,
		},
		{
			"Mar21_2020",
			fixtures["Mar21_2020"],
			`
			चैत ८, २०७६
 Su Mo Tu We Th Fr Sa 
                   १
 २  ३  ४  ५  ६  ७  ८
 ९  १० ११ १२ १३ १४ १५
 १६ १७ १८ १९ २० २१ २२
 २३ २४ २५ २६ २७ २८ २९
 ३०
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
