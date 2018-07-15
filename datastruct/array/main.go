package main

import "fmt"
//数组在函数中传递是值传递
func main(){
	arr:=[3]int{1,2,3}
	fmt.Println(arr)
	changeValue(arr)
	fmt.Println(arr)
}

func changeValue(arr [3]int){
	arr[0] = 2
}
