package main

import (
	"flag"
	"log"

	message "github.com/TheBromo/gochat/client/message"
)

var (
	//addr = flag.String("addr", "localhost:50051", "the address to connect to")
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	flag.Parse()

	log.Println("starting client")

	message.StartReceiver(*port)
}
