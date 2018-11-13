package main

import "fmt"

/*slice 遍历删除示例*/

func main() {
	//定义一个年龄列表
	ageList := []int{1, 3, 7, 7, 8, 2, 5}

	//遍历删除6岁以下的
	for i := 0; i < len(ageList); {
		if ageList[i] < 6 {
			ageList = append(ageList[:i], ageList[i+1:]...)
		} else {
			i++
		}
	}
	fmt.Printf("after del:%v", ageList)
}
