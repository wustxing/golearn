package main

import (
	"github.com/0990/golearn/web/httpshutdown/server"
	"log"
	"time"
)

func main() {
	s := &server.Server{}
	s.Init()

	go func() {
		s.Run()
	}()

	log.Println("app start")
	time.Sleep(time.Second * 10)
	log.Println("server start shutdown")
	s.ShutDown()
	log.Println("app end")
}
