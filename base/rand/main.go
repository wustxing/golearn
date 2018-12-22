package main

import (
	"fmt"
	"time"
)

const times = 10000

type User struct {
	c *time.Ticker
}

func main() {
	//println(util.Rand63())
	//t1 := time.Now()
	//defer func() {
	//	elapsed := time.Since(t1)
	//	fmt.Println(elapsed)
	//}()
	//
	//var wg sync.WaitGroup
	//wg.Add(times)
	//for i := 0; i < times; i++ {
	//	go func() {
	//		util.RandNum(100000)
	//		wg.Done()
	//	}()
	//}
	//wg.Wait()
	user := new(User)
	user.c = time.NewTicker(time.Second)

	go func() {
		for range user.c.C {
			fmt.Printf("hello")
		}
	}()
	time.Sleep(time.Second * 10)
}
