package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/TheBromo/goWebService/common/chat"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedChatServiceServer
}

func (c *server) RegisterRegister(context.Context, *pb.Login) (*pb.Login, error) {
	return nil, nil
}

func (c *server) Unregister(context.Context, *pb.Logout) (*pb.Logout, error) {
	return nil, nil
}

func (c *server) HandleMessage(e pb.ChatService_HandleMessageServer) error {
	return nil
}

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
