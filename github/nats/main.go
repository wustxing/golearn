package main

import (
	"fmt"
	"github.com/nats-io/go-nats"
	"time"
)

func main() {
	conn, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		panic(err)
	}
	now := time.Now()
	conn.Subscribe("hello", func(msg *nats.Msg) {
		fmt.Println("request", time.Since(now))
	})
	conn.Publish("hello", []byte("hidddd"))
	time.Sleep(time.Hour)
}
