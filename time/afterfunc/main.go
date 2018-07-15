package main

import (
	"time"
	"fmt"
)

func main(){
	timer:=time.AfterFunc(2*time.Second,func(){
		fmt.Println("after come")
	})

	time.AfterFunc(time.Second,func(){
		fmt.Println("cancel after")
		timer.Stop()
	})

	time.Sleep(time.Second*10)
	fmt.Println("main end")
}
