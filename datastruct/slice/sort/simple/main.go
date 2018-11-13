package simple

import (
	"fmt"
	"sort"
)

/*slice 简单排序示例*/

func main() {
	//定义一个年龄列表
	ageList := []int{1, 3, 7, 7, 8, 2, 5}

	//排序
	sort.Slice(ageList, func(i, j int) bool {
		return ageList[i] < ageList[j]
	})
	fmt.Printf("after sort:%v", ageList)
}
