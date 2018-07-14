package main

import (
	"reflect"
	"fmt"
	"github.com/0990/golearn/data/dirtyMaptyMap/util"
)

func set(name string,valueMustBePtr interface{}){
	v:=reflect.ValueOf(valueMustBePtr)
	fmt.Println(v)
	if v.Kind()!=reflect.Ptr{
		panic("Dirty value must be ptr")
	}

	elem := v.Elem()
	ori:=elem.Interface()

	copy:=reflect.New(elem.Type()).Interface()
	err:=util.DeepCopy(copy,ori)
	if err!=nil{
		panic(err)
	}

	ret := reflect.ValueOf(copy).Elem().Interface()
	fmt.Println(elem,ori,copy,ret)
}
func main(){
	//str:="hello"
	by := []byte("golang")
	set("name",&by)
}
