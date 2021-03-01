package main

import (
	"fmt"
	nats "github.com/nats-io/nats.go"
	"time"
)

var url = "nats://10.225.136.238:5222,nats://10.225.136.238:5223"

func main() {
	conn, err := nats.Connect(url, nats.ReconnectBufSize(-1), nats.MaxReconnects(-1))
	if err != nil {
		panic(err)
	}
	//now := time.Now()
	_,err=conn.Subscribe("topic_login_event", func(msg *nats.Msg) {
		fmt.Println("login:", string(msg.Data))
	})
	fmt.Println(err)

	_,err=conn.Subscribe("topic_match_event", func(msg *nats.Msg) {
		fmt.Println("match:", string(msg.Data))
	})

	fmt.Println(err)

	_,err=conn.Subscribe("topic_pay_event", func(msg *nats.Msg) {
		fmt.Println("pay:", string(msg.Data))
	})

	fmt.Println(err)
	//for {
	//	time.Sleep(time.Millisecond * 1)
	//	err := conn.Publish("hello", []byte(time.Now().String()))
	//	fmt.Println("publish ", err)
	//}

	time.Sleep(time.Hour)
}
