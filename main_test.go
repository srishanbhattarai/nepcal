package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestShowDate(t *testing.T) {
	b := bytes.NewBuffer([]byte(""))
	stdout = b
	defer func() {
		stdout = os.Stdout
	}()

	tests := []struct {
		name     string
		t        time.Time
		expected string
	}{
		{
			"case-1",
			time.Date(2018, time.May, 17, 0, 0, 0, 0, time.UTC),
			"जेठ 3, 2075\n",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			showDate(test.t)
			assert.Equal(t, test.expected, b.String())
		})
	}
}

func TestShowCal(t *testing.T) {
	b := bytes.NewBuffer([]byte(""))
	stdout = b
	defer func() {
		stdout = os.Stdout
	}()

	tests := []struct {
		name     string
		t        time.Time
		expected string
	}{
		{
			"case-1",
			time.Date(2018, time.May, 17, 0, 0, 0, 0, time.UTC),
			`
			जेठ 3, 2075
 Su Mo Tu We Th Fr Sa 
       1  2  3  4  5
 6  7  8  9  10 11 12
 13 14 15 16 17 18 19
 20 21 22 23 24 25 26
 27 28 29 30 31 32
		`,
		},
	}

	clean := func(op string) string {
		woNewLines := strings.Trim(op, "\n")
		woSpaces := strings.TrimSpace(woNewLines)

		return woSpaces
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			showCal(test.t)
			assert.Equal(t, clean(test.expected), clean(b.String()))
		})
	}
}
