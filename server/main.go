package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/TheBromo/goWebService/common/chat"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

var (
	port    = flag.Int("port", 50051, "The server port")
	clients = make(map[net.Addr]string)
)

type server struct {
	pb.UnimplementedChatServiceServer
}

func (c *server) RegisterRegister(context context.Context, message *pb.Login) (*emptypb.Empty, error) {
	username := message.GetUsername()

	p, ok := peer.FromContext(context)

	if ok {
		clients[p.Addr] = username
		log.Printf("new user registered : %s, address %s \n", username, p.Addr.String())
		return &emptypb.Empty{}, nil
	} else {
		log.Println("can't retrieve user address")
		return nil, errors.New("can't retrieve user address")
	}
}

func (c *server) Unregister(context context.Context, empty *emptypb.Empty) (*emptypb.Empty, error) {
	p, _ := peer.FromContext(context)
	_, ok := clients[p.Addr]
	if ok {
		delete(clients, p.Addr)
		return &emptypb.Empty{}, nil
	} else {
		log.Println("user not registered")
		return nil, errors.New("user not registered")
	}
}

func (c *server) HandleMessage(context context.Context, message *pb.Message) (*emptypb.Empty, error) {
	for k := range clients {
		log.Println(k)
	}
	return nil, nil
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
