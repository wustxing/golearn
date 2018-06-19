package main

import "fmt"

func main(){
	lists:=make([]int,0)
	lists=append(lists,1,2,3)
	fmt.Println(lists)
}