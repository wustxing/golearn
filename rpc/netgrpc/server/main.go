package main

import (
	"github.com/0990/golearn/rpc/netgrpc"
	"log"
	"net"
	"net/rpc"
)

type server struct {
}

func (s *server) SayHello(input *helloworld.HelloRequest, output *helloworld.HelloReply) error {
	output.Message = input.Name + "hi"
	return nil
}

func main() {

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("Accept　error:", err)
	}

	s := rpc.DefaultServer
	helloworld.RegisterGreeter(s, new(server))
	rpc.ServeConn(conn)
}
