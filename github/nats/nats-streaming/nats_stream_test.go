package main

import (
	"fmt"
	"github.com/nats-io/go-nats-streaming"
	"sync"
	"testing"
	"time"
)

const (
	testStr = `{"area_id":10,"log_info":"eyJsb2dvbklQIjoiMTAuMjI1LjEwLjI0NSIsInVzZXJJRCI6MjN9","log_time":1533645769,"login_from":1,"op_type":101,"param1":1,"param2":10,"tbl_name":"tbl_login","user_account":"%s","user_level":1}`
)

func TestNatsSteamSend(t *testing.T) {
	sc, err := stan.Connect("test-cluster", "myIDs", stan.NatsURL("nats://localhost:4222"))
	if err != nil {
		panic(err)
	}
	wg := &sync.WaitGroup{}
	ticker := time.NewTicker(time.Millisecond * 100)
	sendTime := 1
	sendCount := 0
	wg.Add(1)
	go func() {
		defer time.Sleep(time.Second * 20)
		defer wg.Done()
		for {
			select {
			case <-ticker.C:
				for i := 0; i < 100000; i++ {
					newStr := fmt.Sprintf(testStr, fmt.Sprintf("%d_%d", sendTime, i))
					err := sc.Publish("log_topic", []byte(newStr))
					if err == nil {
						sendCount++
					}
				}

				sendTime--
				if sendTime <= 0 {
					return
				}
			}
		}
	}()

	wg.Wait()
	fmt.Println("exit main goroutine,sendCount:", sendCount)
}
