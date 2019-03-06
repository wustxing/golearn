package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := time.Now()

	for i := 0; i < 10000; i++ {
		fmt.Println("fsdf")
	}

	elapsed := time.Since(t1)

	fmt.Println("elapsed:", elapsed)
}
