package main

import (
	"fmt"
	nats "github.com/nats-io/nats.go"
	"time"
)

var url = "nats://10.225.136.212:5222,nats://10.225.136.212:5223"

func main() {
	conn, err := nats.Connect(url, nats.ReconnectBufSize(-1), nats.MaxReconnects(-1))
	if err != nil {
		panic(err)
	}
	now := time.Now()
	conn.Subscribe("hello", func(msg *nats.Msg) {
		fmt.Println("request", time.Since(now))
	})
	for {
		time.Sleep(time.Millisecond * 1)
		err := conn.Publish("hello", []byte(time.Now().String()))
		fmt.Println("publish ", err)
	}

	time.Sleep(time.Hour)
}
