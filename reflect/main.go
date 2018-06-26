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

	fmt.Println(reflect.TypeOf(f))
	fmt.Println(reflect.ValueOf(f))

	var x int = 1
	fmt.Println(reflect.TypeOf(x))
	fmt.Println(reflect.ValueOf(x))
	fmt.Println(reflect.ValueOf(x).Kind())
	fmt.Println(reflect.ValueOf(x).Kind()==reflect.Int)

}