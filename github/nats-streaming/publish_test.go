package main

import (
	"fmt"
	stan "github.com/nats-io/stan.go"
	"testing"
	"time"
)

var url = "nats://127.0.0.1:4222"

var pubConn stan.Conn

func TestPublish(t *testing.T) {
	conn, err := stan.Connect("test-cluster", "111", stan.NatsURL(url), stan.SetConnectionLostHandler(ConnectionLostHandler), stan.Pings(5, 2))
	if err != nil {
		fmt.Println(err)
		return
	}
	pubConn = conn
	PublishPerSecond()

	time.Sleep(time.Hour)
}

func PublishPerSecond() {
	var count int32
	//start := time.Now().Unix()
	//go func() {
	//	ticker := time.NewTicker(time.Second * 2)
	//	defer ticker.Stop()
	//	for range ticker.C {
	//		now := time.Now().Unix()
	//		if now-start > 0 {
	//			fmt.Printf("total:%v sendRate:%v \n", count, count/int32(now-start))
	//		}
	//	}
	//}()

	for {
		s := fmt.Sprintf("%d", count)
		_, err := pubConn.PublishAsync("foo", []byte(s), func(s string, err error) {
			if err != nil {
				fmt.Println("send callback error", err)
			}
		})

		if err != nil {
			fmt.Printf("send error:%v,count:%d \n", err, count)
			//time.Sleep(time.Second * 10)
			//fmt.Println("start pub")
		} else {
			fmt.Printf("send success,count:%d \n", count)
		}

		count++
		time.Sleep(time.Second * 2)
		//
		//now := time.Now().Unix()
		//if now%10 == 0 && now-start > 0 {
		//	fmt.Printf("total:%v sendRate:%v \n", count, count/int32(now-start))
		//}

		//err := pubConn.Publish("foo", []byte(s))
		//fmt.Println("publish ", s, err, &pubConn)
		//time.Sleep(time.Millisecond * 10)
	}
}

func ConnectionLostHandler(conn stan.Conn, err error) {
	fmt.Println("conn lose", err)
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for range ticker.C {

		conn1, err := stan.Connect("test-cluster", "111", stan.NatsURL(url), stan.SetConnectionLostHandler(ConnectionLostHandler))
		if err == nil {
			pubConn = conn1
			//a:=conn.NatsConn()
			//conn.
			//go PublishPerSecond(conn1)
			return
		}
	}
}

func TestSubAllAvailable(t *testing.T) {
	sc, err := stan.Connect("test-cluster", "1100", stan.NatsURL("nats://127.0.0.1:5222"))
	if err != nil {
		fmt.Println(err)
		return
	}
	sub, err := sc.Subscribe("foo", func(m *stan.Msg) {
		fmt.Printf("Recieved a message:%s\n", string(m.Data))
	}, stan.DeliverAllAvailable())

	sub.Unsubscribe()
	sc.Close()
}

func TestSubDurable(t *testing.T) {
	sc, err := stan.Connect("test-cluster", "1101", stan.NatsURL(url))
	if err != nil {
		fmt.Println(err)
		return
	}
	var count int32
	var start int64
	sub, err := sc.Subscribe("foo1", func(m *stan.Msg) {
		if start == 0 {
			start = time.Now().Unix()
		}
		count++
		//fmt.Println(string(m.Data))
	})

	fmt.Println(sub.PendingLimits())
	sub.SetPendingLimits(102400, 5*1024*1024)

	ticker := time.NewTicker(time.Second * 2)
	defer ticker.Stop()
	for range ticker.C {
		now := time.Now().Unix()
		if now-start > 0 {
			fmt.Printf("total:%v sendRate:%v \n", count, count/int32(now-start))
		}
	}
	//sub.Unsubscribe()
	sc.Close()
}
