package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Age int32
}

func main() {
	fmt.Println(reflect.TypeOf((*User)(nil)).Elem())
}
