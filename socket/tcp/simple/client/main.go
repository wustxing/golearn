package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", ":9999")
	if err != nil {
		log.Println("dial error", err)
	}
	//	defer conn.Close()
	buf := make([]byte, 100)
	conn.SetReadDeadline(time.Now().Add(time.Second))
	_, err = conn.Read(buf)
	if err != nil {
		e, ok := err.(net.Error)
		fmt.Println(err, e, ok, e.Temporary(), e.Timeout())
	}
	conn.SetReadDeadline(time.Now().Add(3 * time.Second))
	_, err = conn.Read(buf)
	if err != nil {
		e, ok := err.(net.Error)
		fmt.Println(err, e, ok, e.Temporary(), e.Timeout())
	}
	//_, err = conn.Write([]byte("I am socket client"))
	//if err != nil {
	//	fmt.Println(err)
	//}
	time.Sleep(time.Minute)
}
