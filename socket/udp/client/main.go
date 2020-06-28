package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 9090,
		Zone: "",
	})

	if err != nil {
		panic(err)
	}

	_, err = conn.Write([]byte("hello server!"))
	if err != nil {
		panic(err)
	}

	result := make([]byte, 4096)
	n, remoteAddr, err := conn.ReadFromUDP(result)
	if err != nil {
		panic(err)
	}

	fmt.Printf("receive from addr:%v data:%v\n", remoteAddr, string(result[:n]))

}
