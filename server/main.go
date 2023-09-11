package main

import (
	"flag"
	"fmt"
	"io"
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

func (c *server) SendMessage(stream pb.ChatService_SendMessageServer) error {
	go func() {
		for {
			resp, err := stream.Recv()
			mesages <- *resp
			if err == io.EOF {
				return
			}
			if err != nil {
				log.Fatalf("cannot receive %v", err)
				return

			}
		}
	}()
	return nil
}

func (c *server) PollMesssages(username *pb.Username, msgserver pb.ChatService_ReceiveMesssagesServer) error {
	log.Printf("New user joined: %s \n", username.GetUsername())
	
	/*for true {
		if msg.Timestamp.GetNanos() > timestamp.GetNanos() {
			if err := msgserver.Send(&msg); err != nil {
				log.Printf("send error %v", err)
			}
		}
	}
	*/
	msgserver.Context().Done()
	return nil
}
