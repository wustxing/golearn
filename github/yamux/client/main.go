package main

import (
	"github.com/hashicorp/yamux"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8980")
	if err != nil {
		panic(err)
	}
	s, err := yamux.Client(conn, nil)
	if err != nil {
		panic(err)
	}

	s1, err := s.Open()
	if err != nil {
		panic(err)
	}

	s1.Write([]byte("hes1"))
	time.Sleep(time.Second)

	s2, err := s.Open()
	if err != nil {
		panic(err)
	}

	s2.Write([]byte("hes2"))
	time.Sleep(time.Second)

	s1.Write([]byte("hes11"))

	time.Sleep(time.Minute)

	s1.Close()
	s2.Close()
	s.Close()
	conn.Close()
}
