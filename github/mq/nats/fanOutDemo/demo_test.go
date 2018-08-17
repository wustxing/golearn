package main

import (
	"fmt"
	"github.com/nats-io/go-nats"
	"sync"
	"testing"
	"time"
)

const (
	testStr = `{"area_id":10,"log_info":"eyJsb2dvbklQIjoiMTAuMjI1LjEwLjI0NSIsInVzZXJJRCI6MjN9","log_time":1533645769,"login_from":1,"op_type":101,"param1":1,"param2":10,"tbl_name":"tbl_login","user_account":"hello","user_level":1}`
)

var totalSendCnt int32

func Test_benchmark(t *testing.T) {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		panic(err)
	}

	ticker := time.NewTicker(time.Millisecond * 100)
	exitTimer := time.NewTimer(time.Second * 5)
	var wgTest sync.WaitGroup
	wgTest.Add(1)
	go func() {
		defer func() {
			wgTest.Done()
		}()
		for {
			select {
			case <-ticker.C:
				for i := 0; i < 100000; i++ {
					err := nc.Publish("log_topic", []byte(testStr))
					if err == nil {
						totalSendCnt++
					}
				}
			case <-exitTimer.C:
				return
			}
		}
	}()
	wgTest.Wait()
	fmt.Println("sendTotalLen:", totalSendCnt)
}
