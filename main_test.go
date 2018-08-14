package main

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestShowDate(t *testing.T) {
	b := bytes.NewBuffer([]byte(""))

	tests := []struct {
		name     string
		t        time.Time
		expected string
	}{
		{
			"case-1",
			time.Date(2018, time.May, 17, 0, 0, 0, 0, time.UTC),
			"जेठ 3, 2075 बिहिबार\n",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			showDate(b, test.t)
			assert.Equal(t, test.expected, b.String())
		})
	}
}

func TestParseRawDate(t *testing.T) {
	tests := []struct {
		name string
		raw  string
		yy   int
		mm   int
		dd   int
		ok   bool
	}{
		{"valid date", "08-21-1994", 1994, 8, 21, true},
		{"overflow day", "08-35-1994", -1, -1, -1, false},
		{"underflow day", "08-00-1994", -1, -1, -1, false},
		{"overflow month", "14-21-1994", -1, -1, -1, false},
		{"underflow month", "00-21-1994", -1, -1, -1, false},
		{"underflow year", "14-21-199", -1, -1, -1, false},
		{"overflow year", "14-21-19900", -1, -1, -1, false},
		{"inconversibe month", "aa-21-1994", -1, -1, -1, false},
		{"inconversibe day", "08-aa-1994", -1, -1, -1, false},
		{"inconversibe year", "08-21-xyz", -1, -1, -1, false},
		{"underflow number of elements", "08-21", -1, -1, -1, false},
		{"overflwo number of elements", "08-21-1994-01", -1, -1, -1, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mm, dd, yy, ok := parseRawDate(test.raw)

			assert.Equal(t, test.yy, yy)
			assert.Equal(t, test.mm, mm)
			assert.Equal(t, test.dd, dd)
			assert.Equal(t, test.ok, ok)
		})
	}
}

func TestRunCli(t *testing.T) {
	t.Run("shouldn't crash", func(t *testing.T) {
		assert.NotPanics(t, func() {
			runCli()
		})
	})
}
