package main

import (
	"fmt"
	"time"
)

func main() {
	hello()
	fmt.Println("end")
	time.Sleep(time.Second)
}

func hello() {

	panic("sdfsf")
	protect()
}

func protect() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
}
