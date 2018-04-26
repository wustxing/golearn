package main

import (
	"fmt"
	"sync"
	"time"
)

var waitGroup sync.WaitGroup

func main() {
	t1 := time.NewTimer(time.Second * 2)

	<-t1.C

	fmt.Println("Timer1 expired")

	waitGroup.Add(1)
	t2 := time.NewTimer(time.Second * 2)

	go func() {
		<-t2.C
		fmt.Println("Timer2 expired")
		waitGroup.Done()
	}()

	waitGroup.Wait()
}
