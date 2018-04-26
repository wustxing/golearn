package main

import (
	"fmt"
	"reflect"
)

type parent interface {
	SayHi()
}
type struct1 struct{

}

func(s struct1)SayHi(){
	fmt.Println("struct1 say Hi")
}

func(s struct2)SayHi(){
	fmt.Println("struct2 say Hi")
}

type struct2 struct{

}

func SayHello(p parent){
	//p.SayHi()
}

func main(){
	t:=(*struct1)(nil)
	s :=(*int)(nil)
	fmt.Println(t,reflect.TypeOf(t),s,reflect.TypeOf(s))
	SayHello(t)
}


