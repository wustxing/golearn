package main

import (
	"fmt"
	helloworld "github.com/0990/golearn/rpc/netgrpc"
	"log"
)

func main() {
	client, err := helloworld.DialGreeter("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply helloworld.HelloReply
	client.SayHello(&helloworld.HelloRequest{Name: "0990"}, &reply)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}
