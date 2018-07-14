package main

import (
	"sync"
	"fmt"
)

var once sync.Once
var a string
func setup(){
	a = "hello,world"
	fmt.Println("enter setup")
}

func doprint(){
	once.Do(setup)
	print(a)
}

func main(){
	fmt.Println("enter main")
	go doprint()
	go doprint()
	select{}
}
