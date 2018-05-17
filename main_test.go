package main

import (
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestShowDate(t *testing.T) {
	b := bytes.NewBuffer([]byte(""))
	stdout = b

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
