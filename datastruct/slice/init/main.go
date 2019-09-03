package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type Slice struct {
	ptr unsafe.Pointer // Array pointer
	len int            // slice length
	cap int            // slice capacity
}

func main() {
	s1 := new([]int)
	slice1 := (*Slice)(unsafe.Pointer(s1))
	fmt.Println(slice1)

	s2 := make([]int, 0)
	slice2 := (*Slice)(unsafe.Pointer(&s2))
	fmt.Println(slice2)

	s3 := []int{}
	slice3 := (*Slice)(unsafe.Pointer(&s3))
	fmt.Println(slice3)

	var s4 []int
	slice4 := (*Slice)(unsafe.Pointer(&s4))
	fmt.Println(slice4)

	fmt.Println(*s1 == nil, s2 == nil, s3 == nil, s4 == nil)
	fmt.Println(reflect.TypeOf(s1), reflect.TypeOf(s2), reflect.TypeOf(s3), reflect.TypeOf(s4))
}
