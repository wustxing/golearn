package main

import (
	"fmt"
	"reflect"
)

func main() {
	p := new(int32)
	fmt.Println(p)
	fmt.Println(p == nil)
	fmt.Println(*p)
	*p = 1
	fmt.Println(*p)

	ps := new([]int)
	fmt.Println(ps)
	fmt.Println(*ps)
	fmt.Println(ps == nil)
	fmt.Println(*ps == nil)

	var k []int
	fmt.Println(k == nil)
	fmt.Println(reflect.TypeOf(&k).Kind())
	fmt.Println(reflect.TypeOf(k).Kind())
}
