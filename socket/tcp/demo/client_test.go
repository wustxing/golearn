package main

import (
	"encoding/json"
	"fmt"
	"github.com/0990/golearn/socket/tcp/demo/client"
	"testing"
	"time"
)

func Test_Client(t *testing.T) {
	c, err := client.New("127.0.0.1:8300")
	if err != nil {
		t.Error(err)
	}
	user := &User{
		ID:   1,
		Name: "xujialong",
	}
	data, err := json.Marshal(user)
	if err != nil {
		t.Error(err)
	}
	msg := fmt.Sprintf("%s%s", data, "\n")
	resp, err := c.Send(msg)
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)
	time.Sleep(time.Second * 2)
}
