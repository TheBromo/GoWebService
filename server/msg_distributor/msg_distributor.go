package msg_distributor

import (
	"context"
	"sync"

	pb "github.com/TheBromo/gochat/common/chat"
)

var (
	mu sync.Mutex
)

type MsgDistributor struct {
	msg_input chan []pb.Message
	consumers map[context.Context]chan []pb.Message
}

func (MsgDistributor) New() MsgDistributor {
	distributor := MsgDistributor{
		msg_input: make(chan []pb.Message),
		consumers: make(map[context.Context]chan []pb.Message),
	}

	go distributor.handleDistributions()

	return distributor
}

func (md *MsgDistributor) Close() {
	mu.Lock()
	close(md.msg_input)
	for _, v := range md.consumers {
		close(v)
	}
	mu.Unlock()
}

func (md *MsgDistributor) RegisterConsumer(ctx context.Context, consumer chan []pb.Message) {
	mu.Lock()
	md.consumers[ctx] = consumer
	mu.Unlock()
}

func (md *MsgDistributor) DeregisterConsumer(ctx context.Context) {
	mu.Lock()
	close(md.consumers[ctx])
	delete(md.consumers, ctx)
	mu.Unlock()
}

func (md *MsgDistributor) Distribute(messages []pb.Message) error {
	md.msg_input <- messages
	return nil
}

func (md *MsgDistributor) handleDistributions() {
	for {
		msg, open := <-md.msg_input

		if !open {
			return
		}

		mu.Lock()
		for _, v := range md.consumers {
			v <- msg
		}
		mu.Unlock()
	}
}
