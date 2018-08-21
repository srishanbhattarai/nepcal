package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/srishanbhattarai/nepcal/dateconv"

	proto "github.com/srishanbhattarai/nepcal/api/proto"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 9999, "Port to start the server on")
)

func init() {
	flag.Parse()
}

// TODO: Change the package so that it returns the grpc server rather than being a main package itself.
func main() {
	server := grpc.NewServer()

	api := api{
		converter: dateconv.Converter{},
	}
	proto.RegisterNepcalServer(server, api)

	address := fmt.Sprintf(":%d", *port)

	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Could not start listen on port: %s, %s", address, err.Error())
	}

	if err := server.Serve(lis); err == nil {
		log.Println("Listening on port: ", address)
	}
}
