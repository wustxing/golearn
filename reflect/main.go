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

	p:=reflect.ValueOf(&x)
	fmt.Println(p)
	v:=p.Elem()
	v.SetInt(7)

	f =Foo{
		X:"hello",
		Y:1,
	}

	s := reflect.ValueOf(&f).Elem()

	typeofF :=s.Type()

	for i:=0;i<s.NumField();i++{
		f:=s.Field(i)
		fmt.Println(i,typeofF.Field(i).Name,f.Type,f.Interface())
	}

	reflect.ValueOf(x).Call()

}