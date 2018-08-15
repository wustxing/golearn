package main

import (
	"fmt"
	"github.com/nats-io/go-nats"
	"time"
)

const (
	testStr = `{"area_id":10,"log_info":"eyJsb2dvbklQIjoiMTAuMjI1LjEwLjI0NSIsInVzZXJJRCI6MjN9","log_time":1533645769,"login_from":1,"op_type":101,"param1":1,"param2":10,"tbl_name":"tbl_login","user_account":"hello","user_level":1}`
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		panic(err)
	}

	ticker := time.NewTicker(time.Millisecond * 50)

	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Println("public", time.Now())
				for i := 0; i < 100000; i++ {
					nc.Publish("log_topic", []byte(testStr))
				}
			}
		}
	}()

	time.Sleep(time.Second * 5)
}
