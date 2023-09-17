package msg_distributor

import (
	"context"
	"sync"
	"time"

	pb "github.com/TheBromo/gochat/common/chat"
)

var (
	mu sync.Mutex
)

type IMessageDistributor interface {
	Close()
	DeregisterConsumer(ctx context.Context)
	Distribute(messages pb.Message) error
	RegisterConsumer(ctx context.Context, consumer chan []pb.Message)
	handleDistributions()
}

type messageDistributorImpl struct {
	msgInput  chan pb.Message
	consumers map[context.Context]chan []pb.Message
}

func New() *messageDistributorImpl {
	distributor := messageDistributorImpl{
		msgInput:  make(chan pb.Message),
		consumers: make(map[context.Context]chan []pb.Message),
	}

	go distributor.handleDistributions()

	return &distributor
}

func (md *messageDistributorImpl) Close() {
	mu.Lock()
	close(md.msgInput)
	for _, v := range md.consumers {
		close(v)
	}
	mu.Unlock()
}

func (md *messageDistributorImpl) RegisterConsumer(ctx context.Context, consumer chan []pb.Message) {
	mu.Lock()
	md.consumers[ctx] = consumer
	mu.Unlock()
}

func (md *messageDistributorImpl) DeregisterConsumer(ctx context.Context) {
	mu.Lock()
	close(md.consumers[ctx])
	delete(md.consumers, ctx)
	mu.Unlock()
}

func (md *messageDistributorImpl) Distribute(messages *pb.Message) error {
	md.msgInput <- *messages
	return nil
}

func (md *messageDistributorImpl) handleDistributions() {
	ticker := time.NewTicker(100 * time.Millisecond)
	messages := make([]pb.Message, 0)

	for {

		select {
		case <-ticker.C:
			mu.Lock()
			for _, v := range md.consumers {
				v <- messages
			}
			messages = make([]pb.Message, 0)
			mu.Unlock()

		default:
			msg, open := <-md.msgInput

			if !open {
				ticker.Stop()
				return
			}
			messages = append(messages, msg)
		}

	}
}