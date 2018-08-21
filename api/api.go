package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	proto "github.com/srishanbhattarai/nepcal/api/proto"
)

type converter interface {
	ToBS(time.Time) time.Time
	GetBSMonthName(month time.Month) (monthName string, ok bool)
	GetNepWeekday(time.Weekday) (weekday string, ok bool)
}

type api struct {
	converter converter
}

func (api api) convertDate(ctx context.Context, t time.Time) (string, error) {
	yy, mm, dd := t.Date()

	bsyy, bsmm, bsdd := api.converter.ToBS(time.Date(yy, mm, dd, 0, 0, 0, 0, time.UTC)).Date()
	month, monthOk := api.converter.GetBSMonthName(bsmm)
	weekday, weekdayOk := api.converter.GetNepWeekday(t.Weekday())

	if monthOk && weekdayOk {
		date := fmt.Sprintf("%s %d, %d %s\n", month, bsdd, bsyy, weekday)

		fmt.Printf("Returning: %s\n", date)

		return date, nil
	}

	return "", errors.New("Couldn't convert date")
}

func (api api) TodaysDate(ctx context.Context, _ *proto.Void) (*proto.TodaysDateRes, error) {
	t := time.Now()

	date, err := api.convertDate(ctx, t)
	if err != nil {
		return nil, err
	}

	return &proto.TodaysDateRes{
		Date: date,
	}, nil
}

func (api api) ConvADToBS(ctx context.Context, in *proto.ConvADToBSReq) (*proto.ConvADToBSRes, error) {
	date, err := api.convertDate(ctx, time.Date(int(in.GetYy()), time.Month(in.GetMm()), int(in.GetDd()), 0, 0, 0, 0, time.UTC))
	if err != nil {
		return nil, err
	}

	return &proto.ConvADToBSRes{
		Date: date,
	}, nil
}
