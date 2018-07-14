package main

import "fmt"

func main(){
	lists:=make([]int,0)
	lists=append(lists,1,2,3)
	maps :=make(map[int]*int)

	for _,v:=range lists{
		fmt.Println(v)
		maps[v] = &v
	}
	fmt.Println(lists)
	fmt.Println(maps)
}