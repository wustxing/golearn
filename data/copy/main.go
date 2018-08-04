package main

import "fmt"

type Test struct {
	id       int
	ptr      *int
	helloMap map[int32]*int
}

func main() {
	a := 1
	b := 2
	c := 3
	T1 := &Test{
		id:       1,
		ptr:      &a,
		helloMap: make(map[int32]*int),
	}
	T1.helloMap[1] = &b
	temp := *T1
	T2 := &temp
	T2.helloMap = make(map[int32]*int)
	T2.helloMap[2] = &c
	T2.ptr = &c
	T2.id = 6
	fmt.Println(T1, T2)
}
