package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/TheBromo/gochat/common/chat"

	"google.golang.org/grpc"
)

var (
	port                    = flag.Int("port", 50051, "The server port")
	mesages chan pb.Message = make(chan pb.Message)
)

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterChatServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type server struct {
	pb.UnimplementedChatServiceServer
}

func (c *server) PollMesssages(msgserver pb.ChatService_ExchangeMesssagesServer) error {

	//handle input
	go func() {

	}()

	//hande msg distribution
	go func() {

	}()

	msgserver.Context().Done()
	return nil
}
