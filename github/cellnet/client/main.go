package main

import (
	"fmt"
	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/peer"
	_ "github.com/davyxu/cellnet/peer/tcp" // 注册TCP Peer
	"github.com/davyxu/cellnet/proc"
	_ "github.com/davyxu/cellnet/proc/tcp" // 注册TCP Processor
	"github.com/davyxu/cellnet/tests"
)

const peerAddress = "127.0.0.1:17701"

func main() {
	done := make(chan struct{})

	queue := cellnet.NewEventQueue()
	peerIns := peer.NewGenericPeer("tcp.Connector", "client", peerAddress, queue)

	proc.BindProcessorHandler(peerIns, "tcp.ltv", func(ev cellnet.Event) {
		switch msg := ev.Message().(type) {
		case *cellnet.SessionConnected:
			fmt.Println("client connected")
			ev.Session().Send(&tests.TestEchoACK{
				Msg:   "hello",
				Value: 1234,
			})
		case *tests.TestEchoACK:
			fmt.Printf("client recv %+v\n", msg)
			done <- struct{}{}
		case *cellnet.SessionClosed:
			fmt.Println("client closed")
		}
	})
	peerIns.Start()
	queue.StartLoop()
	<-done
}
