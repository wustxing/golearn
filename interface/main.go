package main

import "fmt"

func main(){
	var a interface{}
	var b string
	a = "asss"
	b = a.(string)
	c,ok := a.(int)
	fmt.Println(a,b,c,ok)
}
