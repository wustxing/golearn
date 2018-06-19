package main

import "fmt"

type Test struct{
	id int
}
func main(){
	T1 :=&Test{1}
	T2:=*T1
	//T3:=&T2
	T1.id = 5
	fmt.Println(T1,T2)
}
