package main

import (
	"github.com/nats-io/go-nats"
	"time"
	"fmt"
)

func main(){
	nc,err:=nats.Connect(nats.DefaultURL)
	if err!=nil{
		panic(err)
	}
	nc.Subscribe("foo",func(m *nats.Msg){
		fmt.Printf("Received a message:%s\n",string(m.Data))
	})
	nc.Subscribe("foo",func(m *nats.Msg){
		fmt.Printf("Received hello message:%s\n",string(m.Data))
	})
	nc.Publish("foo",[]byte("Hello world"))



	time.Sleep(time.Second*100)
}
