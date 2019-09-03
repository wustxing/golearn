package main

import (
	"fmt"
	"github.com/AsynkronIT/protoactor-go/log"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "0.0.0.0:9999")
	if err != nil {
		fmt.Println("listen error:", err)
		return
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			break
		}
		go HandleConn(conn)
	}

}

func HandleConn(conn net.Conn) {
	defer conn.Close()
	buffer := make([]byte, 100)
	for {
		//read
		n, err := conn.Read(buffer)
		if err != nil {
			log.Error(err)
			return
		}
		fmt.Println("receive data:", string(buffer[:n]))
	}
}
