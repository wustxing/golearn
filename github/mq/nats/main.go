package main

import (
	"fmt"
	"github.com/nats-io/go-nats"
	"time"
)

const NatsURL = "nats://10.225.136.159:4222"

func main() {
	nc, err := nats.Connect("nats://10.225.136.159:4222")
	if err != nil {
		panic(err)
	}

	//nc.Subscribe("foo", func(m *nats.Msg) {
	//	fmt.Printf("Received a message:%s\n", string(m.Data))
	//})

	nc.Subscribe("foo", func(m *nats.Msg) {
		fmt.Printf("Received hello message:%s\n", string(m.Data))
	})

	nc.Publish("foo", []byte("Hello world"))
	nc.Publish("foo", []byte("xujialong"))

	time.Sleep(time.Hour)
}
