package communication

import (
	"context"
	"log/slog"
	"sync"

	pb "github.com/TheBromo/gochat/common/chat"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	running = true
)

func ConnectToServer(input chan pb.Message, output chan pb.Message, srvAddr string) {

	conn, err := grpc.Dial(srvAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		slog.Error("couldnt connect to server %v", err)
		return
	}
	defer conn.Close()

	client := pb.NewChatServiceClient(conn)

	var wg sync.WaitGroup

	srv, err := client.ExchangeMesssages(context.Background())

	if err != nil {
		slog.Error("couldnt connect to server")
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
		if msg != nil {
			output <- *msg // Only send if msg is not nil
		}
	}
}
