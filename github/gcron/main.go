package main

import (
	"fmt"
	"github.com/roylee0704/gron"
	"time"
)

func main() {
	c := gron.New()
	c.AddFunc(gron.Every(time.Second), func() {
		fmt.Println("hello")
	})
	c.AddFunc(gron.Every(time.Hour).At("00:00"), func() {
		fmt.Println("at schedule")
	})
	c.Start()

	time.Sleep(time.Hour)
}
