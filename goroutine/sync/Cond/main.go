package main

import (
	"fmt"
	"sync"
	"time"
)

var locker = new(sync.Mutex)

var cond = sync.NewCond(locker)

func test(x int) {
	cond.L.Lock()
	fmt.Println("Waiting start...")
	cond.Wait()
	fmt.Println("Waiting end...")
	fmt.Println(x)
	time.Sleep(time.Second)
	cond.L.Unlock()
	fmt.Println("goutine run ...", x)
}

func main() {
	for i := 0; i < 5; i++ {
		go test(i)
	}

	fmt.Println("start all")
	time.Sleep(time.Second * 3)
	fmt.Println("signal one")
	cond.Signal()
	time.Sleep(time.Second * 3)
	fmt.Println("signal one")
	cond.Signal()
	time.Sleep(time.Second * 3)
	fmt.Println("broadcast")
	cond.Broadcast()
	time.Sleep(time.Second * 60)
}
