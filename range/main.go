package main

import "fmt"
//注意：range 中的v只存一个副本
type student struct{
	Name string
	Age int
}

func main(){
	smap:=make(map[int]student)
	lists:=[...]student{
		student{Name:"xu",Age:10},
		student{Name:"jia",Age:2},
	}
	for _,v:=range lists{
		//name :=v
		smap[v.Age] = v
	}
	for k,v:=range smap{
		fmt.Println(k,v)
	}
}
