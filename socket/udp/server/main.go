package main

import (
	"fmt"
	"net"
)

func main() {
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 9090,
		Zone: "",
	})

	if err != nil {
		panic(err)
	}

	for {
		var data [1024]byte
		n, addr, err := listen.ReadFromUDP(data[:])
		if err != nil {
			fmt.Println(err)
			break
		}

		fmt.Printf("addr:%v data:%v count:%v\n", addr, string(data[:n]), n)
		_, err = listen.WriteToUDP([]byte("received success!"), addr)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}
