package main

import "fmt"

//数组在函数中传递是值传递
func main() {
	//arr := [3]int{1, 2, 3}
	//fmt.Println(arr)
	//changeValue(arr)
	//fmt.Println(arr)
	testArr()
	testSlice()
}

func changeValue(arr [3]int) {
	arr[0] = 2
}

//数组这里很奇怪
func testArr() {
	arr := [3]int{1, 2, 3}
	for i, v := range arr {
		if i == 0 {
			arr[0] += 10
			arr[1] += 10
			arr[2] += 10
		}
		fmt.Println(v, arr[i], &v, &arr[i])
	}
}

//切片正常
func testSlice() {
	arr := []int{1, 2, 3}
	for i, v := range arr {
		if i == 0 {
			arr[0] += 10
			arr[1] += 10
			arr[2] += 10
		}
		fmt.Println(v, arr[i], &v, &arr[i])
	}
}
