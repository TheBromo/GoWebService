package common

import "net"

type TextMessage struct{

}

func (e TextMessage) Sender () net.Addr {
	return &net.IPAddr{}
}

func (e TextMessage) Receiver () net.Addr {
	return &net.IPAddr{}
}