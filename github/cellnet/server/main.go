package main

import (
	"fmt"
	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/peer"
	_ "github.com/davyxu/cellnet/peer/tcp" // 注册TCP Peer
	"github.com/davyxu/cellnet/proc"
	_ "github.com/davyxu/cellnet/proc/tcp" // 注册TCP Processor
	"github.com/davyxu/cellnet/tests"
	"os"
	"os/signal"
)

const peerAddress = "127.0.0.1:17701"

func main() {

	queue := cellnet.NewEventQueue()

	perrIns := peer.NewGenericPeer("tcp.Acceptor", "server", peerAddress, queue)

	proc.BindProcessorHandler(perrIns, "tcp.ltv", func(ev cellnet.Event) {
		switch msg := ev.Message().(type) {
		case *cellnet.SessionAccepted:
			fmt.Println("sever accepted")
		case *tests.TestEchoACK:
			fmt.Printf("server recv %+v\n", msg)
			ev.Session().Send(&tests.TestEchoACK{
				Msg:   msg.Msg,
				Value: msg.Value,
			})
		case *cellnet.SessionClosed:
			fmt.Println("session closed:", ev.Session().ID())
		}
	})
	perrIns.Start()
	queue.StartLoop()
	fmt.Println("start success")
	c := make(chan os.Signal, 1)
	signal.Notify(c)
	s := <-c
	fmt.Println("Got signal:", s)
}
