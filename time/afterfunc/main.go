package main

import (
	"time"
	"fmt"
)

func main(){
	time.AfterFunc(time.Second,func(){
		fmt.Println("timer come")
	})
	time.Sleep(time.Second*10)
	fmt.Println("main end")
}
