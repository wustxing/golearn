package main

import (
	"fmt"
	"github.com/hashicorp/yamux"
	"github.com/sirupsen/logrus"
	"net"
)

func Recv(stream net.Conn, id int) {
	fmt.Println("id", id)
	for {
		buf := make([]byte, 10)
		n, err := stream.Read(buf)
		if err != nil {
			logrus.Println(id, err)
			return
		}
		logrus.Println("ID:", id, ",len:", n, string(buf))
	}
}

func main() {
	addr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:8980")
	if err != nil {
		panic(err)
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		panic(err)
	}

	conn, err := l.Accept()
	if err != nil {
		panic(err)
	}

	session, err := yamux.Server(conn, nil)
	if err != nil {
		panic(err)
	}

	id := 1
	for {
		stream, err := session.Accept()
		if err != nil {
			fmt.Println("session over")
			return
		}

		id++
		go Recv(stream, id)
	}
}
