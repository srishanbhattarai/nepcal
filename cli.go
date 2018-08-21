package main

import (
	"context"
	"fmt"
	"io"
	"time"

	proto "github.com/srishanbhattarai/nepcal/api/proto"
	"github.com/urfave/cli"
	"google.golang.org/grpc"
)

// nepcalCli is a struct to hold all the CLI behavior.
// It is a small wrapper around the urfave/cli package to keep things clean.
type nepcalCli struct {
	// TODO(srishan): Replace this with an interface
	service proto.NepcalClient
}

func newCli() nepcalCli {
	conn, err := grpc.Dial(":9999", grpc.WithInsecure())
	if err != nil {
		panic("Couldn't get grpc server: " + err.Error())
	}

	c := proto.NewNepcalClient(conn)

	return nepcalCli{service: c}
}

// Shows the calendar for the current day.
func (nepcalCli) showCalendar(c *cli.Context) {
	cal := newCalendar()
	cal.Render(writer, time.Now())
}

// Shows the date for the provided time. Returns a cli 'action'.
func (nepcalCli) showDate(w io.Writer, t time.Time) func(c *cli.Context) {
	return func(c *cli.Context) {
		showDate(w, t)
	}
}

// Convert AD date to BS date after validation.
func (nc nepcalCli) convADToBS(c *cli.Context) {
	areArgsValid := func() bool {
		if c.NArg() < 1 {
			return false
		}

		_, _, _, ok := parseRawDate(c.Args().First())
		if !ok {
			return false
		}

		return true
	}()

	if !areArgsValid {
		fmt.Println("Please supply a valid date in the format mm-dd-yyyy. Example: `nepcal conv adtobs 08-21-1994`")
		return
	}

	mm, dd, yy, _ := parseRawDate(c.Args().First())

	date, err := nc.service.ConvADToBS(context.Background(), &proto.ConvADToBSReq{
		Dd: int32(dd),
		Mm: int32(mm),
		Yy: int32(yy),
	})

	if err != nil {
		fmt.Printf("Couldn't convert: %s", err.Error())
		return
	}

	fmt.Printf("%s", date.Date)
}

// Not supported yet.
func (nepcalCli) convBSToAD(c *cli.Context) {
	fmt.Fprintln(writer, "Unfortunately BS to AD isn't supported at this time. :(")
}
