package main

import (
	"fmt"
	"github.com/nats-io/go-nats"
	"time"
)

//var natsUrl = "nats://localhost:4222"

var natsUrl = "nats://47.99.219.57:4222"

func main() {
	conn, err := nats.Connect(
		natsUrl,
		nats.MaxReconnects(-1),
		nats.ReconnectWait(2*time.Second),
	)

	if err != nil {
		panic(err)
	}

	conn.Subscribe("report", func(msg *nats.Msg) {
		fmt.Println(string(msg.Data))
	})

	conn.Subscribe("loadbalance", func(msg *nats.Msg) {
		fmt.Println(string(msg.Data))
	})

	time.Sleep(24 * time.Hour)
}
