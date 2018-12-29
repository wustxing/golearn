package main

import "fmt"

func main() {
	intArr := [2]int{1, 2}
	intSlice := intArr[:]
	fmt.Printf("%T,%v", intArr, intArr)
	fmt.Printf("%T,%v", intSlice, intSlice)
}
