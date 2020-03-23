package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan int, 2)
	c <- 1
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for x := range c {
			fmt.Println(x)
		}
		fmt.Println("go func exit")
	}()
	close(c)
	wg.Wait()
	//time.Sleep(time.Second)
	//v, ok := <-c
	//fmt.Println(v, ok)
	//v, ok = <-c
	//fmt.Println(v, ok)
}
