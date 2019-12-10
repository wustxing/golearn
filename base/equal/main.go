package main

import "fmt"

func main() {
	a := 1
	b := 1

	aptr := &a
	bptr := &b
	fmt.Println(aptr == bptr)
}
