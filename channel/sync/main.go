package main

import (
	"fmt"
	"sync"
)

var waitgroup sync.WaitGroup

func AFunction(shownum int) {
	fmt.Println(shownum)
	waitgroup.Done()
}

func main() {
	for i := 0; i < 5; i++ {
		waitgroup.Add(1)
		go AFunction(i)
	}

	waitgroup.Wait()
}
