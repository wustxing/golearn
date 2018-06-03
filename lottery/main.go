package main

import "fmt"

type item struct{
	index int
	rate int
}

type randItem struct{
	index int
	totalRate int
}
func main(){
	intArr:=make([]int,0)
	intArr1:=make([]int,0)
	intArr = append(intArr,1)
	intArr1 = append(intArr1,2)
	intArr1 = append(intArr1,3)
	intArr = append(intArr,intArr1...)
	fmt.Println(intArr)
}
