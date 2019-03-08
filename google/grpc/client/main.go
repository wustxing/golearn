package main

import (
	"context"
	"github.com/0990/golearn/google/grpc"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	address     = "localhost:50051"
	defaultName = "kids"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect:%v", err)
	}

	defer conn.Close()
	c := helloworld.NewGreeterClient(conn)

	name := defaultName
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SayHello(ctx, &helloworld.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet:%v", err)
	}
	log.Printf("Greeting:%s", r.Message)
}
