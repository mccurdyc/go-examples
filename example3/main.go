package main

import (
	"flag"
	"log"

	"github.com/mccurdyc/go-examples/example3/transports/http/server"
)

var (
	serverHost = flag.String("host", "localhost", "server host")
	serverPort = flag.Int("port", 8080, "server port")
)

func init() {
	flag.Parse()
}

func main() {
	s := server.NewService(*serverHost, *serverPort)
	log.Printf("started server on %d\n", *serverPort)
	s.Start()
}
