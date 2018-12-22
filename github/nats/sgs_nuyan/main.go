package main

import (
	"fmt"
	"github.com/nats-io/go-nats"
	"time"
)

func main() {
	conn, err := nats.Connect(
		nats.DefaultURL,
		nats.MaxReconnects(-1),
		nats.ReconnectWait(2*time.Second),
	)

	if err != nil {
		panic(err)
	}

	conn.Subscribe("report", func(msg *nats.Msg) {
		fmt.Println(string(msg.Data))
	})

	time.Sleep(24 * time.Hour)
}
