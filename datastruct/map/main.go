package main

import (
	"fmt"
)

type User struct {
	age int
}

func main() {
	// 先声明map
	var m1 map[*User]int32
	// 再使用make函数创建一个非nil的map，nil map不能赋值
	m1 = make(map[*User]int32)

	u1 := &User{
		age: 1,
	}
	m1[u1] = 1
	u2 := &User{
		age: 1,
	}
	m1[u2] = 1
	fmt.Println(m1)
}
