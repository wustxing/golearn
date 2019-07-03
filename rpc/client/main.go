package main

import (
	"fmt"
	"github.com/0990/golearn/rpc/msg"
	"github.com/gogo/protobuf/proto"
	"github.com/nats-io/go-nats"
	"github.com/sirupsen/logrus"
	"time"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		logrus.Panic(err)
	}

	nc.Subscribe("toclient", func(message *nats.Msg) {
		respRPC := &msg.RespRPC{}
		err := proto.Unmarshal(message.Data, respRPC)
		if err != nil {
			logrus.Println(err)
			return
		}
		resp := &msg.RespHello{}
		err = proto.Unmarshal(respRPC.Data, resp)
		if err != nil {
			logrus.Println(err)
			return
		}
		fmt.Println(resp)
	})

	reqRPC := &msg.ReqRPC{}
	reqRPC.Data, _ = proto.Marshal(&msg.ReqHello{
		Name: "xu",
	})

	data, _ := proto.Marshal(reqRPC)

	nc.Publish("toserver", data)

	time.Sleep(time.Hour)
}
