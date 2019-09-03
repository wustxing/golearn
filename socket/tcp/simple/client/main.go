package main

import (
	"log"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", ":9999")
	if err != nil {
		log.Println("dial error", err)
	}
	defer conn.Close()
	conn.Write([]byte("I am socket client"))
	time.Sleep(time.Second)
}
