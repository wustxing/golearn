package main

import (
	"fmt"
	"reflect"
)

func say(text string){
	fmt.Println(text)
}


func Call(m map[string]interface{},name string,parmas...interface{})(result []reflect.Value){
	f:=reflect.ValueOf(m[name])
	fmt.Println(f)

	in:=make([]reflect.Value,len(parmas))

	for k,param:=range parmas{
		in[k] = reflect.ValueOf(param)
		fmt.Println("k:",k,",valueofparam:",reflect.ValueOf(param))
	}

	result = f.Call(in)
	return
}

func main(){
	var funcMap = make(map[string]interface{})
	funcMap["say"] = say
	Call(funcMap,"say","world")

	fmt.Println(reflect.TypeOf(say).Kind())
}