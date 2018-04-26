package main

import (
	"fmt"

	"github.com/robfig/cron"
)

func main() {
	spec := "0 45 16 * * *"
	c := cron.New()
	c.AddFunc(spec, callYourFunc)
	c.Start()
	fmt.Println("start")
	select {}
	fmt.Println("end")
}

func callYourFunc() {
	fmt.Println("callYourFunc come")
}
