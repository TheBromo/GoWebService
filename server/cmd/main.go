package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"sync"

	pb "github.com/TheBromo/gochat/common/chat"
	"github.com/TheBromo/gochat/server/msg_distributor"

	"google.golang.org/grpc"
)

var (
	port   = flag.Int("port", 50051, "The server port")
	msgDis = msg_distributor.New()
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
	var wg sync.WaitGroup

	//handle input
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-msgserver.Context().Done():
				log.Fatalln(msgserver.Context().Err())
				return
			default:
				message, err := msgserver.Recv()
				if err != nil {
					msgDis.Distribute(message)
				}
			}
		}

	}()

	//hande msg distribution
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			inputC := make(chan []pb.Message)
			msgDis.RegisterConsumer(msgserver.Context(), inputC)
			defer msgDis.DeregisterConsumer(msgserver.Context())

			select {
			case <-msgserver.Context().Done():
				log.Fatalln(msgserver.Context().Err())
				return
			default:
				messages := <-inputC
				for i := 0; i < len(messages); i++ {
					msgserver.Send(&messages[i])
				}
			}
		}
	}()

	wg.Wait()
	msgDis.Close()
	return nil
}
