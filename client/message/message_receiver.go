package message

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/TheBromo/gochat/common/chat"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	pb.UnimplementedTerminalAppServiceServer
}

func (c *Server) HandleMessage(context context.Context, message *pb.Message) (*emptypb.Empty, error) {
	return nil, nil
}

func StartReceiver(port int) {
	log.Println("starting receiver")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return
	}

	s := grpc.NewServer()
	pb.RegisterTerminalAppServiceServer(s, &Server{})
	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
		return
	}
}
