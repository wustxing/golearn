package main

import (
	"fmt"
	"github.com/nats-io/go-nats-streaming"
	"time"
)

func main() {
	sc, err := stan.Connect("test-cluster", "myID", stan.NatsURL("nats://localhost:4222"))

	if err != nil {
		panic(err)
	}

	funSub(sc)
	//sc.Publish("foo", []byte("Hello world1"))
	//su, _ := sc.Subscribe("foo", func(m *stan.Msg) {
	//	log.Printf("Receive:%s\n", string(m.Data))
	//}, stan.DurableName("my-durable"))
	//
	//time.Sleep(time.Second)
	////su.Unsubscribe()
	//su.Close()
	////sc.Close()
	//time.Sleep(time.Second)
	////sc2, err := stan.Connect("test-cluster", "myID", stan.NatsURL("nats://localhost:4222"))
	//
	////if err != nil {
	////	panic(err)
	////}
	//sc.Publish("foo", []byte("Hello world2"))
	//time.Sleep(time.Second)
	//sc.Subscribe("foo", func(m *stan.Msg) {
	//	log.Printf("Receive:%s\n", string(m.Data))
	//}, stan.DurableName("my-durable"))

	time.Sleep(time.Hour)
}

var count int32

func funSub(conn stan.Conn) {
	sub, _ := conn.Subscribe("log_topic", func(msg *stan.Msg) {
		count++
		time.Sleep(time.Millisecond * 1)
	}, stan.DurableName("my-durables"))
	//sub.SetPendingLimits(100000, 100000*1024)
	//sub.Dropped()

	ticker := time.NewTicker(time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				drop, err := sub.Dropped()
				fmt.Printf("receive:%d,drop:%d,err:%v,stats:%v\n", count, drop, err, conn.NatsConn().Stats())
			}
		}
	}()
}
