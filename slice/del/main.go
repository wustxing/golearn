package main

import "fmt"

func main(){
	intSlice:=make([]int,0)

	for i:=0;i<10;i++{
		intSlice = append(intSlice,i)
	}

	fmt.Println(intSlice)
	fmt.Println(intSlice[:1])
	fmt.Println(intSlice[:5])
	fmt.Println(intSlice[6:])

	for i,v:=range intSlice{
		if v==5{
			fmt.Println(i)
			intSlice = append(intSlice[:i],intSlice[i+1:]...)
		}
	}

	fmt.Println(intSlice)

}
