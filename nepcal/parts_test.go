package nepcal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMonth(t *testing.T) {
	// Name(), and Stringer implementations.
	b := Baisakh
	assert.Equal(t, b.Name(), "बैशाख")
	assert.Equal(t, b.String(), "बैशाख")

	// NumDays test
	b = Jestha
	nd, err := b.NumDays(2009)
	assert.NoError(t, err)
	assert.Equal(t, nd, 31)

	nd, err = b.NumDays(2035)
	assert.NoError(t, err)
	assert.Equal(t, nd, 32)

	_, err = b.NumDays(2096)
	assert.Equal(t, err, ErrOutOfBounds)
}

func TestWeekdayStr(t *testing.T) {
	// Name(), and Stringer implementations.
	b := Sunday
	assert.Equal(t, b.Name(), "आइतबार")
	assert.Equal(t, b.String(), "आइतबार")
}

func TestNumeral(t *testing.T) {
	tt := []struct {
		name     string
		num      int
		expected string
	}{
		{
			"single digit",
			1,
			"१",
		},
		{
			"multiple digits",
			759,
			"७५९",
		},
	}

	for _, test := range tt {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, Numeral(test.num).String())
		})
	}
}
