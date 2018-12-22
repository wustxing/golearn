package main

import (
	"fmt"
	"github.com/nats-io/go-nats"
	"sanguosha.com/sgs_nuyan/gameutil"
	"sync"
	"time"
)

type logManager struct {
	nc      *gameutil.NatsClient
	msgChan chan *nats.Msg
	wg      sync.WaitGroup

	receiveCnt int32
}

func main() {
	conn, err := nats.Connect(
		nats.DefaultURL,
		nats.MaxReconnects(-1),
		nats.ReconnectWait(2*time.Second),
	)

	if err != nil {
		panic(err)
	}

	//funSub(conn)
	//printStats(conn)
	test(conn)
	time.Sleep(time.Hour)
}

func chanSub(conn *nats.Conn) {
	msgChan := make(chan *nats.Msg, 100000)
	_, _ = conn.ChanSubscribe("log_topic", msgChan)
	go func() {
		for {
			select {
			case <-msgChan:
				time.Sleep(time.Millisecond * 100)
			}
		}
	}()
}

func test(conn *nats.Conn) {
	conn.Publish("log_topic", []byte("sfsfdsf1"))
	conn.Subscribe("log_topic", func(msg *nats.Msg) {
		fmt.Println(string(msg.Data))
	})
	conn.Publish("log_topic", []byte("sfsfdsf2"))
}

func funSub(conn *nats.Conn) {
	sub, _ := conn.Subscribe("log_topic", func(msg *nats.Msg) {
		time.Sleep(time.Millisecond * 1)
	})
	//sub.SetPendingLimits(100000, 100000*1024)
	sub.Dropped()

	ticker := time.NewTicker(time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				count, _ := sub.Dropped()
				fmt.Println("drop", count)
			}
		}
	}()
}

func printStats(conn *nats.Conn) {
	ticker := time.NewTicker(time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Println(conn.Stats())
			}
		}

	}()
}
