package main

import (
	"encoding/json"
	"fmt"
	"github.com/0990/golearn/socket/tcp/demo/client"
	"testing"
)

func Test_Client(t *testing.T) {
	c, err := client.New("127.0.0.1:5000")
	if err != nil {
		t.Fatal(err)
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


	err = c.Send(msg)
	if err != nil {
		t.Error(err)
	}
}
