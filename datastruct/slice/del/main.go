package main

import "fmt"

func main(){
	intSlice:=make([]int,0)

	for i:=0;i<2;i++{
		intSlice = append(intSlice,i)
	}
	//intSlice = append(intSlice,10)

	//fmt.Println(intSlice)
	//fmt.Println(intSlice[:])
	//fmt.Println(intSlice[:1])
	//fmt.Println(intSlice[0:])
	//
	//fmt.Println(intSlice)
	//for i,v:=range intSlice{
	//	if v==0{
	//
	//		intSlice = append(intSlice[:i],intSlice[i+1:]...)
	//	}
	//}
	//intSlice = intSlice[1:]
	//intSlice = append(intSlice[:0],intSlice[1:]...)
	//copy(intSlice,intSlice[1:])

	//fmt.Println(intSlice)
	addInt(intSlice...)
}

func addInt(ints ...int){
	for _,v:=range ints{
		fmt.Println(v)
	}
}