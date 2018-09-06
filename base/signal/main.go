package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c)
	ticker := time.Tick(time.Second * 1)
	go func() {
		for {
			select {
			case <-ticker:
				fmt.Println("ticker come")
			}
		}
	}()
	s := <-c
	fmt.Println("Got signal:", s)
}
