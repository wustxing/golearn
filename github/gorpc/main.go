package main

import (
	"encoding/gob"
	"fmt"
	"github.com/ankur-anand/simple-go-rpc/src/client"
	"net"
	"time"

	"github.com/ankur-anand/simple-go-rpc/src/server"
)

type User struct {
	Name string
	Age  int
}

var userDB = map[int]User{
	1: User{"xu", 1},
	2: User{"hi", 2},
}

func QueryUser(id int) (User, error) {
	if u, ok := userDB[id]; ok {
		return u, nil
	}

	return User{}, fmt.Errorf("id %d not in user db", id)
}

func main() {
	gob.Register(User{})
	addr := "localhost:3212"
	srv := server.NewServer(addr)
	srv.Register("QueryUser", QueryUser)

	go srv.Run()

	time.Sleep(time.Second)

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		panic(err)
	}

	cli := client.NewClient(conn)
	var Query func(int) (User, error)
	cli.CallRPC("QueryUser", &Query)

	u, err := Query(1)
	if err != nil {
		panic(err)
	}

	fmt.Println(u)

}
