package main

import (
	"fmt"
	"github.com/nats-io/stan.go"
	"time"
)

func main() {
	sc, err := stan.Connect("test-cluster", "11099999", stan.NatsURL("nats://127.0.0.1:4222"), stan.SetConnectionLostHandler(ServerConnectionLostHandler), stan.Pings(5, 2))
	if err != nil {
		fmt.Println(err)
		return
	}
	sc.NatsConn()
	_, err = sc.Subscribe("foo", func(m *stan.Msg) {
		fmt.Printf("Recieved a message:%s\n", string(m.Data))
	}, stan.DurableName("my-durable"))

	time.Sleep(time.Hour)
	//sub.Unsubscribe()
	sc.Close()

}

func ServerConnectionLostHandler(conn stan.Conn, err error) {
	fmt.Println("serverconn lose", err)
}
