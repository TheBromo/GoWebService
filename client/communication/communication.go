package communication

import (
	"context"
	"log"
	"sync"

	pb "github.com/TheBromo/gochat/common/chat"
	"google.golang.org/grpc"
)

var (
	running = true
)

func ConnectToServer(input chan pb.Message, output chan pb.Message, srvAddr string) {

	var opts []grpc.DialOption
	conn, err := grpc.Dial(srvAddr, opts...)
	if err != nil {
		log.Fatal("couldnt connect to server")
		return
	}
	defer conn.Close()

	client := pb.NewChatServiceClient(conn)

	var wg sync.WaitGroup

	srv, err := client.ExchangeMesssages(context.Background())

	if err != nil {
		log.Fatal("couldnt connect to server")
		return
	}

	wg.Add(1)
	go func() {
		send(input, srv)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		receive(input, srv)
		wg.Done()
	}()

	wg.Wait()
	srv.CloseSend()
}

// channel in
func send(input chan pb.Message, srv pb.ChatService_ExchangeMesssagesClient) {
	for running {
		msg, ok := <-input
		running = ok
		srv.Send(&msg)
	}
}

// channel out
func receive(output chan pb.Message, srv pb.ChatService_ExchangeMesssagesClient) {
	for running {
		msg, err := srv.Recv()
		if err != nil {
			running = false
		}
		output <- *msg
	}
}
