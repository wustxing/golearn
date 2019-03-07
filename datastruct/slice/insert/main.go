package main

import (
	"fmt"
)

/*slice 中间插入示例*/
func main() {
	//定义一个年龄列表
	ageList := []int{1, 3, 7}

	insertValue := 5
	//定义要插入的位置，这里表示将5插入1位置（即1，3）之间
	insertIndex := 1

	//保存后部分数据
	tailList := append([]int{}, ageList[insertIndex:]...)

	ageList = append(ageList[:insertIndex], insertValue)
	ageList = append(ageList, tailList...)
	fmt.Println(ageList)
}
