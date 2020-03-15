package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	go func() {
		err := http.ListenAndServe("localhost:5000", http.DefaultServeMux)
		if err != nil {
			fmt.Println(err)
		}
	}()

	time.Sleep(time.Second)

	for j := 0; j < 10; j++ {
		c := make([]int, 0, 100000000)
		for i := 0; i < 100000000; i++ {
			c = append(c, 1)
		}
		time.Sleep(time.Microsecond * 50)
		fmt.Println(j)
	}

	time.Sleep(time.Hour)
}
