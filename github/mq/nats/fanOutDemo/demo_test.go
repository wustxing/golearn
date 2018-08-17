package main

import (
	"github.com/nats-io/go-nats"
	"testing"
	"time"
)

const (
	testStr = `{"area_id":10,"log_info":"eyJsb2dvbklQIjoiMTAuMjI1LjEwLjI0NSIsInVzZXJJRCI6MjN9","log_time":1533645769,"login_from":1,"op_type":101,"param1":1,"param2":10,"tbl_name":"tbl_login","user_account":"hello","user_level":1}`
)

func Test_benchmark(t *testing.T) {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		panic(err)
	}

	ticker := time.NewTicker(time.Millisecond * 100)
	go func() {
		for {
			select {
			case <-ticker.C:
				for i := 0; i < 10000; i++ {
					nc.Publish("log_topic", []byte(testStr))
				}
			}
		}
	}()
	time.Sleep(time.Second * 5)
}
