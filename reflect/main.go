package main

import (
	"fmt"
	"reflect"
)

type Foo struct{
	X string
	Y int
}

func(f Foo) Do(){
	fmt.Printf("X is:%s,Y is:%d\n",f.X,f.Y)
}
func main(){
	f:=Foo{
		"abc",
		123,
	}

	reflect.ValueOf(f).MethodByName("Do").Call([]reflect.Value{})
}