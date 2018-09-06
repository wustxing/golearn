package main

import (
	"fmt"

	"github.com/robfig/cron"
	"time"
)
//每隔5秒执行一次：*/5 * * * * ?
//每隔1分钟执行一次：0 */1 * * * ?
//每天23点执行一次：0 0 23 * * ?
//每天凌晨1点执行一次：0 0 1 * * ?
//每月1号凌晨1点执行一次：0 0 1 1 * ?
//在26分、29分、33分执行一次：0 26,29,33 * * * ?
//每天的0点、13点、18点、21点都执行一次：0 0 0,13,18,21 * * ?
func main() {
	time.Now().Unix()

	spec := "*/60 * * * * ?"
	spec1 := "5 06 19 * * ?"
	c := cron.New()
	c.AddFunc(spec, callYourFunc)
	c.AddFunc(spec1,callYourFunc)
	c.Start()
	for _,v:=range c.Entries(){
		fmt.Println(v.Next.Unix()-time.Now().Unix())
	}

	fmt.Println("start")
	select {}
	fmt.Println("end")
}

func callYourFunc() {
	fmt.Println("callYourFunc come")
}
