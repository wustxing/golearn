package main

import "fmt"

func main() {
	a := 2
	p := &a
	fmt.Printf("pointer:%p,target:%v\n", &p, p)
	test(&a)
}

func test(x *int) {
	fmt.Printf("test pointer:%v,target:%v\n", &x, x)
}
