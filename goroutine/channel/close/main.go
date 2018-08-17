package main

import "fmt"

func main() {
	c := make(chan int, 2)
	c <- 1
	close(c)
	v, ok := <-c
	fmt.Println(v, ok)
	v, ok = <-c
	fmt.Println(v, ok)
}
