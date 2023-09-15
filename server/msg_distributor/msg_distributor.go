package msg_distributor

import (
	"errors"

	pb "github.com/TheBromo/gochat/common/chat"
)

type MsgDistributor struct {
	msg_input chan []pb.Message
	consumers []chan []pb.Message
}

func (MsgDistributor) New() MsgDistributor {
	//TODO init goroutine

	return MsgDistributor{
		msg_input: make(chan []pb.Message),
		consumers: []chan []pb.Message{},
	}
}

func (md *MsgDistributor) Distribute(messages []pb.Message) error {
	if len(md.consumers) == 0 {
		return errors.New("no consumers registered")
	}
	md.msg_input <- messages
	return nil
}
