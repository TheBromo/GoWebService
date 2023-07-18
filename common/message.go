package common

import "net"

type Message interface {
	Sender() net.Addr
	Receiver() net.Addr
}
