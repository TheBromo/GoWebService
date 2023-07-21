package main

import (
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/TheBromo/goWebService/common/chat"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)


type server struct {
	pb.ChatServiceClient
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
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, )
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
