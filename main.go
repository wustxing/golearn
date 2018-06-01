package main

import "fmt"

type hello struct{
	Map1 map[int32]string
}
func main() {
	list:=[5]int{1,2,3,4,5}

	for _,v:=range list{
		fmt.Println(v)
	}

	for i:=len(list)-1;i>=0;i--{
		fmt.Println(list[i])
	}
}
