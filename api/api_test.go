package main

import (
	"context"
	"testing"
	"time"

	proto "github.com/srishanbhattarai/nepcal/api/proto"
	"github.com/stretchr/testify/assert"
)

type mockConverter struct {
	bsTime      time.Time
	bsMonthName struct {
		monthName string
		ok        bool
	}
	nepWeekday struct {
		weekday string
		ok      bool
	}
}

func (m mockConverter) ToBS(_ time.Time) time.Time {
	return m.bsTime
}

func (m mockConverter) GetBSMonthName(_ time.Month) (monthName string, ok bool) {
	return m.bsMonthName.monthName, m.bsMonthName.ok
}

func (m mockConverter) GetNepWeekday(_ time.Weekday) (monthName string, ok bool) {
	return m.nepWeekday.weekday, m.nepWeekday.ok
}

func TestTodaysDate(t *testing.T) {
	req := &proto.Void{}

	bsTime := time.Date(2018, 8, 21, 0, 0, 0, 0, time.UTC)
	conv := mockConverter{
		bsTime: bsTime,
		bsMonthName: struct {
			monthName string
			ok        bool
		}{
			monthName: "nepMonth",
			ok:        true,
		},
		nepWeekday: struct {
			weekday string
			ok      bool
		}{
			weekday: "nepDay",
			ok:      true,
		},
	}
	api := api{converter: conv}

	res, err := api.TodaysDate(context.Background(), req)

	if err != nil {
		t.Fatalf("Expected error to be nil")
	}

	if res == nil {
		t.Fatalf("Expected res to not be nil")
	}

	expectedDate := "nepMonth 21, 2018 nepDay\n"
	assert.Equal(t, expectedDate, res.Date)
}
