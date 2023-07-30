package main

import (
	"flag"
	"log"

	receiver "github.com/TheBromo/goWebService/client/message"
)

var (
	//addr = flag.String("addr", "localhost:50051", "the address to connect to")
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	flag.Parse()

	log.Println("starting client")

	receiver.Start(*port)

}
