package message

import (
	"fmt"
	"log"
	"net"

	pb "github.com/TheBromo/goWebService/common/chat"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedTerminalAppServiceServer
}

func (c *Server) HandleMessage(e pb.TerminalAppService_HandleMessageServer) error {
	return nil
}

func Start(port int) {
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
