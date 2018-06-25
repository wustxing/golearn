package main

import "fmt"

func main(){
	intSlice:=make([]int,0)

	for i:=0;i<10;i++{
		intSlice = append(intSlice,i)
	}
	intSlice = append(intSlice,10)

	fmt.Println(intSlice)
	fmt.Println(intSlice[:])
	fmt.Println(intSlice[:1])
	fmt.Println(intSlice[0:])
	//
	//for i,v:=range intSlice{
	//	if v==5{
	//		fmt.Println(i)
	//		intSlice = append(intSlice[:i],intSlice[i+1:]...)
	//	}
	//}
	intSlice = intSlice[1:]
	//intSlice = append(intSlice[:0],intSlice[1:]...)
	//copy(intSlice,intSlice[1:])

	fmt.Println(intSlice)
}
