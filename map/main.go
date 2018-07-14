package main

import "log"

func main() {
	// 先声明map
	var m1 map[string]string
	// 再使用make函数创建一个非nil的map，nil map不能赋值
	m1 = make(map[string]string)
	// 最后给已声明的map赋值
	m1["a"] = "aa"
	m1["b"] = "bb"
	m1["c"] = "cc"
	log.Println(m1)

	m2 :=m1
	m2["xx"] ="qq"


	log.Println(m1,m2)
}
