package main

import (
	"fmt"
	"github.com/robfig/cron"
	_ "net/http/pprof"
	"time"
)

func main() {
	spec := "00 11 18 * * ?"
	c := cron.New()
	c.AddFunc(spec, func() {
		fmt.Println("call func:", time.Now())
	})
	c.Start()
	time.Sleep(time.Hour)
}
