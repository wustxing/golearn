package main

import (
	"fmt"
	"log"
	"net"
	"time"
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
	time.Sleep(time.Second * 10)
	conn.Close()
	//defer conn.Close()
	buffer := make([]byte, 100)
	for {
		//read
		n, err := conn.Read(buffer)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println("receive data:", string(buffer[:n]))
	}
}
