package main

import (
	"encoding/json"
	"fmt"
	"github.com/0990/golearn/socket/tcp/demo/server"
	"os"
	"os/signal"
)

const Port = 8300

//type User struct {
//	ID   int32  `json:id`
//	Name string `json:name`
//}

type User struct {
	ID   int32
	Name string
}

type Resp struct {
	Code int32 `json:code`
}

func main() {
	address := fmt.Sprintf(":%d", Port)
	svr := server.New(address)
	svr.OnNewMessage(func(c *server.Client, message string) {
		user := &User{}
		json.Unmarshal([]byte(message), user)
		fmt.Println(user)
		resp := &Resp{
			Code: 100,
		}

		data, _ := json.Marshal(resp)
		msg := fmt.Sprintf("%s%s", data, "\n")
		c.Send(msg)
	})

	svr.OnClientConnectionClosed(func(c *server.Client, err error) {
		fmt.Println("connection closed")
	})

	svr.OnNewClient(func(c *server.Client) {
		fmt.Println("new connection open")
	})

	go svr.Listen()

	c := make(chan os.Signal, 1)
	signal.Notify(c)
	<-c
}
