package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/TheBromo/goWebService/common/chat"
	"google.golang.org/grpc"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedTerminalAppServiceServer
}

func (c *server) HandleMessage(e pb.TerminalAppService_HandleMessageServer) error {
	return nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTerminalAppServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
