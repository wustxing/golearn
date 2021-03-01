package main

import (
	"fmt"
	"github.com/nats-io/stan.go"
	"time"
)

func main() {
	sc, err := stan.Connect("test-cluster", "11099999", stan.NatsURL("nats://127.0.0.1:4222"))
	if err != nil {
		fmt.Println(err)
		return
	}
	sc.NatsConn()
	_, err = sc.Subscribe("log_topic", func(m *stan.Msg) {
		fmt.Printf("Recieved a message:%s\n", string(m.Data))
	}, stan.StartWithLastReceived())

	time.Sleep(time.Hour)
	//sub.Unsubscribe()
	sc.Close()

}
